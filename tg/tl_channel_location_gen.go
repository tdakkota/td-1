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

// ChannelLocationEmpty represents TL type `channelLocationEmpty#bfb5ad8b`.
// No location (normal supergroup)
//
// See https://core.telegram.org/constructor/channelLocationEmpty for reference.
type ChannelLocationEmpty struct {
}

// ChannelLocationEmptyTypeID is TL type id of ChannelLocationEmpty.
const ChannelLocationEmptyTypeID = 0xbfb5ad8b

func (c *ChannelLocationEmpty) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChannelLocationEmpty) String() string {
	if c == nil {
		return "ChannelLocationEmpty(nil)"
	}
	type Alias ChannelLocationEmpty
	return fmt.Sprintf("ChannelLocationEmpty%+v", Alias(*c))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *ChannelLocationEmpty) TypeID() uint32 {
	return ChannelLocationEmptyTypeID
}

// SchemaName returns MTProto type name.
func (c *ChannelLocationEmpty) SchemaName() string {
	return "channelLocationEmpty"
}

// Encode implements bin.Encoder.
func (c *ChannelLocationEmpty) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelLocationEmpty#bfb5ad8b as nil")
	}
	b.PutID(ChannelLocationEmptyTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChannelLocationEmpty) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelLocationEmpty#bfb5ad8b to nil")
	}
	if err := b.ConsumeID(ChannelLocationEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode channelLocationEmpty#bfb5ad8b: %w", err)
	}
	return nil
}

// construct implements constructor of ChannelLocationClass.
func (c ChannelLocationEmpty) construct() ChannelLocationClass { return &c }

// Ensuring interfaces in compile-time for ChannelLocationEmpty.
var (
	_ bin.Encoder = &ChannelLocationEmpty{}
	_ bin.Decoder = &ChannelLocationEmpty{}

	_ ChannelLocationClass = &ChannelLocationEmpty{}
)

// ChannelLocation represents TL type `channelLocation#209b82db`.
// Geographical location of supergroup (geogroups)
//
// See https://core.telegram.org/constructor/channelLocation for reference.
type ChannelLocation struct {
	// Geographical location of supergrup
	GeoPoint GeoPointClass `schemaname:"geo_point"`
	// Textual description of the address
	Address string `schemaname:"address"`
}

// ChannelLocationTypeID is TL type id of ChannelLocation.
const ChannelLocationTypeID = 0x209b82db

func (c *ChannelLocation) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.GeoPoint == nil) {
		return false
	}
	if !(c.Address == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChannelLocation) String() string {
	if c == nil {
		return "ChannelLocation(nil)"
	}
	type Alias ChannelLocation
	return fmt.Sprintf("ChannelLocation%+v", Alias(*c))
}

// FillFrom fills ChannelLocation from given interface.
func (c *ChannelLocation) FillFrom(from interface {
	GetGeoPoint() (value GeoPointClass)
	GetAddress() (value string)
}) {
	c.GeoPoint = from.GetGeoPoint()
	c.Address = from.GetAddress()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *ChannelLocation) TypeID() uint32 {
	return ChannelLocationTypeID
}

// SchemaName returns MTProto type name.
func (c *ChannelLocation) SchemaName() string {
	return "channelLocation"
}

// Encode implements bin.Encoder.
func (c *ChannelLocation) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelLocation#209b82db as nil")
	}
	b.PutID(ChannelLocationTypeID)
	if c.GeoPoint == nil {
		return fmt.Errorf("unable to encode channelLocation#209b82db: field geo_point is nil")
	}
	if err := c.GeoPoint.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channelLocation#209b82db: field geo_point: %w", err)
	}
	b.PutString(c.Address)
	return nil
}

// GetGeoPoint returns value of GeoPoint field.
func (c *ChannelLocation) GetGeoPoint() (value GeoPointClass) {
	return c.GeoPoint
}

// GetAddress returns value of Address field.
func (c *ChannelLocation) GetAddress() (value string) {
	return c.Address
}

// Decode implements bin.Decoder.
func (c *ChannelLocation) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelLocation#209b82db to nil")
	}
	if err := b.ConsumeID(ChannelLocationTypeID); err != nil {
		return fmt.Errorf("unable to decode channelLocation#209b82db: %w", err)
	}
	{
		value, err := DecodeGeoPoint(b)
		if err != nil {
			return fmt.Errorf("unable to decode channelLocation#209b82db: field geo_point: %w", err)
		}
		c.GeoPoint = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channelLocation#209b82db: field address: %w", err)
		}
		c.Address = value
	}
	return nil
}

// construct implements constructor of ChannelLocationClass.
func (c ChannelLocation) construct() ChannelLocationClass { return &c }

// Ensuring interfaces in compile-time for ChannelLocation.
var (
	_ bin.Encoder = &ChannelLocation{}
	_ bin.Decoder = &ChannelLocation{}

	_ ChannelLocationClass = &ChannelLocation{}
)

