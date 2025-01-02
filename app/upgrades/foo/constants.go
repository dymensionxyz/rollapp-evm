package drs5

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	authztypes "github.com/cosmos/cosmos-sdk/x/authz"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	hubgenesistypes "github.com/dymensionxyz/dymension-rdk/x/hub-genesis/types"
	hubtypes "github.com/dymensionxyz/dymension-rdk/x/hub/types"
	rollappparamstypes "github.com/dymensionxyz/dymension-rdk/x/rollappparams/types"
	timeupgradetypes "github.com/dymensionxyz/dymension-rdk/x/timeupgrade/types"

	"github.com/dymensionxyz/rollapp-evm/app/upgrades"
	claimstypes "github.com/evmos/evmos/v12/x/claims/types"
)

const (
	UpgradeName = "foo"
)

var Upgrade = upgrades.Upgrade{
	Name:          UpgradeName,
	CreateHandler: CreateUpgradeHandler,
	StoreUpgrades: storetypes.StoreUpgrades{
		Deleted: []string{
			claimstypes.ModuleName,
			"denommetadata", // TODO: check
		},
		Added: []string{
			authztypes.ModuleName,
			feegrant.ModuleName,
			timeupgradetypes.ModuleName,
			hubtypes.ModuleName,
			hubgenesistypes.ModuleName,
			rollappparamstypes.ModuleName,
		},
		Renamed: []storetypes.StoreRename{},
	},
}
