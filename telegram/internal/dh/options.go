package dh

import (
	"io"

	"go.uber.org/zap"

	"github.com/gotd/td/internal/crypto"
)

// Options of State.
type Options struct {
	// Random is random source. Defaults to crypto.
	Random io.Reader
	// Logger is instance of zap.Logger. No logs by default.
	Logger *zap.Logger
}

func (o *Options) setDefaults() {
	if o.Random == nil {
		o.Random = crypto.DefaultRand()
	}
	if o.Logger == nil {
		o.Logger = zap.NewNop()
	}
}
