package drs5from2d

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	hubgenkeeper "github.com/dymensionxyz/dymension-rdk/x/hub-genesis/keeper"
	hubgenesistypes "github.com/dymensionxyz/dymension-rdk/x/hub-genesis/types"
	rollappparamskeeper "github.com/dymensionxyz/dymension-rdk/x/rollappparams/keeper"
	"github.com/dymensionxyz/rollapp-evm/app/upgrades"
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

	/*
		We set PopulateGenesisInfo because the state it writes is used by GetBaseDenom


		We need to set a dummy checksum otherwise PopulateGenesisInfo complains.
		It's not actually used anywhere.
	*/
	k.SetGenesisInfo(ctx, hubgenesistypes.GenesisInfo{GenesisChecksum: "This is a placeholder - only exists for nim and mande due to migration - not real checksum"})
	return k.PopulateGenesisInfo(ctx, nil)
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
