package ante

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	clienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
)

// Mock implementations
type mockAccountKeeper struct {
	accounts map[string]types.AccountI
}

func newMockAccountKeeper() *mockAccountKeeper {
	return &mockAccountKeeper{
		accounts: make(map[string]types.AccountI),
	}
}

func (m *mockAccountKeeper) GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI {
	return m.accounts[addr.String()]
}

func (m *mockAccountKeeper) SetAccount(ctx sdk.Context, acc types.AccountI) {
	m.accounts[acc.GetAddress().String()] = acc
}

func (m *mockAccountKeeper) NewAccountWithAddress(ctx sdk.Context, addr sdk.AccAddress) types.AccountI {
	return authtypes.NewBaseAccountWithAddress(addr)
}

func (m *mockAccountKeeper) GetParams(ctx sdk.Context) types.Params {
	return types.DefaultParams()
}

func (m *mockAccountKeeper) GetModuleAddress(moduleName string) sdk.AccAddress {
	return sdk.AccAddress([]byte(moduleName))
}

// Mock transaction implementation
type mockTx struct {
	msgs    []sdk.Msg
	signers []sdk.AccAddress
	pubKeys []cryptotypes.PubKey
}

func (m mockTx) GetMsgs() []sdk.Msg {
	return m.msgs
}

func (m mockTx) GetSigners() []sdk.AccAddress {
	return m.signers
}

func (m mockTx) GetPubKeys() ([]cryptotypes.PubKey, error) {
	return m.pubKeys, nil
}

func (m mockTx) GetSignaturesV2() ([]signing.SignatureV2, error) {
	return nil, nil
}

func (m mockTx) ValidateBasic() error {
	return nil
}

// Test suite
type CreateAccountDecoratorTestSuite struct {
	suite.Suite
	ctx           sdk.Context
	accountKeeper *mockAccountKeeper
	decorator     CreateAccountDecorator
	tStoreKey     storetypes.StoreKey
}

func TestCreateAccountDecoratorTestSuite(t *testing.T) {
	suite.Run(t, new(CreateAccountDecoratorTestSuite))
}

func (suite *CreateAccountDecoratorTestSuite) SetupTest() {
	// Create a minimal context with transient store
	db := dbm.NewMemDB()
	kvStoreKey := sdk.NewKVStoreKey("test")
	suite.tStoreKey = storetypes.NewTransientStoreKey("ante_transient")

	cms := store.NewCommitMultiStore(db)
	cms.MountStoreWithDB(kvStoreKey, storetypes.StoreTypeIAVL, db)
	cms.MountStoreWithDB(suite.tStoreKey, storetypes.StoreTypeTransient, db)
	err := cms.LoadLatestVersion()
	if err != nil {
		panic(err)
	}

	suite.ctx = sdk.NewContext(cms, tmproto.Header{Height: 1}, false, log.NewNopLogger())
	suite.accountKeeper = newMockAccountKeeper()
	suite.decorator = NewCreateAccountDecorator(suite.accountKeeper, suite.tStoreKey)
}

// Helper to create a mock next handler
func (suite *CreateAccountDecoratorTestSuite) mockNextHandler(called *bool) sdk.AnteHandler {
	return func(ctx sdk.Context, tx sdk.Tx, simulate bool) (sdk.Context, error) {
		*called = true
		return ctx, nil
	}
}

// Helper to create test address and pubkey
func createTestAccount() (sdk.AccAddress, cryptotypes.PubKey) {
	privKey := secp256k1.GenPrivKey()
	pubKey := privKey.PubKey()
	addr := sdk.AccAddress(pubKey.Address())
	return addr, pubKey
}

// Test 1: Free account creation for IBC messages
func (suite *CreateAccountDecoratorTestSuite) TestCreateAccountForIBCMessage() {
	// Create a new account address that doesn't exist yet
	privKey := secp256k1.GenPrivKey()
	pubKey := privKey.PubKey()
	newAddr := sdk.AccAddress(pubKey.Address())
	require.Nil(suite.T(), suite.accountKeeper.GetAccount(suite.ctx, newAddr))

	// Create an IBC message (client update)
	ibcMsg := &clienttypes.MsgUpdateClient{
		Signer: newAddr.String(),
	}

	tx := mockTx{
		msgs:    []sdk.Msg{ibcMsg},
		signers: []sdk.AccAddress{newAddr},
		pubKeys: []cryptotypes.PubKey{pubKey},
	}

	nextCalled := false
	_, err := suite.decorator.AnteHandle(suite.ctx, tx, false, suite.mockNextHandler(&nextCalled))

	// Should succeed and create account
	require.NoError(suite.T(), err)
	require.True(suite.T(), nextCalled)

	// Account should now exist
	acc := suite.accountKeeper.GetAccount(suite.ctx, newAddr)
	require.NotNil(suite.T(), acc)
	require.Equal(suite.T(), newAddr.String(), acc.GetAddress().String())
}

