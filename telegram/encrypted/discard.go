package encrypted

import (
	"context"

	"github.com/go-faster/errors"
	"go.uber.org/multierr"
	"go.uber.org/zap"

	"github.com/gotd/td/tg"
)

func (m *Manager) discardChat(ctx context.Context, id int, deleteHistory bool, reason string) (rErr error) {
	m.logger.Debug("Discard chat",
		zap.Int("id", id),
		zap.Bool("delete_history", deleteHistory),
		zap.String("reason", reason),
	)

	if err := m.storage.Discard(ctx, id); err != nil {
		rErr = multierr.Append(rErr, errors.Wrap(err, "discard chat in storage"))
	}

	if _, err := m.raw.MessagesDiscardEncryption(ctx, &tg.MessagesDiscardEncryptionRequest{
		DeleteHistory: deleteHistory,
		ChatID:        id,
	}); err != nil {
		rErr = multierr.Append(rErr, errors.Wrap(err, "send discard request"))
	}

	return rErr
}

func (m *Manager) DiscardChat(ctx context.Context, id int, deleteHistory bool) (rErr error) {
	return m.discardChat(ctx, id, deleteHistory, "requested by user")
}
