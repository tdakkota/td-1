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

// ReceivedNotifyMessage represents TL type `receivedNotifyMessage#a384b779`.
// Message ID, for which PUSH-notifications were cancelled.
//
// See https://core.telegram.org/constructor/receivedNotifyMessage for reference.
type ReceivedNotifyMessage struct {
	// Message ID, for which PUSH-notifications were canceled
	ID int `schemaname:"id"`
	// Reserved for future use
	Flags int `schemaname:"flags"`
}

// ReceivedNotifyMessageTypeID is TL type id of ReceivedNotifyMessage.
const ReceivedNotifyMessageTypeID = 0xa384b779

func (r *ReceivedNotifyMessage) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.ID == 0) {
		return false
	}
	if !(r.Flags == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReceivedNotifyMessage) String() string {
	if r == nil {
		return "ReceivedNotifyMessage(nil)"
	}
	type Alias ReceivedNotifyMessage
	return fmt.Sprintf("ReceivedNotifyMessage%+v", Alias(*r))
}

// FillFrom fills ReceivedNotifyMessage from given interface.
func (r *ReceivedNotifyMessage) FillFrom(from interface {
	GetID() (value int)
	GetFlags() (value int)
}) {
	r.ID = from.GetID()
	r.Flags = from.GetFlags()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (r *ReceivedNotifyMessage) TypeID() uint32 {
	return ReceivedNotifyMessageTypeID
}

// SchemaName returns MTProto type name.
func (r *ReceivedNotifyMessage) SchemaName() string {
	return "receivedNotifyMessage"
}

// Encode implements bin.Encoder.
func (r *ReceivedNotifyMessage) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode receivedNotifyMessage#a384b779 as nil")
	}
	b.PutID(ReceivedNotifyMessageTypeID)
	b.PutInt(r.ID)
	b.PutInt(r.Flags)
	return nil
}

// GetID returns value of ID field.
func (r *ReceivedNotifyMessage) GetID() (value int) {
	return r.ID
}

// GetFlags returns value of Flags field.
func (r *ReceivedNotifyMessage) GetFlags() (value int) {
	return r.Flags
}

// Decode implements bin.Decoder.
func (r *ReceivedNotifyMessage) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode receivedNotifyMessage#a384b779 to nil")
	}
	if err := b.ConsumeID(ReceivedNotifyMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode receivedNotifyMessage#a384b779: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode receivedNotifyMessage#a384b779: field id: %w", err)
		}
		r.ID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode receivedNotifyMessage#a384b779: field flags: %w", err)
		}
		r.Flags = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ReceivedNotifyMessage.
var (
	_ bin.Encoder = &ReceivedNotifyMessage{}
	_ bin.Decoder = &ReceivedNotifyMessage{}
)
