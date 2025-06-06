package app

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/x/authz"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/dymensionxyz/dymension-rdk/server/proposal"
	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/mux"
	"github.com/rakyll/statik/fs"
	"github.com/spf13/cast"
	dbm "github.com/tendermint/tm-db"

	abci "github.com/tendermint/tendermint/abci/types"
	tmjson "github.com/tendermint/tendermint/libs/json"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"

	"github.com/dymensionxyz/dymension-rdk/server/consensus"

	"github.com/dymensionxyz/rollapp-evm/app/ante"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	nodeservice "github.com/cosmos/cosmos-sdk/client/grpc/node"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store/streaming"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/capability"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	feegrantmodule "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"

	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	"github.com/dymensionxyz/dymension-rdk/x/mint"
	mintkeeper "github.com/dymensionxyz/dymension-rdk/x/mint/keeper"
	minttypes "github.com/dymensionxyz/dymension-rdk/x/mint/types"

	"github.com/dymensionxyz/dymension-rdk/x/timeupgrade"
	timeupgradekeeper "github.com/dymensionxyz/dymension-rdk/x/timeupgrade/keeper"
	timeupgradetypes "github.com/dymensionxyz/dymension-rdk/x/timeupgrade/types"

	"github.com/dymensionxyz/dymension-rdk/x/epochs"
	epochskeeper "github.com/dymensionxyz/dymension-rdk/x/epochs/keeper"
	epochstypes "github.com/dymensionxyz/dymension-rdk/x/epochs/types"

	hubgenesis "github.com/dymensionxyz/dymension-rdk/x/hub-genesis"
	hubgenkeeper "github.com/dymensionxyz/dymension-rdk/x/hub-genesis/keeper"
	hubgentypes "github.com/dymensionxyz/dymension-rdk/x/hub-genesis/types"

	"github.com/dymensionxyz/dymension-rdk/x/dividends"
	dividendskeeper "github.com/dymensionxyz/dymension-rdk/x/dividends/keeper"
	dividendstypes "github.com/dymensionxyz/dymension-rdk/x/dividends/types"

	ibctransfer "github.com/cosmos/ibc-go/v6/modules/apps/transfer"
	ibctransfertypes "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v6/modules/core"
	ibcclient "github.com/cosmos/ibc-go/v6/modules/core/02-client"
	ibcclientclient "github.com/cosmos/ibc-go/v6/modules/core/02-client/client"
	ibcclienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	ibcporttypes "github.com/cosmos/ibc-go/v6/modules/core/05-port/types"
	ibchost "github.com/cosmos/ibc-go/v6/modules/core/24-host"
	ibckeeper "github.com/cosmos/ibc-go/v6/modules/core/keeper"
	ibctestingtypes "github.com/cosmos/ibc-go/v6/testing/types"

	"github.com/dymensionxyz/dymension-rdk/x/rollappparams"
	rollappparamskeeper "github.com/dymensionxyz/dymension-rdk/x/rollappparams/keeper"
	rollappparamstypes "github.com/dymensionxyz/dymension-rdk/x/rollappparams/types"

	srvflags "github.com/evmos/evmos/v12/server/flags"

	rollappevmparams "github.com/dymensionxyz/rollapp-evm/app/params"

	// unnamed import of statik for swagger UI support
	_ "github.com/cosmos/cosmos-sdk/client/docs/statik"

	"github.com/dymensionxyz/dymension-rdk/x/staking"
	stakingkeeper "github.com/dymensionxyz/dymension-rdk/x/staking/keeper"

	"github.com/dymensionxyz/dymension-rdk/x/sequencers"
	seqkeeper "github.com/dymensionxyz/dymension-rdk/x/sequencers/keeper"
	seqtypes "github.com/dymensionxyz/dymension-rdk/x/sequencers/types"

	distr "github.com/dymensionxyz/dymension-rdk/x/dist"
	distrkeeper "github.com/dymensionxyz/dymension-rdk/x/dist/keeper"

	"github.com/evmos/evmos/v12/ethereum/eip712"
	ethermint "github.com/evmos/evmos/v12/types"
	"github.com/evmos/evmos/v12/x/erc20"
	erc20client "github.com/evmos/evmos/v12/x/erc20/client"
	erc20keeper "github.com/evmos/evmos/v12/x/erc20/keeper"
	erc20types "github.com/evmos/evmos/v12/x/erc20/types"
	"github.com/evmos/evmos/v12/x/evm"
	evmkeeper "github.com/evmos/evmos/v12/x/evm/keeper"
	evmtypes "github.com/evmos/evmos/v12/x/evm/types"
	"github.com/evmos/evmos/v12/x/feemarket"
	feemarketkeeper "github.com/evmos/evmos/v12/x/feemarket/keeper"
	feemarkettypes "github.com/evmos/evmos/v12/x/feemarket/types"
	"github.com/evmos/evmos/v12/x/ibc/transfer"
	transferkeeper "github.com/evmos/evmos/v12/x/ibc/transfer/keeper"

	"github.com/dymensionxyz/dymension-rdk/x/denommetadata"
	denommetadatamoduletypes "github.com/dymensionxyz/dymension-rdk/x/denommetadata/types"
	"github.com/dymensionxyz/dymension-rdk/x/hub"
	hubkeeper "github.com/dymensionxyz/dymension-rdk/x/hub/keeper"
	hubtypes "github.com/dymensionxyz/dymension-rdk/x/hub/types"

	// Upgrade handlers

	// Force-load the tracer engines to trigger registration due to Go-Ethereum v1.10.15 changes
	_ "github.com/ethereum/go-ethereum/eth/tracers/js"
	_ "github.com/ethereum/go-ethereum/eth/tracers/native"

	dymintversion "github.com/dymensionxyz/dymint/version"

	"github.com/dymensionxyz/rollapp-evm/app/upgrades"
	drs2 "github.com/dymensionxyz/rollapp-evm/app/upgrades/drs-2"
	drs3 "github.com/dymensionxyz/rollapp-evm/app/upgrades/drs-3"
	drs4 "github.com/dymensionxyz/rollapp-evm/app/upgrades/drs-4"
	drs5 "github.com/dymensionxyz/rollapp-evm/app/upgrades/drs-5"
	drs5from2d "github.com/dymensionxyz/rollapp-evm/app/upgrades/drs-5-from2d"
	drs6 "github.com/dymensionxyz/rollapp-evm/app/upgrades/drs-6"
)

