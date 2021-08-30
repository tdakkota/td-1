package encrypted

import (
	"fmt"

	"github.com/gotd/td/tg"
)

// ChatDiscardedError returned when other parity discarded encrypted chat during creation.
type ChatDiscardedError struct {
	Chat *tg.EncryptedChatDiscarded
}

// Error implements error interface.
func (c *ChatDiscardedError) Error() string {
	return fmt.Sprintf("chat %d discarded", c.Chat.ID)
}
