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

// LangPackLanguage represents TL type `langPackLanguage#eeca5ce3`.
// Identifies a localization pack
//
// See https://core.telegram.org/constructor/langPackLanguage for reference.
type LangPackLanguage struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `schemaname:"flags"`
	// Whether the language pack is official
	Official bool `schemaname:"official"`
	// Is this a localization pack for an RTL language
	Rtl bool `schemaname:"rtl"`
	// Is this a beta localization pack?
	Beta bool `schemaname:"beta"`
	// Language name
	Name string `schemaname:"name"`
	// Language name in the language itself
	NativeName string `schemaname:"native_name"`
	// Language code (pack identifier)
	LangCode string `schemaname:"lang_code"`
	// Identifier of a base language pack; may be empty. If a string is missed in the language pack, then it should be fetched from base language pack. Unsupported in custom language packs
	//
	// Use SetBaseLangCode and GetBaseLangCode helpers.
	BaseLangCode string `schemaname:"base_lang_code"`
	// A language code to be used to apply plural forms. See https://www.unicode.org/cldr/charts/latest/supplemental/language_plural_rules.html¹ for more info
	//
	// Links:
	//  1) https://www.unicode.org/cldr/charts/latest/supplemental/language_plural_rules.html
	PluralCode string `schemaname:"plural_code"`
	// Total number of non-deleted strings from the language pack
	StringsCount int `schemaname:"strings_count"`
	// Total number of translated strings from the language pack
	TranslatedCount int `schemaname:"translated_count"`
	// Link to language translation interface; empty for custom local language packs
	TranslationsURL string `schemaname:"translations_url"`
}

// LangPackLanguageTypeID is TL type id of LangPackLanguage.
const LangPackLanguageTypeID = 0xeeca5ce3

func (l *LangPackLanguage) Zero() bool {
	if l == nil {
		return true
	}
	if !(l.Flags.Zero()) {
		return false
	}
	if !(l.Official == false) {
		return false
	}
	if !(l.Rtl == false) {
		return false
	}
	if !(l.Beta == false) {
		return false
	}
	if !(l.Name == "") {
		return false
	}
	if !(l.NativeName == "") {
		return false
	}
	if !(l.LangCode == "") {
		return false
	}
	if !(l.BaseLangCode == "") {
		return false
	}
	if !(l.PluralCode == "") {
		return false
	}
	if !(l.StringsCount == 0) {
		return false
	}
	if !(l.TranslatedCount == 0) {
		return false
	}
	if !(l.TranslationsURL == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (l *LangPackLanguage) String() string {
	if l == nil {
		return "LangPackLanguage(nil)"
	}
	type Alias LangPackLanguage
	return fmt.Sprintf("LangPackLanguage%+v", Alias(*l))
}

// FillFrom fills LangPackLanguage from given interface.
func (l *LangPackLanguage) FillFrom(from interface {
	GetOfficial() (value bool)
	GetRtl() (value bool)
	GetBeta() (value bool)
	GetName() (value string)
	GetNativeName() (value string)
	GetLangCode() (value string)
	GetBaseLangCode() (value string, ok bool)
	GetPluralCode() (value string)
	GetStringsCount() (value int)
	GetTranslatedCount() (value int)
	GetTranslationsURL() (value string)
}) {
	l.Official = from.GetOfficial()
	l.Rtl = from.GetRtl()
	l.Beta = from.GetBeta()
	l.Name = from.GetName()
	l.NativeName = from.GetNativeName()
	l.LangCode = from.GetLangCode()
	if val, ok := from.GetBaseLangCode(); ok {
		l.BaseLangCode = val
	}

	l.PluralCode = from.GetPluralCode()
	l.StringsCount = from.GetStringsCount()
	l.TranslatedCount = from.GetTranslatedCount()
	l.TranslationsURL = from.GetTranslationsURL()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (l *LangPackLanguage) TypeID() uint32 {
	return LangPackLanguageTypeID
}

// SchemaName returns MTProto type name.
func (l *LangPackLanguage) SchemaName() string {
	return "langPackLanguage"
}

// Encode implements bin.Encoder.
func (l *LangPackLanguage) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode langPackLanguage#eeca5ce3 as nil")
	}
	b.PutID(LangPackLanguageTypeID)
	if !(l.Official == false) {
		l.Flags.Set(0)
	}
	if !(l.Rtl == false) {
		l.Flags.Set(2)
	}
	if !(l.Beta == false) {
		l.Flags.Set(3)
	}
	if !(l.BaseLangCode == "") {
		l.Flags.Set(1)
	}
	if err := l.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode langPackLanguage#eeca5ce3: field flags: %w", err)
	}
	b.PutString(l.Name)
	b.PutString(l.NativeName)
	b.PutString(l.LangCode)
	if l.Flags.Has(1) {
		b.PutString(l.BaseLangCode)
	}
	b.PutString(l.PluralCode)
	b.PutInt(l.StringsCount)
	b.PutInt(l.TranslatedCount)
	b.PutString(l.TranslationsURL)
	return nil
}

