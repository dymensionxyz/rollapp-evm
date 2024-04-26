package ante

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"golang.org/x/exp/slices"
)

// PermissionedURLsDecorator prevents invalid msg types from being executed
type PermissionedURLsDecorator struct {
	hasPermission    func(ctx sdk.Context, accAddr sdk.AccAddress) bool
	permissionedURls []string
}

func NewPermissionedURLsDecorator(hasPermission func(ctx sdk.Context, accAddr sdk.AccAddress) bool, msgTypeURLs []string) PermissionedURLsDecorator {
	return PermissionedURLsDecorator{
		hasPermission:    hasPermission,
		permissionedURls: msgTypeURLs,
	}
}

// AnteHandle rejects vesting messages that signer does not have permission
func (d PermissionedURLsDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	for _, msg := range tx.GetMsgs() {
		if slices.Contains(d.permissionedURls, sdk.MsgTypeURL(msg)) {
			// Check if vesting tx signer is 1
			if len(msg.GetSigners()) != 1 {
				return ctx, errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "expect 1 signer: signers: %v", msg.GetSigners())
			}

			signer := msg.GetSigners()[0]
			if !d.hasPermission(ctx, signer) {
				return ctx, sdkerrors.ErrUnauthorized
			}
		}
	}
	return next(ctx, tx, simulate)
}
