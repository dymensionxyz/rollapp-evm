package utils

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	ethermint "github.com/evmos/evmos/v12/types"
)

// SetBip44CoinType sets the global coin type to be used in hierarchical deterministic wallets.
func SetBip44CoinType(config *sdk.Config) {
	config.SetCoinType(ethermint.Bip44CoinType)
	config.SetPurpose(sdk.Purpose)
}
