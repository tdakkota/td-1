package encrypted

import (
	"context"
	"errors"

	"github.com/gotd/td/tg/e2e"
)

// EnqueuedMessage is a stored message structure.
type EnqueuedMessage struct {
	SeqNo   int
	Message e2e.DecryptedMessageClass
}

// ErrRangeInvalid denotes that range invalid or storage does not contain this messages.
var ErrRangeInvalid = errors.New("range invalid")

// MessageStorage is a storage of message queues.
type MessageStorage interface {
	// Push adds message to the queue.
	Push(ctx context.Context, chatID int, msg EnqueuedMessage) error
	// GetFrom gets all messages by seqNo >= fromSeqNo.
	GetFrom(ctx context.Context, chatID, fromSeqNo, toSeqNo int) ([]EnqueuedMessage, error)
	// DeleteUntil deletes all messages with seqNo <= untilSeqNo.
	DeleteUntil(ctx context.Context, chatID, untilSeqNo int) error
}
