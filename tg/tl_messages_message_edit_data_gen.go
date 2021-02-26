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

// MessagesMessageEditData represents TL type `messages.messageEditData#26b5dde6`.
// Message edit data for media
//
// See https://core.telegram.org/constructor/messages.messageEditData for reference.
type MessagesMessageEditData struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `schemaname:"flags"`
	// Media caption, if the specified media's caption can be edited
	Caption bool `schemaname:"caption"`
}

// MessagesMessageEditDataTypeID is TL type id of MessagesMessageEditData.
const MessagesMessageEditDataTypeID = 0x26b5dde6

func (m *MessagesMessageEditData) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Flags.Zero()) {
		return false
	}
	if !(m.Caption == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessagesMessageEditData) String() string {
	if m == nil {
		return "MessagesMessageEditData(nil)"
	}
	type Alias MessagesMessageEditData
	return fmt.Sprintf("MessagesMessageEditData%+v", Alias(*m))
}

// FillFrom fills MessagesMessageEditData from given interface.
func (m *MessagesMessageEditData) FillFrom(from interface {
	GetCaption() (value bool)
}) {
	m.Caption = from.GetCaption()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (m *MessagesMessageEditData) TypeID() uint32 {
	return MessagesMessageEditDataTypeID
}

// SchemaName returns MTProto type name.
func (m *MessagesMessageEditData) SchemaName() string {
	return "messages.messageEditData"
}

// Encode implements bin.Encoder.
func (m *MessagesMessageEditData) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messages.messageEditData#26b5dde6 as nil")
	}
	b.PutID(MessagesMessageEditDataTypeID)
	if !(m.Caption == false) {
		m.Flags.Set(0)
	}
	if err := m.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.messageEditData#26b5dde6: field flags: %w", err)
	}
	return nil
}

// SetCaption sets value of Caption conditional field.
func (m *MessagesMessageEditData) SetCaption(value bool) {
	if value {
		m.Flags.Set(0)
		m.Caption = true
	} else {
		m.Flags.Unset(0)
		m.Caption = false
	}
}

// GetCaption returns value of Caption conditional field.
func (m *MessagesMessageEditData) GetCaption() (value bool) {
	return m.Flags.Has(0)
}

// Decode implements bin.Decoder.
func (m *MessagesMessageEditData) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messages.messageEditData#26b5dde6 to nil")
	}
	if err := b.ConsumeID(MessagesMessageEditDataTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.messageEditData#26b5dde6: %w", err)
	}
	{
		if err := m.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.messageEditData#26b5dde6: field flags: %w", err)
		}
	}
	m.Caption = m.Flags.Has(0)
	return nil
}

// Ensuring interfaces in compile-time for MessagesMessageEditData.
var (
	_ bin.Encoder = &MessagesMessageEditData{}
	_ bin.Decoder = &MessagesMessageEditData{}
)
