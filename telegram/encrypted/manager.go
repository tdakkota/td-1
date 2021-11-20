package encrypted

import (
	"io"

	"go.uber.org/zap"

	"github.com/gotd/td/telegram/internal/dh"

	"github.com/gotd/td/tg"
)

// Manager manages encrypted chats state.
type Manager struct {
	raw     *tg.Client
	storage Storage

	accept  RequestHandler
	discard DiscardedHandler
	created CreatedHandler
	message MessageHandler

	dh *dh.State

	rand   io.Reader
	logger *zap.Logger
}

// NewManager creates new Manager.
func NewManager(raw *tg.Client, d tg.UpdateDispatcher, opts Options) *Manager {
	opts.setDefaults()

	m := &Manager{
		raw:     raw,
		storage: opts.Storage,
		accept:  opts.Request,
		created: opts.Created,
		discard: opts.Discarded,
		message: opts.Message,
		dh: dh.NewState(raw, dh.Options{
			Random: opts.Random,
			Logger: opts.Logger,
		}),
		rand:   opts.Random,
		logger: opts.Logger,
	}
	m.Register(d)
	return m
}
