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

// Page represents TL type `page#98657f0d`.
// Instant view¹ page
//
// Links:
//  1) https://instantview.telegram.org
//
// See https://core.telegram.org/constructor/page for reference.
type Page struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `schemaname:"flags"`
	// Indicates that not full page preview is available to the client and it will need to fetch full Instant View from the server using messages.getWebPagePreview¹.
	//
	// Links:
	//  1) https://core.telegram.org/method/messages.getWebPagePreview
	Part bool `schemaname:"part"`
	// Whether the page contains RTL text
	Rtl bool `schemaname:"rtl"`
	// Whether this is an IV v2¹ page
	//
	// Links:
	//  1) https://instantview.telegram.org/docs#what-39s-new-in-2-0
	V2 bool `schemaname:"v2"`
	// Original page HTTP URL
	URL string `schemaname:"url"`
	// Page elements (like with HTML elements, only as TL constructors)
	Blocks []PageBlockClass `schemaname:"blocks"`
	// Photos in page
	Photos []PhotoClass `schemaname:"photos"`
	// Media in page
	Documents []DocumentClass `schemaname:"documents"`
	// Viewcount
	//
	// Use SetViews and GetViews helpers.
	Views int `schemaname:"views"`
}

// PageTypeID is TL type id of Page.
const PageTypeID = 0x98657f0d

func (p *Page) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Flags.Zero()) {
		return false
	}
	if !(p.Part == false) {
		return false
	}
	if !(p.Rtl == false) {
		return false
	}
	if !(p.V2 == false) {
		return false
	}
	if !(p.URL == "") {
		return false
	}
	if !(p.Blocks == nil) {
		return false
	}
	if !(p.Photos == nil) {
		return false
	}
	if !(p.Documents == nil) {
		return false
	}
	if !(p.Views == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *Page) String() string {
	if p == nil {
		return "Page(nil)"
	}
	type Alias Page
	return fmt.Sprintf("Page%+v", Alias(*p))
}

// FillFrom fills Page from given interface.
func (p *Page) FillFrom(from interface {
	GetPart() (value bool)
	GetRtl() (value bool)
	GetV2() (value bool)
	GetURL() (value string)
	GetBlocks() (value []PageBlockClass)
	GetPhotos() (value []PhotoClass)
	GetDocuments() (value []DocumentClass)
	GetViews() (value int, ok bool)
}) {
	p.Part = from.GetPart()
	p.Rtl = from.GetRtl()
	p.V2 = from.GetV2()
	p.URL = from.GetURL()
	p.Blocks = from.GetBlocks()
	p.Photos = from.GetPhotos()
	p.Documents = from.GetDocuments()
	if val, ok := from.GetViews(); ok {
		p.Views = val
	}

}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *Page) TypeID() uint32 {
	return PageTypeID
}

// SchemaName returns MTProto type name.
func (p *Page) SchemaName() string {
	return "page"
}

