package ante

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	ibcante "github.com/cosmos/ibc-go/v6/modules/core/ante"
	evmosante "github.com/evmos/evmos/v12/app/ante"
	cosmosante "github.com/evmos/evmos/v12/app/ante/cosmos"
	evmante "github.com/evmos/evmos/v12/app/ante/evm"
)

// NOTE: this function is copied from evmos
func newEVMAnteHandler(options HandlerOptions) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		// outermost AnteDecorator. SetUpContext must be called first
		evmante.NewEthSetUpContextDecorator(options.EvmKeeper),
		// Check eth effective gas price against the node's minimal-gas-prices config
		evmante.NewEthMempoolFeeDecorator(options.EvmKeeper),
		// Check eth effective gas price against the global MinGasPrice
		evmante.NewEthMinGasPriceDecorator(options.FeeMarketKeeper, options.EvmKeeper),
		evmante.NewEthValidateBasicDecorator(options.EvmKeeper),
		evmante.NewEthSigVerificationDecorator(options.EvmKeeper),
		evmante.NewEthAccountVerificationDecorator(options.AccountKeeper, options.EvmKeeper),
		evmante.NewCanTransferDecorator(options.EvmKeeper),
		evmante.NewEthVestingTransactionDecorator(options.AccountKeeper, options.BankKeeper, options.EvmKeeper),
		evmante.NewEthGasConsumeDecorator(options.BankKeeper, options.DistributionKeeper, options.EvmKeeper, options.StakingKeeper, options.MaxTxGasWanted),
		evmante.NewEthIncrementSenderSequenceDecorator(options.AccountKeeper),
		evmante.NewGasWantedDecorator(options.EvmKeeper, options.FeeMarketKeeper),
		// emit eth tx hash and index at the very last ante handler.
		evmante.NewEthEmitEventDecorator(options.EvmKeeper),
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
	return sdk.ChainAnteDecorators(
		cosmosDecorators(
			options,
			func(c *codectypes.Any) bool { return true },                                                       // allows any extension, this is consistent with the evmos behavior
			cosmosante.NewLegacyEip712SigVerificationDecorator(options.AccountKeeper, options.SignModeHandler), // Use old signature verification: uses EIP instead of the cosmos signature validator
		)...,
	)
}

func cosmosDecorators(options HandlerOptions, extensionChecker authante.ExtensionOptionChecker, sigChecker sdk.AnteDecorator) []sdk.AnteDecorator {
	sigGasConsumer := options.SigGasConsumer
	if sigGasConsumer == nil {
		sigGasConsumer = authante.DefaultSigVerificationGasConsumer
	}
	return []sdk.AnteDecorator{
		cosmosante.NewRejectMessagesDecorator(evmosante.BlockedMessages()), // reject MsgEthereumTxs
		cosmosante.NewAuthzLimiterDecorator( // disable the Msg types that cannot be included on an authz.MsgExec msgs field
			evmosante.BlockedMessages()...,
		),
		ante.NewSetUpContextDecorator(),
		ante.NewExtensionOptionsDecorator(extensionChecker),
		ante.NewValidateBasicDecorator(),
		ante.NewTxTimeoutHeightDecorator(),
		NewPermissionedURLsDecorator(
			func(ctx sdk.Context, accAddr sdk.AccAddress) bool {
				return options.hasPermission(ctx, accAddr, vestingtypes.ModuleName)
			}, []string{
				sdk.MsgTypeURL(&vestingtypes.MsgCreateVestingAccount{}),
				sdk.MsgTypeURL(&vestingtypes.MsgCreatePeriodicVestingAccount{}),
				sdk.MsgTypeURL(&vestingtypes.MsgCreatePermanentLockedAccount{}),
			}),
		ante.NewValidateMemoDecorator(options.AccountKeeper),
		cosmosante.NewMinGasPriceDecorator(options.FeeMarketKeeper, options.EvmKeeper),
		ante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		cosmosante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.DistributionKeeper, options.FeegrantKeeper, options.StakingKeeper, options.TxFeeChecker),
		cosmosante.NewVestingDelegationDecorator(options.AccountKeeper, options.StakingKeeper, options.Cdc),
		// SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewSetPubKeyDecorator(options.AccountKeeper),
		ante.NewValidateSigCountDecorator(options.AccountKeeper),
		ante.NewSigGasConsumeDecorator(options.AccountKeeper, options.SigGasConsumer),
		sigChecker,
		ante.NewIncrementSequenceDecorator(options.AccountKeeper),
		ibcante.NewRedundantRelayDecorator(options.IBCKeeper),
		evmante.NewGasWantedDecorator(options.EvmKeeper, options.FeeMarketKeeper),
	}
}
