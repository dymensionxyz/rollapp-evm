package drs2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	rollappparamskeeper "github.com/dymensionxyz/dymension-rdk/x/rollappparams/keeper"
	rollappparamstypes "github.com/dymensionxyz/dymension-rdk/x/rollappparams/types"
	evmkeeper "github.com/evmos/evmos/v12/x/evm/keeper"
	evmtypes "github.com/evmos/evmos/v12/x/evm/types"
)

func CreateUpgradeHandler(
	rpKeeper rollappparamskeeper.Keeper,
	evmKeeper *evmkeeper.Keeper,
	mm *module.Manager,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {

		//migrate rollapp params with missing min-gas-prices and updating drs to 2
		rpKeeper.SetVersion(ctx, uint32(2))
		rpKeeper.SetMinGasPrices(ctx, rollappparamstypes.DefaultParams().MinGasPrices)

		//migrate evm params with missing gasDenom
		evmOldParams := evmKeeper.GetParams(ctx)
		evmOldParams.GasDenom = evmtypes.DefaultParams().GasDenom
		evmKeeper.SetParams(ctx, evmOldParams)

		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}
