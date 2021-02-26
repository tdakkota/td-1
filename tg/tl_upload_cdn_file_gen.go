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

// UploadCdnFileReuploadNeeded represents TL type `upload.cdnFileReuploadNeeded#eea8e46e`.
// The file was cleared from the temporary RAM cache of the CDN¹ and has to be reuploaded.
//
// Links:
//  1) https://core.telegram.org/cdn
//
// See https://core.telegram.org/constructor/upload.cdnFileReuploadNeeded for reference.
type UploadCdnFileReuploadNeeded struct {
	// Request token (see CDN¹)
	//
	// Links:
	//  1) https://core.telegram.org/cdn
	RequestToken []byte `schemaname:"request_token"`
}

// UploadCdnFileReuploadNeededTypeID is TL type id of UploadCdnFileReuploadNeeded.
const UploadCdnFileReuploadNeededTypeID = 0xeea8e46e

func (c *UploadCdnFileReuploadNeeded) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.RequestToken == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *UploadCdnFileReuploadNeeded) String() string {
	if c == nil {
		return "UploadCdnFileReuploadNeeded(nil)"
	}
	type Alias UploadCdnFileReuploadNeeded
	return fmt.Sprintf("UploadCdnFileReuploadNeeded%+v", Alias(*c))
}

// FillFrom fills UploadCdnFileReuploadNeeded from given interface.
func (c *UploadCdnFileReuploadNeeded) FillFrom(from interface {
	GetRequestToken() (value []byte)
}) {
	c.RequestToken = from.GetRequestToken()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *UploadCdnFileReuploadNeeded) TypeID() uint32 {
	return UploadCdnFileReuploadNeededTypeID
}

// SchemaName returns MTProto type name.
func (c *UploadCdnFileReuploadNeeded) SchemaName() string {
	return "upload.cdnFileReuploadNeeded"
}

// Encode implements bin.Encoder.
func (c *UploadCdnFileReuploadNeeded) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode upload.cdnFileReuploadNeeded#eea8e46e as nil")
	}
	b.PutID(UploadCdnFileReuploadNeededTypeID)
	b.PutBytes(c.RequestToken)
	return nil
}

// GetRequestToken returns value of RequestToken field.
func (c *UploadCdnFileReuploadNeeded) GetRequestToken() (value []byte) {
	return c.RequestToken
}

// Decode implements bin.Decoder.
func (c *UploadCdnFileReuploadNeeded) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode upload.cdnFileReuploadNeeded#eea8e46e to nil")
	}
	if err := b.ConsumeID(UploadCdnFileReuploadNeededTypeID); err != nil {
		return fmt.Errorf("unable to decode upload.cdnFileReuploadNeeded#eea8e46e: %w", err)
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode upload.cdnFileReuploadNeeded#eea8e46e: field request_token: %w", err)
		}
		c.RequestToken = value
	}
	return nil
}

// construct implements constructor of UploadCdnFileClass.
func (c UploadCdnFileReuploadNeeded) construct() UploadCdnFileClass { return &c }

// Ensuring interfaces in compile-time for UploadCdnFileReuploadNeeded.
var (
	_ bin.Encoder = &UploadCdnFileReuploadNeeded{}
	_ bin.Decoder = &UploadCdnFileReuploadNeeded{}

	_ UploadCdnFileClass = &UploadCdnFileReuploadNeeded{}
)

// UploadCdnFile represents TL type `upload.cdnFile#a99fca4f`.
// Represent a chunk of a CDN¹ file.
//
// Links:
//  1) https://core.telegram.org/cdn
//
// See https://core.telegram.org/constructor/upload.cdnFile for reference.
type UploadCdnFile struct {
	// The data
	Bytes []byte `schemaname:"bytes"`
}

// UploadCdnFileTypeID is TL type id of UploadCdnFile.
const UploadCdnFileTypeID = 0xa99fca4f

func (c *UploadCdnFile) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Bytes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *UploadCdnFile) String() string {
	if c == nil {
		return "UploadCdnFile(nil)"
	}
	type Alias UploadCdnFile
	return fmt.Sprintf("UploadCdnFile%+v", Alias(*c))
}

// FillFrom fills UploadCdnFile from given interface.
func (c *UploadCdnFile) FillFrom(from interface {
	GetBytes() (value []byte)
}) {
	c.Bytes = from.GetBytes()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *UploadCdnFile) TypeID() uint32 {
	return UploadCdnFileTypeID
}

// SchemaName returns MTProto type name.
func (c *UploadCdnFile) SchemaName() string {
	return "upload.cdnFile"
}

// Encode implements bin.Encoder.
func (c *UploadCdnFile) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode upload.cdnFile#a99fca4f as nil")
	}
	b.PutID(UploadCdnFileTypeID)
	b.PutBytes(c.Bytes)
	return nil
}

// GetBytes returns value of Bytes field.
func (c *UploadCdnFile) GetBytes() (value []byte) {
	return c.Bytes
}

// Decode implements bin.Decoder.
func (c *UploadCdnFile) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode upload.cdnFile#a99fca4f to nil")
	}
	if err := b.ConsumeID(UploadCdnFileTypeID); err != nil {
		return fmt.Errorf("unable to decode upload.cdnFile#a99fca4f: %w", err)
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode upload.cdnFile#a99fca4f: field bytes: %w", err)
		}
		c.Bytes = value
	}
	return nil
}

// construct implements constructor of UploadCdnFileClass.
func (c UploadCdnFile) construct() UploadCdnFileClass { return &c }

// Ensuring interfaces in compile-time for UploadCdnFile.
var (
	_ bin.Encoder = &UploadCdnFile{}
	_ bin.Decoder = &UploadCdnFile{}

	_ UploadCdnFileClass = &UploadCdnFile{}
)

