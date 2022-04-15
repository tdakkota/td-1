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
	"github.com/gotd/td/tdjson"
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
	_ = tdjson.Encoder{}
)

// MessagesSendWebViewResultMessageRequest represents TL type `messages.sendWebViewResultMessage#a4314f5`.
//
// See https://core.telegram.org/method/messages.sendWebViewResultMessage for reference.
type MessagesSendWebViewResultMessageRequest struct {
	// BotQueryID field of MessagesSendWebViewResultMessageRequest.
	BotQueryID string
	// Result field of MessagesSendWebViewResultMessageRequest.
	Result InputBotInlineResultClass
}

// MessagesSendWebViewResultMessageRequestTypeID is TL type id of MessagesSendWebViewResultMessageRequest.
const MessagesSendWebViewResultMessageRequestTypeID = 0xa4314f5

// Ensuring interfaces in compile-time for MessagesSendWebViewResultMessageRequest.
var (
	_ bin.Encoder     = &MessagesSendWebViewResultMessageRequest{}
	_ bin.Decoder     = &MessagesSendWebViewResultMessageRequest{}
	_ bin.BareEncoder = &MessagesSendWebViewResultMessageRequest{}
	_ bin.BareDecoder = &MessagesSendWebViewResultMessageRequest{}
)

func (s *MessagesSendWebViewResultMessageRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.BotQueryID == "") {
		return false
	}
	if !(s.Result == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSendWebViewResultMessageRequest) String() string {
	if s == nil {
		return "MessagesSendWebViewResultMessageRequest(nil)"
	}
	type Alias MessagesSendWebViewResultMessageRequest
	return fmt.Sprintf("MessagesSendWebViewResultMessageRequest%+v", Alias(*s))
}

// FillFrom fills MessagesSendWebViewResultMessageRequest from given interface.
func (s *MessagesSendWebViewResultMessageRequest) FillFrom(from interface {
	GetBotQueryID() (value string)
	GetResult() (value InputBotInlineResultClass)
}) {
	s.BotQueryID = from.GetBotQueryID()
	s.Result = from.GetResult()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSendWebViewResultMessageRequest) TypeID() uint32 {
	return MessagesSendWebViewResultMessageRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSendWebViewResultMessageRequest) TypeName() string {
	return "messages.sendWebViewResultMessage"
}

// TypeInfo returns info about TL type.
func (s *MessagesSendWebViewResultMessageRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.sendWebViewResultMessage",
		ID:   MessagesSendWebViewResultMessageRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "BotQueryID",
			SchemaName: "bot_query_id",
		},
		{
			Name:       "Result",
			SchemaName: "result",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesSendWebViewResultMessageRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendWebViewResultMessage#a4314f5 as nil")
	}
	b.PutID(MessagesSendWebViewResultMessageRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSendWebViewResultMessageRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendWebViewResultMessage#a4314f5 as nil")
	}
	b.PutString(s.BotQueryID)
	if s.Result == nil {
		return fmt.Errorf("unable to encode messages.sendWebViewResultMessage#a4314f5: field result is nil")
	}
	if err := s.Result.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendWebViewResultMessage#a4314f5: field result: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSendWebViewResultMessageRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendWebViewResultMessage#a4314f5 to nil")
	}
	if err := b.ConsumeID(MessagesSendWebViewResultMessageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.sendWebViewResultMessage#a4314f5: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSendWebViewResultMessageRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendWebViewResultMessage#a4314f5 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendWebViewResultMessage#a4314f5: field bot_query_id: %w", err)
		}
		s.BotQueryID = value
	}
	{
		value, err := DecodeInputBotInlineResult(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendWebViewResultMessage#a4314f5: field result: %w", err)
		}
		s.Result = value
	}
	return nil
}

// GetBotQueryID returns value of BotQueryID field.
func (s *MessagesSendWebViewResultMessageRequest) GetBotQueryID() (value string) {
	if s == nil {
		return
	}
	return s.BotQueryID
}

// GetResult returns value of Result field.
func (s *MessagesSendWebViewResultMessageRequest) GetResult() (value InputBotInlineResultClass) {
	if s == nil {
		return
	}
	return s.Result
}

// MessagesSendWebViewResultMessage invokes method messages.sendWebViewResultMessage#a4314f5 returning error if any.
//
// See https://core.telegram.org/method/messages.sendWebViewResultMessage for reference.
func (c *Client) MessagesSendWebViewResultMessage(ctx context.Context, request *MessagesSendWebViewResultMessageRequest) (*WebViewMessageSent, error) {
	var result WebViewMessageSent

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
