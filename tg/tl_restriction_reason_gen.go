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

// RestrictionReason represents TL type `restrictionReason#d072acb4`.
// Restriction reason.
// Contains the reason why access to a certain object must be restricted. Clients are supposed to deny access to the channel if the platform field is equal to all or to the current platform (ios, android, wp, etc.). Platforms can be concatenated (ios-android, ios-wp), unknown platforms are to be ignored. The text is the error message that should be shown to the user.
//
// See https://core.telegram.org/constructor/restrictionReason for reference.
type RestrictionReason struct {
	// Platform identifier (ios, android, wp, all, etc.), can be concatenated with a dash as separator (android-ios, ios-wp, etc)
	Platform string `schemaname:"platform"`
	// Restriction reason (porno, terms, etc.)
	Reason string `schemaname:"reason"`
	// Error message to be shown to the user
	Text string `schemaname:"text"`
}

// RestrictionReasonTypeID is TL type id of RestrictionReason.
const RestrictionReasonTypeID = 0xd072acb4

func (r *RestrictionReason) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Platform == "") {
		return false
	}
	if !(r.Reason == "") {
		return false
	}
	if !(r.Text == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RestrictionReason) String() string {
	if r == nil {
		return "RestrictionReason(nil)"
	}
	type Alias RestrictionReason
	return fmt.Sprintf("RestrictionReason%+v", Alias(*r))
}

// FillFrom fills RestrictionReason from given interface.
func (r *RestrictionReason) FillFrom(from interface {
	GetPlatform() (value string)
	GetReason() (value string)
	GetText() (value string)
}) {
	r.Platform = from.GetPlatform()
	r.Reason = from.GetReason()
	r.Text = from.GetText()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (r *RestrictionReason) TypeID() uint32 {
	return RestrictionReasonTypeID
}

// SchemaName returns MTProto type name.
func (r *RestrictionReason) SchemaName() string {
	return "restrictionReason"
}

// Encode implements bin.Encoder.
func (r *RestrictionReason) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode restrictionReason#d072acb4 as nil")
	}
	b.PutID(RestrictionReasonTypeID)
	b.PutString(r.Platform)
	b.PutString(r.Reason)
	b.PutString(r.Text)
	return nil
}

// GetPlatform returns value of Platform field.
func (r *RestrictionReason) GetPlatform() (value string) {
	return r.Platform
}

// GetReason returns value of Reason field.
func (r *RestrictionReason) GetReason() (value string) {
	return r.Reason
}

// GetText returns value of Text field.
func (r *RestrictionReason) GetText() (value string) {
	return r.Text
}

// Decode implements bin.Decoder.
func (r *RestrictionReason) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode restrictionReason#d072acb4 to nil")
	}
	if err := b.ConsumeID(RestrictionReasonTypeID); err != nil {
		return fmt.Errorf("unable to decode restrictionReason#d072acb4: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode restrictionReason#d072acb4: field platform: %w", err)
		}
		r.Platform = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode restrictionReason#d072acb4: field reason: %w", err)
		}
		r.Reason = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode restrictionReason#d072acb4: field text: %w", err)
		}
		r.Text = value
	}
	return nil
}

// Ensuring interfaces in compile-time for RestrictionReason.
var (
	_ bin.Encoder = &RestrictionReason{}
	_ bin.Decoder = &RestrictionReason{}
)
