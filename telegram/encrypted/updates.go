package encrypted

import (
	"context"
	"fmt"

	"github.com/go-faster/errors"
	"go.uber.org/multierr"
	"go.uber.org/zap"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tg"
	"github.com/gotd/td/tg/e2e"
)

func (m *Manager) Register(d tg.UpdateDispatcher) {
	d.OnEncryption(m.OnEncryption)
	d.OnNewEncryptedMessage(m.OnNewEncryptedMessage)
}

func getMessageAction(msg e2e.DecryptedMessageClass) (e2e.DecryptedMessageActionClass, bool) {
	var action e2e.DecryptedMessageActionClass
	switch msg := msg.(type) {
	case *e2e.DecryptedMessageService:
		action = msg.GetAction()
	case *e2e.DecryptedMessageService8:
		action = msg.GetAction()
	}
	return action, action != nil
}

func (m *Manager) OnNewEncryptedMessage(
	ctx context.Context,
	_ tg.Entities,
	u *tg.UpdateNewEncryptedMessage,
) (rErr error) {
	chatID := u.Message.GetChatID()

	tx, err := m.storage.Acquire(ctx, chatID)
	if err != nil {
		m.logger.Info("Received encrypted message from unknown chat", zap.Int("chat_id", chatID))
		return errors.Errorf("find chat %d: %w", chatID, err)
	}
	defer func() {
		if rErr != nil {
			multierr.AppendInto(&rErr, tx.Close(ctx))
		}
	}()
	chat := tx.Get()

	data, err := chat.decrypt(u.Message.GetBytes())
	if err != nil {
		return errors.Wrap(err, "encrypt")
	}

	layer := e2e.DecryptedMessageLayer{}
	if err := layer.Decode(&bin.Buffer{Buf: data}); err != nil {
		return errors.Wrap(err, "decode layer")
	}

	myIn, myOut := chat.seqNo()
	logger := m.logger.With(
		zap.Int("chat_id", chat.ID),
		zap.Int("his_in_seq", layer.InSeqNo),
		zap.Int("his_out_seq", layer.OutSeqNo),
		zap.Int("my_in_seq", myIn),
		zap.Int("my_out_seq", myOut),
	)

	logger.Debug("Received encrypted message",
		zap.String("type", fmt.Sprintf("%T", layer.Message)),
	)

	switch chat.consumeMessage(layer.InSeqNo, layer.OutSeqNo) {
	case skipMessage:
		logger.Debug("Skip duplicate message")
		return nil
	case fillGap:
		logger.Debug("Fill gap")
	// TODO(tdakkota): request resend.
	case abortChat:
		logger.Debug("Abort chat due to invalid in_seq")
	default:
	}

	switch action, _ := getMessageAction(layer.Message); action := action.(type) {
	case *e2e.DecryptedMessageActionResend:
		if err := m.resendMessages(ctx, action, tx); err != nil {
			return errors.Wrap(err, "resend messages")
		}
	case *e2e.DecryptedMessageActionNotifyLayer:
		if err := m.updateLayer(ctx, action, tx); err != nil {
			return errors.Wrap(err, "update layer")
		}
	// TODO(tdakkota): handle key rotation
	// case *e2e.DecryptedMessageActionRequestKey:
	// case *e2e.DecryptedMessageActionAcceptKey:
	// case *e2e.DecryptedMessageActionAbortKey:
	// case *e2e.DecryptedMessageActionCommitKey:
	default:
		if err := tx.Commit(ctx, chat); err != nil {
			return errors.Errorf("save chat %d: %w", chat.ID, err)
		}
	}

	return m.message(ctx, chat, layer.Message)
}

func (m *Manager) notifyNewChat(id int, c tg.EncryptedChatClass) {
	m.requestsMux.Lock()
	req, ok := m.requests[id]
	delete(m.requests, id)
	m.requestsMux.Unlock()
	if !ok {
		return
	}

	select {
	case req.result <- c:
	default:
	}
}

func (m *Manager) OnEncryption(ctx context.Context, e tg.Entities, update *tg.UpdateEncryption) error {
	switch c := update.Chat.(type) {
	case *tg.EncryptedChat:
		m.logger.Debug("Chat accepted", zap.Int("chat_id", c.ID))
		m.notifyNewChat(c.ID, c)
		return nil
	case *tg.EncryptedChatDiscarded:
		m.logger.Debug("Chat discarded", zap.Int("chat_id", c.ID))
		m.notifyNewChat(c.ID, c)
		return nil
	case *tg.EncryptedChatRequested:
		accepted, err := m.accept(ctx, e, c)
		if err != nil {
			return errors.Wrap(err, "accept handler")
		}

		if accepted {
			chat, err := m.acceptChat(ctx, c)
			if err != nil {
				return errors.Wrap(err, "accept")
			}

			if err := m.created(ctx, chat); err != nil {
				return errors.Wrap(err, "created handler")
			}

			return nil
		}

		return m.discardChat(ctx, c.ID, false, "Rejected by user")
	default:
		m.logger.Warn("Unexpected type", zap.String("type", fmt.Sprintf("%T", c)))
		return nil
	}
}
