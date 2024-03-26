package rae

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/server"
	raeberpcbackend "github.com/dymensionxyz/rollapp-evm/ra_evm_be_rpc/backend"
	"github.com/tendermint/tendermint/libs/log"
)

// RPC namespaces and API version
const (
	DymRollAppEvmBlockExplorerNamespace = "rae"

	ApiVersion = "1.0"
)

// API is the RollApp EVM Block Explorer JSON-RPC.
type API struct {
	ctx     *server.Context
	logger  log.Logger
	backend raeberpcbackend.RollAppEvmBackendI
}

// NewRaeAPI creates an instance of the RollApp EVM Block Explorer API.
func NewRaeAPI(
	ctx *server.Context,
	backend raeberpcbackend.RollAppEvmBackendI,
) *API {
	return &API{
		ctx:     ctx,
		logger:  ctx.Logger.With("api", "rae"),
		backend: backend,
	}
}

func (api *API) Echo(text string) string {
	api.logger.Debug("rae_echo")
	return fmt.Sprintf("hello \"%s\" from RollApp EVM Block Explorer API", text)
}