// Test 2: Free account creation for authz.MsgGrant
func (suite *CreateAccountDecoratorTestSuite) TestCreateAccountForAuthzMsgGrant() {
	newAddr, pubKey := createTestAccount()
	require.Nil(suite.T(), suite.accountKeeper.GetAccount(suite.ctx, newAddr))

	granteeAddr, _ := createTestAccount()
	authzMsg := &authz.MsgGrant{
		Granter: newAddr.String(),
		Grantee: granteeAddr.String(),
	}

	tx := mockTx{
		msgs:    []sdk.Msg{authzMsg},
		signers: []sdk.AccAddress{newAddr},
		pubKeys: []cryptotypes.PubKey{pubKey},
	}

	nextCalled := false
	_, err := suite.decorator.AnteHandle(suite.ctx, tx, false, suite.mockNextHandler(&nextCalled))

	require.NoError(suite.T(), err)
	require.True(suite.T(), nextCalled)

	acc := suite.accountKeeper.GetAccount(suite.ctx, newAddr)
	require.NotNil(suite.T(), acc)
}

// Test 3: Free account creation for feegrant.MsgGrantAllowance
func (suite *CreateAccountDecoratorTestSuite) TestCreateAccountForFeegrantMsgGrantAllowance() {
	newAddr, pubKey := createTestAccount()
	require.Nil(suite.T(), suite.accountKeeper.GetAccount(suite.ctx, newAddr))

	granteeAddr, _ := createTestAccount()
	feegrantMsg := &feegrant.MsgGrantAllowance{
		Granter: newAddr.String(),
		Grantee: granteeAddr.String(),
	}

	tx := mockTx{
		msgs:    []sdk.Msg{feegrantMsg},
		signers: []sdk.AccAddress{newAddr},
		pubKeys: []cryptotypes.PubKey{pubKey},
	}

	nextCalled := false
	_, err := suite.decorator.AnteHandle(suite.ctx, tx, false, suite.mockNextHandler(&nextCalled))

	require.NoError(suite.T(), err)
	require.True(suite.T(), nextCalled)

	acc := suite.accountKeeper.GetAccount(suite.ctx, newAddr)
	require.NotNil(suite.T(), acc)
}

// Test 4: Account exists - should not create new account
func (suite *CreateAccountDecoratorTestSuite) TestExistingAccountNotRecreated() {
	existingAddr, pubKey := createTestAccount()
	existingAcc := suite.accountKeeper.NewAccountWithAddress(suite.ctx, existingAddr)
	suite.accountKeeper.SetAccount(suite.ctx, existingAcc)

	ibcMsg := &clienttypes.MsgUpdateClient{
		Signer: existingAddr.String(),
	}

	tx := mockTx{
		msgs:    []sdk.Msg{ibcMsg},
		signers: []sdk.AccAddress{existingAddr},
		pubKeys: []cryptotypes.PubKey{pubKey},
	}

	nextCalled := false
	_, err := suite.decorator.AnteHandle(suite.ctx, tx, false, suite.mockNextHandler(&nextCalled))

	require.NoError(suite.T(), err)
	require.True(suite.T(), nextCalled)
}

// Test 5: Rate limiting - exceeding max accounts per block
func (suite *CreateAccountDecoratorTestSuite) TestRateLimitingExceedsMax() {
	// Create MaxFreeAccountsPerBlock accounts successfully
	for i := 0; i < int(MaxFreeAccountsPerBlock); i++ {
		addr, pubKey := createTestAccount()

		ibcMsg := &clienttypes.MsgUpdateClient{
			Signer: addr.String(),
		}

		tx := mockTx{
			msgs:    []sdk.Msg{ibcMsg},
			signers: []sdk.AccAddress{addr},
			pubKeys: []cryptotypes.PubKey{pubKey},
		}

		nextCalled := false
		_, err := suite.decorator.AnteHandle(suite.ctx, tx, false, suite.mockNextHandler(&nextCalled))

		require.NoError(suite.T(), err, "Account %d should succeed", i)
		require.True(suite.T(), nextCalled)

		acc := suite.accountKeeper.GetAccount(suite.ctx, addr)
		require.NotNil(suite.T(), acc)
	}

	// The next account creation should fail due to rate limit
	overLimitAddr, overLimitPubKey := createTestAccount()

	ibcMsg := &clienttypes.MsgUpdateClient{
		Signer: overLimitAddr.String(),
	}

	tx := mockTx{
		msgs:    []sdk.Msg{ibcMsg},
		signers: []sdk.AccAddress{overLimitAddr},
		pubKeys: []cryptotypes.PubKey{overLimitPubKey},
	}

	nextCalled := false
	_, err := suite.decorator.AnteHandle(suite.ctx, tx, false, suite.mockNextHandler(&nextCalled))

	require.Error(suite.T(), err)
	require.Contains(suite.T(), err.Error(), "exceeded maximum free account creations per block")
	require.False(suite.T(), nextCalled)

	// Account should NOT have been created
	acc := suite.accountKeeper.GetAccount(suite.ctx, overLimitAddr)
	require.Nil(suite.T(), acc)
}

