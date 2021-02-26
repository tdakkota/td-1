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

// PhoneDiscardGroupCallRequest represents TL type `phone.discardGroupCall#7a777135`.
//
// See https://core.telegram.org/method/phone.discardGroupCall for reference.
type PhoneDiscardGroupCallRequest struct {
	// Call field of PhoneDiscardGroupCallRequest.
	Call InputGroupCall `schemaname:"call"`
}

// PhoneDiscardGroupCallRequestTypeID is TL type id of PhoneDiscardGroupCallRequest.
const PhoneDiscardGroupCallRequestTypeID = 0x7a777135

func (d *PhoneDiscardGroupCallRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Call.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *PhoneDiscardGroupCallRequest) String() string {
	if d == nil {
		return "PhoneDiscardGroupCallRequest(nil)"
	}
	type Alias PhoneDiscardGroupCallRequest
	return fmt.Sprintf("PhoneDiscardGroupCallRequest%+v", Alias(*d))
}

// FillFrom fills PhoneDiscardGroupCallRequest from given interface.
func (d *PhoneDiscardGroupCallRequest) FillFrom(from interface {
	GetCall() (value InputGroupCall)
}) {
	d.Call = from.GetCall()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (d *PhoneDiscardGroupCallRequest) TypeID() uint32 {
	return PhoneDiscardGroupCallRequestTypeID
}

// SchemaName returns MTProto type name.
func (d *PhoneDiscardGroupCallRequest) SchemaName() string {
	return "phone.discardGroupCall"
}

// Encode implements bin.Encoder.
func (d *PhoneDiscardGroupCallRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode phone.discardGroupCall#7a777135 as nil")
	}
	b.PutID(PhoneDiscardGroupCallRequestTypeID)
	if err := d.Call.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.discardGroupCall#7a777135: field call: %w", err)
	}
	return nil
}

// GetCall returns value of Call field.
func (d *PhoneDiscardGroupCallRequest) GetCall() (value InputGroupCall) {
	return d.Call
}

// Decode implements bin.Decoder.
func (d *PhoneDiscardGroupCallRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode phone.discardGroupCall#7a777135 to nil")
	}
	if err := b.ConsumeID(PhoneDiscardGroupCallRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.discardGroupCall#7a777135: %w", err)
	}
	{
		if err := d.Call.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.discardGroupCall#7a777135: field call: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for PhoneDiscardGroupCallRequest.
var (
	_ bin.Encoder = &PhoneDiscardGroupCallRequest{}
	_ bin.Decoder = &PhoneDiscardGroupCallRequest{}
)

// PhoneDiscardGroupCall invokes method phone.discardGroupCall#7a777135 returning error if any.
//
// See https://core.telegram.org/method/phone.discardGroupCall for reference.
func (c *Client) PhoneDiscardGroupCall(ctx context.Context, call InputGroupCall) (UpdatesClass, error) {
	var result UpdatesBox

	request := &PhoneDiscardGroupCallRequest{
		Call: call,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
