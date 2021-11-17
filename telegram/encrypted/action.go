package encrypted

import (
	"context"
	"math/big"

	"github.com/go-faster/errors"
	"go.uber.org/multierr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/gotd/td/internal/crypto"
	"github.com/gotd/td/tg/e2e"
)

// See https://github.com/DrKLO/Telegram/blob/master/TMessagesProj/src/main/java/org/telegram/messenger/SecretChatHelper.java#L1318-L1324.
func (m *Manager) resendMessages(ctx context.Context, action *e2e.DecryptedMessageActionResend, tx ChatTx) error {
	chat := tx.Get()
	m.logger.Debug("Re-sending messages",
		zap.Int("chat_id", chat.ID),
		zap.Int("start_seq_no", action.StartSeqNo),
		zap.Int("end_seq_no", action.EndSeqNo),
	)

	if action.EndSeqNo < chat.InSeq || action.EndSeqNo < action.StartSeqNo {
		return m.discardChat(ctx, chat.ID, false, "Other party requested invalid range")
	}

	if action.StartSeqNo < chat.InSeq {
		action.StartSeqNo = chat.InSeq
	}

	msgs, err := m.messages.GetFrom(ctx, chat.ID, action.StartSeqNo, action.EndSeqNo)
	if err != nil {
		if errors.Is(err, ErrRangeInvalid) {

			return multierr.Append(
				err,
				m.discardChat(ctx, chat.ID, false, "Other party requested invalid range"),
			)
		}
		return errors.Wrap(err, "get message range")
	}

	if err := tx.Commit(ctx, chat); err != nil {
		return errors.Errorf("save chat %d: %w", chat.ID, err)
	}

	for i, msg := range msgs {
		if err := m.send(ctx, chat.ID, msg.Message); err != nil {
			return errors.Errorf("send requested message %d to %d: %w", i, chat.ID, err)
		}
	}

	return nil
}

func (m *Manager) updateLayer(ctx context.Context, action *e2e.DecryptedMessageActionNotifyLayer, tx ChatTx) error {
	chat := tx.Get()
	chatID := chat.ID

	switch newLayer := action.Layer; {
	case newLayer < chat.Layer:
		m.logger.Warn("Other client decreased version",
			zap.Int("chat_id", chatID),
			zap.Int("layer", newLayer),
		)
	case newLayer < minLayer:
		m.logger.Warn("Other client sent too old layer",
			zap.Int("chat_id", chatID),
			zap.Int("layer", newLayer),
			zap.Int("min_layer", minLayer),
		)
	default:
		chat.Layer = action.Layer
	}

	if err := tx.Commit(ctx, chat); err != nil {
		return errors.Errorf("save chat %d: %w", chatID, err)
	}

	if action.Layer < latestLayer {
		return m.sendLayer(ctx, chatID)
	}

	return nil
}

func (m *Manager) abortKeyExchange(ctx context.Context, chatID int, exchangeID int64) error {
	return m.sendAction(ctx, chatID, &e2e.DecryptedMessageActionAbortKey{
		ExchangeID: exchangeID,
	})
}

func (m *Manager) acceptKeyExchange(
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

func (m *Manager) requestKey(ctx context.Context, action *e2e.DecryptedMessageActionRequestKey, tx ChatTx) error {
	chat := tx.Get()
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
			if err := m.abortKeyExchange(ctx, chatID, action.ExchangeID); err != nil {
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

	if err := m.acceptKeyExchange(ctx, chatID, action.ExchangeID, gB, key.IntID()); err != nil {
		return errors.Wrap(err, "accept key")
	}
	return nil
}

func (m *Manager) commitKeyExchange(
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

func (m *Manager) acceptKey(ctx context.Context, action *e2e.DecryptedMessageActionAcceptKey, tx ChatTx) error {
	chat := tx.Get()
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

		if err := m.abortKeyExchange(ctx, chatID, action.ExchangeID); err != nil {
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

	if err := m.commitKeyExchange(ctx, chatID, action.ExchangeID, fingerprint); err != nil {
		return errors.Wrap(err, "commit key")
	}
	return nil
}

func (m *Manager) commitKey(ctx context.Context, action *e2e.DecryptedMessageActionCommitKey, tx ChatTx) error {
	chat := tx.Get()
	chatID := chat.ID

	if action.ExchangeID != chat.ExchangeID ||
		chat.NextKey.IntID() != action.KeyFingerprint {
		m.logger.Warn("Got invalid commit key request", zap.Int("chat_id", chatID))

		chat.resetExchange()
		// We don't need transaction anymore.
		if err := tx.Commit(ctx, chat); err != nil {
			return errors.Errorf("save chat %d: %w", chatID, err)
		}

		if err := m.abortKeyExchange(ctx, chatID, action.ExchangeID); err != nil {
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
