package encrypted

import (
	"context"
	"math/big"

	"github.com/go-faster/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/gotd/td/internal/crypto"
	"github.com/gotd/td/tg/e2e"
)

func (m *Manager) sendRequestKey(
	ctx context.Context,
	chat Chat,
	tx ChatTx,
) error {
	chatID := chat.ID

	a, err := m.dh.Int()
	if err != nil {
		return errors.Wrap(err, "generate a")
	}

	g := chat.GBig
	dhPrime := chat.P
	gA := big.NewInt(0).Exp(g, a, dhPrime)
	chat.GAorB = gA

	exchangeID, err := crypto.RandInt64(m.rand)
	if err != nil {
		return err
	}

	if err := tx.Commit(ctx, chat); err != nil {
		return errors.Errorf("save chat %d: %w", chatID, err)
	}

	return m.sendAction(ctx, chatID, &e2e.DecryptedMessageActionRequestKey{
		ExchangeID: exchangeID,
		GA: gA.Bytes(),
	})
}

func (m *Manager) onRequestKey(
	ctx context.Context,
	action *e2e.DecryptedMessageActionRequestKey,
	chat Chat,
	tx ChatTx,
) error {
	chatID := chat.ID

	storedID := chat.ExchangeID
	if storedID != 0 {
		log := m.logger.With(
			zap.Int("chat_id", chatID),
			zap.Int64("stored_exchange_id", storedID),
			zap.Int64("received_exchange_id", action.ExchangeID),
		)

		if storedID > action.ExchangeID {
			log.Debug("Got older re-keying request")
		} else {
			log.Warn("Got unexpected re-keying request")
			chat.resetExchange()
		}

		// We don't need transaction anymore.
		if err := tx.Commit(ctx, chat); err != nil {
			return errors.Errorf("save chat %d: %w", chatID, err)
		}

		if storedID <= action.ExchangeID {
			if err := m.sendAbortKey(ctx, chatID, action.ExchangeID); err != nil {
				return errors.Wrap(err, "abort key exchange")
			}
		}

		return nil
	}

	b, err := m.dh.Int()
	if err != nil {
		return errors.Wrap(err, "generate b")
	}

	key, gB, err := acceptKey(chat.GBig, chat.P, b, action.GA)
	if err != nil {
		return err
	}

	chat.ExchangeID = action.ExchangeID
	chat.NextKey = key
	chat.GAorB = gB

	// We don't need transaction anymore.
	if err := tx.Commit(ctx, chat); err != nil {
		return errors.Errorf("save chat %d: %w", chatID, err)
	}

	if err := m.sendAcceptKey(ctx, chatID, action.ExchangeID, gB, key.IntID()); err != nil {
		return errors.Wrap(err, "accept key")
	}
	return nil
}

func (m *Manager) sendAcceptKey(
	ctx context.Context,
	chatID int, exchangeID int64,
	gb *big.Int, fingerprint int64,
) error {
	// TODO(tdakkota): move out handlers and helpers from manager?
	return m.sendAction(ctx, chatID, &e2e.DecryptedMessageActionAcceptKey{
		ExchangeID:     exchangeID,
		GB:             gb.Bytes(),
		KeyFingerprint: fingerprint,
	})
}