// Encode implements bin.Encoder.
func (p *Page) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode page#98657f0d as nil")
	}
	b.PutID(PageTypeID)
	if !(p.Part == false) {
		p.Flags.Set(0)
	}
	if !(p.Rtl == false) {
		p.Flags.Set(1)
	}
	if !(p.V2 == false) {
		p.Flags.Set(2)
	}
	if !(p.Views == 0) {
		p.Flags.Set(3)
	}
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode page#98657f0d: field flags: %w", err)
	}
	b.PutString(p.URL)
	b.PutVectorHeader(len(p.Blocks))
	for idx, v := range p.Blocks {
		if v == nil {
			return fmt.Errorf("unable to encode page#98657f0d: field blocks element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode page#98657f0d: field blocks element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(p.Photos))
	for idx, v := range p.Photos {
		if v == nil {
			return fmt.Errorf("unable to encode page#98657f0d: field photos element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode page#98657f0d: field photos element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(p.Documents))
	for idx, v := range p.Documents {
		if v == nil {
			return fmt.Errorf("unable to encode page#98657f0d: field documents element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode page#98657f0d: field documents element with index %d: %w", idx, err)
		}
	}
	if p.Flags.Has(3) {
		b.PutInt(p.Views)
	}
	return nil
}

// SetPart sets value of Part conditional field.
func (p *Page) SetPart(value bool) {
	if value {
		p.Flags.Set(0)
		p.Part = true
	} else {
		p.Flags.Unset(0)
		p.Part = false
	}
}

// GetPart returns value of Part conditional field.
func (p *Page) GetPart() (value bool) {
	return p.Flags.Has(0)
}

// SetRtl sets value of Rtl conditional field.
func (p *Page) SetRtl(value bool) {
	if value {
		p.Flags.Set(1)
		p.Rtl = true
	} else {
		p.Flags.Unset(1)
		p.Rtl = false
	}
}

// GetRtl returns value of Rtl conditional field.
func (p *Page) GetRtl() (value bool) {
	return p.Flags.Has(1)
}

// SetV2 sets value of V2 conditional field.
func (p *Page) SetV2(value bool) {
	if value {
		p.Flags.Set(2)
		p.V2 = true
	} else {
		p.Flags.Unset(2)
		p.V2 = false
	}
}

// GetV2 returns value of V2 conditional field.
func (p *Page) GetV2() (value bool) {
	return p.Flags.Has(2)
}

// GetURL returns value of URL field.
func (p *Page) GetURL() (value string) {
	return p.URL
}

// GetBlocks returns value of Blocks field.
func (p *Page) GetBlocks() (value []PageBlockClass) {
	return p.Blocks
}

// MapBlocks returns field Blocks wrapped in PageBlockClassArray helper.
func (p *Page) MapBlocks() (value PageBlockClassArray) {
	return PageBlockClassArray(p.Blocks)
}

// GetPhotos returns value of Photos field.
func (p *Page) GetPhotos() (value []PhotoClass) {
	return p.Photos
}

// MapPhotos returns field Photos wrapped in PhotoClassArray helper.
func (p *Page) MapPhotos() (value PhotoClassArray) {
	return PhotoClassArray(p.Photos)
}

// GetDocuments returns value of Documents field.
func (p *Page) GetDocuments() (value []DocumentClass) {
	return p.Documents
}

// MapDocuments returns field Documents wrapped in DocumentClassArray helper.
func (p *Page) MapDocuments() (value DocumentClassArray) {
	return DocumentClassArray(p.Documents)
}

// SetViews sets value of Views conditional field.
func (p *Page) SetViews(value int) {
	p.Flags.Set(3)
	p.Views = value
}

// GetViews returns value of Views conditional field and
// boolean which is true if field was set.
func (p *Page) GetViews() (value int, ok bool) {
	if !p.Flags.Has(3) {
		return value, false
	}
	return p.Views, true
}

// Decode implements bin.Decoder.
func (p *Page) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode page#98657f0d to nil")
	}
	if err := b.ConsumeID(PageTypeID); err != nil {
		return fmt.Errorf("unable to decode page#98657f0d: %w", err)
	}
	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode page#98657f0d: field flags: %w", err)
		}
	}
	p.Part = p.Flags.Has(0)
	p.Rtl = p.Flags.Has(1)
	p.V2 = p.Flags.Has(2)
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode page#98657f0d: field url: %w", err)
		}
		p.URL = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode page#98657f0d: field blocks: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodePageBlock(b)
			if err != nil {
				return fmt.Errorf("unable to decode page#98657f0d: field blocks: %w", err)
			}
			p.Blocks = append(p.Blocks, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode page#98657f0d: field photos: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodePhoto(b)
			if err != nil {
				return fmt.Errorf("unable to decode page#98657f0d: field photos: %w", err)
			}
			p.Photos = append(p.Photos, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode page#98657f0d: field documents: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDocument(b)
			if err != nil {
				return fmt.Errorf("unable to decode page#98657f0d: field documents: %w", err)
			}
			p.Documents = append(p.Documents, value)
		}
	}
	if p.Flags.Has(3) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode page#98657f0d: field views: %w", err)
		}
		p.Views = value
	}
	return nil
}

// Ensuring interfaces in compile-time for Page.
var (
	_ bin.Encoder = &Page{}
	_ bin.Decoder = &Page{}
)
