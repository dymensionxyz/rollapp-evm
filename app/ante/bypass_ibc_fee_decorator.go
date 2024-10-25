package ante

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	conntypes "github.com/cosmos/ibc-go/v6/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	distrkeeper "github.com/dymensionxyz/dymension-rdk/x/dist/keeper"
	seqkeeper "github.com/dymensionxyz/dymension-rdk/x/sequencers/keeper"
)

type anteHandler interface {
	AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error)
}

type BypassIBCFeeDecorator struct {
	nextAnte anteHandler
	dk       distrkeeper.Keeper
	sk       seqkeeper.Keeper
}

func NewBypassIBCFeeDecorator(nextAnte anteHandler, dk distrkeeper.Keeper, sk seqkeeper.Keeper) BypassIBCFeeDecorator {
	return BypassIBCFeeDecorator{
		nextAnte: nextAnte,
		dk:       dk,
		sk:       sk,
	}
}

// SKIP FEE DEDUCT and MIN GAS PRICE Ante handlers for IBC relayer messages
func (n BypassIBCFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	// ================ HACK ================
	whitelisted, err := n.isWhitelistedRelayer(ctx, tx.GetMsgs())
	if err != nil {
		return ctx, fmt.Errorf("is whitelisted relayer: %w", err)
	}
	if whitelisted {
		return next(ctx, tx, simulate)
	}
	// ======================================

	// If it's not an IBC Relayer transfer, proceed with the default fee handling
	return n.nextAnte.AnteHandle(ctx, tx, simulate, next)
}

// isIBCRelayerMsg checks if all the messages in the transaction are IBC relayer messages
func isIBCRelayerMsg(msgs []sdk.Msg) bool {
	for _, msg := range msgs {
		switch msg.(type) {
		// IBC Client Messages
		case *clienttypes.MsgCreateClient, *clienttypes.MsgUpdateClient,
			*clienttypes.MsgUpgradeClient, *clienttypes.MsgSubmitMisbehaviour:
		// IBC Connection Messages
		case *conntypes.MsgConnectionOpenInit, *conntypes.MsgConnectionOpenTry,
			*conntypes.MsgConnectionOpenAck, *conntypes.MsgConnectionOpenConfirm:
		// IBC Channel Messages
		case *channeltypes.MsgChannelOpenInit, *channeltypes.MsgChannelOpenTry,
			*channeltypes.MsgChannelOpenAck, *channeltypes.MsgChannelOpenConfirm,
			*channeltypes.MsgChannelCloseInit, *channeltypes.MsgChannelCloseConfirm:
		// IBC Packet Messages
		case *channeltypes.MsgRecvPacket, *channeltypes.MsgAcknowledgement,
			*channeltypes.MsgTimeout, *channeltypes.MsgTimeoutOnClose:
		default:
			return false
		}
	}

	return true
}

// isWhitelistedRelayer checks if all the messages in the transaction are from whitelisted IBC relayer
func (n BypassIBCFeeDecorator) isWhitelistedRelayer(ctx sdk.Context, msgs []sdk.Msg) (bool, error) {
	consAddr := n.dk.GetPreviousProposerConsAddr(ctx)
	wlRelayers, err := n.sk.GetWhitelistedRelayersByConsAddr(ctx, consAddr)
	if err != nil {
		return false, fmt.Errorf("get whitelisted relayers by consensus addr: %w", err)
	}
	wlRelayersMap := make(map[string]struct{}, len(msgs))
	for _, relayerAddr := range wlRelayers.Relayers {
		wlRelayersMap[relayerAddr] = struct{}{}
	}

	for _, msg := range msgs {
		switch msg.(type) {
		// IBC Client Messages
		case *clienttypes.MsgCreateClient, *clienttypes.MsgUpdateClient,
			*clienttypes.MsgUpgradeClient, *clienttypes.MsgSubmitMisbehaviour:

		// IBC Connection Messages
		case *conntypes.MsgConnectionOpenInit, *conntypes.MsgConnectionOpenTry,
			*conntypes.MsgConnectionOpenAck, *conntypes.MsgConnectionOpenConfirm:

		// IBC Channel Messages
		case *channeltypes.MsgChannelOpenInit, *channeltypes.MsgChannelOpenTry,
			*channeltypes.MsgChannelOpenAck, *channeltypes.MsgChannelOpenConfirm,
			*channeltypes.MsgChannelCloseInit, *channeltypes.MsgChannelCloseConfirm:

		// IBC Packet Messages
		case *channeltypes.MsgRecvPacket, *channeltypes.MsgAcknowledgement,
			*channeltypes.MsgTimeout, *channeltypes.MsgTimeoutOnClose:

		default:
			return false, nil
		}

		signers := msg.GetSigners()
		for _, signer := range signers {
			_, ok := wlRelayersMap[signer.String()]
			if !ok {
				return false, nil
			}
		}
	}

	return true, nil
}
