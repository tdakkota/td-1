package encrypted

import (
	"context"
	"io"

	"go.uber.org/zap"

	"github.com/gotd/td/internal/crypto"
	"github.com/gotd/td/tg"
	"github.com/gotd/td/tg/e2e"
)

type (
	// RequestHandler is a chat request event handler type.
	RequestHandler func(context.Context, tg.Entities, *tg.EncryptedChatRequested) (bool, error)

	// DiscardedHandler is a chat discard event handler type.
	DiscardedHandler func(context.Context, tg.Entities, *tg.EncryptedChatDiscarded) error

	// CreatedHandler is a chat creation event handler type.
	CreatedHandler func(context.Context, Chat) error

	// MessageHandler is an encrypted message event handler type.
	MessageHandler func(context.Context, Chat, e2e.DecryptedMessageClass) error
)

// Options is Manager options.
type Options struct {
	// Request is a chat request event handler.
	Request RequestHandler
	// Discarded is a chat discard event handler.
	Discarded DiscardedHandler
	// Created is a chat creation event handler.
	Created CreatedHandler
	// Message is an encrypted message event handler.
	Message MessageHandler
	// Storage is a chat metadata storage. Defaults to InmemoryStorage.
	Storage Storage
	// Random is random source for key generation. Defaults to rand.Reader.
	Random io.Reader
	// Logger is instance of zap.Logger. No logs by default.
	Logger *zap.Logger
}

func (m *Options) setDefaults() {
	if m.Request == nil {
		m.Request = func(context.Context, tg.Entities, *tg.EncryptedChatRequested) (bool, error) {
			// Reject all.
			return false, nil
		}
	}
	if m.Discarded == nil {
		m.Discarded = func(context.Context, tg.Entities, *tg.EncryptedChatDiscarded) error {
			// No-op.
			return nil
		}
	}
	if m.Created == nil {
		m.Created = func(context.Context, Chat) error {
			// No-op.
			return nil
		}
	}
	if m.Message == nil {
		m.Message = func(context.Context, Chat, e2e.DecryptedMessageClass) error {
			// No-op.
			return nil
		}
	}
	if m.Storage == nil {
		m.Storage = NewInmemoryStorage()
	}
	if m.Random == nil {
		m.Random = crypto.DefaultRand()
	}
	if m.Logger == nil {
		m.Logger = zap.NewNop()
	}
}