// Test 6: Cross-block rate limiting - counter resets for new block
func (suite *CreateAccountDecoratorTestSuite) TestRateLimitingResetsAcrossBlocks() {
	// Block 1: Create MaxFreeAccountsPerBlock accounts
	for i := 0; i < int(MaxFreeAccountsPerBlock); i++ {
		addr, pubKey := createTestAccount()

		ibcMsg := &clienttypes.MsgUpdateClient{
			Signer: addr.String(),
		}

		tx := mockTx{
			msgs:    []sdk.Msg{ibcMsg},
			signers: []sdk.AccAddress{addr},
			pubKeys: []cryptotypes.PubKey{pubKey},
		}

		nextCalled := false
		_, err := suite.decorator.AnteHandle(suite.ctx, tx, false, suite.mockNextHandler(&nextCalled))
		require.NoError(suite.T(), err)
	}

	// Verify we've hit the limit in block 1
	overLimitAddr1, overLimitPubKey1 := createTestAccount()
	tx1 := mockTx{
		msgs:    []sdk.Msg{&clienttypes.MsgUpdateClient{Signer: overLimitAddr1.String()}},
		signers: []sdk.AccAddress{overLimitAddr1},
		pubKeys: []cryptotypes.PubKey{overLimitPubKey1},
	}
	_, err := suite.decorator.AnteHandle(suite.ctx, tx1, false, suite.mockNextHandler(new(bool)))
	require.Error(suite.T(), err)
	require.Contains(suite.T(), err.Error(), "exceeded maximum")

	// Simulate a new block by creating a new context with a fresh transient store
	// In real blockchain, transient store is automatically cleared between blocks
	suite.SetupTest() // This recreates the context with a fresh transient store and increments height
	suite.ctx = suite.ctx.WithBlockHeight(2)

	// Block 2: Should be able to create MaxFreeAccountsPerBlock accounts again
	for i := 0; i < int(MaxFreeAccountsPerBlock); i++ {
		addr, pubKey := createTestAccount()

		ibcMsg := &clienttypes.MsgUpdateClient{
			Signer: addr.String(),
		}

		tx := mockTx{
			msgs:    []sdk.Msg{ibcMsg},
			signers: []sdk.AccAddress{addr},
			pubKeys: []cryptotypes.PubKey{pubKey},
		}

		nextCalled := false
		_, err := suite.decorator.AnteHandle(suite.ctx, tx, false, suite.mockNextHandler(&nextCalled))
		require.NoError(suite.T(), err, "Block 2: Account %d should succeed (counter should reset)", i)
	}

	// Verify we hit the limit again in block 2
	overLimitAddr2, overLimitPubKey2 := createTestAccount()
	tx2 := mockTx{
		msgs:    []sdk.Msg{&clienttypes.MsgUpdateClient{Signer: overLimitAddr2.String()}},
		signers: []sdk.AccAddress{overLimitAddr2},
		pubKeys: []cryptotypes.PubKey{overLimitPubKey2},
	}
	_, err = suite.decorator.AnteHandle(suite.ctx, tx2, false, suite.mockNextHandler(new(bool)))
	require.Error(suite.T(), err)
	require.Contains(suite.T(), err.Error(), "exceeded maximum")
}

// Test 7: Non-free message should fail if account doesn't exist
func (suite *CreateAccountDecoratorTestSuite) TestNonFreeMessageFailsForNonExistentAccount() {
	newAddr, pubKey := createTestAccount()
	require.Nil(suite.T(), suite.accountKeeper.GetAccount(suite.ctx, newAddr))

	// Use a non-free message type (e.g., MsgRevoke which is not in the free list)
	authzRevokeMsg := &authz.MsgRevoke{
		Granter: newAddr.String(),
		Grantee: sdk.AccAddress([]byte("grantee123456789")).String(),
	}

	tx := mockTx{
		msgs:    []sdk.Msg{authzRevokeMsg},
		signers: []sdk.AccAddress{newAddr},
		pubKeys: []cryptotypes.PubKey{pubKey},
	}

	nextCalled := false
	_, err := suite.decorator.AnteHandle(suite.ctx, tx, false, suite.mockNextHandler(&nextCalled))

	// Should fail because account doesn't exist and message is not free
	require.Error(suite.T(), err)
	require.False(suite.T(), nextCalled)

	// Account should NOT have been created
	acc := suite.accountKeeper.GetAccount(suite.ctx, newAddr)
	require.Nil(suite.T(), acc)
}