const (
	Name = "rollapp_evm"
)

var (
	AccountAddressPrefix string
	kvstorekeys          = []string{
		authtypes.StoreKey, authzkeeper.StoreKey,
		feegrant.StoreKey, banktypes.StoreKey,
		stakingtypes.StoreKey, seqtypes.StoreKey,
		minttypes.StoreKey, distrtypes.StoreKey,
		govtypes.StoreKey, paramstypes.StoreKey,
		ibchost.StoreKey, upgradetypes.StoreKey,
		epochstypes.StoreKey, hubtypes.StoreKey, hubgentypes.StoreKey,
		ibctransfertypes.StoreKey, capabilitytypes.StoreKey,
		rollappparamstypes.StoreKey,
		timeupgradetypes.ModuleName,
		dividendstypes.StoreKey,
		// evmos keys
		evmtypes.StoreKey,
		feemarkettypes.StoreKey,
		erc20types.StoreKey,
	}
	// Upgrades contains the upgrade handlers for the application
	Upgrades = []upgrades.Upgrade{drs2.Upgrade, drs3.Upgrade, drs4.Upgrade, drs5.Upgrade, drs5from2d.Upgrade, drs6.Upgrade}
)

func getGovProposalHandlers() []govclient.ProposalHandler {
	var govProposalHandlers []govclient.ProposalHandler

	govProposalHandlers = append(govProposalHandlers,
		paramsclient.ProposalHandler,
		distrclient.ProposalHandler,
		upgradeclient.LegacyProposalHandler,
		upgradeclient.LegacyCancelProposalHandler,
		ibcclientclient.UpdateClientProposalHandler,
		ibcclientclient.UpgradeProposalHandler,
		erc20client.RegisterCoinProposalHandler,
		erc20client.RegisterERC20ProposalHandler,
		erc20client.ToggleTokenConversionProposalHandler,
	)

	return govProposalHandlers
}

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string

	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		authzmodule.AppModuleBasic{},
		genutil.AppModuleBasic{},
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		staking.AppModuleBasic{},
		sequencers.AppModuleBasic{},
		mint.AppModuleBasic{},
		epochs.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic(getGovProposalHandlers()),
		params.AppModuleBasic{},
		feegrantmodule.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		ibc.AppModuleBasic{},
		vesting.AppModuleBasic{},
		hubgenesis.AppModuleBasic{},
		hub.AppModuleBasic{},
		timeupgrade.AppModuleBasic{},
		rollappparams.AppModuleBasic{},
		dividends.AppModuleBasic{},

		// Evmos moudles
		evm.AppModuleBasic{},
		feemarket.AppModuleBasic{},
		erc20.AppModuleBasic{},
		transfer.AppModuleBasic{AppModuleBasic: &ibctransfer.AppModuleBasic{}},
	)

	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName:     nil,
		authz.ModuleName:               nil,
		distrtypes.ModuleName:          nil,
		rollappparamstypes.ModuleName:  nil,
		minttypes.ModuleName:           {authtypes.Minter},
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		govtypes.ModuleName:            {authtypes.Burner},
		ibctransfertypes.ModuleName:    {authtypes.Minter, authtypes.Burner},
		evmtypes.ModuleName:            {authtypes.Minter, authtypes.Burner}, // used for secure addition and subtraction of balance using module account
		erc20types.ModuleName:          {authtypes.Minter, authtypes.Burner},
		hubgentypes.ModuleName:         {authtypes.Minter},
		dividendstypes.ModuleName:      nil,
	}

	// module accounts that are allowed to receive tokens
	maccCanReceiveTokens = []string{
		distrtypes.ModuleName,
		hubgentypes.ModuleName,
	}
)

var (
	_ servertypes.Application = (*App)(nil)
	_ simapp.App              = (*App)(nil)
)

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, "."+Name)

	// manually update the power reduction by replacing micro (u) -> atto (a) evmos
	sdk.DefaultPowerReduction = ethermint.PowerReduction
}

