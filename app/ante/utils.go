package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	rdkante "github.com/dymensionxyz/dymension-rdk/server/ante"
)

// isFreeNonIBCMsg returns true for non-IBC messages that are allowed to bypass fees
// unconditionally when they are the only messages in the transaction.
func isFreeNonIBCMsg(m sdk.Msg) bool {
	switch m.(type) {
	case *authz.MsgGrant:
		return true
	case *feegrant.MsgGrantAllowance:
		return true
	default:
		return false
	}
}

// isFreeMsg returns true if the transaction contains only messages that should
// be free and allow account creation without funds.
// This includes both IBC-only messages and specific non-IBC messages like authz.MsgGrant
// and feegrant.MsgGrantAllowance.
func isFreeMsg(msgs ...sdk.Msg) bool {
	// Check if it's IBC-only messages
	if rdkante.IbcOnly(msgs...) {
		return true
	}

	// Check if all messages are free non-IBC messages
	for _, msg := range msgs {
		if !isFreeNonIBCMsg(msg) {
			return false
		}
	}

	return true
}
