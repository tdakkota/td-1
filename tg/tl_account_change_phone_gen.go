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

// AccountChangePhoneRequest represents TL type `account.changePhone#70c32edb`.
// Change the phone number of the current account
//
// See https://core.telegram.org/method/account.changePhone for reference.
type AccountChangePhoneRequest struct {
	// New phone number
	PhoneNumber string `schemaname:"phone_number"`
	// Phone code hash received when calling account.sendChangePhoneCode¹
	//
	// Links:
	//  1) https://core.telegram.org/method/account.sendChangePhoneCode
	PhoneCodeHash string `schemaname:"phone_code_hash"`
	// Phone code received when calling account.sendChangePhoneCode¹
	//
	// Links:
	//  1) https://core.telegram.org/method/account.sendChangePhoneCode
	PhoneCode string `schemaname:"phone_code"`
}

// AccountChangePhoneRequestTypeID is TL type id of AccountChangePhoneRequest.
const AccountChangePhoneRequestTypeID = 0x70c32edb

func (c *AccountChangePhoneRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.PhoneNumber == "") {
		return false
	}
	if !(c.PhoneCodeHash == "") {
		return false
	}
	if !(c.PhoneCode == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *AccountChangePhoneRequest) String() string {
	if c == nil {
		return "AccountChangePhoneRequest(nil)"
	}
	type Alias AccountChangePhoneRequest
	return fmt.Sprintf("AccountChangePhoneRequest%+v", Alias(*c))
}

// FillFrom fills AccountChangePhoneRequest from given interface.
func (c *AccountChangePhoneRequest) FillFrom(from interface {
	GetPhoneNumber() (value string)
	GetPhoneCodeHash() (value string)
	GetPhoneCode() (value string)
}) {
	c.PhoneNumber = from.GetPhoneNumber()
	c.PhoneCodeHash = from.GetPhoneCodeHash()
	c.PhoneCode = from.GetPhoneCode()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *AccountChangePhoneRequest) TypeID() uint32 {
	return AccountChangePhoneRequestTypeID
}

// SchemaName returns MTProto type name.
func (c *AccountChangePhoneRequest) SchemaName() string {
	return "account.changePhone"
}

// Encode implements bin.Encoder.
func (c *AccountChangePhoneRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode account.changePhone#70c32edb as nil")
	}
	b.PutID(AccountChangePhoneRequestTypeID)
	b.PutString(c.PhoneNumber)
	b.PutString(c.PhoneCodeHash)
	b.PutString(c.PhoneCode)
	return nil
}

// GetPhoneNumber returns value of PhoneNumber field.
func (c *AccountChangePhoneRequest) GetPhoneNumber() (value string) {
	return c.PhoneNumber
}

// GetPhoneCodeHash returns value of PhoneCodeHash field.
func (c *AccountChangePhoneRequest) GetPhoneCodeHash() (value string) {
	return c.PhoneCodeHash
}

// GetPhoneCode returns value of PhoneCode field.
func (c *AccountChangePhoneRequest) GetPhoneCode() (value string) {
	return c.PhoneCode
}

// Decode implements bin.Decoder.
func (c *AccountChangePhoneRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode account.changePhone#70c32edb to nil")
	}
	if err := b.ConsumeID(AccountChangePhoneRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.changePhone#70c32edb: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.changePhone#70c32edb: field phone_number: %w", err)
		}
		c.PhoneNumber = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.changePhone#70c32edb: field phone_code_hash: %w", err)
		}
		c.PhoneCodeHash = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.changePhone#70c32edb: field phone_code: %w", err)
		}
		c.PhoneCode = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountChangePhoneRequest.
var (
	_ bin.Encoder = &AccountChangePhoneRequest{}
	_ bin.Decoder = &AccountChangePhoneRequest{}
)

// AccountChangePhone invokes method account.changePhone#70c32edb returning error if any.
// Change the phone number of the current account
//
// Possible errors:
//  400 PHONE_CODE_EMPTY: phone_code is missing
//  400 PHONE_NUMBER_INVALID: The phone number is invalid
//
// See https://core.telegram.org/method/account.changePhone for reference.
func (c *Client) AccountChangePhone(ctx context.Context, request *AccountChangePhoneRequest) (UserClass, error) {
	var result UserBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.User, nil
}
