package backend

import (
	sequencerstypes "github.com/dymensionxyz/dymension-rdk/x/sequencers/types"
)

func (m *RollAppEvmBackend) GetSequencersModuleParams() (*sequencerstypes.Params, error) {
	res, err := m.queryClient.SequencersQueryClient.Params(m.ctx, &sequencerstypes.QueryParamsRequest{})
	if err != nil {
		return nil, err
	}
	return &res.Params, nil
}
