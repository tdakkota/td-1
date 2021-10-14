package encrypted

import (
	"io"
	"sync"

	"go.uber.org/zap"

	"github.com/gotd/td/tg"
)

// Manager manages encrypted chats state.
type Manager struct {
	raw     *tg.Client
	storage ChatStorage

	accept  AcceptHandler
	created CreatedHandler
	message MessageHandler

	requests    map[int]request
	requestsMux sync.Mutex

	cfg    dhConfig
	cfgMux sync.Mutex

	rand   io.Reader
	logger *zap.Logger
}

// NewManager creates new Manager.
func NewManager(raw *tg.Client, d tg.UpdateDispatcher, opts Options) *Manager {
	opts.setDefaults()

	m := &Manager{
		raw:      raw,
		storage:  opts.Storage,
		accept:   opts.Accept,
		created:  opts.Created,
		message:  opts.Message,
		requests: map[int]request{},
		cfg:      dhConfig{},
		rand:     opts.Random,
		logger:   opts.Logger,
	}
	m.Register(d)
	return m
}
