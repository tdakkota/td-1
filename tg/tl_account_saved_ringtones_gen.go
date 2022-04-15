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

// AccountSavedRingtonesNotModified represents TL type `account.savedRingtonesNotModified#fbf6e8b1`.
//
// See https://core.telegram.org/constructor/account.savedRingtonesNotModified for reference.
type AccountSavedRingtonesNotModified struct {
}

// AccountSavedRingtonesNotModifiedTypeID is TL type id of AccountSavedRingtonesNotModified.
const AccountSavedRingtonesNotModifiedTypeID = 0xfbf6e8b1

// construct implements constructor of AccountSavedRingtonesClass.
func (s AccountSavedRingtonesNotModified) construct() AccountSavedRingtonesClass { return &s }

// Ensuring interfaces in compile-time for AccountSavedRingtonesNotModified.
var (
	_ bin.Encoder     = &AccountSavedRingtonesNotModified{}
	_ bin.Decoder     = &AccountSavedRingtonesNotModified{}
	_ bin.BareEncoder = &AccountSavedRingtonesNotModified{}
	_ bin.BareDecoder = &AccountSavedRingtonesNotModified{}

	_ AccountSavedRingtonesClass = &AccountSavedRingtonesNotModified{}
)

func (s *AccountSavedRingtonesNotModified) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *AccountSavedRingtonesNotModified) String() string {
	if s == nil {
		return "AccountSavedRingtonesNotModified(nil)"
	}
	type Alias AccountSavedRingtonesNotModified
	return fmt.Sprintf("AccountSavedRingtonesNotModified%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountSavedRingtonesNotModified) TypeID() uint32 {
	return AccountSavedRingtonesNotModifiedTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountSavedRingtonesNotModified) TypeName() string {
	return "account.savedRingtonesNotModified"
}

// TypeInfo returns info about TL type.
func (s *AccountSavedRingtonesNotModified) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.savedRingtonesNotModified",
		ID:   AccountSavedRingtonesNotModifiedTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *AccountSavedRingtonesNotModified) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.savedRingtonesNotModified#fbf6e8b1 as nil")
	}
	b.PutID(AccountSavedRingtonesNotModifiedTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *AccountSavedRingtonesNotModified) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.savedRingtonesNotModified#fbf6e8b1 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *AccountSavedRingtonesNotModified) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.savedRingtonesNotModified#fbf6e8b1 to nil")
	}
	if err := b.ConsumeID(AccountSavedRingtonesNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode account.savedRingtonesNotModified#fbf6e8b1: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *AccountSavedRingtonesNotModified) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.savedRingtonesNotModified#fbf6e8b1 to nil")
	}
	return nil
}

// AccountSavedRingtones represents TL type `account.savedRingtones#c1e92cc5`.
//
// See https://core.telegram.org/constructor/account.savedRingtones for reference.
type AccountSavedRingtones struct {
	// Hash field of AccountSavedRingtones.
	Hash int64
	// Ringtones field of AccountSavedRingtones.
	Ringtones []DocumentClass
}

// AccountSavedRingtonesTypeID is TL type id of AccountSavedRingtones.
const AccountSavedRingtonesTypeID = 0xc1e92cc5

// construct implements constructor of AccountSavedRingtonesClass.
func (s AccountSavedRingtones) construct() AccountSavedRingtonesClass { return &s }

// Ensuring interfaces in compile-time for AccountSavedRingtones.
var (
	_ bin.Encoder     = &AccountSavedRingtones{}
	_ bin.Decoder     = &AccountSavedRingtones{}
	_ bin.BareEncoder = &AccountSavedRingtones{}
	_ bin.BareDecoder = &AccountSavedRingtones{}

	_ AccountSavedRingtonesClass = &AccountSavedRingtones{}
)

func (s *AccountSavedRingtones) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Hash == 0) {
		return false
	}
	if !(s.Ringtones == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *AccountSavedRingtones) String() string {
	if s == nil {
		return "AccountSavedRingtones(nil)"
	}
	type Alias AccountSavedRingtones
	return fmt.Sprintf("AccountSavedRingtones%+v", Alias(*s))
}

// FillFrom fills AccountSavedRingtones from given interface.
func (s *AccountSavedRingtones) FillFrom(from interface {
	GetHash() (value int64)
	GetRingtones() (value []DocumentClass)
}) {
	s.Hash = from.GetHash()
	s.Ringtones = from.GetRingtones()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountSavedRingtones) TypeID() uint32 {
	return AccountSavedRingtonesTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountSavedRingtones) TypeName() string {
	return "account.savedRingtones"
}

