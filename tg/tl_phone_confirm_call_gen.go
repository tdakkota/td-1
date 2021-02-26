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

// PhoneConfirmCallRequest represents TL type `phone.confirmCall#2efe1722`.
// Complete phone call E2E encryption key exchange »¹
//
// Links:
//  1) https://core.telegram.org/api/end-to-end/voice-calls
//
// See https://core.telegram.org/method/phone.confirmCall for reference.
type PhoneConfirmCallRequest struct {
	// The phone call
	Peer InputPhoneCall `schemaname:"peer"`
	// Parameter for E2E encryption key exchange »¹
	//
	// Links:
	//  1) https://core.telegram.org/api/end-to-end/voice-calls
	GA []byte `schemaname:"g_a"`
	// Key fingerprint
	KeyFingerprint int64 `schemaname:"key_fingerprint"`
	// Phone call settings
	Protocol PhoneCallProtocol `schemaname:"protocol"`
}

// PhoneConfirmCallRequestTypeID is TL type id of PhoneConfirmCallRequest.
const PhoneConfirmCallRequestTypeID = 0x2efe1722

func (c *PhoneConfirmCallRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Peer.Zero()) {
		return false
	}
	if !(c.GA == nil) {
		return false
	}
	if !(c.KeyFingerprint == 0) {
		return false
	}
	if !(c.Protocol.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *PhoneConfirmCallRequest) String() string {
	if c == nil {
		return "PhoneConfirmCallRequest(nil)"
	}
	type Alias PhoneConfirmCallRequest
	return fmt.Sprintf("PhoneConfirmCallRequest%+v", Alias(*c))
}

// FillFrom fills PhoneConfirmCallRequest from given interface.
func (c *PhoneConfirmCallRequest) FillFrom(from interface {
	GetPeer() (value InputPhoneCall)
	GetGA() (value []byte)
	GetKeyFingerprint() (value int64)
	GetProtocol() (value PhoneCallProtocol)
}) {
	c.Peer = from.GetPeer()
	c.GA = from.GetGA()
	c.KeyFingerprint = from.GetKeyFingerprint()
	c.Protocol = from.GetProtocol()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *PhoneConfirmCallRequest) TypeID() uint32 {
	return PhoneConfirmCallRequestTypeID
}

// SchemaName returns MTProto type name.
func (c *PhoneConfirmCallRequest) SchemaName() string {
	return "phone.confirmCall"
}

// Encode implements bin.Encoder.
func (c *PhoneConfirmCallRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode phone.confirmCall#2efe1722 as nil")
	}
	b.PutID(PhoneConfirmCallRequestTypeID)
	if err := c.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.confirmCall#2efe1722: field peer: %w", err)
	}
	b.PutBytes(c.GA)
	b.PutLong(c.KeyFingerprint)
	if err := c.Protocol.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.confirmCall#2efe1722: field protocol: %w", err)
	}
	return nil
}

// GetPeer returns value of Peer field.
func (c *PhoneConfirmCallRequest) GetPeer() (value InputPhoneCall) {
	return c.Peer
}

// GetGA returns value of GA field.
func (c *PhoneConfirmCallRequest) GetGA() (value []byte) {
	return c.GA
}

// GetKeyFingerprint returns value of KeyFingerprint field.
func (c *PhoneConfirmCallRequest) GetKeyFingerprint() (value int64) {
	return c.KeyFingerprint
}

// GetProtocol returns value of Protocol field.
func (c *PhoneConfirmCallRequest) GetProtocol() (value PhoneCallProtocol) {
	return c.Protocol
}

// Decode implements bin.Decoder.
func (c *PhoneConfirmCallRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode phone.confirmCall#2efe1722 to nil")
	}
	if err := b.ConsumeID(PhoneConfirmCallRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.confirmCall#2efe1722: %w", err)
	}
	{
		if err := c.Peer.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.confirmCall#2efe1722: field peer: %w", err)
		}
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode phone.confirmCall#2efe1722: field g_a: %w", err)
		}
		c.GA = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode phone.confirmCall#2efe1722: field key_fingerprint: %w", err)
		}
		c.KeyFingerprint = value
	}
	{
		if err := c.Protocol.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.confirmCall#2efe1722: field protocol: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for PhoneConfirmCallRequest.
var (
	_ bin.Encoder = &PhoneConfirmCallRequest{}
	_ bin.Decoder = &PhoneConfirmCallRequest{}
)

// PhoneConfirmCall invokes method phone.confirmCall#2efe1722 returning error if any.
// Complete phone call E2E encryption key exchange »¹
//
// Links:
//  1) https://core.telegram.org/api/end-to-end/voice-calls
//
// Possible errors:
//  400 CALL_ALREADY_DECLINED: The call was already declined
//  400 CALL_PEER_INVALID: The provided call peer object is invalid
//
// See https://core.telegram.org/method/phone.confirmCall for reference.
func (c *Client) PhoneConfirmCall(ctx context.Context, request *PhoneConfirmCallRequest) (*PhonePhoneCall, error) {
	var result PhonePhoneCall

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
