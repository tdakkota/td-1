package dh

import (
	"context"
	"crypto/rand"
	"io"
	"math/big"
	"sync"

	"github.com/go-faster/errors"
	"go.uber.org/zap"

	"github.com/gotd/td/internal/crypto"
	"github.com/gotd/td/tg"
)

// Config represents Diffie-Hellman key exchange config.
type Config struct {
	G       int
	GBig    *big.Int
	P       *big.Int
	Version int
	set     bool
}

// State stores current DH state.
type State struct {
	cfg    Config
	cfgMux sync.Mutex

	raw  *tg.Client
	rand io.Reader

	logger *zap.Logger
}

// NewState creates new State.
func NewState(raw *tg.Client, opts Options) *State {
	opts.setDefaults()
	return &State{
		raw:    raw,
		rand:   opts.Random,
		logger: opts.Logger,
	}
}

// Init requests the latest DH config and generates new random number.
func (m *State) Init(ctx context.Context) (*big.Int, Config, error) {
	dhCfg, err := m.GetDHConfig(ctx)
	if err != nil {
		return nil, Config{}, err
	}

	a, err := generateBig(m.rand)
	if err != nil {
		return nil, Config{}, errors.Wrap(err, "generate random")
	}

	return a, dhCfg, nil
}

// GetDHConfig requests Diffie-Hellman protocol config.
//
// See https://core.telegram.org/api/end-to-end#sending-a-request.
//
// See https://core.telegram.org/api/end-to-end/video-calls#key-generation.
//
// See https://core.telegram.org/method/messages.getDhConfig.
func (m *State) GetDHConfig(ctx context.Context) (Config, error) {
	m.cfgMux.Lock()
	defer m.cfgMux.Unlock()

	version := 0
	if m.cfg.set {
		version = m.cfg.Version
	}

	d, err := m.raw.MessagesGetDhConfig(ctx, &tg.MessagesGetDhConfigRequest{
		RandomLength: crypto.RSAKeyBits,
		Version:      version,
	})
	if err != nil {
		return Config{}, errors.Wrap(err, "get DH config")
	}

	switch cfg := d.(type) {
	case *tg.MessagesDhConfig:
		p := big.NewInt(0).SetBytes(cfg.P)

		if err := crypto.CheckDH(cfg.G, p); err != nil {
			return Config{}, errors.Wrap(err, "check DH")
		}

		wasSet := m.cfg.set
		m.cfg = Config{
			G:       cfg.G,
			GBig:    big.NewInt(int64(cfg.G)),
			P:       p,
			Version: cfg.Version,
			set:     true,
		}
		if wasSet {
			m.logger.Debug("DH Config updated", zap.Int("version", cfg.Version))
		} else {
			m.logger.Debug("DH Config received", zap.Int("version", cfg.Version))
		}
	case *tg.MessagesDhConfigNotModified:
		if !m.cfg.set {
			return Config{}, errors.Errorf("unexpected type %T", d)
		}
	default:
		return Config{}, errors.Errorf("unexpected type %T", d)
	}

	return m.cfg, nil
}

var randMax = big.NewInt(0).SetBit(big.NewInt(0), crypto.RSAKeyBits, 1)

func generateBig(r io.Reader) (*big.Int, error) {
	a, err := rand.Int(r, randMax)
	if err != nil {
		return nil, errors.Errorf("generate %d-bit number: %w", crypto.RSAKeyBits, err)
	}
	return a, nil
}
