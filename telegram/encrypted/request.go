package encrypted

import (
	"context"
	"math/big"

	"go.uber.org/multierr"
	"go.uber.org/zap"
	"golang.org/x/xerrors"

	"github.com/gotd/td/internal/crypto"
	"github.com/gotd/td/tg"
)

type request struct {
	result chan tg.EncryptedChatClass
}

// RequestChat requests new encrypted chat and returns chat ID.
//
// See https://core.telegram.org/api/end-to-end#sending-a-request.
func (m *Manager) RequestChat(ctx context.Context, user tg.InputUserClass) (int, error) {
	a, dhCfg, err := m.initDH(ctx)
	if err != nil {
		return 0, xerrors.Errorf("init DH: %w", err)
	}

	g := dhCfg.GBig
	dhPrime := dhCfg.P
	gA := big.NewInt(0).Exp(g, a, dhPrime)

	randomID, err := crypto.RandInt64(m.rand)
	if err != nil {
		return 0, xerrors.Errorf("generate random ID: %w", err)
	}

	m.logger.Debug("Request chat", zap.Int64("random_id", randomID))
	m.requestsMux.Lock()
	requested, err := m.raw.MessagesRequestEncryption(ctx, &tg.MessagesRequestEncryptionRequest{
		UserID:   user,
		RandomID: int(randomID),
		GA:       gA.Bytes(),
	})
	if err != nil {
		m.requestsMux.Unlock()
		return 0, xerrors.Errorf("request chat: %w", err)
	}
	chatID := requested.GetID()

	result := make(chan tg.EncryptedChatClass, 1)
	m.requests[chatID] = request{
		result: result,
	}
	m.requestsMux.Unlock()
	defer func() {
		m.requestsMux.Lock()
		delete(m.requests, chatID)
		m.requestsMux.Unlock()
	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case r := <-result:
		switch c := r.(type) {
		case *tg.EncryptedChat:
			gB := big.NewInt(0).SetBytes(c.GAOrB)

			// key := pow(g_b, a) mod dh_prime
			k := crypto.Key{}
			if !crypto.FillBytes(big.NewInt(0).Exp(gB, a, dhPrime), k[:]) {
				return 0, xerrors.New("auth key is too big")
			}
			key := k.WithID()

			if getKeyFingerprint(key) != c.KeyFingerprint {
				err := xerrors.New("key fingerprint mismatch")
				return 0, multierr.Append(err, m.DiscardChat(ctx, chatID, false))
			}

			created := Chat{
				ID:            c.ID,
				AccessHash:    c.AccessHash,
				Layer:         0,
				Date:          c.Date,
				AdminID:       c.AdminID,
				ParticipantID: c.ParticipantID,
				Originator:    true,
				InSeq:         0,
				OutSeq:        0,
				HisInSeq:      0,
				Key:           key,
			}

			if err := m.storage.Save(ctx, created); err != nil {
				return 0, xerrors.Errorf("save chat: %w", err)
			}

			if err := m.sendLayer(ctx, chatID); err != nil {
				return 0, xerrors.Errorf("notify layer: %w", err)
			}

			return chatID, nil
		case *tg.EncryptedChatDiscarded:
			return 0, &ChatDiscardedError{Chat: c}
		default:
			return 0, xerrors.Errorf("unexpected type %T", c)
		}
	}
}
