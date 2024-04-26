package ante

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/client"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	ibcante "github.com/cosmos/ibc-go/v6/modules/core/ante"
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
) sdktypes.AnteHandler {
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
			sdktypes.MsgTypeURL(&evmtypes.MsgEthereumTx{}),
			sdktypes.MsgTypeURL(&vestingtypes.MsgCreateVestingAccount{}),
			sdktypes.MsgTypeURL(&vestingtypes.MsgCreatePeriodicVestingAccount{}),
			sdktypes.MsgTypeURL(&vestingtypes.MsgCreatePermanentLockedAccount{}),
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
	authante.HandlerOptions

	IBCKeeper *ibckeeper.Keeper
}

// NewHandler returns an AnteHandler that checks and increments sequence
// numbers, checks signatures & account numbers, and deducts fees from the first
// signer.
func NewHandler(options HandlerOptions) (sdktypes.AnteHandler, error) {
	// From x/auth/ante.go
	if options.AccountKeeper == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "account keeper is required for ante builder")
	}

	if options.BankKeeper == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "bank keeper is required for ante builder")
	}

	if options.SignModeHandler == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "sign mode handler is required for ante builder")
	}

	return sdktypes.ChainAnteDecorators(Decorators(options)...), nil
}

func Decorators(options HandlerOptions) []sdktypes.AnteDecorator {
	sigGasConsumer := options.SigGasConsumer
	if sigGasConsumer == nil {
		sigGasConsumer = authante.DefaultSigVerificationGasConsumer
	}

	anteDecorators := []sdktypes.AnteDecorator{
		authante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first

		authante.NewExtensionOptionsDecorator(options.ExtensionOptionChecker),

		authante.NewValidateBasicDecorator(),
		authante.NewTxTimeoutHeightDecorator(),

		authante.NewValidateMemoDecorator(options.AccountKeeper),
		authante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		authante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper, options.TxFeeChecker),
		authante.NewSetPubKeyDecorator(options.AccountKeeper), // SetPubKeyDecorator must be called before all signature verification decorators
		authante.NewValidateSigCountDecorator(options.AccountKeeper),
		authante.NewSigGasConsumeDecorator(options.AccountKeeper, sigGasConsumer),
		authante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		authante.NewIncrementSequenceDecorator(options.AccountKeeper),
	}

	anteDecorators = append(anteDecorators, ibcante.NewRedundantRelayDecorator(options.IBCKeeper))

	return anteDecorators
}
