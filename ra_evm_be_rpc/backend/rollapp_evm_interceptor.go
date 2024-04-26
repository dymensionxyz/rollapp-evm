package backend

import (
	berpcbackend "github.com/bcdevtools/block-explorer-rpc-cosmos/be_rpc/backend"
	berpctypes "github.com/bcdevtools/block-explorer-rpc-cosmos/be_rpc/types"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ berpcbackend.RequestInterceptor = (*RollAppEvmRequestInterceptor)(nil)

type RollAppEvmRequestInterceptor struct {
	beRpcBackend       berpcbackend.BackendI
	backend            RollAppEvmBackendI
	defaultInterceptor berpcbackend.RequestInterceptor
}

func NewRollAppEvmRequestInterceptor(
	beRpcBackend berpcbackend.BackendI,
	backend RollAppEvmBackendI,
	defaultInterceptor berpcbackend.RequestInterceptor,
) *RollAppEvmRequestInterceptor {
	return &RollAppEvmRequestInterceptor{
		beRpcBackend:       beRpcBackend,
		backend:            backend,
		defaultInterceptor: defaultInterceptor,
	}
}

func (m *RollAppEvmRequestInterceptor) GetTransactionByHash(hashStr string) (intercepted bool, response berpctypes.GenericBackendResponse, err error) {
	// handled completely by the default interceptor
	return m.defaultInterceptor.GetTransactionByHash(hashStr)
}

func (m *RollAppEvmRequestInterceptor) GetDenomsInformation() (intercepted, append bool, denoms map[string]string, err error) {
	// handled completely by the default interceptor
	return m.defaultInterceptor.GetDenomsInformation()
}

func (m *RollAppEvmRequestInterceptor) GetModuleParams(moduleName string) (intercepted bool, res berpctypes.GenericBackendResponse, err error) {
	var params any

	switch moduleName {
	case "sequencers":
		sequencersParams, errFetch := m.backend.GetSequencersModuleParams()
		if errFetch != nil {
			err = errors.Wrap(errFetch, "failed to get sequencers params")
		} else {
			params = *sequencersParams
		}
	case "hub-genesis":
		hubGenesisParams, errFetch := m.backend.GetHubGenesisModuleParams()
		if errFetch != nil {
			err = errors.Wrap(errFetch, "failed to get hub genesis params")
		} else {
			params = *hubGenesisParams
		}
	default:
		return m.defaultInterceptor.GetModuleParams(moduleName)
	}

	if err != nil {
		return
	}

	res, err = berpctypes.NewGenericBackendResponseFrom(params)
	if err != nil {
		err = status.Error(codes.Internal, errors.Wrap(err, "module params").Error())
		return
	}

	intercepted = true
	return
}

func (m *RollAppEvmRequestInterceptor) GetAccount(accountAddressStr string) (intercepted, append bool, response berpctypes.GenericBackendResponse, err error) {
	// handled completely by the default interceptor
	return m.defaultInterceptor.GetAccount(accountAddressStr)
}