// App extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type App struct {
	*baseapp.BaseApp

	cdc               *codec.LegacyAmino
	appCodec          codec.Codec
	interfaceRegistry types.InterfaceRegistry

	invCheckPeriod uint

	// keys to access the substores
	keys    map[string]*storetypes.KVStoreKey
	tkeys   map[string]*storetypes.TransientStoreKey
	memKeys map[string]*storetypes.MemoryStoreKey

	// keepers
	AccountKeeper       authkeeper.AccountKeeper
	AuthzKeeper         authzkeeper.Keeper
	BankKeeper          bankkeeper.Keeper
	CapabilityKeeper    *capabilitykeeper.Keeper
	StakingKeeper       stakingkeeper.Keeper
	SequencersKeeper    seqkeeper.Keeper
	MintKeeper          mintkeeper.Keeper
	EpochsKeeper        epochskeeper.Keeper
	DistrKeeper         distrkeeper.Keeper
	GovKeeper           govkeeper.Keeper
	HubKeeper           hubkeeper.Keeper
	HubGenesisKeeper    hubgenkeeper.Keeper
	UpgradeKeeper       upgradekeeper.Keeper
	ParamsKeeper        paramskeeper.Keeper
	IBCKeeper           *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	TransferKeeper      transferkeeper.Keeper
	FeeGrantKeeper      feegrantkeeper.Keeper
	TimeUpgradeKeeper   timeupgradekeeper.Keeper
	RollappParamsKeeper rollappparamskeeper.Keeper
	DividendsKeeper     dividendskeeper.Keeper

	// make scoped keepers public for test purposes
	ScopedIBCKeeper      capabilitykeeper.ScopedKeeper
	ScopedTransferKeeper capabilitykeeper.ScopedKeeper

	// Evmos keepers
	EvmKeeper       *evmkeeper.Keeper
	FeeMarketKeeper feemarketkeeper.Keeper
	Erc20Keeper     erc20keeper.Keeper

	// mm is the module manager
	mm *module.Manager

	// sm is the simulation manager
	sm *module.SimulationManager

	// module configurator
	configurator module.Configurator

	consensusMessageAdmissionHandler consensus.AdmissionHandler

	// optionally override the way to query the dymint version
	dymintVersionGetter func() (uint32, error)
}

