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

// AccountConfirmPasswordEmailRequest represents TL type `account.confirmPasswordEmail#8fdf1920`.
// Verify an email to use as 2FA recovery method¹.
//
// Links:
//  1) https://core.telegram.org/api/srp
//
// See https://core.telegram.org/method/account.confirmPasswordEmail for reference.
type AccountConfirmPasswordEmailRequest struct {
	// The phone code that was received after setting a recovery email¹
	//
	// Links:
	//  1) https://core.telegram.org/api/srp#email-verification
	Code string `schemaname:"code"`
}

// AccountConfirmPasswordEmailRequestTypeID is TL type id of AccountConfirmPasswordEmailRequest.
const AccountConfirmPasswordEmailRequestTypeID = 0x8fdf1920

func (c *AccountConfirmPasswordEmailRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Code == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *AccountConfirmPasswordEmailRequest) String() string {
	if c == nil {
		return "AccountConfirmPasswordEmailRequest(nil)"
	}
	type Alias AccountConfirmPasswordEmailRequest
	return fmt.Sprintf("AccountConfirmPasswordEmailRequest%+v", Alias(*c))
}

// FillFrom fills AccountConfirmPasswordEmailRequest from given interface.
func (c *AccountConfirmPasswordEmailRequest) FillFrom(from interface {
	GetCode() (value string)
}) {
	c.Code = from.GetCode()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *AccountConfirmPasswordEmailRequest) TypeID() uint32 {
	return AccountConfirmPasswordEmailRequestTypeID
}

// SchemaName returns MTProto type name.
func (c *AccountConfirmPasswordEmailRequest) SchemaName() string {
	return "account.confirmPasswordEmail"
}

// Encode implements bin.Encoder.
func (c *AccountConfirmPasswordEmailRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode account.confirmPasswordEmail#8fdf1920 as nil")
	}
	b.PutID(AccountConfirmPasswordEmailRequestTypeID)
	b.PutString(c.Code)
	return nil
}

// GetCode returns value of Code field.
func (c *AccountConfirmPasswordEmailRequest) GetCode() (value string) {
	return c.Code
}

// Decode implements bin.Decoder.
func (c *AccountConfirmPasswordEmailRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode account.confirmPasswordEmail#8fdf1920 to nil")
	}
	if err := b.ConsumeID(AccountConfirmPasswordEmailRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.confirmPasswordEmail#8fdf1920: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.confirmPasswordEmail#8fdf1920: field code: %w", err)
		}
		c.Code = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountConfirmPasswordEmailRequest.
var (
	_ bin.Encoder = &AccountConfirmPasswordEmailRequest{}
	_ bin.Decoder = &AccountConfirmPasswordEmailRequest{}
)

// AccountConfirmPasswordEmail invokes method account.confirmPasswordEmail#8fdf1920 returning error if any.
// Verify an email to use as 2FA recovery method¹.
//
// Links:
//  1) https://core.telegram.org/api/srp
//
// Possible errors:
//  400 CODE_INVALID: Code invalid
//  400 EMAIL_HASH_EXPIRED: Email hash expired
//
// See https://core.telegram.org/method/account.confirmPasswordEmail for reference.
func (c *Client) AccountConfirmPasswordEmail(ctx context.Context, code string) (bool, error) {
	var result BoolBox

	request := &AccountConfirmPasswordEmailRequest{
		Code: code,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
