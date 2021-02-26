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

// InputPhotoEmpty represents TL type `inputPhotoEmpty#1cd7bf0d`.
// Empty constructor.
//
// See https://core.telegram.org/constructor/inputPhotoEmpty for reference.
type InputPhotoEmpty struct {
}

// InputPhotoEmptyTypeID is TL type id of InputPhotoEmpty.
const InputPhotoEmptyTypeID = 0x1cd7bf0d

func (i *InputPhotoEmpty) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPhotoEmpty) String() string {
	if i == nil {
		return "InputPhotoEmpty(nil)"
	}
	type Alias InputPhotoEmpty
	return fmt.Sprintf("InputPhotoEmpty%+v", Alias(*i))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputPhotoEmpty) TypeID() uint32 {
	return InputPhotoEmptyTypeID
}

// SchemaName returns MTProto type name.
func (i *InputPhotoEmpty) SchemaName() string {
	return "inputPhotoEmpty"
}

// Encode implements bin.Encoder.
func (i *InputPhotoEmpty) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPhotoEmpty#1cd7bf0d as nil")
	}
	b.PutID(InputPhotoEmptyTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPhotoEmpty) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPhotoEmpty#1cd7bf0d to nil")
	}
	if err := b.ConsumeID(InputPhotoEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPhotoEmpty#1cd7bf0d: %w", err)
	}
	return nil
}

// construct implements constructor of InputPhotoClass.
func (i InputPhotoEmpty) construct() InputPhotoClass { return &i }

// Ensuring interfaces in compile-time for InputPhotoEmpty.
var (
	_ bin.Encoder = &InputPhotoEmpty{}
	_ bin.Decoder = &InputPhotoEmpty{}

	_ InputPhotoClass = &InputPhotoEmpty{}
)

// InputPhoto represents TL type `inputPhoto#3bb3b94a`.
// Defines a photo for further interaction.
//
// See https://core.telegram.org/constructor/inputPhoto for reference.
type InputPhoto struct {
	// Photo identifier
	ID int64 `schemaname:"id"`
	// access_hash value from the photo¹ constructor
	//
	// Links:
	//  1) https://core.telegram.org/constructor/photo
	AccessHash int64 `schemaname:"access_hash"`
	// File reference¹
	//
	// Links:
	//  1) https://core.telegram.org/api/file_reference
	FileReference []byte `schemaname:"file_reference"`
}

// InputPhotoTypeID is TL type id of InputPhoto.
const InputPhotoTypeID = 0x3bb3b94a

func (i *InputPhoto) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ID == 0) {
		return false
	}
	if !(i.AccessHash == 0) {
		return false
	}
	if !(i.FileReference == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPhoto) String() string {
	if i == nil {
		return "InputPhoto(nil)"
	}
	type Alias InputPhoto
	return fmt.Sprintf("InputPhoto%+v", Alias(*i))
}

// FillFrom fills InputPhoto from given interface.
func (i *InputPhoto) FillFrom(from interface {
	GetID() (value int64)
	GetAccessHash() (value int64)
	GetFileReference() (value []byte)
}) {
	i.ID = from.GetID()
	i.AccessHash = from.GetAccessHash()
	i.FileReference = from.GetFileReference()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputPhoto) TypeID() uint32 {
	return InputPhotoTypeID
}

// SchemaName returns MTProto type name.
func (i *InputPhoto) SchemaName() string {
	return "inputPhoto"
}

// Encode implements bin.Encoder.
func (i *InputPhoto) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPhoto#3bb3b94a as nil")
	}
	b.PutID(InputPhotoTypeID)
	b.PutLong(i.ID)
	b.PutLong(i.AccessHash)
	b.PutBytes(i.FileReference)
	return nil
}

// GetID returns value of ID field.
func (i *InputPhoto) GetID() (value int64) {
	return i.ID
}

// GetAccessHash returns value of AccessHash field.
func (i *InputPhoto) GetAccessHash() (value int64) {
	return i.AccessHash
}

// GetFileReference returns value of FileReference field.
func (i *InputPhoto) GetFileReference() (value []byte) {
	return i.FileReference
}

// Decode implements bin.Decoder.
func (i *InputPhoto) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPhoto#3bb3b94a to nil")
	}
	if err := b.ConsumeID(InputPhotoTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPhoto#3bb3b94a: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputPhoto#3bb3b94a: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputPhoto#3bb3b94a: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode inputPhoto#3bb3b94a: field file_reference: %w", err)
		}
		i.FileReference = value
	}
	return nil
}

// construct implements constructor of InputPhotoClass.
func (i InputPhoto) construct() InputPhotoClass { return &i }

// Ensuring interfaces in compile-time for InputPhoto.
var (
	_ bin.Encoder = &InputPhoto{}
	_ bin.Decoder = &InputPhoto{}

	_ InputPhotoClass = &InputPhoto{}
)

