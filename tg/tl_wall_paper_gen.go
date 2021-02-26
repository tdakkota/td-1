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

// WallPaper represents TL type `wallPaper#a437c3ed`.
// Wallpaper settings.
//
// See https://core.telegram.org/constructor/wallPaper for reference.
type WallPaper struct {
	// Identifier
	ID int64 `schemaname:"id"`
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `schemaname:"flags"`
	// Creator of the wallpaper
	Creator bool `schemaname:"creator"`
	// Whether this is the default wallpaper
	Default bool `schemaname:"default"`
	// Pattern
	Pattern bool `schemaname:"pattern"`
	// Dark mode
	Dark bool `schemaname:"dark"`
	// Access hash
	AccessHash int64 `schemaname:"access_hash"`
	// Unique wallpaper ID
	Slug string `schemaname:"slug"`
	// The actual wallpaper
	Document DocumentClass `schemaname:"document"`
	// Wallpaper settings
	//
	// Use SetSettings and GetSettings helpers.
	Settings WallPaperSettings `schemaname:"settings"`
}

// WallPaperTypeID is TL type id of WallPaper.
const WallPaperTypeID = 0xa437c3ed

func (w *WallPaper) Zero() bool {
	if w == nil {
		return true
	}
	if !(w.ID == 0) {
		return false
	}
	if !(w.Flags.Zero()) {
		return false
	}
	if !(w.Creator == false) {
		return false
	}
	if !(w.Default == false) {
		return false
	}
	if !(w.Pattern == false) {
		return false
	}
	if !(w.Dark == false) {
		return false
	}
	if !(w.AccessHash == 0) {
		return false
	}
	if !(w.Slug == "") {
		return false
	}
	if !(w.Document == nil) {
		return false
	}
	if !(w.Settings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (w *WallPaper) String() string {
	if w == nil {
		return "WallPaper(nil)"
	}
	type Alias WallPaper
	return fmt.Sprintf("WallPaper%+v", Alias(*w))
}

// FillFrom fills WallPaper from given interface.
func (w *WallPaper) FillFrom(from interface {
	GetID() (value int64)
	GetCreator() (value bool)
	GetDefault() (value bool)
	GetPattern() (value bool)
	GetDark() (value bool)
	GetAccessHash() (value int64)
	GetSlug() (value string)
	GetDocument() (value DocumentClass)
	GetSettings() (value WallPaperSettings, ok bool)
}) {
	w.ID = from.GetID()
	w.Creator = from.GetCreator()
	w.Default = from.GetDefault()
	w.Pattern = from.GetPattern()
	w.Dark = from.GetDark()
	w.AccessHash = from.GetAccessHash()
	w.Slug = from.GetSlug()
	w.Document = from.GetDocument()
	if val, ok := from.GetSettings(); ok {
		w.Settings = val
	}

}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (w *WallPaper) TypeID() uint32 {
	return WallPaperTypeID
}

// SchemaName returns MTProto type name.
func (w *WallPaper) SchemaName() string {
	return "wallPaper"
}

// Encode implements bin.Encoder.
func (w *WallPaper) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode wallPaper#a437c3ed as nil")
	}
	b.PutID(WallPaperTypeID)
	if !(w.Creator == false) {
		w.Flags.Set(0)
	}
	if !(w.Default == false) {
		w.Flags.Set(1)
	}
	if !(w.Pattern == false) {
		w.Flags.Set(3)
	}
	if !(w.Dark == false) {
		w.Flags.Set(4)
	}
	if !(w.Settings.Zero()) {
		w.Flags.Set(2)
	}
	b.PutLong(w.ID)
	if err := w.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode wallPaper#a437c3ed: field flags: %w", err)
	}
	b.PutLong(w.AccessHash)
	b.PutString(w.Slug)
	if w.Document == nil {
		return fmt.Errorf("unable to encode wallPaper#a437c3ed: field document is nil")
	}
	if err := w.Document.Encode(b); err != nil {
		return fmt.Errorf("unable to encode wallPaper#a437c3ed: field document: %w", err)
	}
	if w.Flags.Has(2) {
		if err := w.Settings.Encode(b); err != nil {
			return fmt.Errorf("unable to encode wallPaper#a437c3ed: field settings: %w", err)
		}
	}
	return nil
}

// GetID returns value of ID field.
func (w *WallPaper) GetID() (value int64) {
	return w.ID
}

// SetCreator sets value of Creator conditional field.
func (w *WallPaper) SetCreator(value bool) {
	if value {
		w.Flags.Set(0)
		w.Creator = true
	} else {
		w.Flags.Unset(0)
		w.Creator = false
	}
}

