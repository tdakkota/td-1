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

// ReqPqRequest represents TL type `req_pq#60469778`.
type ReqPqRequest struct {
	// Nonce field of ReqPqRequest.
	Nonce bin.Int128 `schemaname:"nonce"`
}

// ReqPqRequestTypeID is TL type id of ReqPqRequest.
const ReqPqRequestTypeID = 0x60469778

func (r *ReqPqRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Nonce == bin.Int128{}) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReqPqRequest) String() string {
	if r == nil {
		return "ReqPqRequest(nil)"
	}
	type Alias ReqPqRequest
	return fmt.Sprintf("ReqPqRequest%+v", Alias(*r))
}

// FillFrom fills ReqPqRequest from given interface.
func (r *ReqPqRequest) FillFrom(from interface {
	GetNonce() (value bin.Int128)
}) {
	r.Nonce = from.GetNonce()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (r *ReqPqRequest) TypeID() uint32 {
	return ReqPqRequestTypeID
}

// SchemaName returns MTProto type name.
func (r *ReqPqRequest) SchemaName() string {
	return "req_pq"
}

// Encode implements bin.Encoder.
func (r *ReqPqRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode req_pq#60469778 as nil")
	}
	b.PutID(ReqPqRequestTypeID)
	b.PutInt128(r.Nonce)
	return nil
}

// GetNonce returns value of Nonce field.
func (r *ReqPqRequest) GetNonce() (value bin.Int128) {
	return r.Nonce
}

// Decode implements bin.Decoder.
func (r *ReqPqRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode req_pq#60469778 to nil")
	}
	if err := b.ConsumeID(ReqPqRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode req_pq#60469778: %w", err)
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode req_pq#60469778: field nonce: %w", err)
		}
		r.Nonce = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ReqPqRequest.
var (
	_ bin.Encoder = &ReqPqRequest{}
	_ bin.Decoder = &ReqPqRequest{}
)

// ReqPq invokes method req_pq#60469778 returning error if any.
func (c *Client) ReqPq(ctx context.Context, nonce bin.Int128) (*ResPQ, error) {
	var result ResPQ

	request := &ReqPqRequest{
		Nonce: nonce,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
