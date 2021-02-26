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

// PaymentsSavedInfo represents TL type `payments.savedInfo#fb8fe43c`.
// Saved server-side order information
//
// See https://core.telegram.org/constructor/payments.savedInfo for reference.
type PaymentsSavedInfo struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `schemaname:"flags"`
	// Whether the user has some saved payment credentials
	HasSavedCredentials bool `schemaname:"has_saved_credentials"`
	// Saved server-side order information
	//
	// Use SetSavedInfo and GetSavedInfo helpers.
	SavedInfo PaymentRequestedInfo `schemaname:"saved_info"`
}

// PaymentsSavedInfoTypeID is TL type id of PaymentsSavedInfo.
const PaymentsSavedInfoTypeID = 0xfb8fe43c

func (s *PaymentsSavedInfo) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.HasSavedCredentials == false) {
		return false
	}
	if !(s.SavedInfo.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *PaymentsSavedInfo) String() string {
	if s == nil {
		return "PaymentsSavedInfo(nil)"
	}
	type Alias PaymentsSavedInfo
	return fmt.Sprintf("PaymentsSavedInfo%+v", Alias(*s))
}

// FillFrom fills PaymentsSavedInfo from given interface.
func (s *PaymentsSavedInfo) FillFrom(from interface {
	GetHasSavedCredentials() (value bool)
	GetSavedInfo() (value PaymentRequestedInfo, ok bool)
}) {
	s.HasSavedCredentials = from.GetHasSavedCredentials()
	if val, ok := from.GetSavedInfo(); ok {
		s.SavedInfo = val
	}

}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *PaymentsSavedInfo) TypeID() uint32 {
	return PaymentsSavedInfoTypeID
}

// SchemaName returns MTProto type name.
func (s *PaymentsSavedInfo) SchemaName() string {
	return "payments.savedInfo"
}

// Encode implements bin.Encoder.
func (s *PaymentsSavedInfo) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode payments.savedInfo#fb8fe43c as nil")
	}
	b.PutID(PaymentsSavedInfoTypeID)
	if !(s.HasSavedCredentials == false) {
		s.Flags.Set(1)
	}
	if !(s.SavedInfo.Zero()) {
		s.Flags.Set(0)
	}
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.savedInfo#fb8fe43c: field flags: %w", err)
	}
	if s.Flags.Has(0) {
		if err := s.SavedInfo.Encode(b); err != nil {
			return fmt.Errorf("unable to encode payments.savedInfo#fb8fe43c: field saved_info: %w", err)
		}
	}
	return nil
}

// SetHasSavedCredentials sets value of HasSavedCredentials conditional field.
func (s *PaymentsSavedInfo) SetHasSavedCredentials(value bool) {
	if value {
		s.Flags.Set(1)
		s.HasSavedCredentials = true
	} else {
		s.Flags.Unset(1)
		s.HasSavedCredentials = false
	}
}

// GetHasSavedCredentials returns value of HasSavedCredentials conditional field.
func (s *PaymentsSavedInfo) GetHasSavedCredentials() (value bool) {
	return s.Flags.Has(1)
}

// SetSavedInfo sets value of SavedInfo conditional field.
func (s *PaymentsSavedInfo) SetSavedInfo(value PaymentRequestedInfo) {
	s.Flags.Set(0)
	s.SavedInfo = value
}

// GetSavedInfo returns value of SavedInfo conditional field and
// boolean which is true if field was set.
func (s *PaymentsSavedInfo) GetSavedInfo() (value PaymentRequestedInfo, ok bool) {
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.SavedInfo, true
}

// Decode implements bin.Decoder.
func (s *PaymentsSavedInfo) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode payments.savedInfo#fb8fe43c to nil")
	}
	if err := b.ConsumeID(PaymentsSavedInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.savedInfo#fb8fe43c: %w", err)
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.savedInfo#fb8fe43c: field flags: %w", err)
		}
	}
	s.HasSavedCredentials = s.Flags.Has(1)
	if s.Flags.Has(0) {
		if err := s.SavedInfo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.savedInfo#fb8fe43c: field saved_info: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for PaymentsSavedInfo.
var (
	_ bin.Encoder = &PaymentsSavedInfo{}
	_ bin.Decoder = &PaymentsSavedInfo{}
)
