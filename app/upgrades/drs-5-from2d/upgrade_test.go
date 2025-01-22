package drs5from2d_test

import (
	"testing"

	up "github.com/dymensionxyz/rollapp-evm/app/upgrades/drs-5-from2d"
	erc20types "github.com/evmos/evmos/v12/x/erc20/types"
	abci "github.com/tendermint/tendermint/abci/types"

	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/dymensionxyz/rollapp-evm/app"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestBeginBlocker(t *testing.T) {

	app, _ := app.SetupWithOneValidator(t)
	app.SetDymintVersionGetter(func() (uint32, error) {
		return uint32(up.DRS), nil
	})
	h := int64(42)
	ctx := app.NewUncachedContext(true, types.Header{
		Height:  h - 1,
		ChainID: "testchain_9000-1",
	})
	err := app.UpgradeKeeper.ScheduleUpgrade(ctx, upgradetypes.Plan{
		Name:   up.UpgradeName,
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
		// simulate the upgrade process
		ctx = ctx.WithBlockHeight(h)
		app.BeginBlocker(ctx, abci.RequestBeginBlock{})
	})

	// Check if the upgrade was applied correctly
	state := app.HubGenesisKeeper.GetState(ctx)
	require.True(t, state.OutboundTransfersEnabled, "OutboundTransfersEnabled should be true")
	require.False(t, state.InFlight, "InFlight should be false")
	require.NotNil(t, state.HubPortAndChannel, "HubPortAndChannel should not be nil")
	require.Equal(t, "transfer", state.HubPortAndChannel.Port, "Port should be 'transfer'")
	require.Equal(t, "channel-0", state.HubPortAndChannel.Channel, "Channel should be 'channel-0'")

	p := app.RollappParamsKeeper.GetParams(ctx)
	require.Equal(t, up.DRS, p.DrsVersion, "Version should be set to DRS")
	require.Equal(t, up.DA, p.Da, "Version should be set to DRS")

	evmParams := app.EvmKeeper.GetParams(ctx)
	require.Equal(t, evmParams.GasDenom, evmParams.EvmDenom, "GasDenom should be set to EvmDenom")

	erc20Params := app.Erc20Keeper.GetParams(ctx)
	require.Equal(t, erc20Params.RegistrationFee, erc20types.DefaultRegistrationFee, "RegistrationFee should be set to DefaultRegistrationFee")
}
