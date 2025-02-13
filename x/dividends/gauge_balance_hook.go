package dividends

import (
	"fmt"
	"strings"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	dividendskeeper "github.com/dymensionxyz/dymension-rdk/x/dividends/keeper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/evmos/v12/contracts"
	erc20keeper "github.com/evmos/evmos/v12/x/erc20/keeper"
	erc20types "github.com/evmos/evmos/v12/x/erc20/types"
)

// GetGaugeBalanceFunc returns a function that calculates the balance of a gauge address.
//  1. Iterate through all the denoms
//  2. If the denom is a ERC20 denom (starts with 'erc20' prefix), start ERC20 flow
//  3. Get the token pair of this denom from x/erc20 module
//  4. Using the token pair, get the respective ERC20 balance of the gauge address
//  5. Convert all the ERC20 tokens to cosmos address of the gauge
//  6. Return the cosmos balance of the gauge
//  7. Distribute the rewards from the gauge address. Later, users will need to convert
//     the cosmos balance to ERC20 balance after claiming the rewards
//  8. If the denom is a native denom, return the native balance of the gauge address
func GetGaugeBalanceFunc(
	erc20Keeper erc20keeper.Keeper,
	dividendsKeeper dividendskeeper.Keeper,
) dividendskeeper.GetGaugeBalanceFunc {
	return func(ctx sdk.Context, address sdk.AccAddress, denoms []string) sdk.Coins {
		for _, denom := range denoms {
			erc20Addr, err := ParseERC20Denom(denom)
			if err != nil {
				// If the denom is not an ERC20 denom, continue to the next denom
				// It's okay, no need to log this
				continue
			}

			// Get the token pair of this denom from x/erc20 module
			tokenPairID := erc20Keeper.GetTokenPairID(ctx, denom)
			tokenPair, found := erc20Keeper.GetTokenPair(ctx, tokenPairID)
			if !found {
				// If the token pair is not found, continue to the next denom
				dividendsKeeper.Logger(ctx).
					With("address", address.String()).
					With("denom", denom).
					Error("token pair not found for denom")
				continue
			}

			// Get the respective ERC20 balance of the gauge address
			erc20 := contracts.ERC20MinterBurnerDecimalsContract.ABI
			contract := tokenPair.GetERC20Contract()
			balanceToken := erc20Keeper.BalanceOf(ctx, erc20, contract, common.BytesToAddress(address.Bytes()))

			if balanceToken == nil || len(balanceToken.Bits()) == 0 {
				// If the balance is not found, continue to the next denom
				dividendsKeeper.Logger(ctx).
					With("address", address.String()).
					With("denom", denom).
					Info("gauge does not have any ERC20 tokens")
				continue
			}

			// Convert all the ERC20 tokens to cosmos address of the gauge
			err = erc20Keeper.TryConvertErc20Sdk(ctx, address, address, erc20Addr.Hex(), math.NewIntFromBigInt(balanceToken))
			if err != nil {
				// If the conversion fails, continue to the next denom
				dividendsKeeper.Logger(ctx).
					With("address", address.String()).
					With("denom", denom).
					Error("failed to convert ERC20 to cosmos")
				continue
			}

			// Now the gauge has ERC20 tokens as cosmos coins on its balance
		}

		return dividendsKeeper.GetBalanceFunc()(ctx, address, denoms)
	}
}

func ParseERC20Denom(denom string) (common.Address, error) {
	denomSplit := strings.SplitN(denom, "/", 2)

	if len(denomSplit) != 2 || denomSplit[0] != erc20types.ModuleName {
		return common.Address{}, fmt.Errorf("invalid denom %s: denomination should be prefixed with the format 'erc20/", denom)
	}

	return common.HexToAddress(denomSplit[1]), nil
}
