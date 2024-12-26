package drs2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	evmkeeper "github.com/evmos/evmos/v12/x/evm/keeper"

	rollappparamskeeper "github.com/dymensionxyz/dymension-rdk/x/rollappparams/keeper"
	rollappparamstypes "github.com/dymensionxyz/dymension-rdk/x/rollappparams/types"
)

func CreateUpgradeHandler(
	rpKeeper rollappparamskeeper.Keeper,
	evmKeeper *evmkeeper.Keeper,
	mm *module.Manager,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		if err := HandleUpgrade(ctx, rpKeeper, evmKeeper); err != nil {
			return nil, err
		}
		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}

func HandleUpgrade(
	ctx sdk.Context,
	rpKeeper rollappparamskeeper.Keeper,
	evmKeeper *evmkeeper.Keeper,
) error {
	// upgrade drs to 2
	if err := rpKeeper.SetVersion(ctx, uint32(2)); err != nil {
		return err
	}

	// migrate rollapp params with missing min-gas-prices
	if err := rpKeeper.SetMinGasPrices(ctx, rollappparamstypes.DefaultParams().MinGasPrices); err != nil {
		return err
	}

	// migrate evm params with missing gasDenom
	evmParams := evmKeeper.GetParams(ctx)
	evmParams.GasDenom = evmParams.EvmDenom

	if err := evmKeeper.SetParams(ctx, evmParams); err != nil {
		return err
	}

	return nil
}
