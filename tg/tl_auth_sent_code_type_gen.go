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

// AuthSentCodeTypeApp represents TL type `auth.sentCodeTypeApp#3dbb5986`.
// The code was sent through the telegram app
//
// See https://core.telegram.org/constructor/auth.sentCodeTypeApp for reference.
type AuthSentCodeTypeApp struct {
	// Length of the code in bytes
	Length int `schemaname:"length"`
}

// AuthSentCodeTypeAppTypeID is TL type id of AuthSentCodeTypeApp.
const AuthSentCodeTypeAppTypeID = 0x3dbb5986

func (s *AuthSentCodeTypeApp) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Length == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *AuthSentCodeTypeApp) String() string {
	if s == nil {
		return "AuthSentCodeTypeApp(nil)"
	}
	type Alias AuthSentCodeTypeApp
	return fmt.Sprintf("AuthSentCodeTypeApp%+v", Alias(*s))
}

// FillFrom fills AuthSentCodeTypeApp from given interface.
func (s *AuthSentCodeTypeApp) FillFrom(from interface {
	GetLength() (value int)
}) {
	s.Length = from.GetLength()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *AuthSentCodeTypeApp) TypeID() uint32 {
	return AuthSentCodeTypeAppTypeID
}

// SchemaName returns MTProto type name.
func (s *AuthSentCodeTypeApp) SchemaName() string {
	return "auth.sentCodeTypeApp"
}

// Encode implements bin.Encoder.
func (s *AuthSentCodeTypeApp) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode auth.sentCodeTypeApp#3dbb5986 as nil")
	}
	b.PutID(AuthSentCodeTypeAppTypeID)
	b.PutInt(s.Length)
	return nil
}

// GetLength returns value of Length field.
func (s *AuthSentCodeTypeApp) GetLength() (value int) {
	return s.Length
}

// Decode implements bin.Decoder.
func (s *AuthSentCodeTypeApp) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode auth.sentCodeTypeApp#3dbb5986 to nil")
	}
	if err := b.ConsumeID(AuthSentCodeTypeAppTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.sentCodeTypeApp#3dbb5986: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode auth.sentCodeTypeApp#3dbb5986: field length: %w", err)
		}
		s.Length = value
	}
	return nil
}

// construct implements constructor of AuthSentCodeTypeClass.
func (s AuthSentCodeTypeApp) construct() AuthSentCodeTypeClass { return &s }

// Ensuring interfaces in compile-time for AuthSentCodeTypeApp.
var (
	_ bin.Encoder = &AuthSentCodeTypeApp{}
	_ bin.Decoder = &AuthSentCodeTypeApp{}

	_ AuthSentCodeTypeClass = &AuthSentCodeTypeApp{}
)

// AuthSentCodeTypeSms represents TL type `auth.sentCodeTypeSms#c000bba2`.
// The code was sent via SMS
//
// See https://core.telegram.org/constructor/auth.sentCodeTypeSms for reference.
type AuthSentCodeTypeSms struct {
	// Length of the code in bytes
	Length int `schemaname:"length"`
}

// AuthSentCodeTypeSmsTypeID is TL type id of AuthSentCodeTypeSms.
const AuthSentCodeTypeSmsTypeID = 0xc000bba2

func (s *AuthSentCodeTypeSms) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Length == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *AuthSentCodeTypeSms) String() string {
	if s == nil {
		return "AuthSentCodeTypeSms(nil)"
	}
	type Alias AuthSentCodeTypeSms
	return fmt.Sprintf("AuthSentCodeTypeSms%+v", Alias(*s))
}

// FillFrom fills AuthSentCodeTypeSms from given interface.
func (s *AuthSentCodeTypeSms) FillFrom(from interface {
	GetLength() (value int)
}) {
	s.Length = from.GetLength()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *AuthSentCodeTypeSms) TypeID() uint32 {
	return AuthSentCodeTypeSmsTypeID
}

// SchemaName returns MTProto type name.
func (s *AuthSentCodeTypeSms) SchemaName() string {
	return "auth.sentCodeTypeSms"
}

// Encode implements bin.Encoder.
func (s *AuthSentCodeTypeSms) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode auth.sentCodeTypeSms#c000bba2 as nil")
	}
	b.PutID(AuthSentCodeTypeSmsTypeID)
	b.PutInt(s.Length)
	return nil
}

