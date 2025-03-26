package drs10

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	rollappparamskeeper "github.com/dymensionxyz/dymension-rdk/x/rollappparams/keeper"
	"github.com/dymensionxyz/rollapp-evm/app/upgrades"
	drs6 "github.com/dymensionxyz/rollapp-evm/app/upgrades/drs-6"
	evmkeeper "github.com/evmos/evmos/v12/x/evm/keeper"
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

		return vmap, nil
	}
}

func HandleUpgrade(ctx sdk.Context, rpKeeper rollappparamskeeper.Keeper, evmKeeper *evmkeeper.Keeper) error {
	if rpKeeper.Version(ctx) < 6 {
		// first run drs-6 migration
		if err := drs6.HandleUpgrade(ctx, rpKeeper, evmKeeper); err != nil {
			return err
		}
	}
	// upgrade drs to 7
	if err := rpKeeper.SetVersion(ctx, DRS); err != nil {
		return err
	}
	return nil
}
