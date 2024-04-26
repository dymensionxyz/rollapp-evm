package ante

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
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

func newCosmosAnteHandler(options HandlerOptions) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		cosmosDecorators(
			options,
			options.ExtensionOptionChecker, // make sure there are no extension options
			ante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler), // Use modern signature verification
		)...,
	)
}

// Deprecated: NewLegacyCosmosAnteHandlerEip712 creates an AnteHandler to process legacy EIP-712
// transactions, as defined by the presence of an ExtensionOptionsWeb3Tx extension.
func newLegacyCosmosAnteHandlerEip712(options HandlerOptions) sdk.AnteHandler {
	// TODO: do we need anything extra? See hub decorators https://github.com/dymensionxyz/dymension/blob/7b27f5ff6c7ae499bac708a3a1d5975686e54dd7/app/ante/handlers.go#L38-L39
	return sdk.ChainAnteDecorators(
		cosmosDecorators(
			options,
			func(c *codectypes.Any) bool {
				return true
			},
			ethante.NewLegacyEip712SigVerificationDecorator(options.AccountKeeper, options.SignModeHandler), // Use old signature verification: uses EIP instead of the cosmos signature validator
		)...,
	)
}

func cosmosDecorators(options HandlerOptions, extensionChecker authante.ExtensionOptionChecker, sigChecker sdk.AnteDecorator) []sdk.AnteDecorator {
	sigGasConsumer := options.SigGasConsumer
	if sigGasConsumer == nil {
		sigGasConsumer = authante.DefaultSigVerificationGasConsumer
	}
	// TODO: do we need anything extra? See hub decorators
	//  https://github.com/dymensionxyz/dymension/blob/7b27f5ff6c7ae499bac708a3a1d5975686e54dd7/app/ante/handlers.go#L79-L80
	//  https://github.com/dymensionxyz/dymension/blob/7b27f5ff6c7ae499bac708a3a1d5975686e54dd7/app/ante/handlers.go#L38-L39
	return []sdk.AnteDecorator{
		ethante.RejectMessagesDecorator{}, // reject MsgEthereumTxs
		// disable the Msg types that cannot be included on an authz.MsgExec msgs field
		ethante.NewAuthzLimiterDecorator(options.DisabledAuthzMsgs),
		ante.NewSetUpContextDecorator(),
		ante.NewExtensionOptionsDecorator(extensionChecker),
		NewPermissionedURLsDecorator(
			func(ctx sdk.Context, accAddr sdk.AccAddress) bool {
				return options.hasPermission(ctx, accAddr, vestingtypes.ModuleName)
			}, []string{ // TODO: can it go here?
				sdk.MsgTypeURL(&vestingtypes.MsgCreateVestingAccount{}),
				sdk.MsgTypeURL(&vestingtypes.MsgCreatePeriodicVestingAccount{}),
				sdk.MsgTypeURL(&vestingtypes.MsgCreatePermanentLockedAccount{}),
			}),
		ante.NewValidateBasicDecorator(),
		ante.NewTxTimeoutHeightDecorator(),
		ethante.NewMinGasPriceDecorator(options.FeeMarketKeeper, options.EvmKeeper),
		ante.NewValidateMemoDecorator(options.AccountKeeper),
		ante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		ante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper, options.TxFeeChecker),
		// SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewSetPubKeyDecorator(options.AccountKeeper),
		ante.NewValidateSigCountDecorator(options.AccountKeeper),
		ante.NewSigGasConsumeDecorator(options.AccountKeeper, sigGasConsumer),
		sigChecker,
		ante.NewIncrementSequenceDecorator(options.AccountKeeper),
		ibcante.NewRedundantRelayDecorator(options.IBCKeeper),
		ethante.NewGasWantedDecorator(options.EvmKeeper, options.FeeMarketKeeper),
	}
}
