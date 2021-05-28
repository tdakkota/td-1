package telegram

import (
	"context"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tg"
)

// InvokeFunc implements tg.Invoker as function.
type InvokeFunc func(ctx context.Context, input bin.Encoder, output bin.Decoder) error

// Invoke implements tg.Invoker.
func (i InvokeFunc) Invoke(ctx context.Context, input bin.Encoder, output bin.Decoder) error {
	return i(ctx, input, output)
}

// Middleware returns new InvokeFunc for next invoker.
type Middleware interface {
	Handle(next tg.Invoker) InvokeFunc
}

// MiddlewareFunc implements Middleware as function.
type MiddlewareFunc func(next tg.Invoker) InvokeFunc

// Handle implements Middleware.
func (m MiddlewareFunc) Handle(next tg.Invoker) InvokeFunc { return m(next) }

func chainMiddlewares(invoker tg.Invoker, chain ...Middleware) tg.Invoker {
	if len(chain) == 0 {
		return invoker
	}
	for i := len(chain) - 1; i >= 0; i-- {
		invoker = chain[i].Handle(invoker)
	}
	return invoker
}
