package drs2

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/dymensionxyz/rollapp-evm/app/upgrades"
	claimstypes "github.com/evmos/evmos/v12/x/claims/types"
)

const (
	UpgradeName = "drs-2"
)

var Upgrade = upgrades.Upgrade{
	Name:          UpgradeName,
	CreateHandler: CreateUpgradeHandler,
	StoreUpgrades: storetypes.StoreUpgrades{
		Deleted: []string{claimstypes.ModuleName},
	},
}