// NewRollapp returns a reference to an initialized blockchain app
func NewRollapp(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	skipUpgradeHeights map[int64]bool,
	homePath string,
	invCheckPeriod uint,
	encodingConfig rollappevmparams.EncodingConfig,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) *App {
	appCodec := encodingConfig.Codec
	cdc := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry

	eip712.SetEncodingConfig(rollappevmparams.EncodingAsSimapp(encodingConfig))

	// NOTE we use custom transaction decoder that supports the sdk.Tx interface instead of sdk.StdTx
	bApp := baseapp.NewBaseApp(
		Name,
		logger,
		db,
		encodingConfig.TxConfig.TxDecoder(),
		baseAppOptions...,
	)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)

	keys := sdk.NewKVStoreKeys(
		kvstorekeys...,
	)

	// Add the EVM transient store key
	tkeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey, evmtypes.TransientKey, feemarkettypes.TransientKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	// load state streaming if enabled
	if _, _, err := streaming.LoadStreamingServices(bApp, appOpts, appCodec, keys); err != nil {
		panic("failed to load state streaming services: " + err.Error())
	}

	app := &App{
		BaseApp:           bApp,
		cdc:               cdc,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		invCheckPeriod:    invCheckPeriod,
		keys:              keys,
		tkeys:             tkeys,
		memKeys:           memKeys,
	}

	app.ParamsKeeper = initParamsKeeper(
		appCodec,
		cdc,
		keys[paramstypes.StoreKey],
		tkeys[paramstypes.TStoreKey],
	)

	// set the BaseApp's parameter store
	bApp.SetParamStore(app.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramstypes.ConsensusParamsKeyTable()))
	app.CapabilityKeeper = capabilitykeeper.NewKeeper(appCodec, keys[capabilitytypes.StoreKey], memKeys[capabilitytypes.MemStoreKey])
	// Applications that wish to enforce statically created ScopedKeepers should call `Seal` after creating
	// their scoped modules in `NewApp` with `ScopeToModule`
	scopedIBCKeeper := app.CapabilityKeeper.ScopeToModule(ibchost.ModuleName)
	scopedTransferKeeper := app.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)
	app.CapabilityKeeper.Seal()

	// add keepers
	app.EpochsKeeper = *epochskeeper.NewKeeper(appCodec, keys[epochstypes.StoreKey])

	// use custom Ethermint account for contracts
	app.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec,
		keys[authtypes.StoreKey],
		app.GetSubspace(authtypes.ModuleName),
		ethermint.ProtoAccount,
		maccPerms,
		sdk.GetConfig().GetBech32AccountAddrPrefix(),
	)

	app.AuthzKeeper = authzkeeper.NewKeeper(
		keys[authzkeeper.StoreKey],
		appCodec,
		app.MsgServiceRouter(),
		app.AccountKeeper,
	)

	app.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec,
		keys[banktypes.StoreKey],
		app.AccountKeeper,
		app.GetSubspace(banktypes.ModuleName),
		app.BlockedAddrs(),
	)

	stakingKeeper := stakingkeeper.NewKeeper(
		appCodec,
		keys[stakingtypes.StoreKey],
		app.AccountKeeper,
		app.BankKeeper,
		&app.Erc20Keeper,
		app.GetSubspace(stakingtypes.ModuleName),
	)

	app.MintKeeper = mintkeeper.NewKeeper(
		appCodec,
		keys[minttypes.StoreKey],
		app.GetSubspace(minttypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		app.EpochsKeeper,
		authtypes.FeeCollectorName,
	)
	app.MintKeeper.SetHooks(
		minttypes.NewMultiMintHooks(
			// insert mint hooks receivers here
		),
	)

	app.DistrKeeper = distrkeeper.NewKeeper(
		appCodec, keys[distrtypes.StoreKey], app.GetSubspace(distrtypes.ModuleName), app.AccountKeeper, app.BankKeeper,
		&stakingKeeper, &app.SequencersKeeper, &app.Erc20Keeper, authtypes.FeeCollectorName,
	)

	app.FeeGrantKeeper = feegrantkeeper.NewKeeper(appCodec, keys[feegrant.StoreKey], app.AccountKeeper)
	app.UpgradeKeeper = upgradekeeper.NewKeeper(skipUpgradeHeights, keys[upgradetypes.StoreKey], appCodec, homePath, app.BaseApp, authtypes.NewModuleAddress(govtypes.ModuleName).String())

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	app.StakingKeeper = *stakingKeeper.SetHooks(app.DistrKeeper.Hooks())

	app.EpochsKeeper.SetHooks(
		epochstypes.NewMultiEpochHooks(
			// insert epoch hooks receivers here
			app.MintKeeper.Hooks(),
		),
	)
	app.RollappParamsKeeper = rollappparamskeeper.NewKeeper(
		app.GetSubspace(rollappparamstypes.ModuleName),
	)

	app.DividendsKeeper = dividendskeeper.NewKeeper(
		appCodec,
		keys[dividendstypes.StoreKey],
		app.StakingKeeper,
		app.AccountKeeper,
		app.DistrKeeper,
		app.BankKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	app.SequencersKeeper = *seqkeeper.NewKeeper(
		appCodec,
		keys[seqtypes.StoreKey],
		app.GetSubspace(seqtypes.ModuleName),
		authtypes.NewModuleAddress(seqtypes.ModuleName).String(),
		app.AccountKeeper,
		app.RollappParamsKeeper,
		app.UpgradeKeeper,
		[]seqkeeper.AccountBumpFilterFunc{
			shouldBumpEvmAccountSequence,
		},
	)

	app.TimeUpgradeKeeper = timeupgradekeeper.NewKeeper(
		appCodec, keys[timeupgradetypes.StoreKey], authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	// ... other modules keepers
	tracer := cast.ToString(appOpts.Get(srvflags.EVMTracer))

	// Create Ethermint keepers
	app.FeeMarketKeeper = feemarketkeeper.NewKeeper(
		appCodec, authtypes.NewModuleAddress(govtypes.ModuleName),
		keys[feemarkettypes.StoreKey],
		tkeys[feemarkettypes.TransientKey],
		app.GetSubspace(feemarkettypes.ModuleName),
	)

	app.EvmKeeper = evmkeeper.NewKeeper(
		appCodec, keys[evmtypes.StoreKey], tkeys[evmtypes.TransientKey], authtypes.NewModuleAddress(govtypes.ModuleName),
		app.AccountKeeper, app.BankKeeper,
		EVMWrappedSeqKeeper{app.SequencersKeeper},
		app.FeeMarketKeeper,
		tracer, app.GetSubspace(evmtypes.ModuleName),
	)

	app.Erc20Keeper = erc20keeper.NewKeeper(
		keys[erc20types.StoreKey], appCodec, authtypes.NewModuleAddress(govtypes.ModuleName),
		app.AccountKeeper, app.BankKeeper, app.EvmKeeper,
	)

	// Register a custom balance getter to handle ERC20 tokens sent as dividends
	app.DividendsKeeper.SetErc20Keeper(app.Erc20Keeper)

	// Create IBC Keeper
	app.IBCKeeper = ibckeeper.NewKeeper(
		appCodec, keys[ibchost.StoreKey], app.GetSubspace(ibchost.ModuleName), app.SequencersKeeper, app.UpgradeKeeper, scopedIBCKeeper,
	)

	// Register the proposal types
	// Deprecated: Avoid adding new handlers, instead use the new proposal flow
	// by granting the governance module the right to execute the message.
	// See: https://github.com/cosmos/cosmos-sdk/blob/release/v0.46.x/x/gov/spec/01_concepts.md#proposal-messages
	govRouter := govv1beta1.NewRouter()
	govRouter.
		AddRoute(govtypes.RouterKey, govv1beta1.ProposalHandler).
		AddRoute(paramproposal.RouterKey, proposal.NewCustomParamChangeProposalHandler(app.ParamsKeeper, app.BankKeeper)).
		AddRoute(distrtypes.RouterKey, distr.NewCommunityPoolSpendProposalHandler(app.DistrKeeper)).
		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(app.UpgradeKeeper)).
		AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientProposalHandler(app.IBCKeeper.ClientKeeper)).
		AddRoute(erc20types.RouterKey, erc20.NewErc20ProposalHandler(&app.Erc20Keeper))

	govConfig := govtypes.DefaultConfig()
	govKeeper := govkeeper.NewKeeper(
		appCodec, keys[govtypes.StoreKey], app.GetSubspace(govtypes.ModuleName), app.AccountKeeper, app.BankKeeper,
		&stakingKeeper, govRouter, app.MsgServiceRouter(), govConfig,
	)

	app.GovKeeper = *govKeeper.SetHooks(
		govtypes.NewMultiGovHooks(
			// register the governance hooks
		),
	)

	app.EvmKeeper = app.EvmKeeper.SetHooks(
		evmkeeper.NewMultiEvmHooks(
			app.Erc20Keeper.Hooks(),
		),
	)

	app.HubGenesisKeeper = hubgenkeeper.NewKeeper(
		appCodec,
		keys[hubgentypes.StoreKey],
		app.GetSubspace(hubgentypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		app.MintKeeper,
		app.IBCKeeper.ChannelKeeper,
	)

	app.HubKeeper = hubkeeper.NewKeeper(
		appCodec,
		keys[hubtypes.StoreKey],
	)

	var ics4Wrapper ibcporttypes.ICS4Wrapper
	// The IBC tranfer submit is wrapped with the following middlewares:
	// - denom metadata middleware
	ics4Wrapper = denommetadata.NewICS4Wrapper(
		app.IBCKeeper.ChannelKeeper,
		app.HubKeeper,
		app.BankKeeper,
		app.HubGenesisKeeper.GetState,
	)
	// - genesis bridge - IBC transfer disabled until genesis bridge protocol completes
	ics4Wrapper = hubgenkeeper.NewICS4Wrapper(ics4Wrapper, app.HubGenesisKeeper)

	app.TransferKeeper = transferkeeper.NewKeeper(
		appCodec,
		keys[ibctransfertypes.StoreKey],
		app.GetSubspace(ibctransfertypes.ModuleName),
		ics4Wrapper,
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		scopedTransferKeeper,
		app.Erc20Keeper, // Add ERC20 Keeper for ERC20 transfers
	)

	// create IBC module from top to bottom of stack
	var transferStack ibcporttypes.IBCModule
	transferStack = transfer.NewIBCModule(app.TransferKeeper)
	transferStack = denommetadata.NewIBCModule(
		transferStack,
		app.BankKeeper,
		app.TransferKeeper,
		app.HubKeeper,
		denommetadatamoduletypes.NewMultiDenommetadataHooks(
			erc20keeper.NewERC20ContractRegistrationHook(app.Erc20Keeper),
		),
	)

	transferStack = erc20.NewIBCMiddleware(app.Erc20Keeper, transferStack)
	transferStack = hubgenkeeper.NewIBCModule(
		transferStack,
		app.HubGenesisKeeper,
		app.BankKeeper,
	)

	// Create static IBC router, add transfer route, then set and seal it
	ibcRouter := ibcporttypes.NewRouter()
	ibcRouter.AddRoute(ibctransfertypes.ModuleName, transferStack)
	app.IBCKeeper.SetRouter(ibcRouter)

	/**** Module Options ****/
	// used for x/mint v2 migrator. it's a direct access to the params store for x/mint
	// this required as we need to access same subspace with different KeyTable
	mintParamsDirectSubspace := paramstypes.NewSubspace(appCodec, cdc, keys[paramstypes.StoreKey], keys[paramstypes.TStoreKey], minttypes.ModuleName)

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.

	modules := []module.AppModule{
		genutil.NewAppModule(
			app.AccountKeeper, app.StakingKeeper, app.BaseApp.DeliverTx,
			encodingConfig.TxConfig,
		),
		auth.NewAppModule(appCodec, app.AccountKeeper, nil),
		authzmodule.NewAppModule(appCodec, app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		vesting.NewAppModule(app.AccountKeeper, app.BankKeeper),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		gov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper),
		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper, app.BankKeeper, mintParamsDirectSubspace),
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		sequencers.NewAppModule(app.SequencersKeeper),
		epochs.NewAppModule(appCodec, app.EpochsKeeper),
		params.NewAppModule(app.ParamsKeeper),
		ibc.NewAppModule(app.IBCKeeper),
		upgrade.NewAppModule(app.UpgradeKeeper),
		hubgenesis.NewAppModule(appCodec, app.HubGenesisKeeper),
		hub.NewAppModule(appCodec, app.HubKeeper),
		timeupgrade.NewAppModule(app.TimeUpgradeKeeper, app.UpgradeKeeper),
		rollappparams.NewAppModule(appCodec, app.RollappParamsKeeper),
		dividends.NewAppModule(app.DividendsKeeper),
		// Ethermint app modules
		evm.NewAppModule(app.EvmKeeper, app.AccountKeeper, app.GetSubspace(evmtypes.ModuleName)),
		feemarket.NewAppModule(app.FeeMarketKeeper, app.GetSubspace(feemarkettypes.ModuleName)),
		// Evmos app modules
		transfer.NewAppModule(app.TransferKeeper),
		erc20.NewAppModule(app.Erc20Keeper, app.AccountKeeper, app.GetSubspace(erc20types.ModuleName)),
	}

	app.mm = module.NewManager(modules...)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	// NOTE: capability module's beginblocker must come before any modules using capabilities (e.g. IBC)
	beginBlockersList := []string{
		upgradetypes.ModuleName,
		timeupgradetypes.ModuleName,
		capabilitytypes.ModuleName,
		minttypes.ModuleName,
		feemarkettypes.ModuleName,
		evmtypes.ModuleName,
		distrtypes.ModuleName,
		stakingtypes.ModuleName,
		seqtypes.ModuleName,
		vestingtypes.ModuleName,
		ibchost.ModuleName,
		ibctransfertypes.ModuleName,
		authtypes.ModuleName,
		authz.ModuleName,
		banktypes.ModuleName,
		govtypes.ModuleName,
		erc20types.ModuleName,
		genutiltypes.ModuleName,
		feegrant.ModuleName,
		epochstypes.ModuleName,
		paramstypes.ModuleName,
		hubgentypes.ModuleName,
		hubtypes.ModuleName,
		rollappparamstypes.ModuleName,
		dividendstypes.ModuleName,
	}
	app.mm.SetOrderBeginBlockers(beginBlockersList...)

	endBlockersList := []string{
		govtypes.ModuleName,
		stakingtypes.ModuleName,
		evmtypes.ModuleName,
		seqtypes.ModuleName,
		feemarkettypes.ModuleName,
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		authz.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		vestingtypes.ModuleName,
		minttypes.ModuleName,
		erc20types.ModuleName,
		genutiltypes.ModuleName,
		feegrant.ModuleName,
		epochstypes.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		timeupgradetypes.ModuleName,
		ibchost.ModuleName,
		ibctransfertypes.ModuleName,
		hubgentypes.ModuleName,
		hubtypes.ModuleName,
		rollappparamstypes.ModuleName,
		dividendstypes.ModuleName,
	}
	app.mm.SetOrderEndBlockers(endBlockersList...)

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: The genutils module must also occur after auth so that it can access the params from auth.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	initGenesisList := []string{
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		authz.ModuleName,
		banktypes.ModuleName,
		evmtypes.ModuleName,
		feemarkettypes.ModuleName,
		distrtypes.ModuleName,
		stakingtypes.ModuleName,
		seqtypes.ModuleName,
		vestingtypes.ModuleName,
		epochstypes.ModuleName,
		govtypes.ModuleName,
		minttypes.ModuleName,
		ibchost.ModuleName,
		genutiltypes.ModuleName,
		erc20types.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		timeupgradetypes.ModuleName,
		ibctransfertypes.ModuleName,
		feegrant.ModuleName,
		hubgentypes.ModuleName,
		hubtypes.ModuleName,
		rollappparamstypes.ModuleName,
		dividendstypes.ModuleName,
	}
	app.mm.SetOrderInitGenesis(initGenesisList...)

	app.mm.RegisterRoutes(app.Router(), app.QueryRouter(), encodingConfig.Amino)
	app.configurator = module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter())
	app.mm.RegisterServices(app.configurator)

	// RegisterUpgradeHandlers is used for registering any on-chain upgrades.
	// Make sure it's called after `app.mm` and `app.configurator` are set.
	// app.RegisterUpgradeHandlers()

	// add test gRPC service for testing gRPC queries in isolation
	testdata.RegisterQueryServer(app.GRPCQueryRouter(), testdata.QueryImpl{})

	// create the simulation manager and define the order of the modules for deterministic simulations
	//
	// NOTE: this is not required apps that don't use the simulator for fuzz testing
	// transactions
	overrideModules := map[string]module.AppModuleSimulation{
		authtypes.ModuleName: auth.NewAppModule(app.appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts),
	}
	app.sm = module.NewSimulationManagerFromAppModules(app.mm.Modules, overrideModules)

	app.sm.RegisterStoreDecoders()

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

	// initialize BaseApp
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)
	app.setupUpgradeHandlers()

	maxGasWanted := cast.ToUint64(appOpts.Get(srvflags.EVMMaxTxGasWanted))
	h := ante.MustCreateHandler(
		app.appCodec,
		encodingConfig.TxConfig,
		maxGasWanted,
		func(ctx sdk.Context, accAddr sdk.AccAddress, perm string) bool {
			return true
			/*
				TODO:
					We had a plan to use the sequencers module to manager permissions, but that idea was changed
					For now, we just assume the only account with permission is the denom one
					We will eventually replace with something more substantial
				TODO:
					The denom one was ripped out https://github.com/dymensionxyz/dymension-rdk/pull/433/files#diff-2caeed9462180cba822eeaff485f2bb87c9c9464040fb65f0f5dcac66fb0e18fL58-L67
			*/
		},
		app.AccountKeeper,
		app.StakingKeeper,
		app.BankKeeper,
		app.FeeMarketKeeper,
		app.EvmKeeper,
		app.Erc20Keeper,
		app.IBCKeeper,
		app.DistrKeeper,
		app.SequencersKeeper,
		app.FeeGrantKeeper,
		app.AuthzKeeper,
		app.RollappParamsKeeper,
	)
	app.SetAnteHandler(h)

	postHandler := ante.NewPostHandler(ante.PostHandlerOptions{
		ERC20Keeper:        app.Erc20Keeper,
		BankKeeper:         app.BankKeeper,
		DistributionKeeper: app.DistrKeeper,
	})
	app.SetPostHandler(postHandler)

	// Admission handler for consensus messages
	app.setAdmissionHandler(consensus.AllowedMessagesHandler([]string{
		proto.MessageName(new(seqtypes.ConsensusMsgUpsertSequencer)),
		proto.MessageName(new(seqtypes.MsgBumpAccountSequences)),
		proto.MessageName(new(seqtypes.MsgUpgradeDRS)),
	}))

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			err = errorsmod.Wrap(err, "new rollapp: load latest version")
			tmos.Exit(err.Error())
		}
	}

	app.ScopedIBCKeeper = scopedIBCKeeper
	app.ScopedTransferKeeper = scopedTransferKeeper

	return app
}

