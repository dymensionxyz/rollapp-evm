package app

import (
	"encoding/json"
	"testing"
	"time"

	version "github.com/dymensionxyz/dymint/version"

	appcodec "github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/testutil/mock"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/tendermint/tendermint/crypto/encoding"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	types2 "github.com/tendermint/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/dymensionxyz/dymension-rdk/testutil/utils"
	rollappparamstypes "github.com/dymensionxyz/dymension-rdk/x/rollappparams/types"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
)

const TestChainID = "testchain_9000-1"

func setup(withGenesis bool, invCheckPeriod uint) (*App, GenesisState) {
	db := dbm.NewMemDB()

	app := NewRollapp(
		log.NewNopLogger(),
		db,
		nil,
		true,
		map[int64]bool{},
		DefaultNodeHome,
		invCheckPeriod,
		MakeEncodingConfig(),
		simapp.EmptyAppOptions{},
	)

	if withGenesis {
		return app, NewDefaultGenesisState(app.appCodec)
	}

	return app, GenesisState{}
}

// SetupWithOneValidator initializes a new App. A Nop logger is set in App.
func SetupWithOneValidator(t *testing.T) (*App, authtypes.AccountI) {
	t.Helper()

	privVal := mock.NewPV()
	pubKey, err := privVal.GetPubKey()
	require.NoError(t, err)

	// create validator set with single validator
	validator := types2.NewValidator(pubKey, 1)
	valSet := types2.NewValidatorSet([]*types2.Validator{validator})

	// generate genesis account
	senderPrivKey := secp256k1.GenPrivKey()
	acc := authtypes.NewBaseAccount(senderPrivKey.PubKey().Address().Bytes(), senderPrivKey.PubKey(), 0, 0)
	balance := banktypes.Balance{
		Address: acc.GetAddress().String(),
		Coins:   sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100000000000000))),
	}

	app := SetupWithGenesisValSet(t, valSet, []authtypes.GenesisAccount{acc}, balance)

	return app, acc
}

// SetupWithGenesisValSet initializes a new SimApp with a validator set and genesis accounts
// that also act as delegators. For simplicity, each validator is bonded with a delegation
// of one consensus engine unit in the default token of the simapp from first genesis
// account. A Nop logger is set in SimApp.
func SetupWithGenesisValSet(t *testing.T, valSet *types2.ValidatorSet, genAccs []authtypes.GenesisAccount, balances ...banktypes.Balance) *App {
	t.Helper()

	app, genesisState := setup(true, 5)
	genesisState = genesisStateWithValSet(t, app, genesisState, valSet, genAccs, balances...)

	version.DRS = "1"
	genesisState = setRollappVersion(app.appCodec, genesisState, 1)
	denomMD := banktypes.Metadata{
		Description: "Stake token",
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    "stake",
				Exponent: 18,
			},
		},
		Base:    "stake",
		Display: "stake",
	}

	genesisState = addDenomToBankModule(app.appCodec, genesisState, denomMD)

	stateBytes, err := json.MarshalIndent(genesisState, "", " ")
	require.NoError(t, err)

	proto, err := encoding.PubKeyToProto(valSet.Validators[0].PubKey)
	require.NoError(t, err)

	// init chain will set the validator set and initialize the genesis accounts
	app.InitChain(
		abci.RequestInitChain{
			Validators: []abci.ValidatorUpdate{
				{
					PubKey: proto,
					Power:  valSet.Validators[0].VotingPower,
				},
			},
			ConsensusParams: utils.DefaultConsensusParams,
			AppStateBytes:   stateBytes,
			ChainId:         TestChainID,
			GenesisChecksum: "abcdef",
		},
	)

	// commit genesis changes
	app.Commit()
	app.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{
		ChainID:            TestChainID,
		Height:             app.LastBlockHeight() + 1,
		AppHash:            app.LastCommitID().Hash,
		ValidatorsHash:     valSet.Hash(),
		NextValidatorsHash: valSet.Hash(),
	}})

	return app
}

