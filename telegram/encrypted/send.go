package encrypted

import (
	"context"
	"io"

	"github.com/go-faster/errors"
	"go.uber.org/multierr"
	"go.uber.org/zap"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/internal/crypto"
	"github.com/gotd/td/tg"
	"github.com/gotd/td/tg/e2e"
)

const (
	minLayer    = 46
	latestLayer = 101
)

func (m *Manager) Send(ctx context.Context, chatID int, msg e2e.DecryptedMessageClass) error {
	return m.send(ctx, chatID, msg)
}

func (m *Manager) sendLayer(ctx context.Context, chatID int) error {
	randomID, err := crypto.RandInt64(m.rand)
	if err != nil {
		return errors.Wrap(err, "generate random_id")
	}

	return m.send(ctx, chatID, &e2e.DecryptedMessageService{
		RandomID: randomID,
		Action: &e2e.DecryptedMessageActionNotifyLayer{
			Layer: latestLayer,
		},
	})
}

func (m *Manager) send(ctx context.Context, chatID int, msg e2e.DecryptedMessageClass) (rErr error) {
	logger := m.logger.With(zap.Int("chat_id", chatID))

	randomBytes := make([]byte, 32)
	if _, err := io.ReadFull(m.rand, randomBytes); err != nil {
		return errors.Wrap(err, "read random bytes")
	}

	tx, err := m.storage.Acquire(ctx, chatID)
	if err != nil {
		return errors.Wrap(err, "acquire")
	}
	defer func() {
		if rErr != nil {
			multierr.AppendInto(&rErr, tx.Close(ctx))
		}
	}()
	chat := tx.Get()

	layer := chat.Layer
	if layer == 0 {
		layer = minLayer
	}

	inSeq, outSeq := chat.nextMessage()
	data := e2e.DecryptedMessageLayer{
		RandomBytes: randomBytes,
		Layer:       layer,
		InSeqNo:     inSeq,
		OutSeqNo:    outSeq,
		Message:     msg,
	}

	if err := m.messages.Push(ctx, chatID, EnqueuedMessage{
		SeqNo:   outSeq,
		Message: msg,
	}); err != nil {
		return errors.Wrap(err, "push message")
	}

	logger.Debug("Send encrypted message",
		zap.Int("in_seq", data.InSeqNo),
		zap.Int("out_seq", data.OutSeqNo),
	)

	if err := tx.Commit(ctx, chat); err != nil {
		return errors.Wrap(err, "save chat")
	}

	if _, err := m.sendRaw(ctx, chat, false, &data); err != nil {
		return err
	}

	return nil
}

func (m *Manager) sendRaw(
	ctx context.Context,
	e Chat,
	silent bool,
	msg bin.Encoder,
) (tg.MessagesSentEncryptedMessageClass, error) {
	b := bin.Buffer{}

	if err := msg.Encode(&b); err != nil {
		return nil, errors.Wrap(err, "encode")
	}

	encrypted, err := e.encrypt(m.rand, b.Buf)
	if err != nil {
		return nil, errors.Wrap(err, "encrypt")
	}

	randomID, err := crypto.RandInt64(m.rand)
	if err != nil {
		return nil, errors.Wrap(err, "generate random_id")
	}

	r, err := m.raw.MessagesSendEncrypted(ctx, &tg.MessagesSendEncryptedRequest{
		Silent: silent,
		Peer: tg.InputEncryptedChat{
			ChatID:     e.ID,
			AccessHash: e.AccessHash,
		},
		RandomID: randomID,
		Data:     encrypted,
	})
	if err != nil {
		return nil, errors.Wrap(err, "send encrypted")
	}

	return r, nil
}
