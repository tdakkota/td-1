package encrypted

import (
	"context"
	"math/big"

	"github.com/go-faster/errors"
	"go.uber.org/zap"

	"github.com/gotd/td/internal/crypto"
	"github.com/gotd/td/tg"
)

// acceptChat generates key and accepts chat request.
//
// See https://core.telegram.org/api/end-to-end#accepting-a-request.
func (m *Manager) acceptChat(ctx context.Context, req *tg.EncryptedChatRequested) (Chat, error) {
	m.logger.Debug("Accept chat", zap.Int("id", req.ID))

	b, dhCfg, err := m.dh.Init(ctx)
	if err != nil {
		return Chat{}, errors.Wrap(err, "init DH")
	}

	g := dhCfg.GBig
	dhPrime := dhCfg.P
	// g_b := pow(g, b) mod dh_prime
	gB := big.NewInt(0).Exp(g, b, dhPrime)

	gA := big.NewInt(0).SetBytes(req.GA)
	// key := pow(g_a, b) mod dh_prime
	k := crypto.Key{}

	if !crypto.FillBytes(big.NewInt(0).Exp(gA, b, dhPrime), k[:]) {
		return Chat{}, errors.New("auth key is too big")
	}
	key := k.WithID()

	c, err := m.raw.MessagesAcceptEncryption(ctx, &tg.MessagesAcceptEncryptionRequest{
		Peer: tg.InputEncryptedChat{
			ChatID:     req.ID,
			AccessHash: req.AccessHash,
		},
		GB:             gB.Bytes(),
		KeyFingerprint: getKeyFingerprint(key),
	})
	if err != nil {
		return Chat{}, err
	}

	switch chat := c.(type) {
	case *tg.EncryptedChat:
		var accepted Chat
		accepted.init(chat, false, key, dhCfg)

		if err := m.storage.Save(ctx, accepted); err != nil {
			return Chat{}, errors.Wrap(err, "save chat")
		}

		if err := m.sendLayer(ctx, accepted.ID); err != nil {
			return Chat{}, errors.Wrap(err, "notify layer")
		}

		return accepted, nil
	case *tg.EncryptedChatDiscarded:
		return Chat{}, &ChatDiscardedError{Chat: chat}
	default:
		return Chat{}, errors.Errorf("unexpected type %T", chat)
	}
}
