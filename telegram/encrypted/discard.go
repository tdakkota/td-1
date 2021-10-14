package encrypted

import (
	"context"

	"go.uber.org/multierr"
	"go.uber.org/zap"
	"golang.org/x/xerrors"

	"github.com/gotd/td/tg"
)

func (m *Manager) discardChat(ctx context.Context, id int, deleteHistory bool) (rErr error) {
	m.logger.Debug("Discard chat", zap.Int("id", id))

	if err := m.storage.Discard(ctx, id); err != nil {
		rErr = multierr.Append(rErr, xerrors.Errorf("discard chat in storage: %w", err))
	}

	if _, err := m.raw.MessagesDiscardEncryption(ctx, &tg.MessagesDiscardEncryptionRequest{
		DeleteHistory: deleteHistory,
		ChatID:        id,
	}); err != nil {
		rErr = multierr.Append(rErr, xerrors.Errorf("send discard request: %w", err))
	}

	return nil
}

func (m *Manager) DiscardChat(ctx context.Context, id int, deleteHistory bool) (rErr error) {
	return m.discardChat(ctx, id, deleteHistory)
}
