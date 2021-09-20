// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
)

// InputUserEmpty represents TL type `inputUserEmpty#b98886cf`.
// Empty constructor, does not define a user.
//
// See https://core.telegram.org/constructor/inputUserEmpty for reference.
type InputUserEmpty struct {
}

// InputUserEmptyTypeID is TL type id of InputUserEmpty.
const InputUserEmptyTypeID = 0xb98886cf

// construct implements constructor of InputUserClass.
func (i InputUserEmpty) construct() InputUserClass { return &i }

// Ensuring interfaces in compile-time for InputUserEmpty.
var (
	_ bin.Encoder     = &InputUserEmpty{}
	_ bin.Decoder     = &InputUserEmpty{}
	_ bin.BareEncoder = &InputUserEmpty{}
	_ bin.BareDecoder = &InputUserEmpty{}

	_ InputUserClass = &InputUserEmpty{}
)

func (i *InputUserEmpty) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputUserEmpty) String() string {
	if i == nil {
		return "InputUserEmpty(nil)"
	}
	type Alias InputUserEmpty
	return fmt.Sprintf("InputUserEmpty%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputUserEmpty) TypeID() uint32 {
	return InputUserEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*InputUserEmpty) TypeName() string {
	return "inputUserEmpty"
}

// TypeInfo returns info about TL type.
func (i *InputUserEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputUserEmpty",
		ID:   InputUserEmptyTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputUserEmpty) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputUserEmpty#b98886cf as nil")
	}
	b.PutID(InputUserEmptyTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputUserEmpty) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputUserEmpty#b98886cf as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputUserEmpty) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputUserEmpty#b98886cf to nil")
	}
	if err := b.ConsumeID(InputUserEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode inputUserEmpty#b98886cf: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputUserEmpty) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputUserEmpty#b98886cf to nil")
	}
	return nil
}

// InputUserSelf represents TL type `inputUserSelf#f7c1b13f`.
// Defines the current user.
//
// See https://core.telegram.org/constructor/inputUserSelf for reference.
type InputUserSelf struct {
}

// InputUserSelfTypeID is TL type id of InputUserSelf.
const InputUserSelfTypeID = 0xf7c1b13f

// construct implements constructor of InputUserClass.
func (i InputUserSelf) construct() InputUserClass { return &i }

// Ensuring interfaces in compile-time for InputUserSelf.
var (
	_ bin.Encoder     = &InputUserSelf{}
	_ bin.Decoder     = &InputUserSelf{}
	_ bin.BareEncoder = &InputUserSelf{}
	_ bin.BareDecoder = &InputUserSelf{}

	_ InputUserClass = &InputUserSelf{}
)

func (i *InputUserSelf) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputUserSelf) String() string {
	if i == nil {
		return "InputUserSelf(nil)"
	}
	type Alias InputUserSelf
	return fmt.Sprintf("InputUserSelf%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputUserSelf) TypeID() uint32 {
	return InputUserSelfTypeID
}

// TypeName returns name of type in TL schema.
func (*InputUserSelf) TypeName() string {
	return "inputUserSelf"
}

// TypeInfo returns info about TL type.
func (i *InputUserSelf) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputUserSelf",
		ID:   InputUserSelfTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputUserSelf) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputUserSelf#f7c1b13f as nil")
	}
	b.PutID(InputUserSelfTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputUserSelf) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputUserSelf#f7c1b13f as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputUserSelf) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputUserSelf#f7c1b13f to nil")
	}
	if err := b.ConsumeID(InputUserSelfTypeID); err != nil {
		return fmt.Errorf("unable to decode inputUserSelf#f7c1b13f: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputUserSelf) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputUserSelf#f7c1b13f to nil")
	}
	return nil
}

// InputUser represents TL type `inputUser#f21158c6`.
// Defines a user for further interaction.
//
// See https://core.telegram.org/constructor/inputUser for reference.
type InputUser struct {
	// User identifier
	UserID int64
	// access_hash value from the user¹ constructor
	//
	// Links:
	//  1) https://core.telegram.org/constructor/user
	AccessHash int64
}

