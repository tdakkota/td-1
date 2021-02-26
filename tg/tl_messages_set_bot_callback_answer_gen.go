// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is
var _ = sort.Ints

// MessagesSetBotCallbackAnswerRequest represents TL type `messages.setBotCallbackAnswer#d58f130a`.
// Set the callback answer to a user button press (bots only)
//
// See https://core.telegram.org/method/messages.setBotCallbackAnswer for reference.
type MessagesSetBotCallbackAnswerRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `schemaname:"flags"`
	// Whether to show the message as a popup instead of a toast notification
	Alert bool `schemaname:"alert"`
	// Query ID
	QueryID int64 `schemaname:"query_id"`
	// Popup to show
	//
	// Use SetMessage and GetMessage helpers.
	Message string `schemaname:"message"`
	// URL to open
	//
	// Use SetURL and GetURL helpers.
	URL string `schemaname:"url"`
	// Cache validity
	CacheTime int `schemaname:"cache_time"`
}

// MessagesSetBotCallbackAnswerRequestTypeID is TL type id of MessagesSetBotCallbackAnswerRequest.
const MessagesSetBotCallbackAnswerRequestTypeID = 0xd58f130a

func (s *MessagesSetBotCallbackAnswerRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Alert == false) {
		return false
	}
	if !(s.QueryID == 0) {
		return false
	}
	if !(s.Message == "") {
		return false
	}
	if !(s.URL == "") {
		return false
	}
	if !(s.CacheTime == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSetBotCallbackAnswerRequest) String() string {
	if s == nil {
		return "MessagesSetBotCallbackAnswerRequest(nil)"
	}
	type Alias MessagesSetBotCallbackAnswerRequest
	return fmt.Sprintf("MessagesSetBotCallbackAnswerRequest%+v", Alias(*s))
}

// FillFrom fills MessagesSetBotCallbackAnswerRequest from given interface.
func (s *MessagesSetBotCallbackAnswerRequest) FillFrom(from interface {
	GetAlert() (value bool)
	GetQueryID() (value int64)
	GetMessage() (value string, ok bool)
	GetURL() (value string, ok bool)
	GetCacheTime() (value int)
}) {
	s.Alert = from.GetAlert()
	s.QueryID = from.GetQueryID()
	if val, ok := from.GetMessage(); ok {
		s.Message = val
	}

	if val, ok := from.GetURL(); ok {
		s.URL = val
	}

	s.CacheTime = from.GetCacheTime()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *MessagesSetBotCallbackAnswerRequest) TypeID() uint32 {
	return MessagesSetBotCallbackAnswerRequestTypeID
}

// SchemaName returns MTProto type name.
func (s *MessagesSetBotCallbackAnswerRequest) SchemaName() string {
	return "messages.setBotCallbackAnswer"
}

// Encode implements bin.Encoder.
func (s *MessagesSetBotCallbackAnswerRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.setBotCallbackAnswer#d58f130a as nil")
	}
	b.PutID(MessagesSetBotCallbackAnswerRequestTypeID)
	if !(s.Alert == false) {
		s.Flags.Set(1)
	}
	if !(s.Message == "") {
		s.Flags.Set(0)
	}
	if !(s.URL == "") {
		s.Flags.Set(2)
	}
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setBotCallbackAnswer#d58f130a: field flags: %w", err)
	}
	b.PutLong(s.QueryID)
	if s.Flags.Has(0) {
		b.PutString(s.Message)
	}
	if s.Flags.Has(2) {
		b.PutString(s.URL)
	}
	b.PutInt(s.CacheTime)
	return nil
}

// SetAlert sets value of Alert conditional field.
func (s *MessagesSetBotCallbackAnswerRequest) SetAlert(value bool) {
	if value {
		s.Flags.Set(1)
		s.Alert = true
	} else {
		s.Flags.Unset(1)
		s.Alert = false
	}
}

// GetAlert returns value of Alert conditional field.
func (s *MessagesSetBotCallbackAnswerRequest) GetAlert() (value bool) {
	return s.Flags.Has(1)
}

// GetQueryID returns value of QueryID field.
func (s *MessagesSetBotCallbackAnswerRequest) GetQueryID() (value int64) {
	return s.QueryID
}

// SetMessage sets value of Message conditional field.
func (s *MessagesSetBotCallbackAnswerRequest) SetMessage(value string) {
	s.Flags.Set(0)
	s.Message = value
}

// GetMessage returns value of Message conditional field and
// boolean which is true if field was set.
func (s *MessagesSetBotCallbackAnswerRequest) GetMessage() (value string, ok bool) {
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.Message, true
}

// SetURL sets value of URL conditional field.
func (s *MessagesSetBotCallbackAnswerRequest) SetURL(value string) {
	s.Flags.Set(2)
	s.URL = value
}

// GetURL returns value of URL conditional field and
// boolean which is true if field was set.
func (s *MessagesSetBotCallbackAnswerRequest) GetURL() (value string, ok bool) {
	if !s.Flags.Has(2) {
		return value, false
	}
	return s.URL, true
}

// GetCacheTime returns value of CacheTime field.
func (s *MessagesSetBotCallbackAnswerRequest) GetCacheTime() (value int) {
	return s.CacheTime
}

// Decode implements bin.Decoder.
func (s *MessagesSetBotCallbackAnswerRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.setBotCallbackAnswer#d58f130a to nil")
	}
	if err := b.ConsumeID(MessagesSetBotCallbackAnswerRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.setBotCallbackAnswer#d58f130a: %w", err)
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.setBotCallbackAnswer#d58f130a: field flags: %w", err)
		}
	}
	s.Alert = s.Flags.Has(1)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.setBotCallbackAnswer#d58f130a: field query_id: %w", err)
		}
		s.QueryID = value
	}
	if s.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.setBotCallbackAnswer#d58f130a: field message: %w", err)
		}
		s.Message = value
	}
	if s.Flags.Has(2) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.setBotCallbackAnswer#d58f130a: field url: %w", err)
		}
		s.URL = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.setBotCallbackAnswer#d58f130a: field cache_time: %w", err)
		}
		s.CacheTime = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesSetBotCallbackAnswerRequest.
var (
	_ bin.Encoder = &MessagesSetBotCallbackAnswerRequest{}
	_ bin.Decoder = &MessagesSetBotCallbackAnswerRequest{}
)

// MessagesSetBotCallbackAnswer invokes method messages.setBotCallbackAnswer#d58f130a returning error if any.
// Set the callback answer to a user button press (bots only)
//
// Possible errors:
//  400 QUERY_ID_INVALID: The query ID is invalid
//  400 URL_INVALID: Invalid URL provided
//
// See https://core.telegram.org/method/messages.setBotCallbackAnswer for reference.
// Can be used by bots.
func (c *Client) MessagesSetBotCallbackAnswer(ctx context.Context, request *MessagesSetBotCallbackAnswerRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
