package types

import (
	"github.com/cosmos/cosmos-sdk/types/tx"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	epochstypes "github.com/dymensionxyz/dymension-rdk/x/epochs/types"
	hubgentypes "github.com/dymensionxyz/dymension-rdk/x/hub-genesis/types"
	sequencerstypes "github.com/dymensionxyz/dymension-rdk/x/sequencers/types"
	evmtypes "github.com/evmos/evmos/v12/x/evm/types"

	"github.com/cosmos/cosmos-sdk/client"
)

// QueryClient defines a gRPC Client used for:
//   - Transaction simulation
type QueryClient struct {
	tx.ServiceClient

	BankQueryClient       banktypes.QueryClient
	EvmQueryClient        evmtypes.QueryClient
	SequencersQueryClient sequencerstypes.QueryClient
	EpochQueryClient      epochstypes.QueryClient
	HubGenesisQueryClient hubgentypes.QueryClient
}

// NewQueryClient creates a new gRPC query client
func NewQueryClient(clientCtx client.Context) *QueryClient {
	return &QueryClient{
		ServiceClient:         tx.NewServiceClient(clientCtx),
		BankQueryClient:       banktypes.NewQueryClient(clientCtx),
		EvmQueryClient:        evmtypes.NewQueryClient(clientCtx),
		SequencersQueryClient: sequencerstypes.NewQueryClient(clientCtx),
		EpochQueryClient:      epochstypes.NewQueryClient(clientCtx),
		HubGenesisQueryClient: hubgentypes.NewQueryClient(clientCtx),
	}
}