// SetOfficial sets value of Official conditional field.
func (l *LangPackLanguage) SetOfficial(value bool) {
	if value {
		l.Flags.Set(0)
		l.Official = true
	} else {
		l.Flags.Unset(0)
		l.Official = false
	}
}

// GetOfficial returns value of Official conditional field.
func (l *LangPackLanguage) GetOfficial() (value bool) {
	return l.Flags.Has(0)
}

// SetRtl sets value of Rtl conditional field.
func (l *LangPackLanguage) SetRtl(value bool) {
	if value {
		l.Flags.Set(2)
		l.Rtl = true
	} else {
		l.Flags.Unset(2)
		l.Rtl = false
	}
}

// GetRtl returns value of Rtl conditional field.
func (l *LangPackLanguage) GetRtl() (value bool) {
	return l.Flags.Has(2)
}

// SetBeta sets value of Beta conditional field.
func (l *LangPackLanguage) SetBeta(value bool) {
	if value {
		l.Flags.Set(3)
		l.Beta = true
	} else {
		l.Flags.Unset(3)
		l.Beta = false
	}
}

// GetBeta returns value of Beta conditional field.
func (l *LangPackLanguage) GetBeta() (value bool) {
	return l.Flags.Has(3)
}

// GetName returns value of Name field.
func (l *LangPackLanguage) GetName() (value string) {
	return l.Name
}

// GetNativeName returns value of NativeName field.
func (l *LangPackLanguage) GetNativeName() (value string) {
	return l.NativeName
}

// GetLangCode returns value of LangCode field.
func (l *LangPackLanguage) GetLangCode() (value string) {
	return l.LangCode
}

// SetBaseLangCode sets value of BaseLangCode conditional field.
func (l *LangPackLanguage) SetBaseLangCode(value string) {
	l.Flags.Set(1)
	l.BaseLangCode = value
}

// GetBaseLangCode returns value of BaseLangCode conditional field and
// boolean which is true if field was set.
func (l *LangPackLanguage) GetBaseLangCode() (value string, ok bool) {
	if !l.Flags.Has(1) {
		return value, false
	}
	return l.BaseLangCode, true
}

// GetPluralCode returns value of PluralCode field.
func (l *LangPackLanguage) GetPluralCode() (value string) {
	return l.PluralCode
}

// GetStringsCount returns value of StringsCount field.
func (l *LangPackLanguage) GetStringsCount() (value int) {
	return l.StringsCount
}

// GetTranslatedCount returns value of TranslatedCount field.
func (l *LangPackLanguage) GetTranslatedCount() (value int) {
	return l.TranslatedCount
}

// GetTranslationsURL returns value of TranslationsURL field.
func (l *LangPackLanguage) GetTranslationsURL() (value string) {
	return l.TranslationsURL
}

// Decode implements bin.Decoder.
func (l *LangPackLanguage) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode langPackLanguage#eeca5ce3 to nil")
	}
	if err := b.ConsumeID(LangPackLanguageTypeID); err != nil {
		return fmt.Errorf("unable to decode langPackLanguage#eeca5ce3: %w", err)
	}
	{
		if err := l.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode langPackLanguage#eeca5ce3: field flags: %w", err)
		}
	}
	l.Official = l.Flags.Has(0)
	l.Rtl = l.Flags.Has(2)
	l.Beta = l.Flags.Has(3)
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langPackLanguage#eeca5ce3: field name: %w", err)
		}
		l.Name = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langPackLanguage#eeca5ce3: field native_name: %w", err)
		}
		l.NativeName = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langPackLanguage#eeca5ce3: field lang_code: %w", err)
		}
		l.LangCode = value
	}
	if l.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langPackLanguage#eeca5ce3: field base_lang_code: %w", err)
		}
		l.BaseLangCode = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langPackLanguage#eeca5ce3: field plural_code: %w", err)
		}
		l.PluralCode = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode langPackLanguage#eeca5ce3: field strings_count: %w", err)
		}
		l.StringsCount = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode langPackLanguage#eeca5ce3: field translated_count: %w", err)
		}
		l.TranslatedCount = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langPackLanguage#eeca5ce3: field translations_url: %w", err)
		}
		l.TranslationsURL = value
	}
	return nil
}

// Ensuring interfaces in compile-time for LangPackLanguage.
var (
	_ bin.Encoder = &LangPackLanguage{}
	_ bin.Decoder = &LangPackLanguage{}
)