func (m *Manager) onAcceptKey(
	ctx context.Context,
	action *e2e.DecryptedMessageActionAcceptKey,
	chat Chat,
	tx ChatTx,
) error {
	chatID := chat.ID

	abort := func(err error, level zapcore.Level, msg string, fields ...zap.Field) error {
		if ce := m.logger.With(zap.Int("chat_id", chatID)).Check(level, msg); ce != nil {
			ce.Write(fields...)
		}

		chat.resetExchange()
		// We don't need transaction anymore.
		if err := tx.Commit(ctx, chat); err != nil {
			return errors.Errorf("save chat %d: %w", chatID, err)
		}

		if err := m.sendAbortKey(ctx, chatID, action.ExchangeID); err != nil {
			return errors.Wrap(err, "abort key exchange")
		}

		return err
	}

	// ActionAcceptKey should be sent only during key rotation.
	//
	// If exchange ID is wrong, drop the exchange.
	if action.ExchangeID != chat.ExchangeID {
		return abort(
			nil,
			zap.DebugLevel, "Got invalid accept key event",
			zap.Int64("exchange_id", action.ExchangeID),
		)
	}

	if chat.GAorB == nil {
		return abort(
			errors.New("unexpected nil g_a or g_b"),
			zap.WarnLevel, "Got nil g_a or g_b: check storage implementation",
		)
	}

	gB := big.NewInt(0).SetBytes(action.GB)

	// key := pow(g_b, a) mod dh_prime
	k := crypto.Key{}
	if !crypto.FillBytes(big.NewInt(0).Exp(gB, chat.GAorB, chat.P), k[:]) {
		return errors.New("auth key is too big")
	}
	key := k.WithID()
	fingerprint := key.IntID()

	if fingerprint != action.KeyFingerprint {
		return abort(
			nil,
			zap.WarnLevel, "Got wrong fingerprint during exchange",
		)
	}

	chat.Key = key
	// TODO(tdakkota): is this correct?
	chat.GAorB = nil

	// We don't need transaction anymore.
	if err := tx.Commit(ctx, chat); err != nil {
		return errors.Errorf("save chat %d: %w", chatID, err)
	}

	if err := m.sendCommitKey(ctx, chatID, action.ExchangeID, fingerprint); err != nil {
		return errors.Wrap(err, "commit key")
	}
	return nil
}

func (m *Manager) sendCommitKey(
	ctx context.Context,
	chatID int,
	exchangeID, fingerprint int64,
) error {
	// TODO(tdakkota): move out handlers and helpers from manager?
	return m.sendAction(ctx, chatID, &e2e.DecryptedMessageActionCommitKey{
		ExchangeID:     exchangeID,
		KeyFingerprint: fingerprint,
	})
}

func (m *Manager) onCommitKey(
	ctx context.Context,
	action *e2e.DecryptedMessageActionCommitKey,
	chat Chat,
	tx ChatTx,
) error {
	chatID := chat.ID

	if action.ExchangeID != chat.ExchangeID ||
		chat.NextKey.IntID() != action.KeyFingerprint {
		m.logger.Warn("Got invalid commit key request", zap.Int("chat_id", chatID))

		chat.resetExchange()
		// We don't need transaction anymore.
		if err := tx.Commit(ctx, chat); err != nil {
			return errors.Errorf("save chat %d: %w", chatID, err)
		}

		if err := m.sendAbortKey(ctx, chatID, action.ExchangeID); err != nil {
			return errors.Wrap(err, "abort key exchange")
		}

		return nil
	}

	// Commit key.
	chat.Key = chat.NextKey
	chat.resetExchange()

	if err := tx.Commit(ctx, chat); err != nil {
		return errors.Errorf("save chat %d: %w", chatID, err)
	}

	// Send a noop message to complete key exchange.
	if err := m.sendNoop(ctx, chatID); err != nil {
		return errors.Wrap(err, "send noop")
	}

	return nil
}

func (m *Manager) sendAbortKey(ctx context.Context, chatID int, exchangeID int64) error {
	return m.sendAction(ctx, chatID, &e2e.DecryptedMessageActionAbortKey{
		ExchangeID: exchangeID,
	})
}

func (m *Manager) onAbortKey(
	ctx context.Context,
	action *e2e.DecryptedMessageActionAbortKey,
	chat Chat,
	tx ChatTx,
) error {
	chatID := chat.ID

	if action.ExchangeID == chat.ExchangeID {
		chat.resetExchange()
	}

	if err := tx.Commit(ctx, chat); err != nil {
		return errors.Errorf("save chat %d: %w", chatID, err)
	}
	return nil
}
