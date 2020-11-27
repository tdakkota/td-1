package telegram

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"go.uber.org/zap"
	"golang.org/x/xerrors"

	"github.com/ernado/td/bin"
	"github.com/ernado/td/internal/mt"
	"github.com/ernado/td/internal/proto"
)

func (c *Client) handlePong(b *bin.Buffer) error {
	var pong mt.Pong
	if err := pong.Decode(b); err != nil {
		return xerrors.Errorf("failed to decode: %x", err)
	}
	c.log.Info("Pong")

	c.pingMux.Lock()
	f, ok := c.ping[pong.PingID]
	c.pingMux.Unlock()
	if ok {
		f()
	}
	return nil
}

func (c *Client) handleSessionCreated(b *bin.Buffer) error {
	var ns mt.NewSessionCreated
	if err := ns.Decode(b); err != nil {
		return xerrors.Errorf("failed to decode: %x", err)
	}
	c.log.Info("Session created")
	return nil
}

func (c *Client) handleUnknown(b *bin.Buffer) error {
	// Can't process unknown type.
	id, err := b.PeekID()
	if err != nil {
		return err
	}
	c.log.With(
		zap.String("type_id", fmt.Sprintf("0x%x", id)),
	).Warn("Unknown type id")

	return nil
}

func (c *Client) handleMessage(b *bin.Buffer) error {
	id, err := b.PeekID()
	if err != nil {
		// Empty body.
		return xerrors.Errorf("failed to determine message type: %w", err)
	}
	switch id {
	case mt.BadMsgNotificationTypeID, mt.BadServerSaltTypeID:
		return c.handleBadMsg(b)
	case proto.MessageContainerTypeID:
		return c.processBatch(b)
	case mt.NewSessionCreatedTypeID:
		return c.handleSessionCreated(b)
	case proto.ResultTypeID:
		return c.handleResult(b)
	case mt.PongTypeID:
		return c.handlePong(b)
	default:
		return c.handleUnknown(b)
	}
}

func (c *Client) processBatch(b *bin.Buffer) error {
	var container proto.MessageContainer
	if err := container.Decode(b); err != nil {
		return xerrors.Errorf("failed to decode container: %w", err)
	}
	for _, msg := range container.Messages {
		if err := c.processContainerMessage(msg); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) processContainerMessage(msg proto.Message) error {
	b := &bin.Buffer{Buf: msg.Body}
	return c.handleMessage(b)
}

func (c *Client) read(ctx context.Context, b *bin.Buffer) error {
	b.Reset()
	defer func() {
		// Reset deadline.
		_ = c.conn.SetReadDeadline(time.Time{})
	}()
	if err := c.conn.SetReadDeadline(c.deadline(ctx)); err != nil {
		return xerrors.Errorf("failed to set read deadline: %w", err)
	}
	if err := proto.ReadIntermediate(c.conn, b); err != nil {
		return xerrors.Errorf("failed to read intermediate: %w", err)
	}
	if err := c.checkProtocolError(b); err != nil {
		return xerrors.Errorf("protocol error: %w", err)
	}

	// Decrypting.
	encMessage := &proto.EncryptedMessage{}
	if err := encMessage.Decode(b); err != nil {
		return xerrors.Errorf("failed to decode encrypted message: %w", err)
	}
	msg, err := c.decryptData(encMessage)
	if err != nil {
		return xerrors.Errorf("failed to decrypt: %w", err)
	}

	// Buffer now contains plaintext message payload.
	b.ResetTo(msg.MessageDataWithPadding[:msg.MessageDataLen])
	if err := c.handleMessage(b); err != nil {
		return xerrors.Errorf("failed to handle message: %w", err)
	}

	return nil
}

func (c *Client) readLoop(ctx context.Context) {
	b := new(bin.Buffer)
	log := c.log.Named("read")
	log.Debug("Read loop started")

	for {
		err := c.read(ctx, b)
		if err == nil {
			// Reading ok.
			log.Debug("Read message")
			continue
		}
		if errors.Is(err, io.EOF) {
			// Nothing was received.
			continue
		}

		// Handling possible errors.
		log.With(zap.Error(err)).Error("Read returned error")
	}
}
