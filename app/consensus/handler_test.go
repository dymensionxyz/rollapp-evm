package consensus

import (
	types3 "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/dymensionxyz/dymension-rdk/x/sequencers/types"
	"github.com/gogo/protobuf/proto"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
)

// MockMessage is a mock implementation of sdk.Msg
type MockMessage struct {
	Name string
}

func (m MockMessage) Reset()                       {}
func (m MockMessage) String() string               { return m.Name }
func (m MockMessage) ProtoMessage()                {}
func (m MockMessage) ValidateBasic() error         { return nil }
func (m MockMessage) GetSigners() []sdk.AccAddress { return nil }

func TestMapAdmissionHandler(t *testing.T) {
	allowedMessages := []string{
		proto.MessageName(&types.MsgCreateSequencer{}),
		proto.MessageName(&types.MsgUpdateSequencer{}),
	}

	handler := MapAdmissionHandler(allowedMessages)

	tests := []struct {
		name    string
		message sdk.Msg
		wantErr bool
	}{
		{
			name:    "Allowed message 1",
			message: &types.MsgCreateSequencer{},
			wantErr: false,
		},
		{
			name:    "Allowed message 2",
			message: &types.MsgUpdateSequencer{},
			wantErr: false,
		},
		{
			name:    "Not allowed message",
			message: &types3.MsgSend{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := handler(sdk.Context{}, tt.message)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "is not allowed")
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
