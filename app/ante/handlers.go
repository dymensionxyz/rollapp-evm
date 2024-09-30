package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	sdkvestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	conntypes "github.com/cosmos/ibc-go/v6/modules/core/03-connection/types"
	ibcante "github.com/cosmos/ibc-go/v6/modules/core/ante"
	cosmosante "github.com/evmos/evmos/v12/app/ante/cosmos"
	evmante "github.com/evmos/evmos/v12/app/ante/evm"
	evmtypes "github.com/evmos/evmos/v12/x/evm/types"
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
		// we intentionally omit the eth vesting transaction decorator
		evmante.NewEthGasConsumeDecorator(options.BankKeeper, options.DistributionKeeper, options.EvmKeeper, options.StakingKeeper, options.MaxTxGasWanted),
		evmante.NewEthIncrementSenderSequenceDecorator(options.AccountKeeper),
		evmante.NewGasWantedDecorator(options.EvmKeeper, options.FeeMarketKeeper),
		// emit eth tx hash and index at the very last ante handler.
		evmante.NewEthEmitEventDecorator(options.EvmKeeper),
	)
}

func cosmosHandler(options HandlerOptions, sigChecker sdk.AnteDecorator) sdk.AnteHandler {
	sigGasConsumer := options.SigGasConsumer
	if sigGasConsumer == nil {
		sigGasConsumer = authante.DefaultSigVerificationGasConsumer
	}
	// only override the modern sig checker, and preserve the legacy one
	if _, ok := sigChecker.(authante.SigVerificationDecorator); ok {
		sigChecker = NewSigCheckDecorator(options.AccountKeeper, options.SignModeHandler)
	}
	return sdk.ChainAnteDecorators(
		cosmosante.NewRejectMessagesDecorator(
			[]string{
				sdk.MsgTypeURL(&evmtypes.MsgEthereumTx{}),
				sdk.MsgTypeURL(&conntypes.MsgConnectionOpenInit{}), // don't let any connection open from the Rollapp side (it's still possible from the other side)
			},
		),
		cosmosante.NewAuthzLimiterDecorator( // disable the Msg types that cannot be included on an authz.MsgExec msgs field
			sdk.MsgTypeURL(&evmtypes.MsgEthereumTx{}),
			sdk.MsgTypeURL(&sdkvestingtypes.MsgCreateVestingAccount{}),
			sdk.MsgTypeURL(&sdkvestingtypes.MsgCreatePermanentLockedAccount{}),
			sdk.MsgTypeURL(&sdkvestingtypes.MsgCreatePeriodicVestingAccount{}),
		),
		ante.NewSetUpContextDecorator(),
		ante.NewExtensionOptionsDecorator(options.ExtensionOptionChecker),
		ante.NewValidateBasicDecorator(),
		ante.NewTxTimeoutHeightDecorator(),
		NewPermissionedURLsDecorator(
			func(ctx sdk.Context, accAddr sdk.AccAddress) bool {
				return options.hasPermission(ctx, accAddr, vestingtypes.ModuleName)
			}, []string{
				sdk.MsgTypeURL(&vestingtypes.MsgCreateVestingAccount{}),
				sdk.MsgTypeURL(&vestingtypes.MsgCreatePermanentLockedAccount{}),
				sdk.MsgTypeURL(&vestingtypes.MsgCreatePeriodicVestingAccount{}),
			}),
		ante.NewValidateMemoDecorator(options.AccountKeeper),
		NewBypassIBCFeeDecorator(cosmosante.NewMinGasPriceDecorator(options.FeeMarketKeeper, options.EvmKeeper)),
		NewCreateAccountDecorator(options.AccountKeeper),
		ante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		NewBypassIBCFeeDecorator(cosmosante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.DistributionKeeper, options.FeegrantKeeper, options.StakingKeeper, options.TxFeeChecker)),
		// SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewSetPubKeyDecorator(options.AccountKeeper),
		ante.NewValidateSigCountDecorator(options.AccountKeeper),
		ante.NewSigGasConsumeDecorator(options.AccountKeeper, sigGasConsumer),
		sigChecker,
		ante.NewIncrementSequenceDecorator(options.AccountKeeper),
		ibcante.NewRedundantRelayDecorator(options.IBCKeeper),
		evmante.NewGasWantedDecorator(options.EvmKeeper, options.FeeMarketKeeper),
	)
}
