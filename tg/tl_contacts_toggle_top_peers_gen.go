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

// ContactsToggleTopPeersRequest represents TL type `contacts.toggleTopPeers#8514bdda`.
// Enable/disable top peers¹
//
// Links:
//  1) https://core.telegram.org/api/top-rating
//
// See https://core.telegram.org/method/contacts.toggleTopPeers for reference.
type ContactsToggleTopPeersRequest struct {
	// Enable/disable
	Enabled bool `schemaname:"enabled"`
}

// ContactsToggleTopPeersRequestTypeID is TL type id of ContactsToggleTopPeersRequest.
const ContactsToggleTopPeersRequestTypeID = 0x8514bdda

func (t *ContactsToggleTopPeersRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Enabled == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ContactsToggleTopPeersRequest) String() string {
	if t == nil {
		return "ContactsToggleTopPeersRequest(nil)"
	}
	type Alias ContactsToggleTopPeersRequest
	return fmt.Sprintf("ContactsToggleTopPeersRequest%+v", Alias(*t))
}

// FillFrom fills ContactsToggleTopPeersRequest from given interface.
func (t *ContactsToggleTopPeersRequest) FillFrom(from interface {
	GetEnabled() (value bool)
}) {
	t.Enabled = from.GetEnabled()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (t *ContactsToggleTopPeersRequest) TypeID() uint32 {
	return ContactsToggleTopPeersRequestTypeID
}

// SchemaName returns MTProto type name.
func (t *ContactsToggleTopPeersRequest) SchemaName() string {
	return "contacts.toggleTopPeers"
}

// Encode implements bin.Encoder.
func (t *ContactsToggleTopPeersRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode contacts.toggleTopPeers#8514bdda as nil")
	}
	b.PutID(ContactsToggleTopPeersRequestTypeID)
	b.PutBool(t.Enabled)
	return nil
}

// GetEnabled returns value of Enabled field.
func (t *ContactsToggleTopPeersRequest) GetEnabled() (value bool) {
	return t.Enabled
}

// Decode implements bin.Decoder.
func (t *ContactsToggleTopPeersRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode contacts.toggleTopPeers#8514bdda to nil")
	}
	if err := b.ConsumeID(ContactsToggleTopPeersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.toggleTopPeers#8514bdda: %w", err)
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.toggleTopPeers#8514bdda: field enabled: %w", err)
		}
		t.Enabled = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsToggleTopPeersRequest.
var (
	_ bin.Encoder = &ContactsToggleTopPeersRequest{}
	_ bin.Decoder = &ContactsToggleTopPeersRequest{}
)

// ContactsToggleTopPeers invokes method contacts.toggleTopPeers#8514bdda returning error if any.
// Enable/disable top peers¹
//
// Links:
//  1) https://core.telegram.org/api/top-rating
//
// See https://core.telegram.org/method/contacts.toggleTopPeers for reference.
func (c *Client) ContactsToggleTopPeers(ctx context.Context, enabled bool) (bool, error) {
	var result BoolBox

	request := &ContactsToggleTopPeersRequest{
		Enabled: enabled,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