func (app *App) setAdmissionHandler(handler consensus.AdmissionHandler) {
	app.consensusMessageAdmissionHandler = handler
}

// Name returns the name of the App
func (app *App) Name() string { return app.BaseApp.Name() }

// BeginBlocker application updates every begin block
func (app *App) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	consensusResponses := consensus.ProcessConsensusMessages(ctx, app.appCodec, app.consensusMessageAdmissionHandler, app.MsgServiceRouter(), req.ConsensusMessages)

	resp := app.mm.BeginBlock(ctx, req)
	resp.ConsensusMessagesResponses = consensusResponses

	drsVersion, err := dymintversion.GetDRSVersion()
	if app.dymintVersionGetter != nil {
		drsVersion, err = app.dymintVersionGetter()
	}
	if err != nil {
		panic(fmt.Errorf("Unable to get DRS version from binary: %w", err))
	}
	if drsVersion != app.RollappParamsKeeper.Version(ctx) {
		panic(fmt.Errorf("DRS version mismatch. rollapp DRS version: %d binary DRS version:%d", app.RollappParamsKeeper.Version(ctx), drsVersion))
	}
	return resp
}

// EndBlocker application updates every end block
func (app *App) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	rollappparams := app.RollappParamsKeeper.GetParams(ctx)
	abciEndBlockResponse := app.mm.EndBlock(ctx, req)
	abciEndBlockResponse.RollappParamUpdates = &abci.RollappParams{
		Da:         rollappparams.Da,
		DrsVersion: rollappparams.DrsVersion,
	}
	return abciEndBlockResponse
}

