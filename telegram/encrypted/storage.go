package encrypted

import (
	"context"
	"errors"
	"sync"

	"github.com/gotd/td/tg/e2e"
)

// ChatTx is a chat transaction.
type ChatTx interface {
	// Get returns chat state.
	Get() Chat
	// Commit applies changes.
	Commit(ctx context.Context, ch Chat) error
	// Close closes transaction.
	Close(ctx context.Context) error
}

// Storage stores encrypted chats.
type Storage interface {
	// Save saves chat to storage.
	Save(ctx context.Context, chat Chat) error
	// Acquire creates new chat transaction.
	Acquire(ctx context.Context, id int) (ChatTx, error)
	// FindByID finds chat by ID.
	FindByID(ctx context.Context, id int) (Chat, error)
	// Discard marks chat as discarded.
	Discard(ctx context.Context, id int) error

	// Push adds message to the queue.
	Push(ctx context.Context, chatID int, msg EnqueuedMessage) error
	// GetFrom gets all messages by seqNo >= fromSeqNo.
	GetFrom(ctx context.Context, chatID, fromSeqNo, toSeqNo int) ([]EnqueuedMessage, error)
	// DeleteUntil deletes all messages with seqNo <= untilSeqNo.
	DeleteUntil(ctx context.Context, chatID, untilSeqNo int) error
}

// EnqueuedMessage is a stored message structure.
type EnqueuedMessage struct {
	SeqNo   int
	Message e2e.DecryptedMessageClass
}

var (
	// ErrRangeInvalid denotes that range invalid or storage does not contain this messages.
	ErrRangeInvalid = errors.New("range invalid")

	// ErrChatNotFound returned when storage does not contain chat with given ID.
	ErrChatNotFound = errors.New("chat not found")
)

var _ Storage = (*InmemoryStorage)(nil)

type inmemoryChat struct {
	mux  sync.Mutex
	chat Chat
}

func (i *inmemoryChat) tx() *inmemoryTx {
	i.mux.Lock()
	return &inmemoryTx{
		guard: i,
	}
}

type inmemoryTx struct {
	guard *inmemoryChat
	once  sync.Once
}

func (i *inmemoryTx) Get() Chat {
	return i.guard.chat
}

func (i *inmemoryTx) Commit(ctx context.Context, ch Chat) error {
	i.guard.chat = ch
	i.unlock()
	return nil
}

func (i *inmemoryTx) Close(ctx context.Context) error {
	i.unlock()
	return nil
}

func (i *inmemoryTx) unlock() {
	i.once.Do(i.guard.mux.Unlock)
}

// InmemoryStorage is an in-memory implementation of ChatStorage.
type InmemoryStorage struct {
	chats map[int]*inmemoryChat
	mux   sync.Mutex

	queues    map[int][]EnqueuedMessage
	queuesMux sync.Mutex
}

// NewInmemoryStorage creates new InmemoryStorage.
func NewInmemoryStorage() *InmemoryStorage {
	return &InmemoryStorage{
		chats:  map[int]*inmemoryChat{},
		queues: map[int][]EnqueuedMessage{},
	}
}

// Acquire finds chat and locks by ID.
func (i *InmemoryStorage) Acquire(ctx context.Context, id int) (ChatTx, error) {
	i.mux.Lock()
	guard, ok := i.chats[id]
	if !ok {
		i.mux.Unlock()
		return nil, ErrChatNotFound
	}
	i.mux.Unlock()

	return guard.tx(), nil
}

// Save saves chat to storage.
func (i *InmemoryStorage) Save(ctx context.Context, chat Chat) error {
	i.mux.Lock()
	guard, ok := i.chats[chat.ID]
	if !ok {
		i.chats[chat.ID] = &inmemoryChat{
			chat: chat,
		}
		i.mux.Unlock()
		return nil
	}
	i.mux.Unlock()

	guard.mux.Lock()
	guard.chat = chat
	guard.mux.Unlock()
	return nil
}

// FindByID finds chat by ID.
func (i *InmemoryStorage) FindByID(ctx context.Context, id int) (Chat, error) {
	i.mux.Lock()
	guard, ok := i.chats[id]
	if !ok {
		i.mux.Unlock()
		return Chat{}, ErrChatNotFound
	}
	i.mux.Unlock()

	guard.mux.Lock()
	chat := guard.chat
	guard.mux.Unlock()

	return chat, nil
}

// Discard deletes chat from storage.
func (i *InmemoryStorage) Discard(ctx context.Context, id int) error {
	i.mux.Lock()
	defer i.mux.Unlock()

	delete(i.chats, id)
	return nil
}

// Push adds message to the queue.
func (i *InmemoryStorage) Push(ctx context.Context, chatID int, msg EnqueuedMessage) error {
	i.queuesMux.Lock()
	defer i.queuesMux.Unlock()

	i.queues[chatID] = append(i.queues[chatID], msg)
	return nil
}

// GetFrom gets all messages by seqNo >= fromSeqNo.
func (i *InmemoryStorage) GetFrom(ctx context.Context, chatID, fromSeqNo, toSeqNo int) (r []EnqueuedMessage, _ error) {
	i.queuesMux.Lock()
	defer i.queuesMux.Unlock()

	queue, ok := i.queues[chatID]
	if !ok {
		return nil, ErrChatNotFound
	}

	for _, msg := range queue {
		if msg.SeqNo >= fromSeqNo {
			r = append(r, msg)
		}
	}

	return r, nil
}

// DeleteUntil deletes all messages with seqNo <= untilSeqNo.
func (i *InmemoryStorage) DeleteUntil(ctx context.Context, chatID, untilSeqNo int) error {
	i.queuesMux.Lock()
	defer i.queuesMux.Unlock()
	queue, ok := i.queues[chatID]
	if !ok {
		return nil
	}
	n := 0
	for _, msg := range queue {
		if msg.SeqNo <= untilSeqNo {
			queue[n] = msg
			n++
		}
	}
	i.queues[chatID] = queue[:n]
	return nil
}
