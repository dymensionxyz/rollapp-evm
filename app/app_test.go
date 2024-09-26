package app

import (
	"fmt"
	prototypes "github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/proto/tendermint/types"
	"testing"
)

func TestBeginBlocker(t *testing.T) {
	app := Setup(t, true)
	ctx := app.NewUncachedContext(true, types.Header{
		Height:  1,
		ChainID: "testchain_9000-1",
	})

	// Simula mensajes de consenso si es necesario
	consensusMsgs := []*prototypes.Any{
		// Agrega tus mensajes aquí
	}

	req := abci.RequestBeginBlock{
		Header: types.Header{
			Height:  1,
			Time:    ctx.BlockTime(),
			ChainID: "testchain_9000-1",
		},
		LastCommitInfo:      abci.LastCommitInfo{},
		ByzantineValidators: []abci.Evidence{},
		ConsensusMessages:   consensusMsgs,
	}

	res := app.BeginBlocker(ctx, req)

	// Verifica el resultado (ajusta según tus necesidades)
	require.NotNil(t, res)
	fmt.Println("Response:", res)

	// Verifica cambios en el estado o eventos
	// Por ejemplo:
	// expectedValue := ...
	// actualValue := app.YourKeeper.GetSomething(ctx)
	// require.Equal(t, expectedValue, actualValue)
}
