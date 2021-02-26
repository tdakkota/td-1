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

// UpdatesDifferenceEmpty represents TL type `updates.differenceEmpty#5d75a138`.
// No events.
//
// See https://core.telegram.org/constructor/updates.differenceEmpty for reference.
type UpdatesDifferenceEmpty struct {
	// Current date
	Date int `schemaname:"date"`
	// Number of sent updates
	Seq int `schemaname:"seq"`
}

// UpdatesDifferenceEmptyTypeID is TL type id of UpdatesDifferenceEmpty.
const UpdatesDifferenceEmptyTypeID = 0x5d75a138

func (d *UpdatesDifferenceEmpty) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Date == 0) {
		return false
	}
	if !(d.Seq == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *UpdatesDifferenceEmpty) String() string {
	if d == nil {
		return "UpdatesDifferenceEmpty(nil)"
	}
	type Alias UpdatesDifferenceEmpty
	return fmt.Sprintf("UpdatesDifferenceEmpty%+v", Alias(*d))
}

// FillFrom fills UpdatesDifferenceEmpty from given interface.
func (d *UpdatesDifferenceEmpty) FillFrom(from interface {
	GetDate() (value int)
	GetSeq() (value int)
}) {
	d.Date = from.GetDate()
	d.Seq = from.GetSeq()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (d *UpdatesDifferenceEmpty) TypeID() uint32 {
	return UpdatesDifferenceEmptyTypeID
}

// SchemaName returns MTProto type name.
func (d *UpdatesDifferenceEmpty) SchemaName() string {
	return "updates.differenceEmpty"
}

// Encode implements bin.Encoder.
func (d *UpdatesDifferenceEmpty) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode updates.differenceEmpty#5d75a138 as nil")
	}
	b.PutID(UpdatesDifferenceEmptyTypeID)
	b.PutInt(d.Date)
	b.PutInt(d.Seq)
	return nil
}

// GetDate returns value of Date field.
func (d *UpdatesDifferenceEmpty) GetDate() (value int) {
	return d.Date
}

// GetSeq returns value of Seq field.
func (d *UpdatesDifferenceEmpty) GetSeq() (value int) {
	return d.Seq
}

// Decode implements bin.Decoder.
func (d *UpdatesDifferenceEmpty) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode updates.differenceEmpty#5d75a138 to nil")
	}
	if err := b.ConsumeID(UpdatesDifferenceEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode updates.differenceEmpty#5d75a138: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.differenceEmpty#5d75a138: field date: %w", err)
		}
		d.Date = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.differenceEmpty#5d75a138: field seq: %w", err)
		}
		d.Seq = value
	}
	return nil
}

// construct implements constructor of UpdatesDifferenceClass.
func (d UpdatesDifferenceEmpty) construct() UpdatesDifferenceClass { return &d }

// Ensuring interfaces in compile-time for UpdatesDifferenceEmpty.
var (
	_ bin.Encoder = &UpdatesDifferenceEmpty{}
	_ bin.Decoder = &UpdatesDifferenceEmpty{}

	_ UpdatesDifferenceClass = &UpdatesDifferenceEmpty{}
)

// UpdatesDifference represents TL type `updates.difference#f49ca0`.
// Full list of occurred events.
//
// See https://core.telegram.org/constructor/updates.difference for reference.
type UpdatesDifference struct {
	// List of new messages
	NewMessages []MessageClass `schemaname:"new_messages"`
	// List of new encrypted secret chat messages
	NewEncryptedMessages []EncryptedMessageClass `schemaname:"new_encrypted_messages"`
	// List of updates
	OtherUpdates []UpdateClass `schemaname:"other_updates"`
	// List of chats mentioned in events
	Chats []ChatClass `schemaname:"chats"`
	// List of users mentioned in events
	Users []UserClass `schemaname:"users"`
	// Current state
	State UpdatesState `schemaname:"state"`
}

// UpdatesDifferenceTypeID is TL type id of UpdatesDifference.
const UpdatesDifferenceTypeID = 0xf49ca0

