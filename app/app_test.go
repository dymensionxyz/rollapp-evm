package app

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	types2 "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/gogo/protobuf/proto"
	prototypes "github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestBeginBlocker(t *testing.T) {
	app, valAccount := SetupWithOneValidator(t)
	ctx := app.NewUncachedContext(true, types.Header{
		Height:  1,
		ChainID: "testchain_9000-1",
	})

	bankSend := &types2.MsgSend{
		FromAddress: valAccount.GetAddress().String(),
		ToAddress:   valAccount.GetAddress().String(),
		Amount:      sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(100))),
	}
	msgBz, err := proto.Marshal(bankSend)
	require.NoError(t, err)

	goodMessage := &prototypes.Any{
		TypeUrl: proto.MessageName(&types2.MsgSend{}),
		Value:   msgBz,
	}

	testCases := []struct {
		name          string
		consensusMsgs []*prototypes.Any
		expectError   bool
	}{
		{
			name: "ValidConsensusMessage",
			consensusMsgs: []*prototypes.Any{
				goodMessage,
			},
			expectError: false,
		},
		{
			name: "InvalidUnpackMessage",
			consensusMsgs: []*prototypes.Any{
				{
					TypeUrl: "/path.to.InvalidMsg",
					Value:   []byte("invalid unpack data"),
				},
			},
			expectError: true,
		},
		{
			name: "InvalidExecutionMessage",
			consensusMsgs: []*prototypes.Any{
				{
					TypeUrl: "/path.to.ExecErrorMsg",
					Value:   []byte("execution error data"),
				},
			},
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := abci.RequestBeginBlock{
				Header: types.Header{
					Height:  1,
					Time:    ctx.BlockTime(),
					ChainID: "testchain_9000-1",
				},
				LastCommitInfo:      abci.LastCommitInfo{},
				ByzantineValidators: []abci.Evidence{},
				ConsensusMessages:   tc.consensusMsgs,
			}

			res := app.BeginBlocker(ctx, req)
			require.NotNil(t, res)

			if tc.expectError {
				require.NotEmpty(t, res.ConsensusMessagesResponses)
				for _, response := range res.ConsensusMessagesResponses {
					_, isError := response.Response.(*abci.ConsensusMessageResponse_Error)
					require.True(t, isError, "Expected an error response but got a success")
				}
			} else {
				require.NotEmpty(t, res.ConsensusMessagesResponses)
				for _, response := range res.ConsensusMessagesResponses {
					_, isOk := response.Response.(*abci.ConsensusMessageResponse_Ok)
					require.True(t, isOk, "Expected a success response but got an error")
				}
			}
		})
	}
}
