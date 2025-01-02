package threed

import (
	"testing"
	"time"

	abci "github.com/cometbft/cometbft/abci/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/dymensionxyz/rollapp-evm/app"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestBeginBlocker(t *testing.T) {
	app, _ := app.SetupWithOneValidator(t)
	h := int64(42)
	ctx := app.NewUncachedContext(true, types.Header{
		Height:  h - 1,
		ChainID: "testchain_9000-1",
	})
	err := app.UpgradeKeeper.ScheduleUpgrade(ctx, upgradetypes.Plan{
		Name:   UpgradeName,
		Height: h,
	})
	require.NoError(t, err)

	// simulate the upgrade process not panic.
	require.NotPanics(t, func() {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("Upgrade panicked: %v", r)
			}
		}()
		// simulate the upgrade process.
		ctx = ctx.WithBlockHeight(h)
		app.BeginBlocker(ctx, abci.RequestBeginBlock{})
	})
}
