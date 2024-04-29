package ante

import (
	"fmt"
	"runtime/debug"

	"github.com/cosmos/cosmos-sdk/codec"

	errorsmod "cosmossdk.io/errors"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	ibckeeper "github.com/cosmos/ibc-go/v6/modules/core/keeper"
	evmosante "github.com/evmos/evmos/v12/app/ante"
	evmosanteevm "github.com/evmos/evmos/v12/app/ante/evm"
	anteutils "github.com/evmos/evmos/v12/app/ante/utils"
	evmostypes "github.com/evmos/evmos/v12/types"
	evmtypes "github.com/evmos/evmos/v12/x/evm/types"
	evmosvestingtypes "github.com/evmos/evmos/v12/x/vesting/types"
	tmlog "github.com/tendermint/tendermint/libs/log"
)

type HasPermission = func(ctx sdk.Context, accAddr sdk.AccAddress, perm string) bool

func MustCreateHandler(
	codec codec.BinaryCodec,
	accountKeeper evmtypes.AccountKeeper,
	stakingKeeper evmosvestingtypes.StakingKeeper,
	bankKeeper evmtypes.BankKeeper,
	ibcKeeper *ibckeeper.Keeper,
	feeMarketKeeper evmosanteevm.FeeMarketKeeper,
	evmKeeper evmosanteevm.EVMKeeper,
	txConfig client.TxConfig,
	maxGasWanted uint64,
	hasPermission HasPermission,
	distrKeeper anteutils.DistributionKeeper,
) sdk.AnteHandler {
	ethOpts := evmosante.HandlerOptions{
		Cdc:                    codec,
		AccountKeeper:          accountKeeper,
		BankKeeper:             bankKeeper,
		ExtensionOptionChecker: evmostypes.HasDynamicFeeExtensionOption,
		EvmKeeper:              evmKeeper,
		StakingKeeper:          stakingKeeper,
		FeegrantKeeper:         nil,
		DistributionKeeper:     distrKeeper,
		IBCKeeper:              ibcKeeper,
		FeeMarketKeeper:        feeMarketKeeper,
		SignModeHandler:        txConfig.SignModeHandler(),
		SigGasConsumer:         evmosante.SigVerificationGasConsumer,
		MaxTxGasWanted:         maxGasWanted,
		TxFeeChecker:           evmosanteevm.NewDynamicFeeChecker(evmKeeper),
	}

	opts := HandlerOptions{
		HandlerOptions: ethOpts,
		hasPermission:  hasPermission,
	}

	h, err := NewHandler(opts)
	if err != nil {
		panic(fmt.Errorf("new ante handler: %w", err))
	}
	return h
}

// HandlerOptions are the options required for constructing a default SDK AnteHandler.
type HandlerOptions struct {
	evmosante.HandlerOptions
	hasPermission HasPermission
}

func (o HandlerOptions) validate() error {
	/*
		First check the eth stuff - the validate method is not exported so this is copy-pasted
	*/
	if o.AccountKeeper == nil {
		return errorsmod.Wrap(sdkerrors.ErrLogic, "account keeper missing")
	}
	if o.BankKeeper == nil {
		return errorsmod.Wrap(sdkerrors.ErrLogic, "bank keeper missing")
	}
	if o.SignModeHandler == nil {
		return errorsmod.Wrap(sdkerrors.ErrLogic, "sign mode handler missing")
	}
	if o.FeeMarketKeeper == nil {
		return errorsmod.Wrap(sdkerrors.ErrLogic, "fee market keeper missing")
	}
	if o.EvmKeeper == nil {
		return errorsmod.Wrap(sdkerrors.ErrLogic, "evm keeper missing")
	}

	/*
	 Our stuff
	*/
	if o.hasPermission == nil {
		return errorsmod.Wrap(sdkerrors.ErrLogic, "permission checker missing")
	}
	if o.IBCKeeper == nil {
		return errorsmod.Wrap(sdkerrors.ErrLogic, "IBC keeper missing")
	}
	return nil
}

func NewHandler(options HandlerOptions) (sdk.AnteHandler, error) {
	if err := options.validate(); err != nil {
		return nil, fmt.Errorf("options validate: %w", err)
	}

	return func(
		ctx sdk.Context, tx sdk.Tx, sim bool,
	) (newCtx sdk.Context, err error) {
		var anteHandler sdk.AnteHandler

		defer Recover(ctx.Logger(), &err)

		txWithExtensions, ok := tx.(authante.HasExtensionOptionsTx)
		if ok {
			opts := txWithExtensions.GetExtensionOptions()
			if len(opts) > 0 {
				switch typeURL := opts[0].GetTypeUrl(); typeURL {
				case "/ethermint.evm.v1.ExtensionOptionsEthereumTx":
					// handle as *evmtypes.MsgEthereumTx. It will get checked by the EVM handler to make sure it is.
					anteHandler = newEVMAnteHandler(options)
				case "/ethermint.types.v1.ExtensionOptionsWeb3Tx":
					// Deprecated: Handle as normal Cosmos SDK tx, except signature is checked for Legacy EIP712 representation
					anteHandler = newLegacyCosmosAnteHandlerEip712(options)
				case "/ethermint.types.v1.ExtensionOptionDynamicFeeTx": // TODO: can delete?
					// cosmos-sdk tx with dynamic fee extension
					anteHandler = newCosmosAnteHandler(options)
				default:
					return ctx, errorsmod.Wrapf(
						sdkerrors.ErrUnknownExtensionOptions,
						"rejecting tx with unsupported extension option: %s", typeURL,
					)
				}

				return anteHandler(ctx, tx, sim)
			}
		}

		// handle as totally normal Cosmos SDK tx
		switch tx.(type) {
		case sdk.Tx:
			anteHandler = newCosmosAnteHandler(options)
		default:
			return ctx, errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "invalid transaction type: %T", tx)
		}

		return anteHandler(ctx, tx, sim)
	}, nil
}

func Recover(logger tmlog.Logger, err *error) {
	if r := recover(); r != nil {
		*err = errorsmod.Wrapf(sdkerrors.ErrPanic, "%v", r)

		if e, ok := r.(error); ok {
			logger.Error(
				"ante handler panicked",
				"error", e,
				"stack trace", string(debug.Stack()),
			)
		} else {
			logger.Error(
				"ante handler panicked",
				"recover", fmt.Sprintf("%v", r),
			)
		}
	}
}