func genesisStateWithValSet(t *testing.T,
	app *App, genesisState GenesisState,
	valSet *types2.ValidatorSet, genAccs []authtypes.GenesisAccount,
	balances ...banktypes.Balance,
) GenesisState {
	// set genesis accounts
	authGenesis := authtypes.NewGenesisState(authtypes.DefaultParams(), genAccs)
	genesisState[authtypes.ModuleName] = app.AppCodec().MustMarshalJSON(authGenesis)

	validators := make([]stakingtypes.Validator, 0, len(valSet.Validators))
	delegations := make([]stakingtypes.Delegation, 0, len(valSet.Validators))

	bondAmt := sdk.DefaultPowerReduction

	for _, val := range valSet.Validators {
		pk, err := codec.FromTmPubKeyInterface(val.PubKey)
		require.NoError(t, err)
		pkAny, err := codectypes.NewAnyWithValue(pk)
		require.NoError(t, err)
		validator := stakingtypes.Validator{
			OperatorAddress:   sdk.ValAddress(val.Address).String(),
			ConsensusPubkey:   pkAny,
			Jailed:            false,
			Status:            stakingtypes.Bonded,
			Tokens:            bondAmt,
			DelegatorShares:   sdk.OneDec(),
			Description:       stakingtypes.Description{},
			UnbondingHeight:   int64(0),
			UnbondingTime:     time.Unix(0, 0).UTC(),
			Commission:        stakingtypes.NewCommission(sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec()),
			MinSelfDelegation: sdk.ZeroInt(),
		}
		validators = append(validators, validator)
		delegations = append(delegations, stakingtypes.NewDelegation(genAccs[0].GetAddress(), val.Address.Bytes(), sdk.OneDec()))

	}
	// set validators and delegations
	stakingGenesis := stakingtypes.NewGenesisState(stakingtypes.DefaultParams(), validators, delegations)
	genesisState[stakingtypes.ModuleName] = app.AppCodec().MustMarshalJSON(stakingGenesis)

	totalSupply := sdk.NewCoins()
	for _, b := range balances {
		// add genesis acc tokens to total supply
		totalSupply = totalSupply.Add(b.Coins...)
	}

	for range delegations {
		// add delegated tokens to total supply
		totalSupply = totalSupply.Add(sdk.NewCoin(sdk.DefaultBondDenom, bondAmt))
	}

	// add bonded amount to bonded pool module account
	balances = append(balances, banktypes.Balance{
		Address: authtypes.NewModuleAddress(stakingtypes.BondedPoolName).String(),
		Coins:   sdk.Coins{sdk.NewCoin(sdk.DefaultBondDenom, bondAmt)},
	})

	// update total supply
	bankGenesis := banktypes.NewGenesisState(banktypes.DefaultGenesisState().Params, balances, totalSupply, []banktypes.Metadata{})
	genesisState[banktypes.ModuleName] = app.AppCodec().MustMarshalJSON(bankGenesis)

	return genesisState
}

func setRollappVersion(appCodec appcodec.Codec, genesisState GenesisState, version uint32) GenesisState {
	var rollappParamsGenesis rollappparamstypes.GenesisState
	if genesisState["rollappparams"] != nil {
		appCodec.MustUnmarshalJSON(genesisState["rollappparams"], &rollappParamsGenesis)
	} else {
		rollappParamsGenesis = rollappparamstypes.GenesisState{
			Params: rollappparamstypes.Params{
				DrsVersion: version,
			},
		}
	}

	rollappParamsGenesis.Params.DrsVersion = version

	genesisState["rollappparams"] = appCodec.MustMarshalJSON(&rollappParamsGenesis)

	return genesisState
}

func addDenomToBankModule(appCodec appcodec.Codec, genesisState GenesisState, denomMD banktypes.Metadata) GenesisState {
	var bankGenesis banktypes.GenesisState
	if genesisState["bank"] != nil {
		appCodec.MustUnmarshalJSON(genesisState["bank"], &bankGenesis)
	} else {
		bankGenesis = *banktypes.DefaultGenesisState()
	}

	bankGenesis.DenomMetadata = append(bankGenesis.DenomMetadata, denomMD)

	genesisState["bank"] = appCodec.MustMarshalJSON(&bankGenesis)

	return genesisState
}
