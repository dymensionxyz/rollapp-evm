package drs5from2d

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	authztypes "github.com/cosmos/cosmos-sdk/x/authz"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	hubgenesistypes "github.com/dymensionxyz/dymension-rdk/x/hub-genesis/types"
	hubtypes "github.com/dymensionxyz/dymension-rdk/x/hub/types"
	rollappparamstypes "github.com/dymensionxyz/dymension-rdk/x/rollappparams/types"
	timeupgradetypes "github.com/dymensionxyz/dymension-rdk/x/timeupgrade/types"
	claimstypes "github.com/evmos/evmos/v12/x/claims/types"

	"github.com/dymensionxyz/rollapp-evm/app/upgrades"
)

const (
	UpgradeName        = "drs-5-from2D"
	DRS         uint32 = 5
	DA          string = "celestia"
)

var Upgrade = upgrades.Upgrade{
	Name:          UpgradeName,
	CreateHandler: CreateUpgradeHandler,
	StoreUpgrades: storetypes.StoreUpgrades{
		Deleted: []string{
			claimstypes.ModuleName,
			"denommetadata",
		},
		Added: []string{
			authztypes.ModuleName,
			feegrant.ModuleName,
			timeupgradetypes.ModuleName,
			hubtypes.ModuleName,
			hubgenesistypes.ModuleName,
			rollappparamstypes.ModuleName,
		},
	},
}
