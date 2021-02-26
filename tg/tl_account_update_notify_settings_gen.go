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

// AccountUpdateNotifySettingsRequest represents TL type `account.updateNotifySettings#84be5b93`.
// Edits notification settings from a given user/group, from all users/all groups.
//
// See https://core.telegram.org/method/account.updateNotifySettings for reference.
type AccountUpdateNotifySettingsRequest struct {
	// Notification source
	Peer InputNotifyPeerClass `schemaname:"peer"`
	// Notification settings
	Settings InputPeerNotifySettings `schemaname:"settings"`
}

// AccountUpdateNotifySettingsRequestTypeID is TL type id of AccountUpdateNotifySettingsRequest.
const AccountUpdateNotifySettingsRequestTypeID = 0x84be5b93

func (u *AccountUpdateNotifySettingsRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Peer == nil) {
		return false
	}
	if !(u.Settings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *AccountUpdateNotifySettingsRequest) String() string {
	if u == nil {
		return "AccountUpdateNotifySettingsRequest(nil)"
	}
	type Alias AccountUpdateNotifySettingsRequest
	return fmt.Sprintf("AccountUpdateNotifySettingsRequest%+v", Alias(*u))
}

// FillFrom fills AccountUpdateNotifySettingsRequest from given interface.
func (u *AccountUpdateNotifySettingsRequest) FillFrom(from interface {
	GetPeer() (value InputNotifyPeerClass)
	GetSettings() (value InputPeerNotifySettings)
}) {
	u.Peer = from.GetPeer()
	u.Settings = from.GetSettings()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (u *AccountUpdateNotifySettingsRequest) TypeID() uint32 {
	return AccountUpdateNotifySettingsRequestTypeID
}

// SchemaName returns MTProto type name.
func (u *AccountUpdateNotifySettingsRequest) SchemaName() string {
	return "account.updateNotifySettings"
}

// Encode implements bin.Encoder.
func (u *AccountUpdateNotifySettingsRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updateNotifySettings#84be5b93 as nil")
	}
	b.PutID(AccountUpdateNotifySettingsRequestTypeID)
	if u.Peer == nil {
		return fmt.Errorf("unable to encode account.updateNotifySettings#84be5b93: field peer is nil")
	}
	if err := u.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.updateNotifySettings#84be5b93: field peer: %w", err)
	}
	if err := u.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.updateNotifySettings#84be5b93: field settings: %w", err)
	}
	return nil
}

// GetPeer returns value of Peer field.
func (u *AccountUpdateNotifySettingsRequest) GetPeer() (value InputNotifyPeerClass) {
	return u.Peer
}

// GetSettings returns value of Settings field.
func (u *AccountUpdateNotifySettingsRequest) GetSettings() (value InputPeerNotifySettings) {
	return u.Settings
}

// Decode implements bin.Decoder.
func (u *AccountUpdateNotifySettingsRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updateNotifySettings#84be5b93 to nil")
	}
	if err := b.ConsumeID(AccountUpdateNotifySettingsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.updateNotifySettings#84be5b93: %w", err)
	}
	{
		value, err := DecodeInputNotifyPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.updateNotifySettings#84be5b93: field peer: %w", err)
		}
		u.Peer = value
	}
	{
		if err := u.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.updateNotifySettings#84be5b93: field settings: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountUpdateNotifySettingsRequest.
var (
	_ bin.Encoder = &AccountUpdateNotifySettingsRequest{}
	_ bin.Decoder = &AccountUpdateNotifySettingsRequest{}
)

// AccountUpdateNotifySettings invokes method account.updateNotifySettings#84be5b93 returning error if any.
// Edits notification settings from a given user/group, from all users/all groups.
//
// Possible errors:
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 MSG_ID_INVALID: Invalid message ID provided
//  400 PEER_ID_INVALID: The provided peer id is invalid
//  400 SETTINGS_INVALID: Invalid settings were provided
//
// See https://core.telegram.org/method/account.updateNotifySettings for reference.
func (c *Client) AccountUpdateNotifySettings(ctx context.Context, request *AccountUpdateNotifySettingsRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
