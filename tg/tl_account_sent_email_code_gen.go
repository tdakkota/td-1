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

// AccountSentEmailCode represents TL type `account.sentEmailCode#811f854f`.
// The sent email code
//
// See https://core.telegram.org/constructor/account.sentEmailCode for reference.
type AccountSentEmailCode struct {
	// The email (to which the code was sent) must match this pattern¹
	//
	// Links:
	//  1) https://core.telegram.org/api/pattern
	EmailPattern string `schemaname:"email_pattern"`
	// The length of the verification code
	Length int `schemaname:"length"`
}

// AccountSentEmailCodeTypeID is TL type id of AccountSentEmailCode.
const AccountSentEmailCodeTypeID = 0x811f854f

func (s *AccountSentEmailCode) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.EmailPattern == "") {
		return false
	}
	if !(s.Length == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *AccountSentEmailCode) String() string {
	if s == nil {
		return "AccountSentEmailCode(nil)"
	}
	type Alias AccountSentEmailCode
	return fmt.Sprintf("AccountSentEmailCode%+v", Alias(*s))
}

// FillFrom fills AccountSentEmailCode from given interface.
func (s *AccountSentEmailCode) FillFrom(from interface {
	GetEmailPattern() (value string)
	GetLength() (value int)
}) {
	s.EmailPattern = from.GetEmailPattern()
	s.Length = from.GetLength()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *AccountSentEmailCode) TypeID() uint32 {
	return AccountSentEmailCodeTypeID
}

// SchemaName returns MTProto type name.
func (s *AccountSentEmailCode) SchemaName() string {
	return "account.sentEmailCode"
}

// Encode implements bin.Encoder.
func (s *AccountSentEmailCode) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.sentEmailCode#811f854f as nil")
	}
	b.PutID(AccountSentEmailCodeTypeID)
	b.PutString(s.EmailPattern)
	b.PutInt(s.Length)
	return nil
}

// GetEmailPattern returns value of EmailPattern field.
func (s *AccountSentEmailCode) GetEmailPattern() (value string) {
	return s.EmailPattern
}

// GetLength returns value of Length field.
func (s *AccountSentEmailCode) GetLength() (value int) {
	return s.Length
}

// Decode implements bin.Decoder.
func (s *AccountSentEmailCode) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.sentEmailCode#811f854f to nil")
	}
	if err := b.ConsumeID(AccountSentEmailCodeTypeID); err != nil {
		return fmt.Errorf("unable to decode account.sentEmailCode#811f854f: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.sentEmailCode#811f854f: field email_pattern: %w", err)
		}
		s.EmailPattern = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode account.sentEmailCode#811f854f: field length: %w", err)
		}
		s.Length = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountSentEmailCode.
var (
	_ bin.Encoder = &AccountSentEmailCode{}
	_ bin.Decoder = &AccountSentEmailCode{}
)
