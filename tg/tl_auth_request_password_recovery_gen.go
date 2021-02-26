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

// AuthRequestPasswordRecoveryRequest represents TL type `auth.requestPasswordRecovery#d897bc66`.
// Request recovery code of a 2FA password¹, only for accounts with a recovery email configured².
//
// Links:
//  1) https://core.telegram.org/api/srp
//  2) https://core.telegram.org/api/srp#email-verification
//
// See https://core.telegram.org/method/auth.requestPasswordRecovery for reference.
type AuthRequestPasswordRecoveryRequest struct {
}

// AuthRequestPasswordRecoveryRequestTypeID is TL type id of AuthRequestPasswordRecoveryRequest.
const AuthRequestPasswordRecoveryRequestTypeID = 0xd897bc66

func (r *AuthRequestPasswordRecoveryRequest) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *AuthRequestPasswordRecoveryRequest) String() string {
	if r == nil {
		return "AuthRequestPasswordRecoveryRequest(nil)"
	}
	type Alias AuthRequestPasswordRecoveryRequest
	return fmt.Sprintf("AuthRequestPasswordRecoveryRequest%+v", Alias(*r))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (r *AuthRequestPasswordRecoveryRequest) TypeID() uint32 {
	return AuthRequestPasswordRecoveryRequestTypeID
}

// SchemaName returns MTProto type name.
func (r *AuthRequestPasswordRecoveryRequest) SchemaName() string {
	return "auth.requestPasswordRecovery"
}

// Encode implements bin.Encoder.
func (r *AuthRequestPasswordRecoveryRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode auth.requestPasswordRecovery#d897bc66 as nil")
	}
	b.PutID(AuthRequestPasswordRecoveryRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (r *AuthRequestPasswordRecoveryRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode auth.requestPasswordRecovery#d897bc66 to nil")
	}
	if err := b.ConsumeID(AuthRequestPasswordRecoveryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.requestPasswordRecovery#d897bc66: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for AuthRequestPasswordRecoveryRequest.
var (
	_ bin.Encoder = &AuthRequestPasswordRecoveryRequest{}
	_ bin.Decoder = &AuthRequestPasswordRecoveryRequest{}
)

// AuthRequestPasswordRecovery invokes method auth.requestPasswordRecovery#d897bc66 returning error if any.
// Request recovery code of a 2FA password¹, only for accounts with a recovery email configured².
//
// Links:
//  1) https://core.telegram.org/api/srp
//  2) https://core.telegram.org/api/srp#email-verification
//
// Possible errors:
//  400 PASSWORD_EMPTY: The provided password is empty
//
// See https://core.telegram.org/method/auth.requestPasswordRecovery for reference.
func (c *Client) AuthRequestPasswordRecovery(ctx context.Context) (*AuthPasswordRecovery, error) {
	var result AuthPasswordRecovery

	request := &AuthRequestPasswordRecoveryRequest{}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
