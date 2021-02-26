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

// AccountSaveWallPaperRequest represents TL type `account.saveWallPaper#6c5a5b37`.
// Install/uninstall wallpaper
//
// See https://core.telegram.org/method/account.saveWallPaper for reference.
type AccountSaveWallPaperRequest struct {
	// Wallpaper to save
	Wallpaper InputWallPaperClass `schemaname:"wallpaper"`
	// Uninstall wallpaper?
	Unsave bool `schemaname:"unsave"`
	// Wallpaper settings
	Settings WallPaperSettings `schemaname:"settings"`
}

// AccountSaveWallPaperRequestTypeID is TL type id of AccountSaveWallPaperRequest.
const AccountSaveWallPaperRequestTypeID = 0x6c5a5b37

func (s *AccountSaveWallPaperRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Wallpaper == nil) {
		return false
	}
	if !(s.Unsave == false) {
		return false
	}
	if !(s.Settings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *AccountSaveWallPaperRequest) String() string {
	if s == nil {
		return "AccountSaveWallPaperRequest(nil)"
	}
	type Alias AccountSaveWallPaperRequest
	return fmt.Sprintf("AccountSaveWallPaperRequest%+v", Alias(*s))
}

// FillFrom fills AccountSaveWallPaperRequest from given interface.
func (s *AccountSaveWallPaperRequest) FillFrom(from interface {
	GetWallpaper() (value InputWallPaperClass)
	GetUnsave() (value bool)
	GetSettings() (value WallPaperSettings)
}) {
	s.Wallpaper = from.GetWallpaper()
	s.Unsave = from.GetUnsave()
	s.Settings = from.GetSettings()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *AccountSaveWallPaperRequest) TypeID() uint32 {
	return AccountSaveWallPaperRequestTypeID
}

// SchemaName returns MTProto type name.
func (s *AccountSaveWallPaperRequest) SchemaName() string {
	return "account.saveWallPaper"
}

// Encode implements bin.Encoder.
func (s *AccountSaveWallPaperRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.saveWallPaper#6c5a5b37 as nil")
	}
	b.PutID(AccountSaveWallPaperRequestTypeID)
	if s.Wallpaper == nil {
		return fmt.Errorf("unable to encode account.saveWallPaper#6c5a5b37: field wallpaper is nil")
	}
	if err := s.Wallpaper.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.saveWallPaper#6c5a5b37: field wallpaper: %w", err)
	}
	b.PutBool(s.Unsave)
	if err := s.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.saveWallPaper#6c5a5b37: field settings: %w", err)
	}
	return nil
}

// GetWallpaper returns value of Wallpaper field.
func (s *AccountSaveWallPaperRequest) GetWallpaper() (value InputWallPaperClass) {
	return s.Wallpaper
}

// GetUnsave returns value of Unsave field.
func (s *AccountSaveWallPaperRequest) GetUnsave() (value bool) {
	return s.Unsave
}

// GetSettings returns value of Settings field.
func (s *AccountSaveWallPaperRequest) GetSettings() (value WallPaperSettings) {
	return s.Settings
}

// Decode implements bin.Decoder.
func (s *AccountSaveWallPaperRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.saveWallPaper#6c5a5b37 to nil")
	}
	if err := b.ConsumeID(AccountSaveWallPaperRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.saveWallPaper#6c5a5b37: %w", err)
	}
	{
		value, err := DecodeInputWallPaper(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.saveWallPaper#6c5a5b37: field wallpaper: %w", err)
		}
		s.Wallpaper = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode account.saveWallPaper#6c5a5b37: field unsave: %w", err)
		}
		s.Unsave = value
	}
	{
		if err := s.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.saveWallPaper#6c5a5b37: field settings: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountSaveWallPaperRequest.
var (
	_ bin.Encoder = &AccountSaveWallPaperRequest{}
	_ bin.Decoder = &AccountSaveWallPaperRequest{}
)

// AccountSaveWallPaper invokes method account.saveWallPaper#6c5a5b37 returning error if any.
// Install/uninstall wallpaper
//
// See https://core.telegram.org/method/account.saveWallPaper for reference.
func (c *Client) AccountSaveWallPaper(ctx context.Context, request *AccountSaveWallPaperRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
