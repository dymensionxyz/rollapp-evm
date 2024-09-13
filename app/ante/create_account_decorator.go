package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

type createAccountDecorator struct {
	ak accountKeeper
}

type accountKeeper interface {
	authante.AccountKeeper
	NewAccountWithAddress(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
}

func NewCreateAccountDecorator(ak accountKeeper) createAccountDecorator {
	return createAccountDecorator{ak: ak}
}

const newAccountCtxKeyPrefix = "new-account/"

func CtxKeyNewAccount(acc string) string {
	return newAccountCtxKeyPrefix + acc
}

func (cad createAccountDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	sigTx, ok := tx.(authsigning.SigVerifiableTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "invalid tx type")
	}

	pubkeys, err := sigTx.GetPubKeys()
	if err != nil {
		return ctx, err
	}

	ibcRelayerMsg := isIBCRelayerMsg(tx.GetMsgs())

	for i, pk := range pubkeys {
		if pk == nil {
			continue
		}

		_, err := authante.GetSignerAcc(ctx, cad.ak, sigTx.GetSigners()[i])
		if err != nil {
			// ======= HACK =========================
			// for IBC relayer messages, create an account if it doesn't exist
			if ibcRelayerMsg {
				address := sdk.AccAddress(pk.Address())
				acc := cad.ak.NewAccountWithAddress(ctx, address)
				// inject the new account flag into the context, in order to signal
				// the account creation to the subsequent decorators (sigchecker)
				ctx = ctx.WithValue(CtxKeyNewAccount(address.String()), struct{}{})
				cad.ak.SetAccount(ctx, acc)
			} else {
				return ctx, err
			}
			// ======================================
		}
	}

	return next(ctx, tx, simulate)
}
