package drs7

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	dividendstypes "github.com/dymensionxyz/dymension-rdk/x/dividends/types"
	"github.com/dymensionxyz/rollapp-evm/app/upgrades"
)

const (
	UpgradeName        = "drs-7"
	DRS         uint32 = 7
)

var Upgrade = upgrades.Upgrade{
	Name:          UpgradeName,
	CreateHandler: CreateUpgradeHandler,
	StoreUpgrades: storetypes.StoreUpgrades{
		Added: []string{dividendstypes.StoreKey},
	},
}
