package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	ibcante "github.com/cosmos/ibc-go/v6/modules/core/ante"
	ethante "github.com/evmos/ethermint/app/ante"
)

// NOTE: this function is copied from ethermint
func newEthAnteHandler(options HandlerOptions) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		ethante.NewEthSetUpContextDecorator(options.EvmKeeper),                         // outermost AnteDecorator. SetUpContext must be called first
		ethante.NewEthMempoolFeeDecorator(options.EvmKeeper),                           // Check eth effective gas price against minimal-gas-prices
		ethante.NewEthMinGasPriceDecorator(options.FeeMarketKeeper, options.EvmKeeper), // Check eth effective gas price against the global MinGasPrice
		ethante.NewEthValidateBasicDecorator(options.EvmKeeper),
		ethante.NewEthSigVerificationDecorator(options.EvmKeeper),
		ethante.NewEthAccountVerificationDecorator(options.AccountKeeper, options.EvmKeeper),
		ethante.NewCanTransferDecorator(options.EvmKeeper),
		ethante.NewVirtualFrontierContractDecorator(options.EvmKeeper), // prevent transfer to virtual frontier contract
		ethante.NewEthGasConsumeDecorator(options.EvmKeeper, options.MaxTxGasWanted),
		ethante.NewEthIncrementSenderSequenceDecorator(options.AccountKeeper), // innermost AnteDecorator.
		ethante.NewGasWantedDecorator(options.EvmKeeper, options.FeeMarketKeeper),
		ethante.NewEthEmitEventDecorator(options.EvmKeeper), // emit eth tx hash and index at the very last ante handler.
	)
}

// NOTE: this function is copied from ethermint AND modified by us
func newCosmosAnteHandler(options HandlerOptions) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		ethante.RejectMessagesDecorator{}, // reject MsgEthereumTxs
		// disable the Msg types that cannot be included on an authz.MsgExec msgs field
		ethante.NewAuthzLimiterDecorator(options.DisabledAuthzMsgs),
		ante.NewSetUpContextDecorator(),
		ante.NewExtensionOptionsDecorator(options.ExtensionOptionChecker),
		ante.NewValidateBasicDecorator(),
		ante.NewTxTimeoutHeightDecorator(),
		ethante.NewMinGasPriceDecorator(options.FeeMarketKeeper, options.EvmKeeper),
		ante.NewValidateMemoDecorator(options.AccountKeeper),
		ante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		ante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper, options.TxFeeChecker),
		// SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewSetPubKeyDecorator(options.AccountKeeper),
		ante.NewValidateSigCountDecorator(options.AccountKeeper),
		ante.NewSigGasConsumeDecorator(options.AccountKeeper, options.SigGasConsumer),
		ante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		ante.NewIncrementSequenceDecorator(options.AccountKeeper),
		ibcante.NewRedundantRelayDecorator(options.IBCKeeper),
		ethante.NewGasWantedDecorator(options.EvmKeeper, options.FeeMarketKeeper),
	)
}

// Deprecated: NewLegacyCosmosAnteHandlerEip712 creates an AnteHandler to process legacy EIP-712
// transactions, as defined by the presence of an ExtensionOptionsWeb3Tx extension.
// NOTE: this function is copied from ethermint
func newLegacyCosmosAnteHandlerEip712(options HandlerOptions) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		ethante.RejectMessagesDecorator{}, // reject MsgEthereumTxs
		// disable the Msg types that cannot be included on an authz.MsgExec msgs field
		ethante.NewAuthzLimiterDecorator(options.DisabledAuthzMsgs),
		authante.NewSetUpContextDecorator(),
		authante.NewValidateBasicDecorator(),
		authante.NewTxTimeoutHeightDecorator(),
		ethante.NewMinGasPriceDecorator(options.FeeMarketKeeper, options.EvmKeeper),
		authante.NewValidateMemoDecorator(options.AccountKeeper),
		authante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		authante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper, options.TxFeeChecker),
		// SetPubKeyDecorator must be called before all signature verification decorators
		authante.NewSetPubKeyDecorator(options.AccountKeeper),
		authante.NewValidateSigCountDecorator(options.AccountKeeper),
		authante.NewSigGasConsumeDecorator(options.AccountKeeper, options.SigGasConsumer),
		// Note: signature verification uses EIP instead of the cosmos signature validator
		ethante.NewLegacyEip712SigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		authante.NewIncrementSequenceDecorator(options.AccountKeeper),
		ibcante.NewRedundantRelayDecorator(options.IBCKeeper),
		ethante.NewGasWantedDecorator(options.EvmKeeper, options.FeeMarketKeeper),
	)
}
