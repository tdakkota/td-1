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

// InputWebDocument represents TL type `inputWebDocument#9bed434d`.
// The document
//
// See https://core.telegram.org/constructor/inputWebDocument for reference.
type InputWebDocument struct {
	// Remote document URL to be downloaded using the appropriate method¹
	//
	// Links:
	//  1) https://core.telegram.org/api/files
	URL string `schemaname:"url"`
	// Remote file size
	Size int `schemaname:"size"`
	// Mime type
	MimeType string `schemaname:"mime_type"`
	// Attributes for media types
	Attributes []DocumentAttributeClass `schemaname:"attributes"`
}

// InputWebDocumentTypeID is TL type id of InputWebDocument.
const InputWebDocumentTypeID = 0x9bed434d

func (i *InputWebDocument) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.URL == "") {
		return false
	}
	if !(i.Size == 0) {
		return false
	}
	if !(i.MimeType == "") {
		return false
	}
	if !(i.Attributes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputWebDocument) String() string {
	if i == nil {
		return "InputWebDocument(nil)"
	}
	type Alias InputWebDocument
	return fmt.Sprintf("InputWebDocument%+v", Alias(*i))
}

// FillFrom fills InputWebDocument from given interface.
func (i *InputWebDocument) FillFrom(from interface {
	GetURL() (value string)
	GetSize() (value int)
	GetMimeType() (value string)
	GetAttributes() (value []DocumentAttributeClass)
}) {
	i.URL = from.GetURL()
	i.Size = from.GetSize()
	i.MimeType = from.GetMimeType()
	i.Attributes = from.GetAttributes()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputWebDocument) TypeID() uint32 {
	return InputWebDocumentTypeID
}

// SchemaName returns MTProto type name.
func (i *InputWebDocument) SchemaName() string {
	return "inputWebDocument"
}

// Encode implements bin.Encoder.
func (i *InputWebDocument) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputWebDocument#9bed434d as nil")
	}
	b.PutID(InputWebDocumentTypeID)
	b.PutString(i.URL)
	b.PutInt(i.Size)
	b.PutString(i.MimeType)
	b.PutVectorHeader(len(i.Attributes))
	for idx, v := range i.Attributes {
		if v == nil {
			return fmt.Errorf("unable to encode inputWebDocument#9bed434d: field attributes element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputWebDocument#9bed434d: field attributes element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetURL returns value of URL field.
func (i *InputWebDocument) GetURL() (value string) {
	return i.URL
}

// GetSize returns value of Size field.
func (i *InputWebDocument) GetSize() (value int) {
	return i.Size
}

// GetMimeType returns value of MimeType field.
func (i *InputWebDocument) GetMimeType() (value string) {
	return i.MimeType
}

// GetAttributes returns value of Attributes field.
func (i *InputWebDocument) GetAttributes() (value []DocumentAttributeClass) {
	return i.Attributes
}

// MapAttributes returns field Attributes wrapped in DocumentAttributeClassArray helper.
func (i *InputWebDocument) MapAttributes() (value DocumentAttributeClassArray) {
	return DocumentAttributeClassArray(i.Attributes)
}

// Decode implements bin.Decoder.
func (i *InputWebDocument) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputWebDocument#9bed434d to nil")
	}
	if err := b.ConsumeID(InputWebDocumentTypeID); err != nil {
		return fmt.Errorf("unable to decode inputWebDocument#9bed434d: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputWebDocument#9bed434d: field url: %w", err)
		}
		i.URL = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputWebDocument#9bed434d: field size: %w", err)
		}
		i.Size = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputWebDocument#9bed434d: field mime_type: %w", err)
		}
		i.MimeType = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode inputWebDocument#9bed434d: field attributes: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDocumentAttribute(b)
			if err != nil {
				return fmt.Errorf("unable to decode inputWebDocument#9bed434d: field attributes: %w", err)
			}
			i.Attributes = append(i.Attributes, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for InputWebDocument.
var (
	_ bin.Encoder = &InputWebDocument{}
	_ bin.Decoder = &InputWebDocument{}
)