// InitChainer application update at chain initialization
func (app *App) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState
	if err := tmjson.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(fmt.Errorf("failed to unmarshal genesis state on InitChain: %w", err))
	}

	genesisInfo := app.HubGenesisKeeper.GetGenesisInfo(ctx)
	genesisInfo.GenesisChecksum = req.GenesisChecksum
	app.HubGenesisKeeper.SetGenesisInfo(ctx, genesisInfo)

	// Passing the dymint sequencers to the sequencer module from RequestInitChain
	if len(req.Validators) == 0 {
		panic("Dymint have no sequencers defined on InitChain")
	}

	// Passing the dymint sequencers to the sequencer module from RequestInitChain
	app.SequencersKeeper.MustSetDymintValidatorUpdates(ctx, req.Validators)

	app.UpgradeKeeper.SetModuleVersionMap(ctx, app.mm.GetVersionMap())
	res := app.mm.InitGenesis(ctx, app.appCodec, genesisState)

	// Everything needed for the genesis bridge data should be set during the InitGenesis call,
	// so we query it after and return it in InitChainResponse.
	genesisBridgeData, err := app.HubGenesisKeeper.PrepareGenesisBridgeData(ctx)
	if err != nil {
		panic(fmt.Errorf("failed to prepare genesis bridge data on InitChain: %w", err))
	}
	bz, err := tmjson.Marshal(genesisBridgeData)
	if err != nil {
		panic(fmt.Errorf("failed to marshal genesis bridge data on InitChain: %w", err))
	}
	res.GenesisBridgeDataBytes = bz

	return res
}

