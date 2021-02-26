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

// PhonePhoneCall represents TL type `phone.phoneCall#ec82e140`.
// A VoIP phone call
//
// See https://core.telegram.org/constructor/phone.phoneCall for reference.
type PhonePhoneCall struct {
	// The VoIP phone call
	PhoneCall PhoneCallClass `schemaname:"phone_call"`
	// VoIP phone call participants
	Users []UserClass `schemaname:"users"`
}

// PhonePhoneCallTypeID is TL type id of PhonePhoneCall.
const PhonePhoneCallTypeID = 0xec82e140

func (p *PhonePhoneCall) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.PhoneCall == nil) {
		return false
	}
	if !(p.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PhonePhoneCall) String() string {
	if p == nil {
		return "PhonePhoneCall(nil)"
	}
	type Alias PhonePhoneCall
	return fmt.Sprintf("PhonePhoneCall%+v", Alias(*p))
}

// FillFrom fills PhonePhoneCall from given interface.
func (p *PhonePhoneCall) FillFrom(from interface {
	GetPhoneCall() (value PhoneCallClass)
	GetUsers() (value []UserClass)
}) {
	p.PhoneCall = from.GetPhoneCall()
	p.Users = from.GetUsers()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *PhonePhoneCall) TypeID() uint32 {
	return PhonePhoneCallTypeID
}

// SchemaName returns MTProto type name.
func (p *PhonePhoneCall) SchemaName() string {
	return "phone.phoneCall"
}

// Encode implements bin.Encoder.
func (p *PhonePhoneCall) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode phone.phoneCall#ec82e140 as nil")
	}
	b.PutID(PhonePhoneCallTypeID)
	if p.PhoneCall == nil {
		return fmt.Errorf("unable to encode phone.phoneCall#ec82e140: field phone_call is nil")
	}
	if err := p.PhoneCall.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.phoneCall#ec82e140: field phone_call: %w", err)
	}
	b.PutVectorHeader(len(p.Users))
	for idx, v := range p.Users {
		if v == nil {
			return fmt.Errorf("unable to encode phone.phoneCall#ec82e140: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode phone.phoneCall#ec82e140: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetPhoneCall returns value of PhoneCall field.
func (p *PhonePhoneCall) GetPhoneCall() (value PhoneCallClass) {
	return p.PhoneCall
}

// GetPhoneCallAsNotEmpty returns mapped value of PhoneCall field.
func (p *PhonePhoneCall) GetPhoneCallAsNotEmpty() (NotEmptyPhoneCall, bool) {
	return p.PhoneCall.AsNotEmpty()
}

// GetUsers returns value of Users field.
func (p *PhonePhoneCall) GetUsers() (value []UserClass) {
	return p.Users
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (p *PhonePhoneCall) MapUsers() (value UserClassArray) {
	return UserClassArray(p.Users)
}

// Decode implements bin.Decoder.
func (p *PhonePhoneCall) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode phone.phoneCall#ec82e140 to nil")
	}
	if err := b.ConsumeID(PhonePhoneCallTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.phoneCall#ec82e140: %w", err)
	}
	{
		value, err := DecodePhoneCall(b)
		if err != nil {
			return fmt.Errorf("unable to decode phone.phoneCall#ec82e140: field phone_call: %w", err)
		}
		p.PhoneCall = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode phone.phoneCall#ec82e140: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode phone.phoneCall#ec82e140: field users: %w", err)
			}
			p.Users = append(p.Users, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for PhonePhoneCall.
var (
	_ bin.Encoder = &PhonePhoneCall{}
	_ bin.Decoder = &PhonePhoneCall{}
)