// Test 8: Mixed messages (free + non-free) should not create account
func (suite *CreateAccountDecoratorTestSuite) TestMixedMessagesDoNotCreateAccount() {
	newAddr, pubKey := createTestAccount()
	require.Nil(suite.T(), suite.accountKeeper.GetAccount(suite.ctx, newAddr))

	// Mix authz.MsgGrant (free) with authz.MsgRevoke (non-free)
	authzGrantMsg := &authz.MsgGrant{
		Granter: newAddr.String(),
		Grantee: sdk.AccAddress([]byte("grantee123456789")).String(),
	}
	authzRevokeMsg := &authz.MsgRevoke{
		Granter: newAddr.String(),
		Grantee: sdk.AccAddress([]byte("grantee123456789")).String(),
	}

	tx := mockTx{
		msgs:    []sdk.Msg{authzGrantMsg, authzRevokeMsg},
		signers: []sdk.AccAddress{newAddr, newAddr},
		pubKeys: []cryptotypes.PubKey{pubKey, pubKey},
	}

	nextCalled := false
	_, err := suite.decorator.AnteHandle(suite.ctx, tx, false, suite.mockNextHandler(&nextCalled))

	// Should fail because not all messages are free
	require.Error(suite.T(), err)
	require.False(suite.T(), nextCalled)

	acc := suite.accountKeeper.GetAccount(suite.ctx, newAddr)
	require.Nil(suite.T(), acc)
}

// Test 9: Context value for new account is set correctly
func (suite *CreateAccountDecoratorTestSuite) TestNewAccountContextValueSet() {
	newAddr, pubKey := createTestAccount()

	ibcMsg := &clienttypes.MsgUpdateClient{
		Signer: newAddr.String(),
	}

	tx := mockTx{
		msgs:    []sdk.Msg{ibcMsg},
		signers: []sdk.AccAddress{newAddr},
		pubKeys: []cryptotypes.PubKey{pubKey},
	}

	// Custom next handler that checks for context value
	nextHandler := func(ctx sdk.Context, tx sdk.Tx, simulate bool) (sdk.Context, error) {
		// Check that the context has the new account flag
		_, isNewAcc := ctx.Value(CtxKeyNewAccount(newAddr.String())).(struct{})
		require.True(suite.T(), isNewAcc, "New account flag should be set in context")
		return ctx, nil
	}

	_, err := suite.decorator.AnteHandle(suite.ctx, tx, false, nextHandler)
	require.NoError(suite.T(), err)
}

// Test 10: Rate limiting doesn't affect existing accounts
func (suite *CreateAccountDecoratorTestSuite) TestRateLimitingDoesNotAffectExistingAccounts() {
	// First, hit the rate limit
	for i := 0; i < int(MaxFreeAccountsPerBlock); i++ {
		addr, pubKey := createTestAccount()

		ibcMsg := &clienttypes.MsgUpdateClient{
			Signer: addr.String(),
		}

		tx := mockTx{
			msgs:    []sdk.Msg{ibcMsg},
			signers: []sdk.AccAddress{addr},
			pubKeys: []cryptotypes.PubKey{pubKey},
		}

		_, err := suite.decorator.AnteHandle(suite.ctx, tx, false, suite.mockNextHandler(new(bool)))
		require.NoError(suite.T(), err)
	}

	// Create an existing account
	existingAddr, existingPubKey := createTestAccount()
	existingAcc := suite.accountKeeper.NewAccountWithAddress(suite.ctx, existingAddr)
	suite.accountKeeper.SetAccount(suite.ctx, existingAcc)

	// Transaction from existing account should still work even though rate limit is hit
	ibcMsg := &clienttypes.MsgUpdateClient{
		Signer: existingAddr.String(),
	}

	tx := mockTx{
		msgs:    []sdk.Msg{ibcMsg},
		signers: []sdk.AccAddress{existingAddr},
		pubKeys: []cryptotypes.PubKey{existingPubKey},
	}

	nextCalled := false
	_, err := suite.decorator.AnteHandle(suite.ctx, tx, false, suite.mockNextHandler(&nextCalled))

	// Should succeed because account already exists
	require.NoError(suite.T(), err)
	require.True(suite.T(), nextCalled)
}
