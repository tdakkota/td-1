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

// PhoneSetCallRatingRequest represents TL type `phone.setCallRating#59ead627`.
// Rate a call
//
// See https://core.telegram.org/method/phone.setCallRating for reference.
type PhoneSetCallRatingRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `schemaname:"flags"`
	// Whether the user decided on their own initiative to rate the call
	UserInitiative bool `schemaname:"user_initiative"`
	// The call to rate
	Peer InputPhoneCall `schemaname:"peer"`
	// Rating in 1-5 stars
	Rating int `schemaname:"rating"`
	// An additional comment
	Comment string `schemaname:"comment"`
}

// PhoneSetCallRatingRequestTypeID is TL type id of PhoneSetCallRatingRequest.
const PhoneSetCallRatingRequestTypeID = 0x59ead627

func (s *PhoneSetCallRatingRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.UserInitiative == false) {
		return false
	}
	if !(s.Peer.Zero()) {
		return false
	}
	if !(s.Rating == 0) {
		return false
	}
	if !(s.Comment == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *PhoneSetCallRatingRequest) String() string {
	if s == nil {
		return "PhoneSetCallRatingRequest(nil)"
	}
	type Alias PhoneSetCallRatingRequest
	return fmt.Sprintf("PhoneSetCallRatingRequest%+v", Alias(*s))
}

// FillFrom fills PhoneSetCallRatingRequest from given interface.
func (s *PhoneSetCallRatingRequest) FillFrom(from interface {
	GetUserInitiative() (value bool)
	GetPeer() (value InputPhoneCall)
	GetRating() (value int)
	GetComment() (value string)
}) {
	s.UserInitiative = from.GetUserInitiative()
	s.Peer = from.GetPeer()
	s.Rating = from.GetRating()
	s.Comment = from.GetComment()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *PhoneSetCallRatingRequest) TypeID() uint32 {
	return PhoneSetCallRatingRequestTypeID
}

// SchemaName returns MTProto type name.
func (s *PhoneSetCallRatingRequest) SchemaName() string {
	return "phone.setCallRating"
}

// Encode implements bin.Encoder.
func (s *PhoneSetCallRatingRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode phone.setCallRating#59ead627 as nil")
	}
	b.PutID(PhoneSetCallRatingRequestTypeID)
	if !(s.UserInitiative == false) {
		s.Flags.Set(0)
	}
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.setCallRating#59ead627: field flags: %w", err)
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.setCallRating#59ead627: field peer: %w", err)
	}
	b.PutInt(s.Rating)
	b.PutString(s.Comment)
	return nil
}

// SetUserInitiative sets value of UserInitiative conditional field.
func (s *PhoneSetCallRatingRequest) SetUserInitiative(value bool) {
	if value {
		s.Flags.Set(0)
		s.UserInitiative = true
	} else {
		s.Flags.Unset(0)
		s.UserInitiative = false
	}
}

// GetUserInitiative returns value of UserInitiative conditional field.
func (s *PhoneSetCallRatingRequest) GetUserInitiative() (value bool) {
	return s.Flags.Has(0)
}

// GetPeer returns value of Peer field.
func (s *PhoneSetCallRatingRequest) GetPeer() (value InputPhoneCall) {
	return s.Peer
}

// GetRating returns value of Rating field.
func (s *PhoneSetCallRatingRequest) GetRating() (value int) {
	return s.Rating
}

// GetComment returns value of Comment field.
func (s *PhoneSetCallRatingRequest) GetComment() (value string) {
	return s.Comment
}

// Decode implements bin.Decoder.
func (s *PhoneSetCallRatingRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode phone.setCallRating#59ead627 to nil")
	}
	if err := b.ConsumeID(PhoneSetCallRatingRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.setCallRating#59ead627: %w", err)
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.setCallRating#59ead627: field flags: %w", err)
		}
	}
	s.UserInitiative = s.Flags.Has(0)
	{
		if err := s.Peer.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.setCallRating#59ead627: field peer: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phone.setCallRating#59ead627: field rating: %w", err)
		}
		s.Rating = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode phone.setCallRating#59ead627: field comment: %w", err)
		}
		s.Comment = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PhoneSetCallRatingRequest.
var (
	_ bin.Encoder = &PhoneSetCallRatingRequest{}
	_ bin.Decoder = &PhoneSetCallRatingRequest{}
)

// PhoneSetCallRating invokes method phone.setCallRating#59ead627 returning error if any.
// Rate a call
//
// Possible errors:
//  400 CALL_PEER_INVALID: The provided call peer object is invalid
//
// See https://core.telegram.org/method/phone.setCallRating for reference.
func (c *Client) PhoneSetCallRating(ctx context.Context, request *PhoneSetCallRatingRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
