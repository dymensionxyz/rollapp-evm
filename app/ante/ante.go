package ante

import (
	"fmt"
	"runtime/debug"

	distrkeeper "github.com/dymensionxyz/dymension-rdk/x/dist/keeper"
	seqkeeper "github.com/dymensionxyz/dymension-rdk/x/sequencers/keeper"
	cosmosante "github.com/evmos/evmos/v12/app/ante/cosmos"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	"github.com/cosmos/cosmos-sdk/codec"

	errorsmod "cosmossdk.io/errors"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	ibckeeper "github.com/cosmos/ibc-go/v6/modules/core/keeper"
	evmosante "github.com/evmos/evmos/v12/app/ante"
	evmosanteevm "github.com/evmos/evmos/v12/app/ante/evm"
	evmostypes "github.com/evmos/evmos/v12/types"
	erc20keeper "github.com/evmos/evmos/v12/x/erc20/keeper"
	evmtypes "github.com/evmos/evmos/v12/x/evm/types"
	evmosvestingtypes "github.com/evmos/evmos/v12/x/vesting/types"
	tmlog "github.com/tendermint/tendermint/libs/log"
)

type HasPermission = func(ctx sdk.Context, accAddr sdk.AccAddress, perm string) bool

func MustCreateHandler(codec codec.BinaryCodec,
	txConfig client.TxConfig,
	maxGasWanted uint64,
	hasPermission HasPermission,
	accountKeeper evmtypes.AccountKeeper,
	stakingKeeper evmosvestingtypes.StakingKeeper,
	bankKeeper evmtypes.BankKeeper,
	feeMarketKeeper evmosanteevm.FeeMarketKeeper,
	evmKeeper evmosanteevm.EVMKeeper,
	erc20Keeper erc20keeper.Keeper,
	ibcKeeper *ibckeeper.Keeper,
	distrKeeper distrkeeper.Keeper,
	sequencerKeeper seqkeeper.Keeper,
	feeGrantKeeper authante.FeegrantKeeper,
	authzKeeper evmosanteevm.AuthzKeeper,
) sdk.AnteHandler {
	ethOpts := evmosante.HandlerOptions{
		Cdc:                codec,
		AccountKeeper:      accountKeeper,
		BankKeeper:         bankKeeper,
		EvmKeeper:          evmKeeper,
		StakingKeeper:      stakingKeeper,
		FeegrantKeeper:     feeGrantKeeper,
		DistributionKeeper: distrKeeper,
		ERC20Keeper:        erc20Keeper,
		IBCKeeper:          ibcKeeper,
		FeeMarketKeeper:    feeMarketKeeper,
		SignModeHandler:    txConfig.SignModeHandler(),
		SigGasConsumer:     evmosante.SigVerificationGasConsumer,
		MaxTxGasWanted:     maxGasWanted,
		TxFeeChecker:       evmosanteevm.NewDynamicFeeChecker(evmKeeper),
		AuthzKeeper:        authzKeeper,
	}

	opts := HandlerOptions{
		HandlerOptions:   ethOpts,
		hasPermission:    hasPermission,
		DistrKeeper:      distrKeeper,
		SequencersKeeper: sequencerKeeper,
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
	hasPermission    HasPermission
	DistrKeeper      distrkeeper.Keeper
	SequencersKeeper seqkeeper.Keeper
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
	if o.ERC20Keeper == nil {
		return errorsmod.Wrap(sdkerrors.ErrLogic, "erc20 keeper missing")
	}
	if o.DistributionKeeper == nil {
		return errorsmod.Wrap(sdkerrors.ErrLogic, "distribution keeper missing")
	}
	if o.StakingKeeper == nil {
		return errorsmod.Wrap(sdkerrors.ErrLogic, "staking keeper missing")
	}
	if o.FeegrantKeeper == nil {
		return errorsmod.Wrap(sdkerrors.ErrLogic, "feegrant keeper missing")
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
					options.ExtensionOptionChecker = func(c *codectypes.Any) bool {
						_, ok := c.GetCachedValue().(*evmostypes.ExtensionOptionsWeb3Tx)
						return ok
					}
					anteHandler = cosmosHandler(
						options,
						// nolint:staticcheck
						cosmosante.NewLegacyEip712SigVerificationDecorator(options.AccountKeeper, options.SignModeHandler), // Use old signature verification: uses EIP instead of the cosmos signature validator
					)
				case "/ethermint.types.v1.ExtensionOptionDynamicFeeTx": // TODO: can delete?
					// cosmos-sdk tx with dynamic fee extension
					options.ExtensionOptionChecker = evmostypes.HasDynamicFeeExtensionOption
					anteHandler = cosmosHandler(
						options,
						authante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler), // Use modern signature verification
					)
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
			// we reject any extension
			anteHandler = cosmosHandler(
				options,
				authante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler), // Use modern signature verification
			)
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
