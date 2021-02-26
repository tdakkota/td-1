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

// MaskCoords represents TL type `maskCoords#aed6dbb2`.
// Position on a photo where a mask should be placed
// The n position indicates where the mask should be placed:
//
// See https://core.telegram.org/constructor/maskCoords for reference.
type MaskCoords struct {
	// Part of the face, relative to which the mask should be placed
	N int `schemaname:"n"`
	// Shift by X-axis measured in widths of the mask scaled to the face size, from left to right. (For example, -1.0 will place the mask just to the left of the default mask position)
	X float64 `schemaname:"x"`
	// Shift by Y-axis measured in widths of the mask scaled to the face size, from left to right. (For example, -1.0 will place the mask just to the left of the default mask position)
	Y float64 `schemaname:"y"`
	// Mask scaling coefficient. (For example, 2.0 means a doubled size)
	Zoom float64 `schemaname:"zoom"`
}

// MaskCoordsTypeID is TL type id of MaskCoords.
const MaskCoordsTypeID = 0xaed6dbb2

func (m *MaskCoords) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.N == 0) {
		return false
	}
	if !(m.X == 0) {
		return false
	}
	if !(m.Y == 0) {
		return false
	}
	if !(m.Zoom == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MaskCoords) String() string {
	if m == nil {
		return "MaskCoords(nil)"
	}
	type Alias MaskCoords
	return fmt.Sprintf("MaskCoords%+v", Alias(*m))
}

// FillFrom fills MaskCoords from given interface.
func (m *MaskCoords) FillFrom(from interface {
	GetN() (value int)
	GetX() (value float64)
	GetY() (value float64)
	GetZoom() (value float64)
}) {
	m.N = from.GetN()
	m.X = from.GetX()
	m.Y = from.GetY()
	m.Zoom = from.GetZoom()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (m *MaskCoords) TypeID() uint32 {
	return MaskCoordsTypeID
}

// SchemaName returns MTProto type name.
func (m *MaskCoords) SchemaName() string {
	return "maskCoords"
}

// Encode implements bin.Encoder.
func (m *MaskCoords) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode maskCoords#aed6dbb2 as nil")
	}
	b.PutID(MaskCoordsTypeID)
	b.PutInt(m.N)
	b.PutDouble(m.X)
	b.PutDouble(m.Y)
	b.PutDouble(m.Zoom)
	return nil
}

// GetN returns value of N field.
func (m *MaskCoords) GetN() (value int) {
	return m.N
}

// GetX returns value of X field.
func (m *MaskCoords) GetX() (value float64) {
	return m.X
}

// GetY returns value of Y field.
func (m *MaskCoords) GetY() (value float64) {
	return m.Y
}

// GetZoom returns value of Zoom field.
func (m *MaskCoords) GetZoom() (value float64) {
	return m.Zoom
}

// Decode implements bin.Decoder.
func (m *MaskCoords) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode maskCoords#aed6dbb2 to nil")
	}
	if err := b.ConsumeID(MaskCoordsTypeID); err != nil {
		return fmt.Errorf("unable to decode maskCoords#aed6dbb2: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode maskCoords#aed6dbb2: field n: %w", err)
		}
		m.N = value
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode maskCoords#aed6dbb2: field x: %w", err)
		}
		m.X = value
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode maskCoords#aed6dbb2: field y: %w", err)
		}
		m.Y = value
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode maskCoords#aed6dbb2: field zoom: %w", err)
		}
		m.Zoom = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MaskCoords.
var (
	_ bin.Encoder = &MaskCoords{}
	_ bin.Decoder = &MaskCoords{}
)
