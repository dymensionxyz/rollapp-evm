package drs5

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	rollappparamskeeper "github.com/dymensionxyz/dymension-rdk/x/rollappparams/keeper"
	evmkeeper "github.com/evmos/evmos/v12/x/evm/keeper"

	drs4 "github.com/dymensionxyz/rollapp-evm/app/upgrades/drs-4"
)

func CreateUpgradeHandler(
	rpKeeper rollappparamskeeper.Keeper,
	evmKeeper *evmkeeper.Keeper,
	mm *module.Manager,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		if err := HandleUpgrade(ctx, rpKeeper, evmKeeper); err != nil {
			return nil, err
		}
		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}

func HandleUpgrade(ctx sdk.Context, rpKeeper rollappparamskeeper.Keeper, evmKeeper *evmkeeper.Keeper) error {
	if rpKeeper.Version(ctx) < 4 {
		// first run drs-4 migration
		if err := drs4.HandleUpgrade(ctx, rpKeeper, evmKeeper); err != nil {
			return err
		}
	}
	// upgrade drs to 5
	if err := rpKeeper.SetVersion(ctx, uint32(5)); err != nil {
		return err
	}
	return nil
}
