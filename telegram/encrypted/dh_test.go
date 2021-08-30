package encrypted

import (
	"context"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gotd/td/internal/crypto"
	"github.com/gotd/td/internal/testutil"
	"github.com/gotd/td/tg"
	"github.com/gotd/td/tgmock"
)

var testP = func() []byte {
	p, err := hex.DecodeString("C71CAEB9C6B1C9048E6C522F70F13F73980D40238E3E21C14934D037563D930F" +
		"48198A0AA7C14058229493D22530F4DBFA336F6E0AC925139543AED44CCE7C37" +
		"20FD51F69458705AC68CD4FE6B6B13ABDC9746512969328454F18FAF8C595F64" +
		"2477FE96BB2A941D5BCD1D4AC8CC49880708FA9B378E3C4F3A9060BEE67CF9A4" +
		"A4A695811051907E162753B56B0F6B410DBA74D8A84B2A14B3144E0EF1284754" +
		"FD17ED950D5965B4B9DD46582DB1178D169C6BC465B0D6FF9CA3928FEF5B9AE4" +
		"E418FC15E83EBEA0F87FA9FF5EED70050DED2849F47BF959D956850CE929851F" +
		"0D8115F635B105EE2E4E15D04B2454BF6F4FADF034B10403119CD8E3B92FCC5B")
	if err != nil {
		panic(err)
	}
	return p
}()

func TestManager_initDH(t *testing.T) {
	ctx := context.Background()
	invoker := tgmock.New(t)

	m := Manager{
		raw:    tg.NewClient(invoker),
		rand:   crypto.DefaultRand(),
		logger: zap.NewNop(),
	}
	t.Run("RPCError", func(t *testing.T) {
		invoker.ExpectCall(&tg.MessagesGetDhConfigRequest{
			Version:      0,
			RandomLength: crypto.RSAKeyBits,
		}).ThenErr(testutil.TestError())
		_, _, err := m.initDH(ctx)
		require.Error(t, err)
	})
	t.Run("UnexpectedNotModified", func(t *testing.T) {
		invoker.ExpectCall(&tg.MessagesGetDhConfigRequest{
			Version:      0,
			RandomLength: crypto.RSAKeyBits,
		}).ThenResult(&tg.MessagesDhConfigNotModified{})
		_, _, err := m.initDH(ctx)
		require.Error(t, err)
	})
	t.Run("InvalidP", func(t *testing.T) {
		invoker.ExpectCall(&tg.MessagesGetDhConfigRequest{
			Version:      0,
			RandomLength: crypto.RSAKeyBits,
		}).ThenResult(&tg.MessagesDhConfig{})
		_, _, err := m.initDH(ctx)
		require.Error(t, err)
	})
	t.Run("OK", func(t *testing.T) {
		a := require.New(t)
		g := 3

		invoker.ExpectCall(&tg.MessagesGetDhConfigRequest{
			Version:      0,
			RandomLength: crypto.RSAKeyBits,
		}).ThenResult(&tg.MessagesDhConfig{
			G:       g,
			P:       testP,
			Version: 1,
			Random:  nil,
		})

		_, cfg, err := m.initDH(ctx)
		a.NoError(err)
		a.Equal(g, cfg.G)
		a.Equal(int64(g), cfg.GBig.Int64())
		a.Equal(1, cfg.Version)
	})
}
