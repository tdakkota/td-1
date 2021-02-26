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

// HelpTermsOfService represents TL type `help.termsOfService#780a0310`.
// Info about the latest telegram Terms Of Service
//
// See https://core.telegram.org/constructor/help.termsOfService for reference.
type HelpTermsOfService struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `schemaname:"flags"`
	// Whether a prompt must be showed to the user, in order to accept the new terms.
	Popup bool `schemaname:"popup"`
	// ID of the new terms
	ID DataJSON `schemaname:"id"`
	// Text of the new terms
	Text string `schemaname:"text"`
	// Message entities for styled text¹
	//
	// Links:
	//  1) https://core.telegram.org/api/entities
	Entities []MessageEntityClass `schemaname:"entities"`
	// Minimum age required to sign up to telegram, the user must confirm that they is older than the minimum age.
	//
	// Use SetMinAgeConfirm and GetMinAgeConfirm helpers.
	MinAgeConfirm int `schemaname:"min_age_confirm"`
}

// HelpTermsOfServiceTypeID is TL type id of HelpTermsOfService.
const HelpTermsOfServiceTypeID = 0x780a0310

func (t *HelpTermsOfService) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Flags.Zero()) {
		return false
	}
	if !(t.Popup == false) {
		return false
	}
	if !(t.ID.Zero()) {
		return false
	}
	if !(t.Text == "") {
		return false
	}
	if !(t.Entities == nil) {
		return false
	}
	if !(t.MinAgeConfirm == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *HelpTermsOfService) String() string {
	if t == nil {
		return "HelpTermsOfService(nil)"
	}
	type Alias HelpTermsOfService
	return fmt.Sprintf("HelpTermsOfService%+v", Alias(*t))
}

// FillFrom fills HelpTermsOfService from given interface.
func (t *HelpTermsOfService) FillFrom(from interface {
	GetPopup() (value bool)
	GetID() (value DataJSON)
	GetText() (value string)
	GetEntities() (value []MessageEntityClass)
	GetMinAgeConfirm() (value int, ok bool)
}) {
	t.Popup = from.GetPopup()
	t.ID = from.GetID()
	t.Text = from.GetText()
	t.Entities = from.GetEntities()
	if val, ok := from.GetMinAgeConfirm(); ok {
		t.MinAgeConfirm = val
	}

}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (t *HelpTermsOfService) TypeID() uint32 {
	return HelpTermsOfServiceTypeID
}

// SchemaName returns MTProto type name.
func (t *HelpTermsOfService) SchemaName() string {
	return "help.termsOfService"
}

// Encode implements bin.Encoder.
func (t *HelpTermsOfService) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode help.termsOfService#780a0310 as nil")
	}
	b.PutID(HelpTermsOfServiceTypeID)
	if !(t.Popup == false) {
		t.Flags.Set(0)
	}
	if !(t.MinAgeConfirm == 0) {
		t.Flags.Set(1)
	}
	if err := t.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode help.termsOfService#780a0310: field flags: %w", err)
	}
	if err := t.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode help.termsOfService#780a0310: field id: %w", err)
	}
	b.PutString(t.Text)
	b.PutVectorHeader(len(t.Entities))
	for idx, v := range t.Entities {
		if v == nil {
			return fmt.Errorf("unable to encode help.termsOfService#780a0310: field entities element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode help.termsOfService#780a0310: field entities element with index %d: %w", idx, err)
		}
	}
	if t.Flags.Has(1) {
		b.PutInt(t.MinAgeConfirm)
	}
	return nil
}

// SetPopup sets value of Popup conditional field.
func (t *HelpTermsOfService) SetPopup(value bool) {
	if value {
		t.Flags.Set(0)
		t.Popup = true
	} else {
		t.Flags.Unset(0)
		t.Popup = false
	}
}

// GetPopup returns value of Popup conditional field.
func (t *HelpTermsOfService) GetPopup() (value bool) {
	return t.Flags.Has(0)
}

// GetID returns value of ID field.
func (t *HelpTermsOfService) GetID() (value DataJSON) {
	return t.ID
}

// GetText returns value of Text field.
func (t *HelpTermsOfService) GetText() (value string) {
	return t.Text
}

// GetEntities returns value of Entities field.
func (t *HelpTermsOfService) GetEntities() (value []MessageEntityClass) {
	return t.Entities
}

// MapEntities returns field Entities wrapped in MessageEntityClassArray helper.
func (t *HelpTermsOfService) MapEntities() (value MessageEntityClassArray) {
	return MessageEntityClassArray(t.Entities)
}

// SetMinAgeConfirm sets value of MinAgeConfirm conditional field.
func (t *HelpTermsOfService) SetMinAgeConfirm(value int) {
	t.Flags.Set(1)
	t.MinAgeConfirm = value
}

// GetMinAgeConfirm returns value of MinAgeConfirm conditional field and
// boolean which is true if field was set.
func (t *HelpTermsOfService) GetMinAgeConfirm() (value int, ok bool) {
	if !t.Flags.Has(1) {
		return value, false
	}
	return t.MinAgeConfirm, true
}

// Decode implements bin.Decoder.
func (t *HelpTermsOfService) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode help.termsOfService#780a0310 to nil")
	}
	if err := b.ConsumeID(HelpTermsOfServiceTypeID); err != nil {
		return fmt.Errorf("unable to decode help.termsOfService#780a0310: %w", err)
	}
	{
		if err := t.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode help.termsOfService#780a0310: field flags: %w", err)
		}
	}
	t.Popup = t.Flags.Has(0)
	{
		if err := t.ID.Decode(b); err != nil {
			return fmt.Errorf("unable to decode help.termsOfService#780a0310: field id: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.termsOfService#780a0310: field text: %w", err)
		}
		t.Text = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode help.termsOfService#780a0310: field entities: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode help.termsOfService#780a0310: field entities: %w", err)
			}
			t.Entities = append(t.Entities, value)
		}
	}
	if t.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode help.termsOfService#780a0310: field min_age_confirm: %w", err)
		}
		t.MinAgeConfirm = value
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpTermsOfService.
var (
	_ bin.Encoder = &HelpTermsOfService{}
	_ bin.Decoder = &HelpTermsOfService{}
)
