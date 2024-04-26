package backend

import (
	hubgentypes "github.com/dymensionxyz/dymension-rdk/x/hub-genesis/types"
)

func (m *RollAppEvmBackend) GetHubGenesisModuleParams() (*hubgentypes.Params, error) {
	res, err := m.queryClient.HubGenesisQueryClient.Params(m.ctx, &hubgentypes.QueryParamsRequest{})
	if err != nil {
		return nil, err
	}
	return &res.Params, nil
}
