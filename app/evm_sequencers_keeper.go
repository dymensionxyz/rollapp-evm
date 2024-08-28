package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	seqkeeper "github.com/dymensionxyz/dymension-rdk/x/sequencers/keeper"
	evmtypes "github.com/evmos/evmos/v12/x/evm/types"
)

type EVMWrappedSeqKeeper struct {
	seqkeeper.Keeper
}

func (w EVMWrappedSeqKeeper) GetValidatorByConsAddr(ctx sdk.Context, consAddr sdk.ConsAddress) (validator stakingtypes.Validator, found bool) {
	ret, ok := w.Keeper.GetValidatorByConsAddr(ctx, consAddr)
	if !ok {
		return stakingtypes.Validator{}, true
	}
	return ret, true
}

var _ evmtypes.StakingKeeper = &EVMWrappedSeqKeeper{}