// InputUserTypeID is TL type id of InputUser.
const InputUserTypeID = 0xf21158c6

// construct implements constructor of InputUserClass.
func (i InputUser) construct() InputUserClass { return &i }

// Ensuring interfaces in compile-time for InputUser.
var (
	_ bin.Encoder     = &InputUser{}
	_ bin.Decoder     = &InputUser{}
	_ bin.BareEncoder = &InputUser{}
	_ bin.BareDecoder = &InputUser{}

	_ InputUserClass = &InputUser{}
)

func (i *InputUser) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.UserID == 0) {
		return false
	}
	if !(i.AccessHash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputUser) String() string {
	if i == nil {
		return "InputUser(nil)"
	}
	type Alias InputUser
	return fmt.Sprintf("InputUser%+v", Alias(*i))
}

// FillFrom fills InputUser from given interface.
func (i *InputUser) FillFrom(from interface {
	GetUserID() (value int64)
	GetAccessHash() (value int64)
}) {
	i.UserID = from.GetUserID()
	i.AccessHash = from.GetAccessHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputUser) TypeID() uint32 {
	return InputUserTypeID
}

// TypeName returns name of type in TL schema.
func (*InputUser) TypeName() string {
	return "inputUser"
}

// TypeInfo returns info about TL type.
func (i *InputUser) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputUser",
		ID:   InputUserTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "AccessHash",
			SchemaName: "access_hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputUser) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputUser#f21158c6 as nil")
	}
	b.PutID(InputUserTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputUser) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputUser#f21158c6 as nil")
	}
	b.PutLong(i.UserID)
	b.PutLong(i.AccessHash)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputUser) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputUser#f21158c6 to nil")
	}
	if err := b.ConsumeID(InputUserTypeID); err != nil {
		return fmt.Errorf("unable to decode inputUser#f21158c6: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputUser) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputUser#f21158c6 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputUser#f21158c6: field user_id: %w", err)
		}
		i.UserID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputUser#f21158c6: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// GetUserID returns value of UserID field.
func (i *InputUser) GetUserID() (value int64) {
	return i.UserID
}

// GetAccessHash returns value of AccessHash field.
func (i *InputUser) GetAccessHash() (value int64) {
	return i.AccessHash
}

// InputUserFromMessage represents TL type `inputUserFromMessage#1da448e2`.
// Defines a min¹ user that was seen in a certain message of a certain chat.
//
// Links:
//  1) https://core.telegram.org/api/min
//
// See https://core.telegram.org/constructor/inputUserFromMessage for reference.
type InputUserFromMessage struct {
	// The chat where the user was seen
	Peer InputPeerClass
	// The message ID
	MsgID int
	// The identifier of the user that was seen
	UserID int64
}

// InputUserFromMessageTypeID is TL type id of InputUserFromMessage.
const InputUserFromMessageTypeID = 0x1da448e2

// construct implements constructor of InputUserClass.
func (i InputUserFromMessage) construct() InputUserClass { return &i }

// Ensuring interfaces in compile-time for InputUserFromMessage.
var (
	_ bin.Encoder     = &InputUserFromMessage{}
	_ bin.Decoder     = &InputUserFromMessage{}
	_ bin.BareEncoder = &InputUserFromMessage{}
	_ bin.BareDecoder = &InputUserFromMessage{}

	_ InputUserClass = &InputUserFromMessage{}
)

func (i *InputUserFromMessage) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Peer == nil) {
		return false
	}
	if !(i.MsgID == 0) {
		return false
	}
	if !(i.UserID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputUserFromMessage) String() string {
	if i == nil {
		return "InputUserFromMessage(nil)"
	}
	type Alias InputUserFromMessage
	return fmt.Sprintf("InputUserFromMessage%+v", Alias(*i))
}