// UploadCdnFileClass represents upload.CdnFile generic type.
//
// See https://core.telegram.org/type/upload.CdnFile for reference.
//
// Example:
//  g, err := tg.DecodeUploadCdnFile(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.UploadCdnFileReuploadNeeded: // upload.cdnFileReuploadNeeded#eea8e46e
//  case *tg.UploadCdnFile: // upload.cdnFile#a99fca4f
//  default: panic(v)
//  }
type UploadCdnFileClass interface {
	bin.Encoder
	bin.Decoder
	construct() UploadCdnFileClass

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

// DecodeUploadCdnFile implements binary de-serialization for UploadCdnFileClass.
func DecodeUploadCdnFile(buf *bin.Buffer) (UploadCdnFileClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case UploadCdnFileReuploadNeededTypeID:
		// Decoding upload.cdnFileReuploadNeeded#eea8e46e.
		v := UploadCdnFileReuploadNeeded{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UploadCdnFileClass: %w", err)
		}
		return &v, nil
	case UploadCdnFileTypeID:
		// Decoding upload.cdnFile#a99fca4f.
		v := UploadCdnFile{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UploadCdnFileClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode UploadCdnFileClass: %w", bin.NewUnexpectedID(id))
	}
}

// UploadCdnFile boxes the UploadCdnFileClass providing a helper.
type UploadCdnFileBox struct {
	CdnFile UploadCdnFileClass
}

// Decode implements bin.Decoder for UploadCdnFileBox.
func (b *UploadCdnFileBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode UploadCdnFileBox to nil")
	}
	v, err := DecodeUploadCdnFile(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.CdnFile = v
	return nil
}

// Encode implements bin.Encode for UploadCdnFileBox.
func (b *UploadCdnFileBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.CdnFile == nil {
		return fmt.Errorf("unable to encode UploadCdnFileClass as nil")
	}
	return b.CdnFile.Encode(buf)
}

// UploadCdnFileClassArray is adapter for slice of UploadCdnFileClass.
type UploadCdnFileClassArray []UploadCdnFileClass

// Sort sorts slice of UploadCdnFileClass.
func (s UploadCdnFileClassArray) Sort(less func(a, b UploadCdnFileClass) bool) UploadCdnFileClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of UploadCdnFileClass.
func (s UploadCdnFileClassArray) SortStable(less func(a, b UploadCdnFileClass) bool) UploadCdnFileClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of UploadCdnFileClass.
func (s UploadCdnFileClassArray) Retain(keep func(x UploadCdnFileClass) bool) UploadCdnFileClassArray {
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
func (s UploadCdnFileClassArray) First() (v UploadCdnFileClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UploadCdnFileClassArray) Last() (v UploadCdnFileClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UploadCdnFileClassArray) PopFirst() (v UploadCdnFileClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero UploadCdnFileClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UploadCdnFileClassArray) Pop() (v UploadCdnFileClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsUploadCdnFileReuploadNeeded returns copy with only UploadCdnFileReuploadNeeded constructors.
func (s UploadCdnFileClassArray) AsUploadCdnFileReuploadNeeded() (to UploadCdnFileReuploadNeededArray) {
	for _, elem := range s {
		value, ok := elem.(*UploadCdnFileReuploadNeeded)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsUploadCdnFile returns copy with only UploadCdnFile constructors.
func (s UploadCdnFileClassArray) AsUploadCdnFile() (to UploadCdnFileArray) {
	for _, elem := range s {
		value, ok := elem.(*UploadCdnFile)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// UploadCdnFileReuploadNeededArray is adapter for slice of UploadCdnFileReuploadNeeded.
type UploadCdnFileReuploadNeededArray []UploadCdnFileReuploadNeeded

// Sort sorts slice of UploadCdnFileReuploadNeeded.
func (s UploadCdnFileReuploadNeededArray) Sort(less func(a, b UploadCdnFileReuploadNeeded) bool) UploadCdnFileReuploadNeededArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of UploadCdnFileReuploadNeeded.
func (s UploadCdnFileReuploadNeededArray) SortStable(less func(a, b UploadCdnFileReuploadNeeded) bool) UploadCdnFileReuploadNeededArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of UploadCdnFileReuploadNeeded.
func (s UploadCdnFileReuploadNeededArray) Retain(keep func(x UploadCdnFileReuploadNeeded) bool) UploadCdnFileReuploadNeededArray {
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
func (s UploadCdnFileReuploadNeededArray) First() (v UploadCdnFileReuploadNeeded, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UploadCdnFileReuploadNeededArray) Last() (v UploadCdnFileReuploadNeeded, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UploadCdnFileReuploadNeededArray) PopFirst() (v UploadCdnFileReuploadNeeded, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero UploadCdnFileReuploadNeeded
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UploadCdnFileReuploadNeededArray) Pop() (v UploadCdnFileReuploadNeeded, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// UploadCdnFileArray is adapter for slice of UploadCdnFile.
type UploadCdnFileArray []UploadCdnFile

// Sort sorts slice of UploadCdnFile.
func (s UploadCdnFileArray) Sort(less func(a, b UploadCdnFile) bool) UploadCdnFileArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of UploadCdnFile.
func (s UploadCdnFileArray) SortStable(less func(a, b UploadCdnFile) bool) UploadCdnFileArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of UploadCdnFile.
func (s UploadCdnFileArray) Retain(keep func(x UploadCdnFile) bool) UploadCdnFileArray {
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
func (s UploadCdnFileArray) First() (v UploadCdnFile, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UploadCdnFileArray) Last() (v UploadCdnFile, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UploadCdnFileArray) PopFirst() (v UploadCdnFile, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero UploadCdnFile
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UploadCdnFileArray) Pop() (v UploadCdnFile, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
