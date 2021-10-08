package telegram

import (
	"context"
	"net"
	"strings"
	"syscall"

	"go.uber.org/zap"
	"golang.org/x/xerrors"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tg"
	"github.com/gotd/td/tgerr"
)

// API returns *tg.Client for calling raw MTProto methods.
func (c *Client) API() *tg.Client {
	return c.tg
}

// Invoke invokes raw MTProto RPC method. It sends input and decodes result
// into output.
func (c *Client) Invoke(ctx context.Context, input bin.Encoder, output bin.Decoder) error {
	return c.invoker.Invoke(ctx, input, output)
}

// invokeDirect directly invokes RPC method, automatically handling datacenter redirects.
func (c *Client) invokeDirect(ctx context.Context, input bin.Encoder, output bin.Decoder) error {
	if err := c.invokeConn(ctx, input, output); err != nil {
		if isNetworkError(err) {
			if c.acquireRestart.CAS(false, true) {
				defer c.acquireRestart.Store(false)
				c.log.Debug("Restarting connection due to network error", zap.Error(err))
				c.resetReady()
				c.restart <- struct{}{}
			}

			c.log.Debug("Waiting for reconnection")
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c.ready.Ready():
				c.log.Debug("Connection restarted, re-invoking request")
			}

			return c.invokeDirect(ctx, input, output)
		}

		// Handling datacenter migration request.
		if rpcErr, ok := tgerr.As(err); ok && strings.HasSuffix(rpcErr.Type, "_MIGRATE") {
			targetDC := rpcErr.Argument
			log := c.log.With(
				zap.String("error_type", rpcErr.Type),
				zap.Int("target_dc", targetDC),
			)
			// If migration error is FILE_MIGRATE or STATS_MIGRATE, then the method
			// called by authorized client, so we should try to transfer auth to new DC
			// and create new connection.
			if rpcErr.IsOneOf("FILE_MIGRATE", "STATS_MIGRATE") {
				log.Debug("Invoking on target DC")
				return c.invokeSub(ctx, targetDC, input, output)
			}

			// Otherwise we should change primary DC.
			log.Info("Migrating to target DC")
			return c.invokeMigrate(ctx, targetDC, input, output)
		}

		return err
	}

	return nil
}

// invokeConn directly invokes RPC call on primary connection without any
// additional handling.
func (c *Client) invokeConn(ctx context.Context, input bin.Encoder, output bin.Decoder) error {
	c.connMux.Lock()
	conn := c.conn
	c.connMux.Unlock()

	return conn.Invoke(ctx, input, output)
}

func isNetworkError(err error) bool {
	if xerrors.Is(err, syscall.EPIPE) {
		return true
	}

	if err, ok := err.(net.Error); ok && err.Timeout() {
		return true
	}

	return false
}
