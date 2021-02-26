// Code generated by gotdgen, DO NOT EDIT.

package td

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

// String represents TL type `string#b5286e24`.
//
// See https://localhost:80/doc/constructor/string for reference.
type String struct {
}

// StringTypeID is TL type id of String.
const StringTypeID = 0xb5286e24

func (s *String) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *String) String() string {
	if s == nil {
		return "String(nil)"
	}
	type Alias String
	return fmt.Sprintf("String%+v", Alias(*s))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *String) TypeID() uint32 {
	return StringTypeID
}

// SchemaName returns MTProto type name.
func (s *String) SchemaName() string {
	return "string"
}

// Encode implements bin.Encoder.
func (s *String) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode string#b5286e24 as nil")
	}
	b.PutID(StringTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *String) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode string#b5286e24 to nil")
	}
	if err := b.ConsumeID(StringTypeID); err != nil {
		return fmt.Errorf("unable to decode string#b5286e24: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for String.
var (
	_ bin.Encoder = &String{}
	_ bin.Decoder = &String{}
)
