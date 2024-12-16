package drs3

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	rollappparamskeeper "github.com/dymensionxyz/dymension-rdk/x/rollappparams/keeper"
	evmkeeper "github.com/evmos/evmos/v12/x/evm/keeper"

	drs2 "github.com/dymensionxyz/rollapp-evm/app/upgrades/drs-2"
)

func CreateUpgradeHandler(
	rpKeeper rollappparamskeeper.Keeper,
	evmKeeper *evmkeeper.Keeper,
	mm *module.Manager,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		if rpKeeper.GetParams(ctx).DrsVersion == 1 {
			// first run drs-2 migration
			if err := drs2.HandleUpgrade(ctx, rpKeeper, evmKeeper); err != nil {
				return nil, err
			}
		}
		// upgrade drs to 3
		err := rpKeeper.SetVersion(ctx, uint32(3))
		if err != nil {
			return nil, err
		}
		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}
