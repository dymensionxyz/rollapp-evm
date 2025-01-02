package foo

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	hubgenkeeper "github.com/dymensionxyz/dymension-rdk/x/hub-genesis/keeper"
	"github.com/dymensionxyz/rollapp-evm/app/upgrades"
	"github.com/tendermint/tendermint/libs/log"
)

func CreateUpgradeHandler(
	kk upgrades.UpgradeKeepers,
	mm *module.Manager,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		l := ctx.Logger().With("upgrade", UpgradeName)

		migrateHubGenesis(ctx, l, kk.HubgenK)

		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}

func migrateHubGenesis(ctx sdk.Context, l log.Logger, k hubgenkeeper.Keeper) {
	
}
