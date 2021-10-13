package encrypted

import (
	"context"
	"errors"
	"sync"
)

// ChatTx is a chat transaction.
type ChatTx interface {
	Get() Chat
	Commit(ctx context.Context, ch Chat) error
	Rollback(ctx context.Context) error
}

// ChatStorage contains encrypted chats.
type ChatStorage interface {
	Save(ctx context.Context, chat Chat) error
	Acquire(ctx context.Context, id int) (ChatTx, error)
	FindByID(ctx context.Context, id int) (Chat, error)
	Discard(ctx context.Context, id int) error
}

// ErrChatNotFound returned when storage does not contain chat with given ID.
var ErrChatNotFound = errors.New("chat not found")

var _ ChatStorage = (*InmemoryStorage)(nil)

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

func (i *inmemoryTx) Rollback(ctx context.Context) error {
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
}

// NewInmemoryStorage creates new InmemoryStorage.
func NewInmemoryStorage() *InmemoryStorage {
	return &InmemoryStorage{
		chats: map[int]*inmemoryChat{},
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
