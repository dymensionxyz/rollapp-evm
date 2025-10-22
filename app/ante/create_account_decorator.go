package ante

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

type CreateAccountDecorator struct {
	ak accountKeeper
}

type accountKeeper interface {
	authante.AccountKeeper
	NewAccountWithAddress(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
}

func NewCreateAccountDecorator(ak accountKeeper) CreateAccountDecorator {
	return CreateAccountDecorator{ak: ak}
}

const newAccountCtxKeyPrefix = "new-account/"

func CtxKeyNewAccount(acc string) string {
	return newAccountCtxKeyPrefix + acc
}

func (cad CreateAccountDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	sigTx, ok := tx.(authsigning.SigVerifiableTx)
	if !ok {
		return ctx, errorsmod.Wrap(sdkerrors.ErrTxDecode, "invalid tx type")
	}

	pubkeys, err := sigTx.GetPubKeys()
	if err != nil {
		return ctx, err
	}

	freeMsg := isFreeMsg(tx.GetMsgs()...)

	for i, pk := range pubkeys {
		if pk == nil {
			continue
		}

		_, err := authante.GetSignerAcc(ctx, cad.ak, sigTx.GetSigners()[i])
		if err != nil {
			// ======= HACK =========================
			// for free messages (IBC relayer or special non-IBC messages like authz.MsgGrant, feegrant.MsgGrantAllowance),
			// create an account if it doesn't exist
			if freeMsg {
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
