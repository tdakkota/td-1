// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
)

// MessagesMigrateChatRequest represents TL type `messages.migrateChat#a2875319`.
// Turn a legacy group into a supergroup¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/messages.migrateChat for reference.
type MessagesMigrateChatRequest struct {
	// Legacy group to migrate
	ChatID int64
}

// MessagesMigrateChatRequestTypeID is TL type id of MessagesMigrateChatRequest.
const MessagesMigrateChatRequestTypeID = 0xa2875319

// Ensuring interfaces in compile-time for MessagesMigrateChatRequest.
var (
	_ bin.Encoder     = &MessagesMigrateChatRequest{}
	_ bin.Decoder     = &MessagesMigrateChatRequest{}
	_ bin.BareEncoder = &MessagesMigrateChatRequest{}
	_ bin.BareDecoder = &MessagesMigrateChatRequest{}
)

func (m *MessagesMigrateChatRequest) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessagesMigrateChatRequest) String() string {
	if m == nil {
		return "MessagesMigrateChatRequest(nil)"
	}
	type Alias MessagesMigrateChatRequest
	return fmt.Sprintf("MessagesMigrateChatRequest%+v", Alias(*m))
}

// FillFrom fills MessagesMigrateChatRequest from given interface.
func (m *MessagesMigrateChatRequest) FillFrom(from interface {
	GetChatID() (value int64)
}) {
	m.ChatID = from.GetChatID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesMigrateChatRequest) TypeID() uint32 {
	return MessagesMigrateChatRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesMigrateChatRequest) TypeName() string {
	return "messages.migrateChat"
}

// TypeInfo returns info about TL type.
func (m *MessagesMigrateChatRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.migrateChat",
		ID:   MessagesMigrateChatRequestTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessagesMigrateChatRequest) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messages.migrateChat#a2875319 as nil")
	}
	b.PutID(MessagesMigrateChatRequestTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessagesMigrateChatRequest) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messages.migrateChat#a2875319 as nil")
	}
	b.PutLong(m.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessagesMigrateChatRequest) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messages.migrateChat#a2875319 to nil")
	}
	if err := b.ConsumeID(MessagesMigrateChatRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.migrateChat#a2875319: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessagesMigrateChatRequest) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messages.migrateChat#a2875319 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.migrateChat#a2875319: field chat_id: %w", err)
		}
		m.ChatID = value
	}
	return nil
}

// GetChatID returns value of ChatID field.
func (m *MessagesMigrateChatRequest) GetChatID() (value int64) {
	return m.ChatID
}

// MessagesMigrateChat invokes method messages.migrateChat#a2875319 returning error if any.
// Turn a legacy group into a supergroup¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// Possible errors:
//  403 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//  400 CHAT_ID_INVALID: The provided chat id is invalid
//  400 PEER_ID_INVALID: The provided peer id is invalid
//
// See https://core.telegram.org/method/messages.migrateChat for reference.
func (c *Client) MessagesMigrateChat(ctx context.Context, chatid int64) (UpdatesClass, error) {
	var result UpdatesBox

	request := &MessagesMigrateChatRequest{
		ChatID: chatid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