// FillFrom fills InputUserFromMessage from given interface.
func (i *InputUserFromMessage) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetMsgID() (value int)
	GetUserID() (value int64)
}) {
	i.Peer = from.GetPeer()
	i.MsgID = from.GetMsgID()
	i.UserID = from.GetUserID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputUserFromMessage) TypeID() uint32 {
	return InputUserFromMessageTypeID
}

// TypeName returns name of type in TL schema.
func (*InputUserFromMessage) TypeName() string {
	return "inputUserFromMessage"
}

// TypeInfo returns info about TL type.
func (i *InputUserFromMessage) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputUserFromMessage",
		ID:   InputUserFromMessageTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "MsgID",
			SchemaName: "msg_id",
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputUserFromMessage) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputUserFromMessage#1da448e2 as nil")
	}
	b.PutID(InputUserFromMessageTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputUserFromMessage) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputUserFromMessage#1da448e2 as nil")
	}
	if i.Peer == nil {
		return fmt.Errorf("unable to encode inputUserFromMessage#1da448e2: field peer is nil")
	}
	if err := i.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputUserFromMessage#1da448e2: field peer: %w", err)
	}
	b.PutInt(i.MsgID)
	b.PutLong(i.UserID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputUserFromMessage) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputUserFromMessage#1da448e2 to nil")
	}
	if err := b.ConsumeID(InputUserFromMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode inputUserFromMessage#1da448e2: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputUserFromMessage) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputUserFromMessage#1da448e2 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputUserFromMessage#1da448e2: field peer: %w", err)
		}
		i.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputUserFromMessage#1da448e2: field msg_id: %w", err)
		}
		i.MsgID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputUserFromMessage#1da448e2: field user_id: %w", err)
		}
		i.UserID = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (i *InputUserFromMessage) GetPeer() (value InputPeerClass) {
	return i.Peer
}

// GetMsgID returns value of MsgID field.
func (i *InputUserFromMessage) GetMsgID() (value int) {
	return i.MsgID
}

// GetUserID returns value of UserID field.
func (i *InputUserFromMessage) GetUserID() (value int64) {
	return i.UserID
}

// InputUserClass represents InputUser generic type.
//
// See https://core.telegram.org/type/InputUser for reference.
//
// Example:
//  g, err := tg.DecodeInputUser(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.InputUserEmpty: // inputUserEmpty#b98886cf
//  case *tg.InputUserSelf: // inputUserSelf#f7c1b13f
//  case *tg.InputUser: // inputUser#f21158c6
//  case *tg.InputUserFromMessage: // inputUserFromMessage#1da448e2
//  default: panic(v)
//  }
type InputUserClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputUserClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeInputUser implements binary de-serialization for InputUserClass.
func DecodeInputUser(buf *bin.Buffer) (InputUserClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputUserEmptyTypeID:
		// Decoding inputUserEmpty#b98886cf.
		v := InputUserEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputUserClass: %w", err)
		}
		return &v, nil
	case InputUserSelfTypeID:
		// Decoding inputUserSelf#f7c1b13f.
		v := InputUserSelf{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputUserClass: %w", err)
		}
		return &v, nil
	case InputUserTypeID:
		// Decoding inputUser#f21158c6.
		v := InputUser{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputUserClass: %w", err)
		}
		return &v, nil
	case InputUserFromMessageTypeID:
		// Decoding inputUserFromMessage#1da448e2.
		v := InputUserFromMessage{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputUserClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputUserClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputUser boxes the InputUserClass providing a helper.
type InputUserBox struct {
	InputUser InputUserClass
}

// Decode implements bin.Decoder for InputUserBox.
func (b *InputUserBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputUserBox to nil")
	}
	v, err := DecodeInputUser(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputUser = v
	return nil
}

// Encode implements bin.Encode for InputUserBox.
func (b *InputUserBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputUser == nil {
		return fmt.Errorf("unable to encode InputUserClass as nil")
	}
	return b.InputUser.Encode(buf)
}
