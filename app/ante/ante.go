package ante

import (
	"cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/client"
	types2 "github.com/cosmos/cosmos-sdk/types"
	errors2 "github.com/cosmos/cosmos-sdk/types/errors"
	ante2 "github.com/cosmos/cosmos-sdk/x/auth/ante"
	types4 "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	ante3 "github.com/cosmos/ibc-go/v6/modules/core/ante"
	"github.com/cosmos/ibc-go/v6/modules/core/keeper"
	"github.com/evmos/ethermint/app/ante"
	types3 "github.com/evmos/ethermint/types"
	"github.com/evmos/ethermint/x/evm/types"
)

func MustCreateAnteHandler(
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	iBCKeeper *keeper.Keeper,
	feeMarketKeeper ante.FeeMarketKeeper,
	evmKeeper ante.EVMKeeper,
	feeGrantKeeper ante2.FeegrantKeeper,
	txConfig client.TxConfig,
	maxGasWanted uint64,
) types2.AnteHandler {
	options := ante.HandlerOptions{
		AccountKeeper:          accountKeeper,
		BankKeeper:             bankKeeper,
		SignModeHandler:        txConfig.SignModeHandler(),
		EvmKeeper:              evmKeeper,
		FeegrantKeeper:         feeGrantKeeper,
		IBCKeeper:              iBCKeeper,
		FeeMarketKeeper:        feeMarketKeeper,
		SigGasConsumer:         ante.DefaultSigVerificationGasConsumer,
		MaxTxGasWanted:         maxGasWanted,
		ExtensionOptionChecker: types3.HasDynamicFeeExtensionOption,
		TxFeeChecker:           ante.NewDynamicFeeChecker(evmKeeper),
		DisabledAuthzMsgs: []string{
			types2.MsgTypeURL(&types.MsgEthereumTx{}),
			types2.MsgTypeURL(&types4.MsgCreateVestingAccount{}),
			types2.MsgTypeURL(&types4.MsgCreatePeriodicVestingAccount{}),
			types2.MsgTypeURL(&types4.MsgCreatePermanentLockedAccount{}),
		},
	}
	handler, err := ante.NewAnteHandler(options)
	if err != nil {
		panic(err)
	}
	return handler
}

// HandlerOptions are the options required for constructing a default SDK AnteHandler.
type HandlerOptions struct {
	ante2.HandlerOptions

	IBCKeeper *keeper.Keeper
}

func GetAnteDecorators(options HandlerOptions) []types2.AnteDecorator {
	sigGasConsumer := options.SigGasConsumer
	if sigGasConsumer == nil {
		sigGasConsumer = ante2.DefaultSigVerificationGasConsumer
	}

	anteDecorators := []types2.AnteDecorator{
		ante2.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first

		ante2.NewExtensionOptionsDecorator(options.ExtensionOptionChecker),

		ante2.NewValidateBasicDecorator(),
		ante2.NewTxTimeoutHeightDecorator(),

		ante2.NewValidateMemoDecorator(options.AccountKeeper),
		ante2.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		ante2.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper, options.TxFeeChecker),
		ante2.NewSetPubKeyDecorator(options.AccountKeeper), // SetPubKeyDecorator must be called before all signature verification decorators
		ante2.NewValidateSigCountDecorator(options.AccountKeeper),
		ante2.NewSigGasConsumeDecorator(options.AccountKeeper, sigGasConsumer),
		ante2.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		ante2.NewIncrementSequenceDecorator(options.AccountKeeper),
	}

	anteDecorators = append(anteDecorators, ante3.NewRedundantRelayDecorator(options.IBCKeeper))

	return anteDecorators
}

// NewAnteHandler returns an AnteHandler that checks and increments sequence
// numbers, checks signatures & account numbers, and deducts fees from the first
// signer.
func NewAnteHandler(options HandlerOptions) (types2.AnteHandler, error) {
	// From x/auth/ante.go
	if options.AccountKeeper == nil {
		return nil, errors.Wrap(errors2.ErrLogic, "account keeper is required for ante builder")
	}

	if options.BankKeeper == nil {
		return nil, errors.Wrap(errors2.ErrLogic, "bank keeper is required for ante builder")
	}

	if options.SignModeHandler == nil {
		return nil, errors.Wrap(errors2.ErrLogic, "sign mode handler is required for ante builder")
	}

	return types2.ChainAnteDecorators(GetAnteDecorators(options)...), nil
}
