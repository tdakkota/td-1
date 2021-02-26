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

// AuthBindTempAuthKeyRequest represents TL type `auth.bindTempAuthKey#cdd42a05`.
// Binds a temporary authorization key temp_auth_key_id to the permanent authorization key perm_auth_key_id. Each permanent key may only be bound to one temporary key at a time, binding a new temporary key overwrites the previous one.
// For more information, see Perfect Forward Secrecy¹.
//
// Links:
//  1) https://core.telegram.org/api/pfs
//
// See https://core.telegram.org/method/auth.bindTempAuthKey for reference.
type AuthBindTempAuthKeyRequest struct {
	// Permanent auth_key_id to bind to
	PermAuthKeyID int64 `schemaname:"perm_auth_key_id"`
	// Random long from Binding message contents¹
	//
	// Links:
	//  1) https://core.telegram.org#binding-message-contents
	Nonce int64 `schemaname:"nonce"`
	// Unix timestamp to invalidate temporary key, see Binding message contents¹
	//
	// Links:
	//  1) https://core.telegram.org#binding-message-contents
	ExpiresAt int `schemaname:"expires_at"`
	// See Generating encrypted_message¹
	//
	// Links:
	//  1) https://core.telegram.org#generating-encrypted-message
	EncryptedMessage []byte `schemaname:"encrypted_message"`
}

// AuthBindTempAuthKeyRequestTypeID is TL type id of AuthBindTempAuthKeyRequest.
const AuthBindTempAuthKeyRequestTypeID = 0xcdd42a05

func (b *AuthBindTempAuthKeyRequest) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.PermAuthKeyID == 0) {
		return false
	}
	if !(b.Nonce == 0) {
		return false
	}
	if !(b.ExpiresAt == 0) {
		return false
	}
	if !(b.EncryptedMessage == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *AuthBindTempAuthKeyRequest) String() string {
	if b == nil {
		return "AuthBindTempAuthKeyRequest(nil)"
	}
	type Alias AuthBindTempAuthKeyRequest
	return fmt.Sprintf("AuthBindTempAuthKeyRequest%+v", Alias(*b))
}

// FillFrom fills AuthBindTempAuthKeyRequest from given interface.
func (b *AuthBindTempAuthKeyRequest) FillFrom(from interface {
	GetPermAuthKeyID() (value int64)
	GetNonce() (value int64)
	GetExpiresAt() (value int)
	GetEncryptedMessage() (value []byte)
}) {
	b.PermAuthKeyID = from.GetPermAuthKeyID()
	b.Nonce = from.GetNonce()
	b.ExpiresAt = from.GetExpiresAt()
	b.EncryptedMessage = from.GetEncryptedMessage()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (b *AuthBindTempAuthKeyRequest) TypeID() uint32 {
	return AuthBindTempAuthKeyRequestTypeID
}

// SchemaName returns MTProto type name.
func (b *AuthBindTempAuthKeyRequest) SchemaName() string {
	return "auth.bindTempAuthKey"
}

// Encode implements bin.Encoder.
func (b *AuthBindTempAuthKeyRequest) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode auth.bindTempAuthKey#cdd42a05 as nil")
	}
	buf.PutID(AuthBindTempAuthKeyRequestTypeID)
	buf.PutLong(b.PermAuthKeyID)
	buf.PutLong(b.Nonce)
	buf.PutInt(b.ExpiresAt)
	buf.PutBytes(b.EncryptedMessage)
	return nil
}

// GetPermAuthKeyID returns value of PermAuthKeyID field.
func (b *AuthBindTempAuthKeyRequest) GetPermAuthKeyID() (value int64) {
	return b.PermAuthKeyID
}

// GetNonce returns value of Nonce field.
func (b *AuthBindTempAuthKeyRequest) GetNonce() (value int64) {
	return b.Nonce
}

// GetExpiresAt returns value of ExpiresAt field.
func (b *AuthBindTempAuthKeyRequest) GetExpiresAt() (value int) {
	return b.ExpiresAt
}

// GetEncryptedMessage returns value of EncryptedMessage field.
func (b *AuthBindTempAuthKeyRequest) GetEncryptedMessage() (value []byte) {
	return b.EncryptedMessage
}

// Decode implements bin.Decoder.
func (b *AuthBindTempAuthKeyRequest) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode auth.bindTempAuthKey#cdd42a05 to nil")
	}
	if err := buf.ConsumeID(AuthBindTempAuthKeyRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.bindTempAuthKey#cdd42a05: %w", err)
	}
	{
		value, err := buf.Long()
		if err != nil {
			return fmt.Errorf("unable to decode auth.bindTempAuthKey#cdd42a05: field perm_auth_key_id: %w", err)
		}
		b.PermAuthKeyID = value
	}
	{
		value, err := buf.Long()
		if err != nil {
			return fmt.Errorf("unable to decode auth.bindTempAuthKey#cdd42a05: field nonce: %w", err)
		}
		b.Nonce = value
	}
	{
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode auth.bindTempAuthKey#cdd42a05: field expires_at: %w", err)
		}
		b.ExpiresAt = value
	}
	{
		value, err := buf.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode auth.bindTempAuthKey#cdd42a05: field encrypted_message: %w", err)
		}
		b.EncryptedMessage = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AuthBindTempAuthKeyRequest.
var (
	_ bin.Encoder = &AuthBindTempAuthKeyRequest{}
	_ bin.Decoder = &AuthBindTempAuthKeyRequest{}
)

// AuthBindTempAuthKey invokes method auth.bindTempAuthKey#cdd42a05 returning error if any.
// Binds a temporary authorization key temp_auth_key_id to the permanent authorization key perm_auth_key_id. Each permanent key may only be bound to one temporary key at a time, binding a new temporary key overwrites the previous one.
// For more information, see Perfect Forward Secrecy¹.
//
// Links:
//  1) https://core.telegram.org/api/pfs
//
// Possible errors:
//  400 ENCRYPTED_MESSAGE_INVALID: Encrypted message is incorrect
//  400 INPUT_REQUEST_TOO_LONG: The request is too big
//  400 TEMP_AUTH_KEY_ALREADY_BOUND: The passed temporary key is already bound to another perm_auth_key_id
//  400 TEMP_AUTH_KEY_EMPTY: The request was not performed with a temporary authorization key
//
// See https://core.telegram.org/method/auth.bindTempAuthKey for reference.
// Can be used by bots.
func (c *Client) AuthBindTempAuthKey(ctx context.Context, request *AuthBindTempAuthKeyRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
