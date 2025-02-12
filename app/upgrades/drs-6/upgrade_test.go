package drs6_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting/exported"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/dymensionxyz/rollapp-evm/app"
	up "github.com/dymensionxyz/rollapp-evm/app/upgrades/drs-6"
)

func TestBeginBlocker(t *testing.T) {
	app, _ := app.SetupWithOneValidator(t)
	app.SetDymintVersionGetter(func() (uint32, error) {
		return up.DRS, nil
	})
	h := int64(42)
	ctx := app.NewUncachedContext(true, types.Header{
		Height:  h - 1,
		ChainID: "testchain_9000-1",
	})

	// create two vesting accounts
	acc1Address := "cosmos1hd6fsrvnz6qkp87s3u86ludegq97agxsdkwzyh"
	acc2Address := "cosmos1gu6y2a0ffteesyeyeesk23082c6998xyzmt9mz"
	now := time.Now()
	nowSecs := now.Unix()
	nowMillis := now.UnixMilli()

	addr1, err := sdk.AccAddressFromBech32(acc1Address)
	require.NoError(t, err)

	addr2, err := sdk.AccAddressFromBech32(acc2Address)
	require.NoError(t, err)

	base1Acc := authtypes.NewBaseAccount(addr1, nil, 0, 0)
	base2Acc := authtypes.NewBaseAccount(addr2, nil, 0, 0)

	base1Vesting := vestingtypes.NewBaseVestingAccount(base1Acc, nil, nowSecs)   // EndTime in seconds
	base2Vesting := vestingtypes.NewBaseVestingAccount(base2Acc, nil, nowMillis) // EndTime in milliseconds

	acc1Vesting := vestingtypes.NewContinuousVestingAccountRaw(base1Vesting, nowSecs)   // StartTime in seconds
	acc2Vesting := vestingtypes.NewContinuousVestingAccountRaw(base2Vesting, nowMillis) // StartTime in milliseconds

	app.AccountKeeper.SetAccount(ctx, acc1Vesting)
	app.AccountKeeper.SetAccount(ctx, acc2Vesting)

	err = app.UpgradeKeeper.ScheduleUpgrade(ctx, upgradetypes.Plan{
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
	app.AccountKeeper.IterateAccounts(ctx, func(acc authtypes.AccountI) bool {
		accVesting, ok := acc.(exported.VestingAccount)
		if !ok {
			return false
		}

		contVesting, ok := accVesting.(*vestingtypes.ContinuousVestingAccount)
		if !ok {
			return false
		}

		require.Equalf(t, contVesting.StartTime, nowSecs, "StartTime should be exactly %d", nowSecs)
		require.Equalf(t, contVesting.EndTime, nowSecs, "EndTime should be exactly %d", nowSecs)
		return false
	})

	p := app.RollappParamsKeeper.GetParams(ctx)
	require.Equal(t, up.DRS, p.DrsVersion, "Version should be set to DRS")
}
