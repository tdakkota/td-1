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
	"github.com/gotd/td/tdjson"
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
	_ = tdjson.Encoder{}
)

// MessagesTranscribedAudio represents TL type `messages.transcribedAudio#93752c52`.
// Transcribed text from a voice message
//
// See https://core.telegram.org/constructor/messages.transcribedAudio for reference.
type MessagesTranscribedAudio struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether the transcription is partial because audio transcription is still in progress,
	// if set the user may receive further updateTranscribedAudio¹ updates with the updated
	// transcription.
	//
	// Links:
	//  1) https://core.telegram.org/constructor/updateTranscribedAudio
	Pending bool
	// Transcription ID
	TranscriptionID int64
	// Transcripted text
	Text string
}

// MessagesTranscribedAudioTypeID is TL type id of MessagesTranscribedAudio.
const MessagesTranscribedAudioTypeID = 0x93752c52

// Ensuring interfaces in compile-time for MessagesTranscribedAudio.
var (
	_ bin.Encoder     = &MessagesTranscribedAudio{}
	_ bin.Decoder     = &MessagesTranscribedAudio{}
	_ bin.BareEncoder = &MessagesTranscribedAudio{}
	_ bin.BareDecoder = &MessagesTranscribedAudio{}
)

func (t *MessagesTranscribedAudio) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Flags.Zero()) {
		return false
	}
	if !(t.Pending == false) {
		return false
	}
	if !(t.TranscriptionID == 0) {
		return false
	}
	if !(t.Text == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *MessagesTranscribedAudio) String() string {
	if t == nil {
		return "MessagesTranscribedAudio(nil)"
	}
	type Alias MessagesTranscribedAudio
	return fmt.Sprintf("MessagesTranscribedAudio%+v", Alias(*t))
}

// FillFrom fills MessagesTranscribedAudio from given interface.
func (t *MessagesTranscribedAudio) FillFrom(from interface {
	GetPending() (value bool)
	GetTranscriptionID() (value int64)
	GetText() (value string)
}) {
	t.Pending = from.GetPending()
	t.TranscriptionID = from.GetTranscriptionID()
	t.Text = from.GetText()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesTranscribedAudio) TypeID() uint32 {
	return MessagesTranscribedAudioTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesTranscribedAudio) TypeName() string {
	return "messages.transcribedAudio"
}

// TypeInfo returns info about TL type.
func (t *MessagesTranscribedAudio) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.transcribedAudio",
		ID:   MessagesTranscribedAudioTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Pending",
			SchemaName: "pending",
			Null:       !t.Flags.Has(0),
		},
		{
			Name:       "TranscriptionID",
			SchemaName: "transcription_id",
		},
		{
			Name:       "Text",
			SchemaName: "text",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (t *MessagesTranscribedAudio) SetFlags() {
	if !(t.Pending == false) {
		t.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (t *MessagesTranscribedAudio) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode messages.transcribedAudio#93752c52 as nil")
	}
	b.PutID(MessagesTranscribedAudioTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *MessagesTranscribedAudio) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode messages.transcribedAudio#93752c52 as nil")
	}
	t.SetFlags()
	if err := t.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.transcribedAudio#93752c52: field flags: %w", err)
	}
	b.PutLong(t.TranscriptionID)
	b.PutString(t.Text)
	return nil
}

// Decode implements bin.Decoder.
func (t *MessagesTranscribedAudio) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode messages.transcribedAudio#93752c52 to nil")
	}
	if err := b.ConsumeID(MessagesTranscribedAudioTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.transcribedAudio#93752c52: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *MessagesTranscribedAudio) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode messages.transcribedAudio#93752c52 to nil")
	}
	{
		if err := t.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.transcribedAudio#93752c52: field flags: %w", err)
		}
	}
	t.Pending = t.Flags.Has(0)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.transcribedAudio#93752c52: field transcription_id: %w", err)
		}
		t.TranscriptionID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.transcribedAudio#93752c52: field text: %w", err)
		}
		t.Text = value
	}
	return nil
}

// SetPending sets value of Pending conditional field.
func (t *MessagesTranscribedAudio) SetPending(value bool) {
	if value {
		t.Flags.Set(0)
		t.Pending = true
	} else {
		t.Flags.Unset(0)
		t.Pending = false
	}
}

// GetPending returns value of Pending conditional field.
func (t *MessagesTranscribedAudio) GetPending() (value bool) {
	if t == nil {
		return
	}
	return t.Flags.Has(0)
}

// GetTranscriptionID returns value of TranscriptionID field.
func (t *MessagesTranscribedAudio) GetTranscriptionID() (value int64) {
	if t == nil {
		return
	}
	return t.TranscriptionID
}

// GetText returns value of Text field.
func (t *MessagesTranscribedAudio) GetText() (value string) {
	if t == nil {
		return
	}
	return t.Text
}