// GetLength returns value of Length field.
func (s *AuthSentCodeTypeSms) GetLength() (value int) {
	return s.Length
}

// Decode implements bin.Decoder.
func (s *AuthSentCodeTypeSms) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode auth.sentCodeTypeSms#c000bba2 to nil")
	}
	if err := b.ConsumeID(AuthSentCodeTypeSmsTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.sentCodeTypeSms#c000bba2: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode auth.sentCodeTypeSms#c000bba2: field length: %w", err)
		}
		s.Length = value
	}
	return nil
}

// construct implements constructor of AuthSentCodeTypeClass.
func (s AuthSentCodeTypeSms) construct() AuthSentCodeTypeClass { return &s }

// Ensuring interfaces in compile-time for AuthSentCodeTypeSms.
var (
	_ bin.Encoder = &AuthSentCodeTypeSms{}
	_ bin.Decoder = &AuthSentCodeTypeSms{}

	_ AuthSentCodeTypeClass = &AuthSentCodeTypeSms{}
)

// AuthSentCodeTypeCall represents TL type `auth.sentCodeTypeCall#5353e5a7`.
// The code will be sent via a phone call: a synthesized voice will tell the user which verification code to input.
//
// See https://core.telegram.org/constructor/auth.sentCodeTypeCall for reference.
type AuthSentCodeTypeCall struct {
	// Length of the verification code
	Length int `schemaname:"length"`
}

// AuthSentCodeTypeCallTypeID is TL type id of AuthSentCodeTypeCall.
const AuthSentCodeTypeCallTypeID = 0x5353e5a7

func (s *AuthSentCodeTypeCall) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Length == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *AuthSentCodeTypeCall) String() string {
	if s == nil {
		return "AuthSentCodeTypeCall(nil)"
	}
	type Alias AuthSentCodeTypeCall
	return fmt.Sprintf("AuthSentCodeTypeCall%+v", Alias(*s))
}

// FillFrom fills AuthSentCodeTypeCall from given interface.
func (s *AuthSentCodeTypeCall) FillFrom(from interface {
	GetLength() (value int)
}) {
	s.Length = from.GetLength()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *AuthSentCodeTypeCall) TypeID() uint32 {
	return AuthSentCodeTypeCallTypeID
}

// SchemaName returns MTProto type name.
func (s *AuthSentCodeTypeCall) SchemaName() string {
	return "auth.sentCodeTypeCall"
}

// Encode implements bin.Encoder.
func (s *AuthSentCodeTypeCall) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode auth.sentCodeTypeCall#5353e5a7 as nil")
	}
	b.PutID(AuthSentCodeTypeCallTypeID)
	b.PutInt(s.Length)
	return nil
}

// GetLength returns value of Length field.
func (s *AuthSentCodeTypeCall) GetLength() (value int) {
	return s.Length
}

// Decode implements bin.Decoder.
func (s *AuthSentCodeTypeCall) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode auth.sentCodeTypeCall#5353e5a7 to nil")
	}
	if err := b.ConsumeID(AuthSentCodeTypeCallTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.sentCodeTypeCall#5353e5a7: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode auth.sentCodeTypeCall#5353e5a7: field length: %w", err)
		}
		s.Length = value
	}
	return nil
}

// construct implements constructor of AuthSentCodeTypeClass.
func (s AuthSentCodeTypeCall) construct() AuthSentCodeTypeClass { return &s }

// Ensuring interfaces in compile-time for AuthSentCodeTypeCall.
var (
	_ bin.Encoder = &AuthSentCodeTypeCall{}
	_ bin.Decoder = &AuthSentCodeTypeCall{}

	_ AuthSentCodeTypeClass = &AuthSentCodeTypeCall{}
)

// AuthSentCodeTypeFlashCall represents TL type `auth.sentCodeTypeFlashCall#ab03c6d9`.
// The code will be sent via a flash phone call, that will be closed immediately. The phone code will then be the phone number itself, just make sure that the phone number matches the specified pattern.
//
// See https://core.telegram.org/constructor/auth.sentCodeTypeFlashCall for reference.
type AuthSentCodeTypeFlashCall struct {
	// pattern¹ to match
	//
	// Links:
	//  1) https://core.telegram.org/api/pattern
	Pattern string `schemaname:"pattern"`
}

