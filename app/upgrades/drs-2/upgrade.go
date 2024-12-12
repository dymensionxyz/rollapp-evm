package drs2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	rollappparamskeeper "github.com/dymensionxyz/dymension-rdk/x/rollappparams/keeper"
	rollappparamstypes "github.com/dymensionxyz/dymension-rdk/x/rollappparams/types"
	evmkeeper "github.com/evmos/evmos/v12/x/evm/keeper"
)

func CreateUpgradeHandler(
	rpKeeper rollappparamskeeper.Keeper,
	evmKeeper *evmkeeper.Keeper,
	mm *module.Manager,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {

		//migrate rollapp params with missing min-gas-prices and updating drs to 2
		err := rpKeeper.SetVersion(ctx, uint32(2))
		if err != nil {
			return nil, err
		}
		err = rpKeeper.SetMinGasPrices(ctx, rollappparamstypes.DefaultParams().MinGasPrices)
		if err != nil {
			return nil, err
		}
		//migrate evm params with missing gasDenom
		evmParams := evmKeeper.GetParams(ctx)
		evmParams.GasDenom = evmParams.EvmDenom
		err = evmKeeper.SetParams(ctx, evmParams)
		if err != nil {
			return nil, err
		}
		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}
