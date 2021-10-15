package encrypted

import (
	"context"

	"go.uber.org/multierr"
	"go.uber.org/zap"
	"golang.org/x/xerrors"

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
		if xerrors.Is(err, ErrRangeInvalid) {

			return multierr.Append(
				err,
				m.discardChat(ctx, chat.ID, false, "Other party requested invalid range"),
			)
		}
		return xerrors.Errorf("get message range: %w", err)
	}

	if err := tx.Commit(ctx, chat); err != nil {
		return xerrors.Errorf("save chat %d: %w", chat.ID, err)
	}

	for i, msg := range msgs {
		if err := m.send(ctx, chat.ID, msg.Message); err != nil {
			return xerrors.Errorf("send requested message %d to %d: %w", i, chat.ID, err)
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
		return xerrors.Errorf("save chat %d: %w", chat.ID, err)
	}

	if action.Layer < latestLayer {
		return m.sendLayer(ctx, chatID)
	}

	return nil
}