// AuthSentCodeTypeFlashCallTypeID is TL type id of AuthSentCodeTypeFlashCall.
const AuthSentCodeTypeFlashCallTypeID = 0xab03c6d9

func (s *AuthSentCodeTypeFlashCall) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Pattern == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *AuthSentCodeTypeFlashCall) String() string {
	if s == nil {
		return "AuthSentCodeTypeFlashCall(nil)"
	}
	type Alias AuthSentCodeTypeFlashCall
	return fmt.Sprintf("AuthSentCodeTypeFlashCall%+v", Alias(*s))
}

// FillFrom fills AuthSentCodeTypeFlashCall from given interface.
func (s *AuthSentCodeTypeFlashCall) FillFrom(from interface {
	GetPattern() (value string)
}) {
	s.Pattern = from.GetPattern()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *AuthSentCodeTypeFlashCall) TypeID() uint32 {
	return AuthSentCodeTypeFlashCallTypeID
}

// SchemaName returns MTProto type name.
func (s *AuthSentCodeTypeFlashCall) SchemaName() string {
	return "auth.sentCodeTypeFlashCall"
}

// Encode implements bin.Encoder.
func (s *AuthSentCodeTypeFlashCall) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode auth.sentCodeTypeFlashCall#ab03c6d9 as nil")
	}
	b.PutID(AuthSentCodeTypeFlashCallTypeID)
	b.PutString(s.Pattern)
	return nil
}

// GetPattern returns value of Pattern field.
func (s *AuthSentCodeTypeFlashCall) GetPattern() (value string) {
	return s.Pattern
}

// Decode implements bin.Decoder.
func (s *AuthSentCodeTypeFlashCall) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode auth.sentCodeTypeFlashCall#ab03c6d9 to nil")
	}
	if err := b.ConsumeID(AuthSentCodeTypeFlashCallTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.sentCodeTypeFlashCall#ab03c6d9: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.sentCodeTypeFlashCall#ab03c6d9: field pattern: %w", err)
		}
		s.Pattern = value
	}
	return nil
}

// construct implements constructor of AuthSentCodeTypeClass.
func (s AuthSentCodeTypeFlashCall) construct() AuthSentCodeTypeClass { return &s }

// Ensuring interfaces in compile-time for AuthSentCodeTypeFlashCall.
var (
	_ bin.Encoder = &AuthSentCodeTypeFlashCall{}
	_ bin.Decoder = &AuthSentCodeTypeFlashCall{}

	_ AuthSentCodeTypeClass = &AuthSentCodeTypeFlashCall{}
)

// AuthSentCodeTypeClass represents auth.SentCodeType generic type.
//
// See https://core.telegram.org/type/auth.SentCodeType for reference.
//
// Example:
//  g, err := tg.DecodeAuthSentCodeType(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.AuthSentCodeTypeApp: // auth.sentCodeTypeApp#3dbb5986
//  case *tg.AuthSentCodeTypeSms: // auth.sentCodeTypeSms#c000bba2
//  case *tg.AuthSentCodeTypeCall: // auth.sentCodeTypeCall#5353e5a7
//  case *tg.AuthSentCodeTypeFlashCall: // auth.sentCodeTypeFlashCall#ab03c6d9
//  default: panic(v)
//  }
type AuthSentCodeTypeClass interface {
	bin.Encoder
	bin.Decoder
	construct() AuthSentCodeTypeClass

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// SchemaName returns MTProto type name.
	SchemaName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeAuthSentCodeType implements binary de-serialization for AuthSentCodeTypeClass.
func DecodeAuthSentCodeType(buf *bin.Buffer) (AuthSentCodeTypeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case AuthSentCodeTypeAppTypeID:
		// Decoding auth.sentCodeTypeApp#3dbb5986.
		v := AuthSentCodeTypeApp{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthSentCodeTypeClass: %w", err)
		}
		return &v, nil
	case AuthSentCodeTypeSmsTypeID:
		// Decoding auth.sentCodeTypeSms#c000bba2.
		v := AuthSentCodeTypeSms{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthSentCodeTypeClass: %w", err)
		}
		return &v, nil
	case AuthSentCodeTypeCallTypeID:
		// Decoding auth.sentCodeTypeCall#5353e5a7.
		v := AuthSentCodeTypeCall{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthSentCodeTypeClass: %w", err)
		}
		return &v, nil
	case AuthSentCodeTypeFlashCallTypeID:
		// Decoding auth.sentCodeTypeFlashCall#ab03c6d9.
		v := AuthSentCodeTypeFlashCall{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthSentCodeTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode AuthSentCodeTypeClass: %w", bin.NewUnexpectedID(id))
	}
}

// AuthSentCodeType boxes the AuthSentCodeTypeClass providing a helper.
type AuthSentCodeTypeBox struct {
	SentCodeType AuthSentCodeTypeClass
}

// Decode implements bin.Decoder for AuthSentCodeTypeBox.
func (b *AuthSentCodeTypeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode AuthSentCodeTypeBox to nil")
	}
	v, err := DecodeAuthSentCodeType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.SentCodeType = v
	return nil
}

