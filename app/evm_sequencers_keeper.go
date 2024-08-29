package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	seqkeeper "github.com/dymensionxyz/dymension-rdk/x/sequencers/keeper"
	evmtypes "github.com/evmos/evmos/v12/x/evm/types"
)

// EVMWrappedSeqKeeper adapts values returned by sequencer keeper. It makes sure a valid validator object is always returned.
// The operator addr of the object will be used in as the EVM 'coinbase' param.
type EVMWrappedSeqKeeper struct {
	seqkeeper.Keeper
}

var _ evmtypes.StakingKeeper = &EVMWrappedSeqKeeper{}

func (w EVMWrappedSeqKeeper) GetValidatorByConsAddr(ctx sdk.Context, consAddr sdk.ConsAddress) (validator stakingtypes.Validator, found bool) {
	ret, ok := w.Keeper.GetValidatorByConsAddr(ctx, consAddr)
	if !ok {
		ret = w.generateValidator()
	}
	return ret, true
}

func (w EVMWrappedSeqKeeper) generateValidator() stakingtypes.Validator {
	pref := sdk.GetConfig().GetBech32ValidatorAddrPrefix()
	operator, _ := bech32.ConvertAndEncode(pref, []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"))
	return stakingtypes.Validator{OperatorAddress: operator}
}