// GetCreator returns value of Creator conditional field.
func (w *WallPaper) GetCreator() (value bool) {
	return w.Flags.Has(0)
}

// SetDefault sets value of Default conditional field.
func (w *WallPaper) SetDefault(value bool) {
	if value {
		w.Flags.Set(1)
		w.Default = true
	} else {
		w.Flags.Unset(1)
		w.Default = false
	}
}

// GetDefault returns value of Default conditional field.
func (w *WallPaper) GetDefault() (value bool) {
	return w.Flags.Has(1)
}

// SetPattern sets value of Pattern conditional field.
func (w *WallPaper) SetPattern(value bool) {
	if value {
		w.Flags.Set(3)
		w.Pattern = true
	} else {
		w.Flags.Unset(3)
		w.Pattern = false
	}
}

// GetPattern returns value of Pattern conditional field.
func (w *WallPaper) GetPattern() (value bool) {
	return w.Flags.Has(3)
}

// SetDark sets value of Dark conditional field.
func (w *WallPaper) SetDark(value bool) {
	if value {
		w.Flags.Set(4)
		w.Dark = true
	} else {
		w.Flags.Unset(4)
		w.Dark = false
	}
}

// GetDark returns value of Dark conditional field.
func (w *WallPaper) GetDark() (value bool) {
	return w.Flags.Has(4)
}

// GetAccessHash returns value of AccessHash field.
func (w *WallPaper) GetAccessHash() (value int64) {
	return w.AccessHash
}

// GetSlug returns value of Slug field.
func (w *WallPaper) GetSlug() (value string) {
	return w.Slug
}

// GetDocument returns value of Document field.
func (w *WallPaper) GetDocument() (value DocumentClass) {
	return w.Document
}

// SetSettings sets value of Settings conditional field.
func (w *WallPaper) SetSettings(value WallPaperSettings) {
	w.Flags.Set(2)
	w.Settings = value
}

// GetSettings returns value of Settings conditional field and
// boolean which is true if field was set.
func (w *WallPaper) GetSettings() (value WallPaperSettings, ok bool) {
	if !w.Flags.Has(2) {
		return value, false
	}
	return w.Settings, true
}

// Decode implements bin.Decoder.
func (w *WallPaper) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode wallPaper#a437c3ed to nil")
	}
	if err := b.ConsumeID(WallPaperTypeID); err != nil {
		return fmt.Errorf("unable to decode wallPaper#a437c3ed: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode wallPaper#a437c3ed: field id: %w", err)
		}
		w.ID = value
	}
	{
		if err := w.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode wallPaper#a437c3ed: field flags: %w", err)
		}
	}
	w.Creator = w.Flags.Has(0)
	w.Default = w.Flags.Has(1)
	w.Pattern = w.Flags.Has(3)
	w.Dark = w.Flags.Has(4)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode wallPaper#a437c3ed: field access_hash: %w", err)
		}
		w.AccessHash = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode wallPaper#a437c3ed: field slug: %w", err)
		}
		w.Slug = value
	}
	{
		value, err := DecodeDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode wallPaper#a437c3ed: field document: %w", err)
		}
		w.Document = value
	}
	if w.Flags.Has(2) {
		if err := w.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode wallPaper#a437c3ed: field settings: %w", err)
		}
	}
	return nil
}

// construct implements constructor of WallPaperClass.
func (w WallPaper) construct() WallPaperClass { return &w }

// Ensuring interfaces in compile-time for WallPaper.
var (
	_ bin.Encoder = &WallPaper{}
	_ bin.Decoder = &WallPaper{}

	_ WallPaperClass = &WallPaper{}
)

// WallPaperNoFile represents TL type `wallPaperNoFile#8af40b25`.
// No file wallpaper
//
// See https://core.telegram.org/constructor/wallPaperNoFile for reference.
type WallPaperNoFile struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `schemaname:"flags"`
	// Whether this is the default wallpaper
	Default bool `schemaname:"default"`
	// Dark mode
	Dark bool `schemaname:"dark"`
	// Wallpaper settings
	//
	// Use SetSettings and GetSettings helpers.
	Settings WallPaperSettings `schemaname:"settings"`
}

// WallPaperNoFileTypeID is TL type id of WallPaperNoFile.
const WallPaperNoFileTypeID = 0x8af40b25

