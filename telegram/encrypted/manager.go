package encrypted

import (
	"io"
	"sync"

	"github.com/gotd/td/telegram/internal/dh"
	"go.uber.org/zap"

	"github.com/gotd/td/tg"
)

// Manager manages encrypted chats state.
type Manager struct {
	raw      *tg.Client
	storage  ChatStorage
	messages MessageStorage

	accept  AcceptHandler
	created CreatedHandler
	message MessageHandler

	requests    map[int]request
	requestsMux sync.Mutex

	dh *dh.State

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
