package ante

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	ibckeeper "github.com/cosmos/ibc-go/v6/modules/core/keeper"
	ethante "github.com/evmos/ethermint/app/ante"
	ethtypes "github.com/evmos/ethermint/types"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
)

func MustCreateHandler(
	accountKeeper evmtypes.AccountKeeper,
	bankKeeper evmtypes.BankKeeper,
	iBCKeeper *ibckeeper.Keeper,
	feeMarketKeeper ethante.FeeMarketKeeper,
	evmKeeper ethante.EVMKeeper,
	feeGrantKeeper authante.FeegrantKeeper,
	txConfig client.TxConfig,
	maxGasWanted uint64,
) sdk.AnteHandler {
	options := ethante.HandlerOptions{
		AccountKeeper:          accountKeeper,
		BankKeeper:             bankKeeper,
		SignModeHandler:        txConfig.SignModeHandler(),
		EvmKeeper:              evmKeeper,
		FeegrantKeeper:         feeGrantKeeper,
		IBCKeeper:              iBCKeeper,
		FeeMarketKeeper:        feeMarketKeeper,
		SigGasConsumer:         ethante.DefaultSigVerificationGasConsumer,
		MaxTxGasWanted:         maxGasWanted,
		ExtensionOptionChecker: ethtypes.HasDynamicFeeExtensionOption,
		TxFeeChecker:           ethante.NewDynamicFeeChecker(evmKeeper),
		DisabledAuthzMsgs: []string{
			sdk.MsgTypeURL(&evmtypes.MsgEthereumTx{}),
			sdk.MsgTypeURL(&vestingtypes.MsgCreateVestingAccount{}),
			sdk.MsgTypeURL(&vestingtypes.MsgCreatePeriodicVestingAccount{}),
			sdk.MsgTypeURL(&vestingtypes.MsgCreatePermanentLockedAccount{}),
		},
	}
	handler, err := ethante.NewAnteHandler(options)
	if err != nil {
		panic(err)
	}
	return handler
}

// HandlerOptions are the options required for constructing a default SDK AnteHandler.
type HandlerOptions struct {
	ethante.HandlerOptions

	hasPermission HasPermission
	IBCKeeper     *ibckeeper.Keeper
}

func (o HandlerOptions) validate() error {
	/*
		First check the eth stuff - the validate method is not exported so this is copy-pasted
	*/
	if o.AccountKeeper == nil {
		return errorsmod.Wrap(sdkerrors.ErrLogic, "account keeper is required for AnteHandler")
	}
	if o.BankKeeper == nil {
		return errorsmod.Wrap(sdkerrors.ErrLogic, "bank keeper is required for AnteHandler")
	}
	if o.SignModeHandler == nil {
		return errorsmod.Wrap(sdkerrors.ErrLogic, "sign mode handler is required for ante builder")
	}
	if o.FeeMarketKeeper == nil {
		return errorsmod.Wrap(sdkerrors.ErrLogic, "fee market keeper is required for AnteHandler")
	}
	if o.EvmKeeper == nil {
		return errorsmod.Wrap(sdkerrors.ErrLogic, "evm keeper is required for AnteHandler")
	}

	/*
		Now our stuff
	*/
	if o.hasPermission == nil {
		return errorsmod.Wrap(sdkerrors.ErrLogic, "permission checker is required for ante builder")
	}
	if o.IBCKeeper == nil {
		return errorsmod.Wrap(sdkerrors.ErrLogic, "IBC keeper is required for ante builder")
	}
	return nil
}

func NewAnteHandler(options HandlerOptions) (sdk.AnteHandler, error) {
	if err := options.validate(); err != nil {
		return nil, err
	}

	return func(
		ctx sdk.Context, tx sdk.Tx, sim bool,
	) (newCtx sdk.Context, err error) {
		var anteHandler sdk.AnteHandler

		defer ethante.Recover(ctx.Logger(), &err)

		txWithExtensions, ok := tx.(authante.HasExtensionOptionsTx)
		if ok {
			opts := txWithExtensions.GetExtensionOptions()
			if len(opts) > 0 {
				switch typeURL := opts[0].GetTypeUrl(); typeURL {
				case "/ethermint.evm.v1.ExtensionOptionsEthereumTx":
					// handle as *evmtypes.MsgEthereumTx. It will get checked by the EVM handler to make sure it is.
					anteHandler = newEthAnteHandler(options)
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
