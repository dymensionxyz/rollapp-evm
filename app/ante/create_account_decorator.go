package ante

import (
	errorsmod "cosmossdk.io/errors"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

const (
	// MaxFreeAccountsPerBlock is the maximum number of free accounts that can be created per block
	MaxFreeAccountsPerBlock = 10

	// freeAccountCountKey is the key used to track the number of free accounts created in the current block
	freeAccountCountKey = "free_account_count"
)

type CreateAccountDecorator struct {
	ak                    accountKeeper
	anteTransientStoreKey storetypes.StoreKey
}

type accountKeeper interface {
	authante.AccountKeeper
	NewAccountWithAddress(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
}

func NewCreateAccountDecorator(ak accountKeeper, anteTransientStoreKey storetypes.StoreKey) CreateAccountDecorator {
	return CreateAccountDecorator{
		ak:                    ak,
		anteTransientStoreKey: anteTransientStoreKey,
	}
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
				// Check rate limit: max free accounts per block
				if err := cad.checkAndIncrementFreeAccountCount(ctx); err != nil {
					return ctx, err
				}

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

// checkAndIncrementFreeAccountCount checks if the free account creation limit has been exceeded
// and increments the counter if not. Returns an error if the limit is exceeded.
func (cad CreateAccountDecorator) checkAndIncrementFreeAccountCount(ctx sdk.Context) error {
	store := ctx.TransientStore(cad.anteTransientStoreKey)
	key := []byte(freeAccountCountKey)

	countBz := store.Get(key)
	count := uint64(0)
	if countBz != nil {
		count = sdk.BigEndianToUint64(countBz)
	}

	if count >= MaxFreeAccountsPerBlock {
		return errorsmod.Wrapf(
			sdkerrors.ErrInvalidRequest,
			"exceeded maximum free account creations per block (%d)", MaxFreeAccountsPerBlock,
		)
	}

	// Increment the counter
	store.Set(key, sdk.Uint64ToBigEndian(count+1))
	return nil
}
