package drs6

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting/exported"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	rollappparamskeeper "github.com/dymensionxyz/dymension-rdk/x/rollappparams/keeper"
	evmkeeper "github.com/evmos/evmos/v12/x/evm/keeper"
	"github.com/evmos/evmos/v12/x/vesting/types"

	"github.com/dymensionxyz/rollapp-evm/app/upgrades"
	drs5 "github.com/dymensionxyz/rollapp-evm/app/upgrades/drs-5"
)

func CreateUpgradeHandler(
	kk upgrades.UpgradeKeepers,
	mm *module.Manager,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		if err := HandleUpgrade(ctx, kk.RpKeeper, kk.EvmKeeper); err != nil {
			return nil, err
		}

		vmap, err := mm.RunMigrations(ctx, configurator, fromVM)
		if err != nil {
			return nil, fmt.Errorf("run migrations: %w", err)
		}

		if err := migrateVestingAccounts(ctx, kk.AccountKeeper); err != nil {
			return nil, fmt.Errorf("migrate vested accounts: %w", err)
		}

		return vmap, nil
	}
}

func HandleUpgrade(ctx sdk.Context, rpKeeper rollappparamskeeper.Keeper, evmKeeper *evmkeeper.Keeper) error {
	if rpKeeper.Version(ctx) < 5 {
		// first run drs-5 migration
		if err := drs5.HandleUpgrade(ctx, rpKeeper, evmKeeper); err != nil {
			return err
		}
	}
	// upgrade drs to 6
	if err := rpKeeper.SetVersion(ctx, DRS); err != nil {
		return err
	}
	return nil
}

const (
	MaxSecs      int64 = 2000000000
	millisToSecs       = 1000
)

// migrateVestingAccounts fixes the vesting accounts' start and end times to be expressed in seconds instead of milliseconds
func migrateVestingAccounts(ctx sdk.Context, ak authkeeper.AccountKeeperI) error {
	ak.IterateAccounts(ctx, func(acc authtypes.AccountI) bool {
		accVesting, ok := acc.(exported.VestingAccount)
		if !ok {
			return false
		}

		updated := false

		if accVesting.GetStartTime() > MaxSecs {
			switch av := accVesting.(type) {
			case *vestingtypes.ContinuousVestingAccount:
				av.StartTime /= millisToSecs
			case *vestingtypes.PeriodicVestingAccount:
				av.StartTime /= millisToSecs
			}
			updated = true
		}

		if accVesting.GetEndTime() > MaxSecs {
			switch av := accVesting.(type) {
			case *types.ClawbackVestingAccount:
				av.EndTime /= millisToSecs
			case *vestingtypes.ContinuousVestingAccount:
				av.EndTime /= millisToSecs
			case *vestingtypes.DelayedVestingAccount:
				av.EndTime /= millisToSecs
			case *vestingtypes.PeriodicVestingAccount:
				av.EndTime /= millisToSecs
			case *vestingtypes.PermanentLockedAccount:
				av.EndTime /= millisToSecs
			}
			updated = true
		}

		if updated {
			ak.SetAccount(ctx, acc)
		}
		return false
	})

	return nil
}
