package ante

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type anteHandler interface {
	AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error)
}

type BypassIBCFeeDecorator struct {
	nextAnte anteHandler
}

func NewBypassIBCFeeDecorator(nextAnte anteHandler) BypassIBCFeeDecorator {
	return BypassIBCFeeDecorator{nextAnte: nextAnte}
}

// SKIP FEE DEDUCT and MIN GAS PRICE Ante handlers for IBC relayer messages
func (n BypassIBCFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	// ======== HACK ================
	if isIBCRelayerMsg(tx.GetMsgs()) {
		return next(ctx, tx, simulate)
	}
	// ==============================

	// If it's not an IBC Relayer transfer, proceed with the default fee handling
	return n.nextAnte.AnteHandle(ctx, tx, simulate, next)
}

// isIBCRelayerMsg checks if all the messages in the transaction are IBC relayer messages
func isIBCRelayerMsg(msgs []sdk.Msg) bool {
	ibcMsgTypes := map[string]struct{}{
		// IBC Client Messages
		"*clienttypes.MsgCreateClient":       {},
		"*clienttypes.MsgUpdateClient":       {},
		"*clienttypes.MsgUpgradeClient":      {},
		"*clienttypes.MsgSubmitMisbehaviour": {},
		// IBC Connection Messages
		"*conntypes.MsgConnectionOpenInit":    {},
		"*conntypes.MsgConnectionOpenTry":     {},
		"*conntypes.MsgConnectionOpenAck":     {},
		"*conntypes.MsgConnectionOpenConfirm": {},
		// IBC Channel Messages
		"*channeltypes.MsgChannelOpenInit":     {},
		"*channeltypes.MsgChannelOpenTry":      {},
		"*channeltypes.MsgChannelOpenAck":      {},
		"*channeltypes.MsgChannelOpenConfirm":  {},
		"*channeltypes.MsgChannelCloseInit":    {},
		"*channeltypes.MsgChannelCloseConfirm": {},
		// IBC Packet Messages
		"*channeltypes.MsgRecvPacket":      {},
		"*channeltypes.MsgAcknowledgement": {},
		"*channeltypes.MsgTimeout":         {},
		"*channeltypes.MsgTimeoutOnClose":  {},
	}

	for _, msg := range msgs {
		if _, ok := ibcMsgTypes[fmt.Sprintf("%T", msg)]; !ok {
			return false
		}
	}

	return true
}
