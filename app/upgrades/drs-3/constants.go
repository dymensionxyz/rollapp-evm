package drs3

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	"github.com/dymensionxyz/rollapp-evm/app/upgrades"
)

const (
	UpgradeName = "drs-3"
)

var Upgrade = upgrades.Upgrade{
	Name:          UpgradeName,
	CreateHandler: CreateUpgradeHandler,
	StoreUpgrades: storetypes.StoreUpgrades{},
}
