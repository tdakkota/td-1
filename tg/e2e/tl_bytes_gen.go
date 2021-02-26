// Code generated by gotdgen, DO NOT EDIT.

package e2e

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

// Bytes represents TL type `bytes#e937bb82`.
//
// See https://core.telegram.org/constructor/bytes for reference.
type Bytes struct {
}

// BytesTypeID is TL type id of Bytes.
const BytesTypeID = 0xe937bb82

func (b *Bytes) Zero() bool {
	if b == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (b *Bytes) String() string {
	if b == nil {
		return "Bytes(nil)"
	}
	type Alias Bytes
	return fmt.Sprintf("Bytes%+v", Alias(*b))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (b *Bytes) TypeID() uint32 {
	return BytesTypeID
}

// SchemaName returns MTProto type name.
func (b *Bytes) SchemaName() string {
	return "bytes"
}

// Encode implements bin.Encoder.
func (b *Bytes) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode bytes#e937bb82 as nil")
	}
	buf.PutID(BytesTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (b *Bytes) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode bytes#e937bb82 to nil")
	}
	if err := buf.ConsumeID(BytesTypeID); err != nil {
		return fmt.Errorf("unable to decode bytes#e937bb82: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for Bytes.
var (
	_ bin.Encoder = &Bytes{}
	_ bin.Decoder = &Bytes{}
)