func (w *WallPaperNoFile) Zero() bool {
	if w == nil {
		return true
	}
	if !(w.Flags.Zero()) {
		return false
	}
	if !(w.Default == false) {
		return false
	}
	if !(w.Dark == false) {
		return false
	}
	if !(w.Settings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (w *WallPaperNoFile) String() string {
	if w == nil {
		return "WallPaperNoFile(nil)"
	}
	type Alias WallPaperNoFile
	return fmt.Sprintf("WallPaperNoFile%+v", Alias(*w))
}

// FillFrom fills WallPaperNoFile from given interface.
func (w *WallPaperNoFile) FillFrom(from interface {
	GetDefault() (value bool)
	GetDark() (value bool)
	GetSettings() (value WallPaperSettings, ok bool)
}) {
	w.Default = from.GetDefault()
	w.Dark = from.GetDark()
	if val, ok := from.GetSettings(); ok {
		w.Settings = val
	}

}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (w *WallPaperNoFile) TypeID() uint32 {
	return WallPaperNoFileTypeID
}

// SchemaName returns MTProto type name.
func (w *WallPaperNoFile) SchemaName() string {
	return "wallPaperNoFile"
}

// Encode implements bin.Encoder.
func (w *WallPaperNoFile) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode wallPaperNoFile#8af40b25 as nil")
	}
	b.PutID(WallPaperNoFileTypeID)
	if !(w.Default == false) {
		w.Flags.Set(1)
	}
	if !(w.Dark == false) {
		w.Flags.Set(4)
	}
	if !(w.Settings.Zero()) {
		w.Flags.Set(2)
	}
	if err := w.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode wallPaperNoFile#8af40b25: field flags: %w", err)
	}
	if w.Flags.Has(2) {
		if err := w.Settings.Encode(b); err != nil {
			return fmt.Errorf("unable to encode wallPaperNoFile#8af40b25: field settings: %w", err)
		}
	}
	return nil
}

// SetDefault sets value of Default conditional field.
func (w *WallPaperNoFile) SetDefault(value bool) {
	if value {
		w.Flags.Set(1)
		w.Default = true
	} else {
		w.Flags.Unset(1)
		w.Default = false
	}
}

// GetDefault returns value of Default conditional field.
func (w *WallPaperNoFile) GetDefault() (value bool) {
	return w.Flags.Has(1)
}

// SetDark sets value of Dark conditional field.
func (w *WallPaperNoFile) SetDark(value bool) {
	if value {
		w.Flags.Set(4)
		w.Dark = true
	} else {
		w.Flags.Unset(4)
		w.Dark = false
	}
}

// GetDark returns value of Dark conditional field.
func (w *WallPaperNoFile) GetDark() (value bool) {
	return w.Flags.Has(4)
}

// SetSettings sets value of Settings conditional field.
func (w *WallPaperNoFile) SetSettings(value WallPaperSettings) {
	w.Flags.Set(2)
	w.Settings = value
}

// GetSettings returns value of Settings conditional field and
// boolean which is true if field was set.
func (w *WallPaperNoFile) GetSettings() (value WallPaperSettings, ok bool) {
	if !w.Flags.Has(2) {
		return value, false
	}
	return w.Settings, true
}

// Decode implements bin.Decoder.
func (w *WallPaperNoFile) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode wallPaperNoFile#8af40b25 to nil")
	}
	if err := b.ConsumeID(WallPaperNoFileTypeID); err != nil {
		return fmt.Errorf("unable to decode wallPaperNoFile#8af40b25: %w", err)
	}
	{
		if err := w.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode wallPaperNoFile#8af40b25: field flags: %w", err)
		}
	}
	w.Default = w.Flags.Has(1)
	w.Dark = w.Flags.Has(4)
	if w.Flags.Has(2) {
		if err := w.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode wallPaperNoFile#8af40b25: field settings: %w", err)
		}
	}
	return nil
}

// construct implements constructor of WallPaperClass.
func (w WallPaperNoFile) construct() WallPaperClass { return &w }

// Ensuring interfaces in compile-time for WallPaperNoFile.
var (
	_ bin.Encoder = &WallPaperNoFile{}
	_ bin.Decoder = &WallPaperNoFile{}

	_ WallPaperClass = &WallPaperNoFile{}
)

// WallPaperClass represents WallPaper generic type.
//
// See https://core.telegram.org/type/WallPaper for reference.
//
// Example:
//  g, err := tg.DecodeWallPaper(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.WallPaper: // wallPaper#a437c3ed
//  case *tg.WallPaperNoFile: // wallPaperNoFile#8af40b25
//  default: panic(v)
//  }
type WallPaperClass interface {
	bin.Encoder
	bin.Decoder
	construct() WallPaperClass

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// SchemaName returns MTProto type name.
	SchemaName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// Whether this is the default wallpaper
	GetDefault() (value bool)
	// Dark mode
	GetDark() (value bool)
	// Wallpaper settings
	GetSettings() (value WallPaperSettings, ok bool)
}

