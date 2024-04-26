package backend

import (
	"context"
	"github.com/bcdevtools/block-explorer-rpc-cosmos/be_rpc/config"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
	hubgentypes "github.com/dymensionxyz/dymension-rdk/x/hub-genesis/types"
	sequencerstypes "github.com/dymensionxyz/dymension-rdk/x/sequencers/types"
	raeberpctypes "github.com/dymensionxyz/rollapp-evm/ra_evm_be_rpc/types"
	"github.com/tendermint/tendermint/libs/log"
)

type RollAppEvmBackendI interface {
	// Misc

	GetSequencersModuleParams() (*sequencerstypes.Params, error)
	GetHubGenesisModuleParams() (*hubgentypes.Params, error)
}

var _ RollAppEvmBackendI = (*RollAppEvmBackend)(nil)

// RollAppEvmBackend implements the RollAppEvmBackendI interface
type RollAppEvmBackend struct {
	ctx         context.Context
	clientCtx   client.Context
	queryClient *raeberpctypes.QueryClient // gRPC query client
	logger      log.Logger
	cfg         config.BeJsonRpcConfig
}

// NewRollAppEvmBackend creates a new RollAppEvmBackend instance for RollApp EVM Block Explorer
func NewRollAppEvmBackend(
	ctx *server.Context,
	logger log.Logger,
	clientCtx client.Context,
) *RollAppEvmBackend {
	appConf, err := config.GetConfig(ctx.Viper)
	if err != nil {
		panic(err)
	}

	return &RollAppEvmBackend{
		ctx:         context.Background(),
		clientCtx:   clientCtx,
		queryClient: raeberpctypes.NewQueryClient(clientCtx),
		logger:      logger.With("module", "rae_be_rpc"),
		cfg:         appConf,
	}
}
