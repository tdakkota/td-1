package encrypted

import (
	"context"
	"math/big"

	"github.com/go-faster/errors"
	"go.uber.org/multierr"
	"go.uber.org/zap"

	"github.com/gotd/td/internal/crypto"
	"github.com/gotd/td/tg"
)

func (m *Manager) confirmChat(ctx context.Context, c *tg.EncryptedChat) (ch Chat, rErr error) {
	tx, err := m.storage.Acquire(ctx, c.ID)
	if err != nil {
		return Chat{}, errors.Wrap(err, "acquire")
	}
	defer func() {
		if rErr != nil {
			multierr.AppendInto(&rErr, tx.Close(ctx))
		}
	}()
	chat := tx.Get()
	chatID := chat.ID

	gB := big.NewInt(0).SetBytes(c.GAOrB)

	// key := pow(g_b, a) mod dh_prime
	k := crypto.Key{}
	if !crypto.FillBytes(big.NewInt(0).Exp(gB, chat.A, chat.P), k[:]) {
		return Chat{}, errors.New("auth key is too big")
	}
	key := k.WithID()

	if key.IntID() != c.KeyFingerprint {
		err := errors.New("key fingerprint mismatch")
		return Chat{}, multierr.Append(err, m.DiscardChat(ctx, chatID, false))
	}

	chat.A = nil
	chat.Key = key
	if err := tx.Commit(ctx, chat); err != nil {
		return Chat{}, errors.Errorf("save chat %d: %w", chatID, err)
	}

	if err := m.sendLayer(ctx, chatID); err != nil {
		return Chat{}, errors.Wrap(err, "notify layer")
	}

	return ch, nil
}

// RequestChat requests new encrypted chat and returns chat ID.
//
// See https://core.telegram.org/api/end-to-end#sending-a-request.
func (m *Manager) RequestChat(ctx context.Context, user tg.InputUserClass) (int, error) {
	a, dhCfg, err := m.dh.Init(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "init DH")
	}

	g := dhCfg.GBig
	dhPrime := dhCfg.P
	gA := big.NewInt(0).Exp(g, a, dhPrime)

	randomID, err := crypto.RandInt64(m.rand)
	if err != nil {
		return 0, errors.Wrap(err, "generate random ID")
	}

	m.logger.Debug("Request chat", zap.Int64("random_id", randomID))
	result, err := m.raw.MessagesRequestEncryption(ctx, &tg.MessagesRequestEncryptionRequest{
		UserID:   user,
		RandomID: int(randomID),
		GA:       gA.Bytes(),
	})
	if err != nil {
		return 0, errors.Wrap(err, "request chat")
	}

	requested, ok := result.(*tg.EncryptedChatWaiting)
	if !ok {
		return 0, errors.Errorf("unexpected type %T", result)
	}

	var chat Chat
	chat.requested(requested, a, dhCfg)

	if err := m.storage.Save(ctx, chat); err != nil {
		return 0, errors.Wrap(err, "save requested chat")
	}

	return result.GetID(), nil
}