// LoadHeight loads a particular height
func (app *App) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *App) ModuleAccountAddrs() map[string]bool {
	ret := make(map[string]bool)
	for acc := range maccPerms {
		ret[authtypes.NewModuleAddress(acc).String()] = true
	}
	return ret
}

// BlockedAddrs returns all the app's module account addresses that are not
// allowed to receive funds. If the map value is true, that account cannot receive funds.
func (app *App) BlockedAddrs() map[string]bool {
	// block all modules by default
	ret := app.ModuleAccountAddrs()
	// delete them if they CAN receive tokens
	for _, acc := range maccCanReceiveTokens {
		delete(ret, authtypes.NewModuleAddress(acc).String())
	}
	return ret
}

// LegacyAmino returns App's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) LegacyAmino() *codec.LegacyAmino {
	return app.cdc
}

// AppCodec returns an app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) AppCodec() codec.Codec {
	return app.appCodec
}

// InterfaceRegistry returns an InterfaceRegistry
func (app *App) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetKey(storeKey string) *storetypes.KVStoreKey {
	return app.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetTKey(storeKey string) *storetypes.TransientStoreKey {
	return app.tkeys[storeKey]
}

// GetMemKey returns the MemStoreKey for the provided mem key.
//
// NOTE: This is solely used for testing purposes.
func (app *App) GetMemKey(storeKey string) *storetypes.MemoryStoreKey {
	return app.memKeys[storeKey]
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// SimulationManager implements the SimulationApp interface
func (app *App) SimulationManager() *module.SimulationManager {
	return app.sm
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *App) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx
	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register new tendermint queries routes from grpc-gateway.
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register node gRPC service for grpc-gateway.
	nodeservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register grpc-gateway routes for all modules.
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// register swagger API from root so that other applications can override easily
	if apiConfig.Swagger {
		RegisterSwaggerAPI(clientCtx, apiSvr.Router)
	}
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *App) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *App) RegisterTendermintService(clientCtx client.Context) {
	tmservice.RegisterTendermintService(
		clientCtx,
		app.BaseApp.GRPCQueryRouter(),
		app.interfaceRegistry,
		app.Query,
	)
}

