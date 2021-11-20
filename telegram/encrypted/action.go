package encrypted

import (
	"context"

	"github.com/go-faster/errors"
	"go.uber.org/multierr"
	"go.uber.org/zap"

	"github.com/gotd/td/tg/e2e"
)

// See https://github.com/DrKLO/Telegram/blob/master/TMessagesProj/src/main/java/org/telegram/messenger/SecretChatHelper.java#L1318-L1324.
func (m *Manager) resendMessages(
	ctx context.Context,
	action *e2e.DecryptedMessageActionResend,
	chat Chat,
	tx ChatTx,
) error {
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

	msgs, err := m.storage.GetFrom(ctx, chat.ID, action.StartSeqNo, action.EndSeqNo)
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

func (m *Manager) updateLayer(
	ctx context.Context,
	action *e2e.DecryptedMessageActionNotifyLayer,
	chat Chat,
	tx ChatTx,
) error {
	chatID := chat.ID

	log := m.logger.With(
		zap.Int("chat_id", chatID),
		zap.Int("new_layer", action.Layer),
	)
	switch newLayer := action.Layer; {
	case newLayer < chat.Layer:
		log.Warn("Other client decreased version")
	case newLayer < minLayer:
		log.Warn("Other client sent too old layer",
			zap.Int("min_layer", minLayer),
		)
	default:
		log.Info("Updating layer", zap.Int("old_layer", chat.Layer))
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