// InputPhotoClass represents InputPhoto generic type.
//
// See https://core.telegram.org/type/InputPhoto for reference.
//
// Example:
//  g, err := tg.DecodeInputPhoto(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.InputPhotoEmpty: // inputPhotoEmpty#1cd7bf0d
//  case *tg.InputPhoto: // inputPhoto#3bb3b94a
//  default: panic(v)
//  }
type InputPhotoClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputPhotoClass

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// SchemaName returns MTProto type name.
	SchemaName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// AsNotEmpty tries to map InputPhotoClass to InputPhoto.
	AsNotEmpty() (*InputPhoto, bool)
}

// AsInputPhotoFileLocation tries to map InputPhoto to InputPhotoFileLocation.
func (i *InputPhoto) AsInputPhotoFileLocation() *InputPhotoFileLocation {
	value := new(InputPhotoFileLocation)
	value.ID = i.GetID()
	value.AccessHash = i.GetAccessHash()
	value.FileReference = i.GetFileReference()

	return value
}

// AsNotEmpty tries to map InputPhotoEmpty to InputPhoto.
func (i *InputPhotoEmpty) AsNotEmpty() (*InputPhoto, bool) {
	return nil, false
}

// AsNotEmpty tries to map InputPhoto to InputPhoto.
func (i *InputPhoto) AsNotEmpty() (*InputPhoto, bool) {
	return i, true
}

// DecodeInputPhoto implements binary de-serialization for InputPhotoClass.
func DecodeInputPhoto(buf *bin.Buffer) (InputPhotoClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputPhotoEmptyTypeID:
		// Decoding inputPhotoEmpty#1cd7bf0d.
		v := InputPhotoEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPhotoClass: %w", err)
		}
		return &v, nil
	case InputPhotoTypeID:
		// Decoding inputPhoto#3bb3b94a.
		v := InputPhoto{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPhotoClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputPhotoClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputPhoto boxes the InputPhotoClass providing a helper.
type InputPhotoBox struct {
	InputPhoto InputPhotoClass
}

// Decode implements bin.Decoder for InputPhotoBox.
func (b *InputPhotoBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputPhotoBox to nil")
	}
	v, err := DecodeInputPhoto(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputPhoto = v
	return nil
}

// Encode implements bin.Encode for InputPhotoBox.
func (b *InputPhotoBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputPhoto == nil {
		return fmt.Errorf("unable to encode InputPhotoClass as nil")
	}
	return b.InputPhoto.Encode(buf)
}

// InputPhotoClassArray is adapter for slice of InputPhotoClass.
type InputPhotoClassArray []InputPhotoClass

// Sort sorts slice of InputPhotoClass.
func (s InputPhotoClassArray) Sort(less func(a, b InputPhotoClass) bool) InputPhotoClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputPhotoClass.
func (s InputPhotoClassArray) SortStable(less func(a, b InputPhotoClass) bool) InputPhotoClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputPhotoClass.
func (s InputPhotoClassArray) Retain(keep func(x InputPhotoClass) bool) InputPhotoClassArray {
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
func (s InputPhotoClassArray) First() (v InputPhotoClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputPhotoClassArray) Last() (v InputPhotoClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputPhotoClassArray) PopFirst() (v InputPhotoClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputPhotoClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputPhotoClassArray) Pop() (v InputPhotoClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsInputPhoto returns copy with only InputPhoto constructors.
func (s InputPhotoClassArray) AsInputPhoto() (to InputPhotoArray) {
	for _, elem := range s {
		value, ok := elem.(*InputPhoto)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyNotEmpty appends only NotEmpty constructors to
// given slice.
func (s InputPhotoClassArray) AppendOnlyNotEmpty(to []*InputPhoto) []*InputPhoto {
	for _, elem := range s {
		value, ok := elem.AsNotEmpty()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsNotEmpty returns copy with only NotEmpty constructors.
func (s InputPhotoClassArray) AsNotEmpty() (to []*InputPhoto) {
	return s.AppendOnlyNotEmpty(to)
}

// FirstAsNotEmpty returns first element of slice (if exists).
func (s InputPhotoClassArray) FirstAsNotEmpty() (v *InputPhoto, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// LastAsNotEmpty returns last element of slice (if exists).
func (s InputPhotoClassArray) LastAsNotEmpty() (v *InputPhoto, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopFirstAsNotEmpty returns element of slice (if exists).
func (s *InputPhotoClassArray) PopFirstAsNotEmpty() (v *InputPhoto, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopAsNotEmpty returns element of slice (if exists).
func (s *InputPhotoClassArray) PopAsNotEmpty() (v *InputPhoto, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// InputPhotoArray is adapter for slice of InputPhoto.
type InputPhotoArray []InputPhoto

// Sort sorts slice of InputPhoto.
func (s InputPhotoArray) Sort(less func(a, b InputPhoto) bool) InputPhotoArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputPhoto.
func (s InputPhotoArray) SortStable(less func(a, b InputPhoto) bool) InputPhotoArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputPhoto.
func (s InputPhotoArray) Retain(keep func(x InputPhoto) bool) InputPhotoArray {
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
func (s InputPhotoArray) First() (v InputPhoto, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputPhotoArray) Last() (v InputPhoto, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputPhotoArray) PopFirst() (v InputPhoto, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputPhoto
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputPhotoArray) Pop() (v InputPhoto, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
