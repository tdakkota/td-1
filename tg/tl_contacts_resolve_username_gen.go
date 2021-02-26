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

// ContactsResolveUsernameRequest represents TL type `contacts.resolveUsername#f93ccba3`.
// Resolve a @username to get peer info
//
// See https://core.telegram.org/method/contacts.resolveUsername for reference.
type ContactsResolveUsernameRequest struct {
	// @username to resolve
	Username string `schemaname:"username"`
}

// ContactsResolveUsernameRequestTypeID is TL type id of ContactsResolveUsernameRequest.
const ContactsResolveUsernameRequestTypeID = 0xf93ccba3

func (r *ContactsResolveUsernameRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Username == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ContactsResolveUsernameRequest) String() string {
	if r == nil {
		return "ContactsResolveUsernameRequest(nil)"
	}
	type Alias ContactsResolveUsernameRequest
	return fmt.Sprintf("ContactsResolveUsernameRequest%+v", Alias(*r))
}

// FillFrom fills ContactsResolveUsernameRequest from given interface.
func (r *ContactsResolveUsernameRequest) FillFrom(from interface {
	GetUsername() (value string)
}) {
	r.Username = from.GetUsername()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (r *ContactsResolveUsernameRequest) TypeID() uint32 {
	return ContactsResolveUsernameRequestTypeID
}

// SchemaName returns MTProto type name.
func (r *ContactsResolveUsernameRequest) SchemaName() string {
	return "contacts.resolveUsername"
}

// Encode implements bin.Encoder.
func (r *ContactsResolveUsernameRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode contacts.resolveUsername#f93ccba3 as nil")
	}
	b.PutID(ContactsResolveUsernameRequestTypeID)
	b.PutString(r.Username)
	return nil
}

// GetUsername returns value of Username field.
func (r *ContactsResolveUsernameRequest) GetUsername() (value string) {
	return r.Username
}

// Decode implements bin.Decoder.
func (r *ContactsResolveUsernameRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode contacts.resolveUsername#f93ccba3 to nil")
	}
	if err := b.ConsumeID(ContactsResolveUsernameRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.resolveUsername#f93ccba3: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.resolveUsername#f93ccba3: field username: %w", err)
		}
		r.Username = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsResolveUsernameRequest.
var (
	_ bin.Encoder = &ContactsResolveUsernameRequest{}
	_ bin.Decoder = &ContactsResolveUsernameRequest{}
)

// ContactsResolveUsername invokes method contacts.resolveUsername#f93ccba3 returning error if any.
// Resolve a @username to get peer info
//
// Possible errors:
//  401 AUTH_KEY_PERM_EMPTY: The temporary auth key must be binded to the permanent auth key to use these methods.
//  400 CONNECTION_DEVICE_MODEL_EMPTY: Device model empty
//  400 CONNECTION_LAYER_INVALID: Layer invalid
//  400 USERNAME_INVALID: The provided username is not valid
//  400 USERNAME_NOT_OCCUPIED: The provided username is not occupied
//
// See https://core.telegram.org/method/contacts.resolveUsername for reference.
// Can be used by bots.
func (c *Client) ContactsResolveUsername(ctx context.Context, username string) (*ContactsResolvedPeer, error) {
	var result ContactsResolvedPeer

	request := &ContactsResolveUsernameRequest{
		Username: username,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
