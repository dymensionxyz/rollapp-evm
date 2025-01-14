package drs5from2d

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	hubgenkeeper "github.com/dymensionxyz/dymension-rdk/x/hub-genesis/keeper"
	hubgenesistypes "github.com/dymensionxyz/dymension-rdk/x/hub-genesis/types"
	rollappparamskeeper "github.com/dymensionxyz/dymension-rdk/x/rollappparams/keeper"
	"github.com/dymensionxyz/rollapp-evm/app/upgrades"
	erc20keeper "github.com/evmos/evmos/v12/x/erc20/keeper"
	erc20types "github.com/evmos/evmos/v12/x/erc20/types"
	evmkeeper "github.com/evmos/evmos/v12/x/evm/keeper"
)

func CreateUpgradeHandler(
	kk upgrades.UpgradeKeepers,
	mm *module.Manager,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {

		if err := migrateHubGenesis(ctx, kk.HubgenK); err != nil {
			return nil, fmt.Errorf("migrate hub genesis: %w", err)
		}

		vmap, err := mm.RunMigrations(ctx, configurator, fromVM)
		if err != nil {
			return nil, fmt.Errorf("run migrations: %w", err)
		}

		// rollappparams is a new module so needs to go after RunMigrations to go after InitGenesis
		if err := migrateRollappParams(ctx, kk.RpKeeper); err != nil {
			return nil, fmt.Errorf("migrate rollapp params: %w", err)
		}

		if err := migrateEvmosParams(ctx, kk.EvmKeeper, kk.Erc20keeper); err != nil {
			return nil, fmt.Errorf("migrate rollapp params: %w", err)
		}

		return vmap, nil
	}
}

func migrateHubGenesis(ctx sdk.Context, k hubgenkeeper.Keeper) error {
	s := k.GetState(ctx)
	s.OutboundTransfersEnabled = true
	s.InFlight = false
	s.HubPortAndChannel = &hubgenesistypes.PortAndChannel{
		Port:    "transfer",
		Channel: "channel-0",
	}
	k.SetState(ctx, s)
	// Note: we won't populate genesis info. It has difficulties, such as missing checksum and denom metadata.
	// It isn't needed for established chains.
	return nil
}

func migrateRollappParams(ctx sdk.Context, k rollappparamskeeper.Keeper) error {
	if err := k.SetVersion(ctx, DRS); err != nil {
		return fmt.Errorf("set version: %w", err)
	}
	if err := k.SetDA(ctx, DA); err != nil {
		return fmt.Errorf("set DA: %w", err)
	}
	// no need to set min gas prices, rollapp can do it when it likes
	return nil
}

// migration v12.1.6-dymension-v0.4.3 -> v12.1.6-dymension-v0.5.0-rc02
func migrateEvmosParams(ctx sdk.Context, evmK *evmkeeper.Keeper, erc20K erc20keeper.Keeper) error {

	{
		p := evmK.GetParams(ctx)
		p.GasDenom = p.EvmDenom

		if err := evmK.SetParams(ctx, p); err != nil {
			return errorsmod.Wrap(err, "evm set params")
		}
	}

	{
		p := erc20K.GetParams(ctx)
		p.RegistrationFee = erc20types.DefaultRegistrationFee
		if err := erc20K.SetParams(ctx, p); err != nil {
			return errorsmod.Wrap(err, "erc20 set params")
		}
	}

	return nil
}