func (app *App) RegisterNodeService(clientCtx client.Context) {
	nodeservice.RegisterNodeService(clientCtx, app.GRPCQueryRouter())
}

// IBC Go TestingApp functions

// GetBaseApp implements the TestingApp interface.
func (app *App) GetBaseApp() *baseapp.BaseApp {
	return app.BaseApp
}

// GetStakingKeeper implements the TestingApp interface.
func (app *App) GetStakingKeeper() ibctestingtypes.StakingKeeper {
	return app.StakingKeeper
}

// GetStakingKeeperSDK implements the TestingApp interface.
func (app *App) GetStakingKeeperSDK() stakingkeeper.Keeper {
	return app.StakingKeeper
}

// GetIBCKeeper implements the TestingApp interface.
func (app *App) GetIBCKeeper() *ibckeeper.Keeper {
	return app.IBCKeeper
}

// GetScopedIBCKeeper implements the TestingApp interface.
func (app *App) GetScopedIBCKeeper() capabilitykeeper.ScopedKeeper {
	return app.ScopedIBCKeeper
}

// GetTxConfig implements the TestingApp interface.
func (app *App) GetTxConfig() client.TxConfig {
	cfg := rollappevmparams.MakeEncodingConfig()
	return cfg.TxConfig
}

// RegisterSwaggerAPI registers swagger route with API Server
func RegisterSwaggerAPI(_ client.Context, rtr *mux.Router) {
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}

	staticServer := http.FileServer(statikFS)
	rtr.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", staticServer))
}

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(appCodec codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key, tkey storetypes.StoreKey) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(stakingtypes.ModuleName)
	paramsKeeper.Subspace(seqtypes.ModuleName)
	paramsKeeper.Subspace(minttypes.ModuleName)
	paramsKeeper.Subspace(epochstypes.ModuleName)
	paramsKeeper.Subspace(distrtypes.ModuleName)
	paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govv1.ParamKeyTable())
	paramsKeeper.Subspace(ibctransfertypes.ModuleName)
	paramsKeeper.Subspace(ibchost.ModuleName)
	paramsKeeper.Subspace(hubgentypes.ModuleName)
	paramsKeeper.Subspace(rollappparamstypes.ModuleName)
	paramsKeeper.Subspace(dividendstypes.ModuleName)

	// ethermint subspaces
	paramsKeeper.Subspace(evmtypes.ModuleName)
	paramsKeeper.Subspace(feemarkettypes.ModuleName)
	// evmos subspaces
	paramsKeeper.Subspace(erc20types.ModuleName)
	return paramsKeeper
}

var evmAccountName = proto.MessageName(&ethermint.EthAccount{})

func shouldBumpEvmAccountSequence(accountProtoName string, account authtypes.AccountI) (bool, error) {
	if accountProtoName != evmAccountName {
		return false, nil
	}

	evmAccount, ok := account.(*ethermint.EthAccount)
	if !ok {
		// this is really unlikely but let's create a nice error.
		return false, fmt.Errorf("account is not an EVM account, but it has the same proto name: %T", account)
	}
	return evmAccount.Type() == ethermint.AccountTypeEOA, nil
}

func (app *App) setupUpgradeHandlers() {
	for _, u := range Upgrades {
		app.setupUpgradeHandler(u)
	}
}

func (app *App) setupUpgradeHandler(u upgrades.Upgrade) {
	app.UpgradeKeeper.SetUpgradeHandler(
		u.Name,
		u.CreateHandler(
			upgrades.UpgradeKeepers{
				RpKeeper:      app.RollappParamsKeeper,
				EvmKeeper:     app.EvmKeeper,
				HubgenK:       app.HubGenesisKeeper,
				Erc20keeper:   app.Erc20Keeper,
				AccountKeeper: app.AccountKeeper,
			},
			app.mm,
			app.configurator,
		),
	)

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Errorf("failed to read upgrade info from disk: %w", err))
	}

	if upgradeInfo.Name == u.Name && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		// configure store loader with the store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &u.StoreUpgrades))

		app.Logger().Info("SetupUpgradeHandler Set store loader.")
	}
}

func (app *App) SetDymintVersionGetter(getter func() (uint32, error)) {
	app.dymintVersionGetter = getter
}
