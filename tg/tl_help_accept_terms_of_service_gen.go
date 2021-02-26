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

// HelpAcceptTermsOfServiceRequest represents TL type `help.acceptTermsOfService#ee72f79a`.
// Accept the new terms of service
//
// See https://core.telegram.org/method/help.acceptTermsOfService for reference.
type HelpAcceptTermsOfServiceRequest struct {
	// ID of terms of service
	ID DataJSON `schemaname:"id"`
}

// HelpAcceptTermsOfServiceRequestTypeID is TL type id of HelpAcceptTermsOfServiceRequest.
const HelpAcceptTermsOfServiceRequestTypeID = 0xee72f79a

func (a *HelpAcceptTermsOfServiceRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.ID.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *HelpAcceptTermsOfServiceRequest) String() string {
	if a == nil {
		return "HelpAcceptTermsOfServiceRequest(nil)"
	}
	type Alias HelpAcceptTermsOfServiceRequest
	return fmt.Sprintf("HelpAcceptTermsOfServiceRequest%+v", Alias(*a))
}

// FillFrom fills HelpAcceptTermsOfServiceRequest from given interface.
func (a *HelpAcceptTermsOfServiceRequest) FillFrom(from interface {
	GetID() (value DataJSON)
}) {
	a.ID = from.GetID()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (a *HelpAcceptTermsOfServiceRequest) TypeID() uint32 {
	return HelpAcceptTermsOfServiceRequestTypeID
}

// SchemaName returns MTProto type name.
func (a *HelpAcceptTermsOfServiceRequest) SchemaName() string {
	return "help.acceptTermsOfService"
}

// Encode implements bin.Encoder.
func (a *HelpAcceptTermsOfServiceRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode help.acceptTermsOfService#ee72f79a as nil")
	}
	b.PutID(HelpAcceptTermsOfServiceRequestTypeID)
	if err := a.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode help.acceptTermsOfService#ee72f79a: field id: %w", err)
	}
	return nil
}

// GetID returns value of ID field.
func (a *HelpAcceptTermsOfServiceRequest) GetID() (value DataJSON) {
	return a.ID
}

// Decode implements bin.Decoder.
func (a *HelpAcceptTermsOfServiceRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode help.acceptTermsOfService#ee72f79a to nil")
	}
	if err := b.ConsumeID(HelpAcceptTermsOfServiceRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.acceptTermsOfService#ee72f79a: %w", err)
	}
	{
		if err := a.ID.Decode(b); err != nil {
			return fmt.Errorf("unable to decode help.acceptTermsOfService#ee72f79a: field id: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpAcceptTermsOfServiceRequest.
var (
	_ bin.Encoder = &HelpAcceptTermsOfServiceRequest{}
	_ bin.Decoder = &HelpAcceptTermsOfServiceRequest{}
)

// HelpAcceptTermsOfService invokes method help.acceptTermsOfService#ee72f79a returning error if any.
// Accept the new terms of service
//
// See https://core.telegram.org/method/help.acceptTermsOfService for reference.
func (c *Client) HelpAcceptTermsOfService(ctx context.Context, id DataJSON) (bool, error) {
	var result BoolBox

	request := &HelpAcceptTermsOfServiceRequest{
		ID: id,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
