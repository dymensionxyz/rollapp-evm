package ante

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	seqkeeper "github.com/dymensionxyz/dymension-rdk/x/sequencers/keeper"
)

// PermissionedVestingDecorator prevents invalid msg types from being executed
type PermissionedVestingDecorator struct {
	sequencerKeeper     seqkeeper.Keeper
	disabledMsgTypeURLs []string
}

func NewPermissionedVestingDecorator(sequencerKeeper seqkeeper.Keeper, msgTypeURLs []string) PermissionedVestingDecorator {
	return PermissionedVestingDecorator{
		sequencerKeeper:     sequencerKeeper,
		disabledMsgTypeURLs: msgTypeURLs,
	}
}

// AnteHandle rejects vesting messages that signer does not have permissions
// to create vesting account.
func (pvd PermissionedVestingDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	for _, msg := range tx.GetMsgs() {
		typeURL := sdk.MsgTypeURL(msg)
		for _, disabledTypeURL := range pvd.disabledMsgTypeURLs {
			if typeURL == disabledTypeURL {
				// Check if vesting tx signer is 1
				if len(msg.GetSigners()) != 1 {
					return ctx, errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signers: %v", msg.GetSigners())
				}

				signer := msg.GetSigners()[0]
				if !pvd.sequencerKeeper.HasPermission(ctx, signer, vestingtypes.ModuleName) {
					return ctx, sdkerrors.ErrUnauthorized
				}
			}
		}
	}
	return next(ctx, tx, simulate)
}