func (d *UpdatesDifference) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.NewMessages == nil) {
		return false
	}
	if !(d.NewEncryptedMessages == nil) {
		return false
	}
	if !(d.OtherUpdates == nil) {
		return false
	}
	if !(d.Chats == nil) {
		return false
	}
	if !(d.Users == nil) {
		return false
	}
	if !(d.State.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *UpdatesDifference) String() string {
	if d == nil {
		return "UpdatesDifference(nil)"
	}
	type Alias UpdatesDifference
	return fmt.Sprintf("UpdatesDifference%+v", Alias(*d))
}

// FillFrom fills UpdatesDifference from given interface.
func (d *UpdatesDifference) FillFrom(from interface {
	GetNewMessages() (value []MessageClass)
	GetNewEncryptedMessages() (value []EncryptedMessageClass)
	GetOtherUpdates() (value []UpdateClass)
	GetChats() (value []ChatClass)
	GetUsers() (value []UserClass)
	GetState() (value UpdatesState)
}) {
	d.NewMessages = from.GetNewMessages()
	d.NewEncryptedMessages = from.GetNewEncryptedMessages()
	d.OtherUpdates = from.GetOtherUpdates()
	d.Chats = from.GetChats()
	d.Users = from.GetUsers()
	d.State = from.GetState()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (d *UpdatesDifference) TypeID() uint32 {
	return UpdatesDifferenceTypeID
}

// SchemaName returns MTProto type name.
func (d *UpdatesDifference) SchemaName() string {
	return "updates.difference"
}

// Encode implements bin.Encoder.
func (d *UpdatesDifference) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode updates.difference#f49ca0 as nil")
	}
	b.PutID(UpdatesDifferenceTypeID)
	b.PutVectorHeader(len(d.NewMessages))
	for idx, v := range d.NewMessages {
		if v == nil {
			return fmt.Errorf("unable to encode updates.difference#f49ca0: field new_messages element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.difference#f49ca0: field new_messages element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(d.NewEncryptedMessages))
	for idx, v := range d.NewEncryptedMessages {
		if v == nil {
			return fmt.Errorf("unable to encode updates.difference#f49ca0: field new_encrypted_messages element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.difference#f49ca0: field new_encrypted_messages element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(d.OtherUpdates))
	for idx, v := range d.OtherUpdates {
		if v == nil {
			return fmt.Errorf("unable to encode updates.difference#f49ca0: field other_updates element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.difference#f49ca0: field other_updates element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(d.Chats))
	for idx, v := range d.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode updates.difference#f49ca0: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.difference#f49ca0: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(d.Users))
	for idx, v := range d.Users {
		if v == nil {
			return fmt.Errorf("unable to encode updates.difference#f49ca0: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.difference#f49ca0: field users element with index %d: %w", idx, err)
		}
	}
	if err := d.State.Encode(b); err != nil {
		return fmt.Errorf("unable to encode updates.difference#f49ca0: field state: %w", err)
	}
	return nil
}

// GetNewMessages returns value of NewMessages field.
func (d *UpdatesDifference) GetNewMessages() (value []MessageClass) {
	return d.NewMessages
}

// MapNewMessages returns field NewMessages wrapped in MessageClassArray helper.
func (d *UpdatesDifference) MapNewMessages() (value MessageClassArray) {
	return MessageClassArray(d.NewMessages)
}

// GetNewEncryptedMessages returns value of NewEncryptedMessages field.
func (d *UpdatesDifference) GetNewEncryptedMessages() (value []EncryptedMessageClass) {
	return d.NewEncryptedMessages
}

// MapNewEncryptedMessages returns field NewEncryptedMessages wrapped in EncryptedMessageClassArray helper.
func (d *UpdatesDifference) MapNewEncryptedMessages() (value EncryptedMessageClassArray) {
	return EncryptedMessageClassArray(d.NewEncryptedMessages)
}

// GetOtherUpdates returns value of OtherUpdates field.
func (d *UpdatesDifference) GetOtherUpdates() (value []UpdateClass) {
	return d.OtherUpdates
}

// MapOtherUpdates returns field OtherUpdates wrapped in UpdateClassArray helper.
func (d *UpdatesDifference) MapOtherUpdates() (value UpdateClassArray) {
	return UpdateClassArray(d.OtherUpdates)
}

// GetChats returns value of Chats field.
func (d *UpdatesDifference) GetChats() (value []ChatClass) {
	return d.Chats
}

// MapChats returns field Chats wrapped in ChatClassArray helper.
func (d *UpdatesDifference) MapChats() (value ChatClassArray) {
	return ChatClassArray(d.Chats)
}

// GetUsers returns value of Users field.
func (d *UpdatesDifference) GetUsers() (value []UserClass) {
	return d.Users
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (d *UpdatesDifference) MapUsers() (value UserClassArray) {
	return UserClassArray(d.Users)
}

// GetState returns value of State field.
func (d *UpdatesDifference) GetState() (value UpdatesState) {
	return d.State
}

// Decode implements bin.Decoder.
func (d *UpdatesDifference) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode updates.difference#f49ca0 to nil")
	}
	if err := b.ConsumeID(UpdatesDifferenceTypeID); err != nil {
		return fmt.Errorf("unable to decode updates.difference#f49ca0: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.difference#f49ca0: field new_messages: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessage(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.difference#f49ca0: field new_messages: %w", err)
			}
			d.NewMessages = append(d.NewMessages, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.difference#f49ca0: field new_encrypted_messages: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeEncryptedMessage(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.difference#f49ca0: field new_encrypted_messages: %w", err)
			}
			d.NewEncryptedMessages = append(d.NewEncryptedMessages, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.difference#f49ca0: field other_updates: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUpdate(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.difference#f49ca0: field other_updates: %w", err)
			}
			d.OtherUpdates = append(d.OtherUpdates, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.difference#f49ca0: field chats: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.difference#f49ca0: field chats: %w", err)
			}
			d.Chats = append(d.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.difference#f49ca0: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.difference#f49ca0: field users: %w", err)
			}
			d.Users = append(d.Users, value)
		}
	}
	{
		if err := d.State.Decode(b); err != nil {
			return fmt.Errorf("unable to decode updates.difference#f49ca0: field state: %w", err)
		}
	}
	return nil
}

// construct implements constructor of UpdatesDifferenceClass.
func (d UpdatesDifference) construct() UpdatesDifferenceClass { return &d }

// Ensuring interfaces in compile-time for UpdatesDifference.
var (
	_ bin.Encoder = &UpdatesDifference{}
	_ bin.Decoder = &UpdatesDifference{}

	_ UpdatesDifferenceClass = &UpdatesDifference{}
)

// UpdatesDifferenceSlice represents TL type `updates.differenceSlice#a8fb1981`.
// Incomplete list of occurred events.
//
// See https://core.telegram.org/constructor/updates.differenceSlice for reference.
type UpdatesDifferenceSlice struct {
	// List of new messgaes
	NewMessages []MessageClass `schemaname:"new_messages"`
	// New messages from the encrypted event sequence¹
	//
	// Links:
	//  1) https://core.telegram.org/api/updates
	NewEncryptedMessages []EncryptedMessageClass `schemaname:"new_encrypted_messages"`
	// List of updates
	OtherUpdates []UpdateClass `schemaname:"other_updates"`
	// List of chats mentioned in events
	Chats []ChatClass `schemaname:"chats"`
	// List of users mentioned in events
	Users []UserClass `schemaname:"users"`
	// Intermediary state
	IntermediateState UpdatesState `schemaname:"intermediate_state"`
}

// UpdatesDifferenceSliceTypeID is TL type id of UpdatesDifferenceSlice.
const UpdatesDifferenceSliceTypeID = 0xa8fb1981

func (d *UpdatesDifferenceSlice) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.NewMessages == nil) {
		return false
	}
	if !(d.NewEncryptedMessages == nil) {
		return false
	}
	if !(d.OtherUpdates == nil) {
		return false
	}
	if !(d.Chats == nil) {
		return false
	}
	if !(d.Users == nil) {
		return false
	}
	if !(d.IntermediateState.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *UpdatesDifferenceSlice) String() string {
	if d == nil {
		return "UpdatesDifferenceSlice(nil)"
	}
	type Alias UpdatesDifferenceSlice
	return fmt.Sprintf("UpdatesDifferenceSlice%+v", Alias(*d))
}

// FillFrom fills UpdatesDifferenceSlice from given interface.
func (d *UpdatesDifferenceSlice) FillFrom(from interface {
	GetNewMessages() (value []MessageClass)
	GetNewEncryptedMessages() (value []EncryptedMessageClass)
	GetOtherUpdates() (value []UpdateClass)
	GetChats() (value []ChatClass)
	GetUsers() (value []UserClass)
	GetIntermediateState() (value UpdatesState)
}) {
	d.NewMessages = from.GetNewMessages()
	d.NewEncryptedMessages = from.GetNewEncryptedMessages()
	d.OtherUpdates = from.GetOtherUpdates()
	d.Chats = from.GetChats()
	d.Users = from.GetUsers()
	d.IntermediateState = from.GetIntermediateState()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (d *UpdatesDifferenceSlice) TypeID() uint32 {
	return UpdatesDifferenceSliceTypeID
}

// SchemaName returns MTProto type name.
func (d *UpdatesDifferenceSlice) SchemaName() string {
	return "updates.differenceSlice"
}

// Encode implements bin.Encoder.
func (d *UpdatesDifferenceSlice) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode updates.differenceSlice#a8fb1981 as nil")
	}
	b.PutID(UpdatesDifferenceSliceTypeID)
	b.PutVectorHeader(len(d.NewMessages))
	for idx, v := range d.NewMessages {
		if v == nil {
			return fmt.Errorf("unable to encode updates.differenceSlice#a8fb1981: field new_messages element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.differenceSlice#a8fb1981: field new_messages element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(d.NewEncryptedMessages))
	for idx, v := range d.NewEncryptedMessages {
		if v == nil {
			return fmt.Errorf("unable to encode updates.differenceSlice#a8fb1981: field new_encrypted_messages element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.differenceSlice#a8fb1981: field new_encrypted_messages element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(d.OtherUpdates))
	for idx, v := range d.OtherUpdates {
		if v == nil {
			return fmt.Errorf("unable to encode updates.differenceSlice#a8fb1981: field other_updates element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.differenceSlice#a8fb1981: field other_updates element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(d.Chats))
	for idx, v := range d.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode updates.differenceSlice#a8fb1981: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.differenceSlice#a8fb1981: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(d.Users))
	for idx, v := range d.Users {
		if v == nil {
			return fmt.Errorf("unable to encode updates.differenceSlice#a8fb1981: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.differenceSlice#a8fb1981: field users element with index %d: %w", idx, err)
		}
	}
	if err := d.IntermediateState.Encode(b); err != nil {
		return fmt.Errorf("unable to encode updates.differenceSlice#a8fb1981: field intermediate_state: %w", err)
	}
	return nil
}

// GetNewMessages returns value of NewMessages field.
func (d *UpdatesDifferenceSlice) GetNewMessages() (value []MessageClass) {
	return d.NewMessages
}

// MapNewMessages returns field NewMessages wrapped in MessageClassArray helper.
func (d *UpdatesDifferenceSlice) MapNewMessages() (value MessageClassArray) {
	return MessageClassArray(d.NewMessages)
}

// GetNewEncryptedMessages returns value of NewEncryptedMessages field.
func (d *UpdatesDifferenceSlice) GetNewEncryptedMessages() (value []EncryptedMessageClass) {
	return d.NewEncryptedMessages
}

// MapNewEncryptedMessages returns field NewEncryptedMessages wrapped in EncryptedMessageClassArray helper.
func (d *UpdatesDifferenceSlice) MapNewEncryptedMessages() (value EncryptedMessageClassArray) {
	return EncryptedMessageClassArray(d.NewEncryptedMessages)
}

// GetOtherUpdates returns value of OtherUpdates field.
func (d *UpdatesDifferenceSlice) GetOtherUpdates() (value []UpdateClass) {
	return d.OtherUpdates
}

// MapOtherUpdates returns field OtherUpdates wrapped in UpdateClassArray helper.
func (d *UpdatesDifferenceSlice) MapOtherUpdates() (value UpdateClassArray) {
	return UpdateClassArray(d.OtherUpdates)
}

// GetChats returns value of Chats field.
func (d *UpdatesDifferenceSlice) GetChats() (value []ChatClass) {
	return d.Chats
}

// MapChats returns field Chats wrapped in ChatClassArray helper.
func (d *UpdatesDifferenceSlice) MapChats() (value ChatClassArray) {
	return ChatClassArray(d.Chats)
}

// GetUsers returns value of Users field.
func (d *UpdatesDifferenceSlice) GetUsers() (value []UserClass) {
	return d.Users
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (d *UpdatesDifferenceSlice) MapUsers() (value UserClassArray) {
	return UserClassArray(d.Users)
}

// GetIntermediateState returns value of IntermediateState field.
func (d *UpdatesDifferenceSlice) GetIntermediateState() (value UpdatesState) {
	return d.IntermediateState
}

// Decode implements bin.Decoder.
func (d *UpdatesDifferenceSlice) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode updates.differenceSlice#a8fb1981 to nil")
	}
	if err := b.ConsumeID(UpdatesDifferenceSliceTypeID); err != nil {
		return fmt.Errorf("unable to decode updates.differenceSlice#a8fb1981: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.differenceSlice#a8fb1981: field new_messages: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessage(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.differenceSlice#a8fb1981: field new_messages: %w", err)
			}
			d.NewMessages = append(d.NewMessages, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.differenceSlice#a8fb1981: field new_encrypted_messages: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeEncryptedMessage(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.differenceSlice#a8fb1981: field new_encrypted_messages: %w", err)
			}
			d.NewEncryptedMessages = append(d.NewEncryptedMessages, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.differenceSlice#a8fb1981: field other_updates: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUpdate(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.differenceSlice#a8fb1981: field other_updates: %w", err)
			}
			d.OtherUpdates = append(d.OtherUpdates, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.differenceSlice#a8fb1981: field chats: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.differenceSlice#a8fb1981: field chats: %w", err)
			}
			d.Chats = append(d.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.differenceSlice#a8fb1981: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.differenceSlice#a8fb1981: field users: %w", err)
			}
			d.Users = append(d.Users, value)
		}
	}
	{
		if err := d.IntermediateState.Decode(b); err != nil {
			return fmt.Errorf("unable to decode updates.differenceSlice#a8fb1981: field intermediate_state: %w", err)
		}
	}
	return nil
}

// construct implements constructor of UpdatesDifferenceClass.
func (d UpdatesDifferenceSlice) construct() UpdatesDifferenceClass { return &d }

// Ensuring interfaces in compile-time for UpdatesDifferenceSlice.
var (
	_ bin.Encoder = &UpdatesDifferenceSlice{}
	_ bin.Decoder = &UpdatesDifferenceSlice{}

	_ UpdatesDifferenceClass = &UpdatesDifferenceSlice{}
)

// UpdatesDifferenceTooLong represents TL type `updates.differenceTooLong#4afe8f6d`.
// The difference is too long¹, and the specified state must be used to refetch updates.
//
// Links:
//  1) https://core.telegram.org/api/updates#recovering-gaps
//
// See https://core.telegram.org/constructor/updates.differenceTooLong for reference.
type UpdatesDifferenceTooLong struct {
	// The new state to use.
	Pts int `schemaname:"pts"`
}

// UpdatesDifferenceTooLongTypeID is TL type id of UpdatesDifferenceTooLong.
const UpdatesDifferenceTooLongTypeID = 0x4afe8f6d

func (d *UpdatesDifferenceTooLong) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Pts == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *UpdatesDifferenceTooLong) String() string {
	if d == nil {
		return "UpdatesDifferenceTooLong(nil)"
	}
	type Alias UpdatesDifferenceTooLong
	return fmt.Sprintf("UpdatesDifferenceTooLong%+v", Alias(*d))
}

// FillFrom fills UpdatesDifferenceTooLong from given interface.
func (d *UpdatesDifferenceTooLong) FillFrom(from interface {
	GetPts() (value int)
}) {
	d.Pts = from.GetPts()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (d *UpdatesDifferenceTooLong) TypeID() uint32 {
	return UpdatesDifferenceTooLongTypeID
}

// SchemaName returns MTProto type name.
func (d *UpdatesDifferenceTooLong) SchemaName() string {
	return "updates.differenceTooLong"
}

// Encode implements bin.Encoder.
func (d *UpdatesDifferenceTooLong) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode updates.differenceTooLong#4afe8f6d as nil")
	}
	b.PutID(UpdatesDifferenceTooLongTypeID)
	b.PutInt(d.Pts)
	return nil
}

// GetPts returns value of Pts field.
func (d *UpdatesDifferenceTooLong) GetPts() (value int) {
	return d.Pts
}

// Decode implements bin.Decoder.
func (d *UpdatesDifferenceTooLong) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode updates.differenceTooLong#4afe8f6d to nil")
	}
	if err := b.ConsumeID(UpdatesDifferenceTooLongTypeID); err != nil {
		return fmt.Errorf("unable to decode updates.differenceTooLong#4afe8f6d: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.differenceTooLong#4afe8f6d: field pts: %w", err)
		}
		d.Pts = value
	}
	return nil
}

// construct implements constructor of UpdatesDifferenceClass.
func (d UpdatesDifferenceTooLong) construct() UpdatesDifferenceClass { return &d }

// Ensuring interfaces in compile-time for UpdatesDifferenceTooLong.
var (
	_ bin.Encoder = &UpdatesDifferenceTooLong{}
	_ bin.Decoder = &UpdatesDifferenceTooLong{}

	_ UpdatesDifferenceClass = &UpdatesDifferenceTooLong{}
)

// UpdatesDifferenceClass represents updates.Difference generic type.
//
// See https://core.telegram.org/type/updates.Difference for reference.
//
// Example:
//  g, err := tg.DecodeUpdatesDifference(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.UpdatesDifferenceEmpty: // updates.differenceEmpty#5d75a138
//  case *tg.UpdatesDifference: // updates.difference#f49ca0
//  case *tg.UpdatesDifferenceSlice: // updates.differenceSlice#a8fb1981
//  case *tg.UpdatesDifferenceTooLong: // updates.differenceTooLong#4afe8f6d
//  default: panic(v)
//  }
type UpdatesDifferenceClass interface {
	bin.Encoder
	bin.Decoder
	construct() UpdatesDifferenceClass

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// SchemaName returns MTProto type name.
	SchemaName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeUpdatesDifference implements binary de-serialization for UpdatesDifferenceClass.
func DecodeUpdatesDifference(buf *bin.Buffer) (UpdatesDifferenceClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case UpdatesDifferenceEmptyTypeID:
		// Decoding updates.differenceEmpty#5d75a138.
		v := UpdatesDifferenceEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UpdatesDifferenceClass: %w", err)
		}
		return &v, nil
	case UpdatesDifferenceTypeID:
		// Decoding updates.difference#f49ca0.
		v := UpdatesDifference{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UpdatesDifferenceClass: %w", err)
		}
		return &v, nil
	case UpdatesDifferenceSliceTypeID:
		// Decoding updates.differenceSlice#a8fb1981.
		v := UpdatesDifferenceSlice{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UpdatesDifferenceClass: %w", err)
		}
		return &v, nil
	case UpdatesDifferenceTooLongTypeID:
		// Decoding updates.differenceTooLong#4afe8f6d.
		v := UpdatesDifferenceTooLong{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UpdatesDifferenceClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode UpdatesDifferenceClass: %w", bin.NewUnexpectedID(id))
	}
}

// UpdatesDifference boxes the UpdatesDifferenceClass providing a helper.
type UpdatesDifferenceBox struct {
	Difference UpdatesDifferenceClass
}

// Decode implements bin.Decoder for UpdatesDifferenceBox.
func (b *UpdatesDifferenceBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode UpdatesDifferenceBox to nil")
	}
	v, err := DecodeUpdatesDifference(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Difference = v
	return nil
}

// Encode implements bin.Encode for UpdatesDifferenceBox.
func (b *UpdatesDifferenceBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Difference == nil {
		return fmt.Errorf("unable to encode UpdatesDifferenceClass as nil")
	}
	return b.Difference.Encode(buf)
}

// UpdatesDifferenceClassArray is adapter for slice of UpdatesDifferenceClass.
type UpdatesDifferenceClassArray []UpdatesDifferenceClass

// Sort sorts slice of UpdatesDifferenceClass.
func (s UpdatesDifferenceClassArray) Sort(less func(a, b UpdatesDifferenceClass) bool) UpdatesDifferenceClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of UpdatesDifferenceClass.
func (s UpdatesDifferenceClassArray) SortStable(less func(a, b UpdatesDifferenceClass) bool) UpdatesDifferenceClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of UpdatesDifferenceClass.
func (s UpdatesDifferenceClassArray) Retain(keep func(x UpdatesDifferenceClass) bool) UpdatesDifferenceClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s UpdatesDifferenceClassArray) First() (v UpdatesDifferenceClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UpdatesDifferenceClassArray) Last() (v UpdatesDifferenceClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UpdatesDifferenceClassArray) PopFirst() (v UpdatesDifferenceClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero UpdatesDifferenceClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UpdatesDifferenceClassArray) Pop() (v UpdatesDifferenceClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsUpdatesDifferenceEmpty returns copy with only UpdatesDifferenceEmpty constructors.
func (s UpdatesDifferenceClassArray) AsUpdatesDifferenceEmpty() (to UpdatesDifferenceEmptyArray) {
	for _, elem := range s {
		value, ok := elem.(*UpdatesDifferenceEmpty)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsUpdatesDifference returns copy with only UpdatesDifference constructors.
func (s UpdatesDifferenceClassArray) AsUpdatesDifference() (to UpdatesDifferenceArray) {
	for _, elem := range s {
		value, ok := elem.(*UpdatesDifference)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsUpdatesDifferenceSlice returns copy with only UpdatesDifferenceSlice constructors.
func (s UpdatesDifferenceClassArray) AsUpdatesDifferenceSlice() (to UpdatesDifferenceSliceArray) {
	for _, elem := range s {
		value, ok := elem.(*UpdatesDifferenceSlice)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsUpdatesDifferenceTooLong returns copy with only UpdatesDifferenceTooLong constructors.
func (s UpdatesDifferenceClassArray) AsUpdatesDifferenceTooLong() (to UpdatesDifferenceTooLongArray) {
	for _, elem := range s {
		value, ok := elem.(*UpdatesDifferenceTooLong)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// UpdatesDifferenceEmptyArray is adapter for slice of UpdatesDifferenceEmpty.
type UpdatesDifferenceEmptyArray []UpdatesDifferenceEmpty

// Sort sorts slice of UpdatesDifferenceEmpty.
func (s UpdatesDifferenceEmptyArray) Sort(less func(a, b UpdatesDifferenceEmpty) bool) UpdatesDifferenceEmptyArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of UpdatesDifferenceEmpty.
func (s UpdatesDifferenceEmptyArray) SortStable(less func(a, b UpdatesDifferenceEmpty) bool) UpdatesDifferenceEmptyArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of UpdatesDifferenceEmpty.
func (s UpdatesDifferenceEmptyArray) Retain(keep func(x UpdatesDifferenceEmpty) bool) UpdatesDifferenceEmptyArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s UpdatesDifferenceEmptyArray) First() (v UpdatesDifferenceEmpty, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UpdatesDifferenceEmptyArray) Last() (v UpdatesDifferenceEmpty, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UpdatesDifferenceEmptyArray) PopFirst() (v UpdatesDifferenceEmpty, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero UpdatesDifferenceEmpty
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UpdatesDifferenceEmptyArray) Pop() (v UpdatesDifferenceEmpty, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByDate sorts slice of UpdatesDifferenceEmpty by Date.
func (s UpdatesDifferenceEmptyArray) SortByDate() UpdatesDifferenceEmptyArray {
	return s.Sort(func(a, b UpdatesDifferenceEmpty) bool {
		return a.GetDate() < b.GetDate()
	})
}

// SortStableByDate sorts slice of UpdatesDifferenceEmpty by Date.
func (s UpdatesDifferenceEmptyArray) SortStableByDate() UpdatesDifferenceEmptyArray {
	return s.SortStable(func(a, b UpdatesDifferenceEmpty) bool {
		return a.GetDate() < b.GetDate()
	})
}

// UpdatesDifferenceArray is adapter for slice of UpdatesDifference.
type UpdatesDifferenceArray []UpdatesDifference

// Sort sorts slice of UpdatesDifference.
func (s UpdatesDifferenceArray) Sort(less func(a, b UpdatesDifference) bool) UpdatesDifferenceArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of UpdatesDifference.
func (s UpdatesDifferenceArray) SortStable(less func(a, b UpdatesDifference) bool) UpdatesDifferenceArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of UpdatesDifference.
func (s UpdatesDifferenceArray) Retain(keep func(x UpdatesDifference) bool) UpdatesDifferenceArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s UpdatesDifferenceArray) First() (v UpdatesDifference, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UpdatesDifferenceArray) Last() (v UpdatesDifference, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UpdatesDifferenceArray) PopFirst() (v UpdatesDifference, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero UpdatesDifference
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UpdatesDifferenceArray) Pop() (v UpdatesDifference, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// UpdatesDifferenceSliceArray is adapter for slice of UpdatesDifferenceSlice.
type UpdatesDifferenceSliceArray []UpdatesDifferenceSlice

// Sort sorts slice of UpdatesDifferenceSlice.
func (s UpdatesDifferenceSliceArray) Sort(less func(a, b UpdatesDifferenceSlice) bool) UpdatesDifferenceSliceArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of UpdatesDifferenceSlice.
func (s UpdatesDifferenceSliceArray) SortStable(less func(a, b UpdatesDifferenceSlice) bool) UpdatesDifferenceSliceArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of UpdatesDifferenceSlice.
func (s UpdatesDifferenceSliceArray) Retain(keep func(x UpdatesDifferenceSlice) bool) UpdatesDifferenceSliceArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s UpdatesDifferenceSliceArray) First() (v UpdatesDifferenceSlice, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UpdatesDifferenceSliceArray) Last() (v UpdatesDifferenceSlice, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UpdatesDifferenceSliceArray) PopFirst() (v UpdatesDifferenceSlice, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero UpdatesDifferenceSlice
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UpdatesDifferenceSliceArray) Pop() (v UpdatesDifferenceSlice, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// UpdatesDifferenceTooLongArray is adapter for slice of UpdatesDifferenceTooLong.
type UpdatesDifferenceTooLongArray []UpdatesDifferenceTooLong

// Sort sorts slice of UpdatesDifferenceTooLong.
func (s UpdatesDifferenceTooLongArray) Sort(less func(a, b UpdatesDifferenceTooLong) bool) UpdatesDifferenceTooLongArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of UpdatesDifferenceTooLong.
func (s UpdatesDifferenceTooLongArray) SortStable(less func(a, b UpdatesDifferenceTooLong) bool) UpdatesDifferenceTooLongArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of UpdatesDifferenceTooLong.
func (s UpdatesDifferenceTooLongArray) Retain(keep func(x UpdatesDifferenceTooLong) bool) UpdatesDifferenceTooLongArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s UpdatesDifferenceTooLongArray) First() (v UpdatesDifferenceTooLong, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UpdatesDifferenceTooLongArray) Last() (v UpdatesDifferenceTooLong, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UpdatesDifferenceTooLongArray) PopFirst() (v UpdatesDifferenceTooLong, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero UpdatesDifferenceTooLong
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UpdatesDifferenceTooLongArray) Pop() (v UpdatesDifferenceTooLong, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