// ChannelLocationClass represents ChannelLocation generic type.
//
// See https://core.telegram.org/type/ChannelLocation for reference.
//
// Example:
//  g, err := tg.DecodeChannelLocation(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.ChannelLocationEmpty: // channelLocationEmpty#bfb5ad8b
//  case *tg.ChannelLocation: // channelLocation#209b82db
//  default: panic(v)
//  }
type ChannelLocationClass interface {
	bin.Encoder
	bin.Decoder
	construct() ChannelLocationClass

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// SchemaName returns MTProto type name.
	SchemaName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// AsNotEmpty tries to map ChannelLocationClass to ChannelLocation.
	AsNotEmpty() (*ChannelLocation, bool)
}

// AsNotEmpty tries to map ChannelLocationEmpty to ChannelLocation.
func (c *ChannelLocationEmpty) AsNotEmpty() (*ChannelLocation, bool) {
	return nil, false
}

// AsNotEmpty tries to map ChannelLocation to ChannelLocation.
func (c *ChannelLocation) AsNotEmpty() (*ChannelLocation, bool) {
	return c, true
}

// DecodeChannelLocation implements binary de-serialization for ChannelLocationClass.
func DecodeChannelLocation(buf *bin.Buffer) (ChannelLocationClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ChannelLocationEmptyTypeID:
		// Decoding channelLocationEmpty#bfb5ad8b.
		v := ChannelLocationEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChannelLocationClass: %w", err)
		}
		return &v, nil
	case ChannelLocationTypeID:
		// Decoding channelLocation#209b82db.
		v := ChannelLocation{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChannelLocationClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ChannelLocationClass: %w", bin.NewUnexpectedID(id))
	}
}

// ChannelLocation boxes the ChannelLocationClass providing a helper.
type ChannelLocationBox struct {
	ChannelLocation ChannelLocationClass
}

// Decode implements bin.Decoder for ChannelLocationBox.
func (b *ChannelLocationBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ChannelLocationBox to nil")
	}
	v, err := DecodeChannelLocation(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ChannelLocation = v
	return nil
}

// Encode implements bin.Encode for ChannelLocationBox.
func (b *ChannelLocationBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ChannelLocation == nil {
		return fmt.Errorf("unable to encode ChannelLocationClass as nil")
	}
	return b.ChannelLocation.Encode(buf)
}

// ChannelLocationClassArray is adapter for slice of ChannelLocationClass.
type ChannelLocationClassArray []ChannelLocationClass

// Sort sorts slice of ChannelLocationClass.
func (s ChannelLocationClassArray) Sort(less func(a, b ChannelLocationClass) bool) ChannelLocationClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ChannelLocationClass.
func (s ChannelLocationClassArray) SortStable(less func(a, b ChannelLocationClass) bool) ChannelLocationClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ChannelLocationClass.
func (s ChannelLocationClassArray) Retain(keep func(x ChannelLocationClass) bool) ChannelLocationClassArray {
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
func (s ChannelLocationClassArray) First() (v ChannelLocationClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ChannelLocationClassArray) Last() (v ChannelLocationClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ChannelLocationClassArray) PopFirst() (v ChannelLocationClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ChannelLocationClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ChannelLocationClassArray) Pop() (v ChannelLocationClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsChannelLocation returns copy with only ChannelLocation constructors.
func (s ChannelLocationClassArray) AsChannelLocation() (to ChannelLocationArray) {
	for _, elem := range s {
		value, ok := elem.(*ChannelLocation)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyNotEmpty appends only NotEmpty constructors to
// given slice.
func (s ChannelLocationClassArray) AppendOnlyNotEmpty(to []*ChannelLocation) []*ChannelLocation {
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
func (s ChannelLocationClassArray) AsNotEmpty() (to []*ChannelLocation) {
	return s.AppendOnlyNotEmpty(to)
}

// FirstAsNotEmpty returns first element of slice (if exists).
func (s ChannelLocationClassArray) FirstAsNotEmpty() (v *ChannelLocation, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// LastAsNotEmpty returns last element of slice (if exists).
func (s ChannelLocationClassArray) LastAsNotEmpty() (v *ChannelLocation, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopFirstAsNotEmpty returns element of slice (if exists).
func (s *ChannelLocationClassArray) PopFirstAsNotEmpty() (v *ChannelLocation, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopAsNotEmpty returns element of slice (if exists).
func (s *ChannelLocationClassArray) PopAsNotEmpty() (v *ChannelLocation, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// ChannelLocationArray is adapter for slice of ChannelLocation.
type ChannelLocationArray []ChannelLocation

// Sort sorts slice of ChannelLocation.
func (s ChannelLocationArray) Sort(less func(a, b ChannelLocation) bool) ChannelLocationArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ChannelLocation.
func (s ChannelLocationArray) SortStable(less func(a, b ChannelLocation) bool) ChannelLocationArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ChannelLocation.
func (s ChannelLocationArray) Retain(keep func(x ChannelLocation) bool) ChannelLocationArray {
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
func (s ChannelLocationArray) First() (v ChannelLocation, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ChannelLocationArray) Last() (v ChannelLocation, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ChannelLocationArray) PopFirst() (v ChannelLocation, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ChannelLocation
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ChannelLocationArray) Pop() (v ChannelLocation, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