// AsInput tries to map WallPaper to InputWallPaper.
func (w *WallPaper) AsInput() *InputWallPaper {
	value := new(InputWallPaper)
	value.ID = w.GetID()
	value.AccessHash = w.GetAccessHash()

	return value
}

// AsInputWallPaperSlug tries to map WallPaper to InputWallPaperSlug.
func (w *WallPaper) AsInputWallPaperSlug() *InputWallPaperSlug {
	value := new(InputWallPaperSlug)
	value.Slug = w.GetSlug()

	return value
}

// DecodeWallPaper implements binary de-serialization for WallPaperClass.
func DecodeWallPaper(buf *bin.Buffer) (WallPaperClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case WallPaperTypeID:
		// Decoding wallPaper#a437c3ed.
		v := WallPaper{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode WallPaperClass: %w", err)
		}
		return &v, nil
	case WallPaperNoFileTypeID:
		// Decoding wallPaperNoFile#8af40b25.
		v := WallPaperNoFile{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode WallPaperClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode WallPaperClass: %w", bin.NewUnexpectedID(id))
	}
}

// WallPaper boxes the WallPaperClass providing a helper.
type WallPaperBox struct {
	WallPaper WallPaperClass
}

// Decode implements bin.Decoder for WallPaperBox.
func (b *WallPaperBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode WallPaperBox to nil")
	}
	v, err := DecodeWallPaper(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.WallPaper = v
	return nil
}

// Encode implements bin.Encode for WallPaperBox.
func (b *WallPaperBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.WallPaper == nil {
		return fmt.Errorf("unable to encode WallPaperClass as nil")
	}
	return b.WallPaper.Encode(buf)
}

// WallPaperClassArray is adapter for slice of WallPaperClass.
type WallPaperClassArray []WallPaperClass

// Sort sorts slice of WallPaperClass.
func (s WallPaperClassArray) Sort(less func(a, b WallPaperClass) bool) WallPaperClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of WallPaperClass.
func (s WallPaperClassArray) SortStable(less func(a, b WallPaperClass) bool) WallPaperClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of WallPaperClass.
func (s WallPaperClassArray) Retain(keep func(x WallPaperClass) bool) WallPaperClassArray {
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
func (s WallPaperClassArray) First() (v WallPaperClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s WallPaperClassArray) Last() (v WallPaperClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *WallPaperClassArray) PopFirst() (v WallPaperClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero WallPaperClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *WallPaperClassArray) Pop() (v WallPaperClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsWallPaper returns copy with only WallPaper constructors.
func (s WallPaperClassArray) AsWallPaper() (to WallPaperArray) {
	for _, elem := range s {
		value, ok := elem.(*WallPaper)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsWallPaperNoFile returns copy with only WallPaperNoFile constructors.
func (s WallPaperClassArray) AsWallPaperNoFile() (to WallPaperNoFileArray) {
	for _, elem := range s {
		value, ok := elem.(*WallPaperNoFile)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// WallPaperArray is adapter for slice of WallPaper.
type WallPaperArray []WallPaper

// Sort sorts slice of WallPaper.
func (s WallPaperArray) Sort(less func(a, b WallPaper) bool) WallPaperArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of WallPaper.
func (s WallPaperArray) SortStable(less func(a, b WallPaper) bool) WallPaperArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of WallPaper.
func (s WallPaperArray) Retain(keep func(x WallPaper) bool) WallPaperArray {
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
func (s WallPaperArray) First() (v WallPaper, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s WallPaperArray) Last() (v WallPaper, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *WallPaperArray) PopFirst() (v WallPaper, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero WallPaper
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *WallPaperArray) Pop() (v WallPaper, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// WallPaperNoFileArray is adapter for slice of WallPaperNoFile.
type WallPaperNoFileArray []WallPaperNoFile

// Sort sorts slice of WallPaperNoFile.
func (s WallPaperNoFileArray) Sort(less func(a, b WallPaperNoFile) bool) WallPaperNoFileArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of WallPaperNoFile.
func (s WallPaperNoFileArray) SortStable(less func(a, b WallPaperNoFile) bool) WallPaperNoFileArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of WallPaperNoFile.
func (s WallPaperNoFileArray) Retain(keep func(x WallPaperNoFile) bool) WallPaperNoFileArray {
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
func (s WallPaperNoFileArray) First() (v WallPaperNoFile, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s WallPaperNoFileArray) Last() (v WallPaperNoFile, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *WallPaperNoFileArray) PopFirst() (v WallPaperNoFile, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero WallPaperNoFile
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *WallPaperNoFileArray) Pop() (v WallPaperNoFile, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
