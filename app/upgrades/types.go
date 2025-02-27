package upgrades

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	hubgenkeeper "github.com/dymensionxyz/dymension-rdk/x/hub-genesis/keeper"
	rollappparamskeeper "github.com/dymensionxyz/dymension-rdk/x/rollappparams/keeper"
	erc20keeper "github.com/evmos/evmos/v12/x/erc20/keeper"
	evmkeeper "github.com/evmos/evmos/v12/x/evm/keeper"
)

type UpgradeKeepers struct {
	RpKeeper      rollappparamskeeper.Keeper
	EvmKeeper     *evmkeeper.Keeper
	Erc20keeper   erc20keeper.Keeper
	HubgenK       hubgenkeeper.Keeper
	AccountKeeper authkeeper.AccountKeeper
}

// Upgrade defines a struct containing necessary fields that a SoftwareUpgradeProposal
// must have written, in order for the state migration to go smoothly.
// An upgrade must implement this struct, and then set it in the app.go.
// The app.go will then define the handler.
type Upgrade struct {
	// Upgrade version name, for the upgrade handler, e.g. `v4`
	Name string

	// CreateHandler defines the function that creates an upgrade handler
	CreateHandler func(
		kk UpgradeKeepers,
		mm *module.Manager,
		configurator module.Configurator,
	) upgradetypes.UpgradeHandler

	// Store upgrades, should be used for any new modules introduced, new modules deleted, or store names renamed.
	StoreUpgrades storetypes.StoreUpgrades
}