// TypeInfo returns info about TL type.
func (s *AccountSavedRingtones) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.savedRingtones",
		ID:   AccountSavedRingtonesTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
		{
			Name:       "Ringtones",
			SchemaName: "ringtones",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *AccountSavedRingtones) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.savedRingtones#c1e92cc5 as nil")
	}
	b.PutID(AccountSavedRingtonesTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *AccountSavedRingtones) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.savedRingtones#c1e92cc5 as nil")
	}
	b.PutLong(s.Hash)
	b.PutVectorHeader(len(s.Ringtones))
	for idx, v := range s.Ringtones {
		if v == nil {
			return fmt.Errorf("unable to encode account.savedRingtones#c1e92cc5: field ringtones element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.savedRingtones#c1e92cc5: field ringtones element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *AccountSavedRingtones) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.savedRingtones#c1e92cc5 to nil")
	}
	if err := b.ConsumeID(AccountSavedRingtonesTypeID); err != nil {
		return fmt.Errorf("unable to decode account.savedRingtones#c1e92cc5: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *AccountSavedRingtones) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.savedRingtones#c1e92cc5 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode account.savedRingtones#c1e92cc5: field hash: %w", err)
		}
		s.Hash = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.savedRingtones#c1e92cc5: field ringtones: %w", err)
		}

		if headerLen > 0 {
			s.Ringtones = make([]DocumentClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDocument(b)
			if err != nil {
				return fmt.Errorf("unable to decode account.savedRingtones#c1e92cc5: field ringtones: %w", err)
			}
			s.Ringtones = append(s.Ringtones, value)
		}
	}
	return nil
}

// GetHash returns value of Hash field.
func (s *AccountSavedRingtones) GetHash() (value int64) {
	if s == nil {
		return
	}
	return s.Hash
}

// GetRingtones returns value of Ringtones field.
func (s *AccountSavedRingtones) GetRingtones() (value []DocumentClass) {
	if s == nil {
		return
	}
	return s.Ringtones
}

// MapRingtones returns field Ringtones wrapped in DocumentClassArray helper.
func (s *AccountSavedRingtones) MapRingtones() (value DocumentClassArray) {
	return DocumentClassArray(s.Ringtones)
}

// AccountSavedRingtonesClassName is schema name of AccountSavedRingtonesClass.
const AccountSavedRingtonesClassName = "account.SavedRingtones"

// AccountSavedRingtonesClass represents account.SavedRingtones generic type.
//
// See https://core.telegram.org/type/account.SavedRingtones for reference.
//
// Example:
//  g, err := tg.DecodeAccountSavedRingtones(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.AccountSavedRingtonesNotModified: // account.savedRingtonesNotModified#fbf6e8b1
//  case *tg.AccountSavedRingtones: // account.savedRingtones#c1e92cc5
//  default: panic(v)
//  }
type AccountSavedRingtonesClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() AccountSavedRingtonesClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// AsModified tries to map AccountSavedRingtonesClass to AccountSavedRingtones.
	AsModified() (*AccountSavedRingtones, bool)
}

// AsModified tries to map AccountSavedRingtonesNotModified to AccountSavedRingtones.
func (s *AccountSavedRingtonesNotModified) AsModified() (*AccountSavedRingtones, bool) {
	return nil, false
}

// AsModified tries to map AccountSavedRingtones to AccountSavedRingtones.
func (s *AccountSavedRingtones) AsModified() (*AccountSavedRingtones, bool) {
	return s, true
}

// DecodeAccountSavedRingtones implements binary de-serialization for AccountSavedRingtonesClass.
func DecodeAccountSavedRingtones(buf *bin.Buffer) (AccountSavedRingtonesClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case AccountSavedRingtonesNotModifiedTypeID:
		// Decoding account.savedRingtonesNotModified#fbf6e8b1.
		v := AccountSavedRingtonesNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AccountSavedRingtonesClass: %w", err)
		}
		return &v, nil
	case AccountSavedRingtonesTypeID:
		// Decoding account.savedRingtones#c1e92cc5.
		v := AccountSavedRingtones{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AccountSavedRingtonesClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode AccountSavedRingtonesClass: %w", bin.NewUnexpectedID(id))
	}
}

// AccountSavedRingtones boxes the AccountSavedRingtonesClass providing a helper.
type AccountSavedRingtonesBox struct {
	SavedRingtones AccountSavedRingtonesClass
}

// Decode implements bin.Decoder for AccountSavedRingtonesBox.
func (b *AccountSavedRingtonesBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode AccountSavedRingtonesBox to nil")
	}
	v, err := DecodeAccountSavedRingtones(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.SavedRingtones = v
	return nil
}

// Encode implements bin.Encode for AccountSavedRingtonesBox.
func (b *AccountSavedRingtonesBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.SavedRingtones == nil {
		return fmt.Errorf("unable to encode AccountSavedRingtonesClass as nil")
	}
	return b.SavedRingtones.Encode(buf)
}