// Encode implements bin.Encode for AuthSentCodeTypeBox.
func (b *AuthSentCodeTypeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.SentCodeType == nil {
		return fmt.Errorf("unable to encode AuthSentCodeTypeClass as nil")
	}
	return b.SentCodeType.Encode(buf)
}

// AuthSentCodeTypeClassArray is adapter for slice of AuthSentCodeTypeClass.
type AuthSentCodeTypeClassArray []AuthSentCodeTypeClass

// Sort sorts slice of AuthSentCodeTypeClass.
func (s AuthSentCodeTypeClassArray) Sort(less func(a, b AuthSentCodeTypeClass) bool) AuthSentCodeTypeClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of AuthSentCodeTypeClass.
func (s AuthSentCodeTypeClassArray) SortStable(less func(a, b AuthSentCodeTypeClass) bool) AuthSentCodeTypeClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of AuthSentCodeTypeClass.
func (s AuthSentCodeTypeClassArray) Retain(keep func(x AuthSentCodeTypeClass) bool) AuthSentCodeTypeClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s AuthSentCodeTypeClassArray) First() (v AuthSentCodeTypeClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s AuthSentCodeTypeClassArray) Last() (v AuthSentCodeTypeClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *AuthSentCodeTypeClassArray) PopFirst() (v AuthSentCodeTypeClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero AuthSentCodeTypeClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *AuthSentCodeTypeClassArray) Pop() (v AuthSentCodeTypeClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsAuthSentCodeTypeApp returns copy with only AuthSentCodeTypeApp constructors.
func (s AuthSentCodeTypeClassArray) AsAuthSentCodeTypeApp() (to AuthSentCodeTypeAppArray) {
	for _, elem := range s {
		value, ok := elem.(*AuthSentCodeTypeApp)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsAuthSentCodeTypeSms returns copy with only AuthSentCodeTypeSms constructors.
func (s AuthSentCodeTypeClassArray) AsAuthSentCodeTypeSms() (to AuthSentCodeTypeSmsArray) {
	for _, elem := range s {
		value, ok := elem.(*AuthSentCodeTypeSms)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsAuthSentCodeTypeCall returns copy with only AuthSentCodeTypeCall constructors.
func (s AuthSentCodeTypeClassArray) AsAuthSentCodeTypeCall() (to AuthSentCodeTypeCallArray) {
	for _, elem := range s {
		value, ok := elem.(*AuthSentCodeTypeCall)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsAuthSentCodeTypeFlashCall returns copy with only AuthSentCodeTypeFlashCall constructors.
func (s AuthSentCodeTypeClassArray) AsAuthSentCodeTypeFlashCall() (to AuthSentCodeTypeFlashCallArray) {
	for _, elem := range s {
		value, ok := elem.(*AuthSentCodeTypeFlashCall)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AuthSentCodeTypeAppArray is adapter for slice of AuthSentCodeTypeApp.
type AuthSentCodeTypeAppArray []AuthSentCodeTypeApp

// Sort sorts slice of AuthSentCodeTypeApp.
func (s AuthSentCodeTypeAppArray) Sort(less func(a, b AuthSentCodeTypeApp) bool) AuthSentCodeTypeAppArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of AuthSentCodeTypeApp.
func (s AuthSentCodeTypeAppArray) SortStable(less func(a, b AuthSentCodeTypeApp) bool) AuthSentCodeTypeAppArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of AuthSentCodeTypeApp.
func (s AuthSentCodeTypeAppArray) Retain(keep func(x AuthSentCodeTypeApp) bool) AuthSentCodeTypeAppArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s AuthSentCodeTypeAppArray) First() (v AuthSentCodeTypeApp, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s AuthSentCodeTypeAppArray) Last() (v AuthSentCodeTypeApp, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *AuthSentCodeTypeAppArray) PopFirst() (v AuthSentCodeTypeApp, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero AuthSentCodeTypeApp
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *AuthSentCodeTypeAppArray) Pop() (v AuthSentCodeTypeApp, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AuthSentCodeTypeSmsArray is adapter for slice of AuthSentCodeTypeSms.
type AuthSentCodeTypeSmsArray []AuthSentCodeTypeSms

// Sort sorts slice of AuthSentCodeTypeSms.
func (s AuthSentCodeTypeSmsArray) Sort(less func(a, b AuthSentCodeTypeSms) bool) AuthSentCodeTypeSmsArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of AuthSentCodeTypeSms.
func (s AuthSentCodeTypeSmsArray) SortStable(less func(a, b AuthSentCodeTypeSms) bool) AuthSentCodeTypeSmsArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of AuthSentCodeTypeSms.
func (s AuthSentCodeTypeSmsArray) Retain(keep func(x AuthSentCodeTypeSms) bool) AuthSentCodeTypeSmsArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s AuthSentCodeTypeSmsArray) First() (v AuthSentCodeTypeSms, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s AuthSentCodeTypeSmsArray) Last() (v AuthSentCodeTypeSms, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *AuthSentCodeTypeSmsArray) PopFirst() (v AuthSentCodeTypeSms, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero AuthSentCodeTypeSms
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *AuthSentCodeTypeSmsArray) Pop() (v AuthSentCodeTypeSms, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AuthSentCodeTypeCallArray is adapter for slice of AuthSentCodeTypeCall.
type AuthSentCodeTypeCallArray []AuthSentCodeTypeCall

// Sort sorts slice of AuthSentCodeTypeCall.
func (s AuthSentCodeTypeCallArray) Sort(less func(a, b AuthSentCodeTypeCall) bool) AuthSentCodeTypeCallArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of AuthSentCodeTypeCall.
func (s AuthSentCodeTypeCallArray) SortStable(less func(a, b AuthSentCodeTypeCall) bool) AuthSentCodeTypeCallArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of AuthSentCodeTypeCall.
func (s AuthSentCodeTypeCallArray) Retain(keep func(x AuthSentCodeTypeCall) bool) AuthSentCodeTypeCallArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s AuthSentCodeTypeCallArray) First() (v AuthSentCodeTypeCall, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s AuthSentCodeTypeCallArray) Last() (v AuthSentCodeTypeCall, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *AuthSentCodeTypeCallArray) PopFirst() (v AuthSentCodeTypeCall, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero AuthSentCodeTypeCall
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *AuthSentCodeTypeCallArray) Pop() (v AuthSentCodeTypeCall, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AuthSentCodeTypeFlashCallArray is adapter for slice of AuthSentCodeTypeFlashCall.
type AuthSentCodeTypeFlashCallArray []AuthSentCodeTypeFlashCall

// Sort sorts slice of AuthSentCodeTypeFlashCall.
func (s AuthSentCodeTypeFlashCallArray) Sort(less func(a, b AuthSentCodeTypeFlashCall) bool) AuthSentCodeTypeFlashCallArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of AuthSentCodeTypeFlashCall.
func (s AuthSentCodeTypeFlashCallArray) SortStable(less func(a, b AuthSentCodeTypeFlashCall) bool) AuthSentCodeTypeFlashCallArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of AuthSentCodeTypeFlashCall.
func (s AuthSentCodeTypeFlashCallArray) Retain(keep func(x AuthSentCodeTypeFlashCall) bool) AuthSentCodeTypeFlashCallArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s AuthSentCodeTypeFlashCallArray) First() (v AuthSentCodeTypeFlashCall, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s AuthSentCodeTypeFlashCallArray) Last() (v AuthSentCodeTypeFlashCall, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *AuthSentCodeTypeFlashCallArray) PopFirst() (v AuthSentCodeTypeFlashCall, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero AuthSentCodeTypeFlashCall
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *AuthSentCodeTypeFlashCallArray) Pop() (v AuthSentCodeTypeFlashCall, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
