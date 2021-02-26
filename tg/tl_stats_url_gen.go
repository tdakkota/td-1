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

// StatsURL represents TL type `statsURL#47a971e0`.
// URL with chat statistics
//
// See https://core.telegram.org/constructor/statsURL for reference.
type StatsURL struct {
	// Chat statistics
	URL string `schemaname:"url"`
}

// StatsURLTypeID is TL type id of StatsURL.
const StatsURLTypeID = 0x47a971e0

func (s *StatsURL) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.URL == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StatsURL) String() string {
	if s == nil {
		return "StatsURL(nil)"
	}
	type Alias StatsURL
	return fmt.Sprintf("StatsURL%+v", Alias(*s))
}

// FillFrom fills StatsURL from given interface.
func (s *StatsURL) FillFrom(from interface {
	GetURL() (value string)
}) {
	s.URL = from.GetURL()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *StatsURL) TypeID() uint32 {
	return StatsURLTypeID
}

// SchemaName returns MTProto type name.
func (s *StatsURL) SchemaName() string {
	return "statsURL"
}

// Encode implements bin.Encoder.
func (s *StatsURL) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode statsURL#47a971e0 as nil")
	}
	b.PutID(StatsURLTypeID)
	b.PutString(s.URL)
	return nil
}

// GetURL returns value of URL field.
func (s *StatsURL) GetURL() (value string) {
	return s.URL
}

// Decode implements bin.Decoder.
func (s *StatsURL) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode statsURL#47a971e0 to nil")
	}
	if err := b.ConsumeID(StatsURLTypeID); err != nil {
		return fmt.Errorf("unable to decode statsURL#47a971e0: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode statsURL#47a971e0: field url: %w", err)
		}
		s.URL = value
	}
	return nil
}

// Ensuring interfaces in compile-time for StatsURL.
var (
	_ bin.Encoder = &StatsURL{}
	_ bin.Decoder = &StatsURL{}
)
