// Code generated by gotdgen, DO NOT EDIT.

package mt

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

// DestroySessionRequest represents TL type `destroy_session#e7512126`.
type DestroySessionRequest struct {
	// SessionID field of DestroySessionRequest.
	SessionID int64 `schemaname:"session_id"`
}

// DestroySessionRequestTypeID is TL type id of DestroySessionRequest.
const DestroySessionRequestTypeID = 0xe7512126

func (d *DestroySessionRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.SessionID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DestroySessionRequest) String() string {
	if d == nil {
		return "DestroySessionRequest(nil)"
	}
	type Alias DestroySessionRequest
	return fmt.Sprintf("DestroySessionRequest%+v", Alias(*d))
}

// FillFrom fills DestroySessionRequest from given interface.
func (d *DestroySessionRequest) FillFrom(from interface {
	GetSessionID() (value int64)
}) {
	d.SessionID = from.GetSessionID()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (d *DestroySessionRequest) TypeID() uint32 {
	return DestroySessionRequestTypeID
}

// SchemaName returns MTProto type name.
func (d *DestroySessionRequest) SchemaName() string {
	return "destroy_session"
}

// Encode implements bin.Encoder.
func (d *DestroySessionRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode destroy_session#e7512126 as nil")
	}
	b.PutID(DestroySessionRequestTypeID)
	b.PutLong(d.SessionID)
	return nil
}

// GetSessionID returns value of SessionID field.
func (d *DestroySessionRequest) GetSessionID() (value int64) {
	return d.SessionID
}

// Decode implements bin.Decoder.
func (d *DestroySessionRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode destroy_session#e7512126 to nil")
	}
	if err := b.ConsumeID(DestroySessionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode destroy_session#e7512126: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode destroy_session#e7512126: field session_id: %w", err)
		}
		d.SessionID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for DestroySessionRequest.
var (
	_ bin.Encoder = &DestroySessionRequest{}
	_ bin.Decoder = &DestroySessionRequest{}
)

// DestroySession invokes method destroy_session#e7512126 returning error if any.
func (c *Client) DestroySession(ctx context.Context, sessionid int64) (DestroySessionResClass, error) {
	var result DestroySessionResBox

	request := &DestroySessionRequest{
		SessionID: sessionid,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.DestroySessionRes, nil
}
