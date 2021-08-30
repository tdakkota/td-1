package encrypted

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gotd/td/internal/crypto"
	"github.com/gotd/td/internal/testutil"
	"github.com/gotd/td/tg"
	"github.com/gotd/td/tgmock"
)

type mockStorage struct {
	InmemoryStorage
}

func (m *mockStorage) Discard(ctx context.Context, id int) error {
	return testutil.TestError()
}

func TestManager_discardChat(t *testing.T) {
	ctx := context.Background()
	a := require.New(t)
	invoker := tgmock.New(t)

	m := Manager{
		raw:    tg.NewClient(invoker),
		rand:   crypto.DefaultRand(),
		logger: zap.NewNop(),
		storage: &mockStorage{
			InmemoryStorage: *NewInmemoryStorage(),
		},
	}

	// Ensure that DiscardChat makes RPC request even if removal from storage failed.
	invoker.ExpectCall(&tg.MessagesDiscardEncryptionRequest{
		DeleteHistory: false,
		ChatID:        10,
	}).ThenTrue()
	a.Error(
		m.DiscardChat(ctx, 10, false),
		"DiscardChat must send discard request even if removal from storage failed",
	)

	m.storage = NewInmemoryStorage()
	invoker.ExpectCall(&tg.MessagesDiscardEncryptionRequest{
		DeleteHistory: false,
		ChatID:        10,
	}).ThenTrue()
	a.NoError(m.DiscardChat(ctx, 10, false))

	invoker.ExpectCall(&tg.MessagesDiscardEncryptionRequest{
		DeleteHistory: false,
		ChatID:        10,
	}).ThenErr(testutil.TestError())
	a.Error(m.DiscardChat(ctx, 10, false))
}
