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

// CodeSettings represents TL type `codeSettings#debebe83`.
// Settings used by telegram servers for sending the confirm code.
// Example implementations: telegram for android¹, tdlib².
//
// Links:
//  1) https://github.com/DrKLO/Telegram/blob/master/TMessagesProj/src/main/java/org/telegram/ui/LoginActivity.java
//  2) https://github.com/tdlib/td/tree/master/td/telegram/SendCodeHelper.cpp
//
// See https://core.telegram.org/constructor/codeSettings for reference.
type CodeSettings struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `schemaname:"flags"`
	// Whether to allow phone verification via phone calls¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/auth
	AllowFlashcall bool `schemaname:"allow_flashcall"`
	// Pass true if the phone number is used on the current device. Ignored if allow_flashcall is not set.
	CurrentNumber bool `schemaname:"current_number"`
	// If a token that will be included in eventually sent SMSs is required: required in newer versions of android, to use the android SMS receiver APIs¹
	//
	// Links:
	//  1) https://developers.google.com/identity/sms-retriever/overview
	AllowAppHash bool `schemaname:"allow_app_hash"`
}

// CodeSettingsTypeID is TL type id of CodeSettings.
const CodeSettingsTypeID = 0xdebebe83

func (c *CodeSettings) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.AllowFlashcall == false) {
		return false
	}
	if !(c.CurrentNumber == false) {
		return false
	}
	if !(c.AllowAppHash == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CodeSettings) String() string {
	if c == nil {
		return "CodeSettings(nil)"
	}
	type Alias CodeSettings
	return fmt.Sprintf("CodeSettings%+v", Alias(*c))
}

// FillFrom fills CodeSettings from given interface.
func (c *CodeSettings) FillFrom(from interface {
	GetAllowFlashcall() (value bool)
	GetCurrentNumber() (value bool)
	GetAllowAppHash() (value bool)
}) {
	c.AllowFlashcall = from.GetAllowFlashcall()
	c.CurrentNumber = from.GetCurrentNumber()
	c.AllowAppHash = from.GetAllowAppHash()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *CodeSettings) TypeID() uint32 {
	return CodeSettingsTypeID
}

// SchemaName returns MTProto type name.
func (c *CodeSettings) SchemaName() string {
	return "codeSettings"
}

// Encode implements bin.Encoder.
func (c *CodeSettings) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode codeSettings#debebe83 as nil")
	}
	b.PutID(CodeSettingsTypeID)
	if !(c.AllowFlashcall == false) {
		c.Flags.Set(0)
	}
	if !(c.CurrentNumber == false) {
		c.Flags.Set(1)
	}
	if !(c.AllowAppHash == false) {
		c.Flags.Set(4)
	}
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode codeSettings#debebe83: field flags: %w", err)
	}
	return nil
}

// SetAllowFlashcall sets value of AllowFlashcall conditional field.
func (c *CodeSettings) SetAllowFlashcall(value bool) {
	if value {
		c.Flags.Set(0)
		c.AllowFlashcall = true
	} else {
		c.Flags.Unset(0)
		c.AllowFlashcall = false
	}
}

// GetAllowFlashcall returns value of AllowFlashcall conditional field.
func (c *CodeSettings) GetAllowFlashcall() (value bool) {
	return c.Flags.Has(0)
}

// SetCurrentNumber sets value of CurrentNumber conditional field.
func (c *CodeSettings) SetCurrentNumber(value bool) {
	if value {
		c.Flags.Set(1)
		c.CurrentNumber = true
	} else {
		c.Flags.Unset(1)
		c.CurrentNumber = false
	}
}

// GetCurrentNumber returns value of CurrentNumber conditional field.
func (c *CodeSettings) GetCurrentNumber() (value bool) {
	return c.Flags.Has(1)
}

// SetAllowAppHash sets value of AllowAppHash conditional field.
func (c *CodeSettings) SetAllowAppHash(value bool) {
	if value {
		c.Flags.Set(4)
		c.AllowAppHash = true
	} else {
		c.Flags.Unset(4)
		c.AllowAppHash = false
	}
}

// GetAllowAppHash returns value of AllowAppHash conditional field.
func (c *CodeSettings) GetAllowAppHash() (value bool) {
	return c.Flags.Has(4)
}

// Decode implements bin.Decoder.
func (c *CodeSettings) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode codeSettings#debebe83 to nil")
	}
	if err := b.ConsumeID(CodeSettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode codeSettings#debebe83: %w", err)
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode codeSettings#debebe83: field flags: %w", err)
		}
	}
	c.AllowFlashcall = c.Flags.Has(0)
	c.CurrentNumber = c.Flags.Has(1)
	c.AllowAppHash = c.Flags.Has(4)
	return nil
}

// Ensuring interfaces in compile-time for CodeSettings.
var (
	_ bin.Encoder = &CodeSettings{}
	_ bin.Decoder = &CodeSettings{}
)
