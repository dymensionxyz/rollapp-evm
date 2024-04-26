package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/config"
	"github.com/cosmos/cosmos-sdk/client/debug"
	"github.com/cosmos/cosmos-sdk/client/flags"
	pruningtypes "github.com/cosmos/cosmos-sdk/pruning/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/ethermint/crypto/hd"

	berpcconfig "github.com/bcdevtools/block-explorer-rpc-cosmos/be_rpc/config"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/server"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/version"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	tmcfg "github.com/tendermint/tendermint/config"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	tmlog "github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	dbm "github.com/tendermint/tm-db"

	rdkserver "github.com/dymensionxyz/dymension-rdk/server"
	rdk_utils "github.com/dymensionxyz/dymension-rdk/utils"
	dymintconf "github.com/dymensionxyz/dymint/config"
	"github.com/dymensionxyz/rollapp-evm/app"
	"github.com/dymensionxyz/rollapp-evm/app/params"
	"github.com/dymensionxyz/rollapp-evm/utils"

	ethermintclient "github.com/evmos/ethermint/client"

	rdk_servertypes "github.com/dymensionxyz/dymension-rdk/server/types"
	rdk_genutilcli "github.com/dymensionxyz/dymension-rdk/x/genutil/client/cli"
	rdk_genutiltypes "github.com/dymensionxyz/dymension-rdk/x/genutil/types"
	evmserver "github.com/evmos/ethermint/server"
	evmconfig "github.com/evmos/ethermint/server/config"
)

const rollappAscii = `
███████ ██    ██ ███    ███     ██████   ██████  ██      ██       █████  ██████  ██████  
██      ██    ██ ████  ████     ██   ██ ██    ██ ██      ██      ██   ██ ██   ██ ██   ██ 
█████   ██    ██ ██ ████ ██     ██████  ██    ██ ██      ██      ███████ ██████  ██████  
██       ██  ██  ██  ██  ██     ██   ██ ██    ██ ██      ██      ██   ██ ██      ██      
███████   ████   ██      ██     ██   ██  ██████  ███████ ███████ ██   ██ ██      ██                                                                                                                                                            
`

// NewRootCmd creates a new root rollappd command. It is called once in the main function.
func NewRootCmd() (*cobra.Command, params.EncodingConfig) {
	encodingConfig := app.MakeEncodingConfig()

	//TODO: refactor to use depinject

	initClientCtx := client.Context{}.
		WithCodec(encodingConfig.Codec).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithInput(os.Stdin).
		WithAccountRetriever(types.AccountRetriever{}).
		WithHomeDir(app.DefaultNodeHome).
		WithKeyringOptions(hd.EthSecp256k1Option()).
		WithViper("ROLLAPP")

	rootCmd := &cobra.Command{
		//TODO: set by code, not in Makefile
		Use:   version.AppName,
		Short: rollappAscii,
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			// set the default command outputs
			cmd.SetOut(cmd.OutOrStdout())
			cmd.SetErr(cmd.ErrOrStderr())

			initClientCtx, err := client.ReadPersistentCommandFlags(initClientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			initClientCtx, err = config.ReadFromClientConfig(initClientCtx)
			if err != nil {
				return err
			}

			if err := client.SetCmdClientContextHandler(initClientCtx, cmd); err != nil {
				return err
			}

			customTMConfig := initTendermintConfig()
			customAppTemplate, customAppConfig := initAppConfig()
			err = server.InterceptConfigsPreRunHandler(
				cmd, customAppTemplate, customAppConfig, customTMConfig,
			)
			if err != nil {
				return err
			}
			serverCtx := server.GetServerContextFromCmd(cmd)

			//create dymint toml config file
			home := serverCtx.Viper.GetString(tmcli.HomeFlag)
			chainID := client.GetClientContextFromCmd(cmd).ChainID
			dymintconf.EnsureRoot(home, dymintconf.DefaultConfig(home, chainID))

			//create Block Explorer Json-RPC toml config file
			berpcconfig.EnsureRoot(home, berpcconfig.DefaultBeJsonRpcConfig())

			// Set config
			sdkconfig := sdk.GetConfig()
			utils.SetBip44CoinType(sdkconfig)
			cfg := serverCtx.Config
			genFile := cfg.GenesisFile()
			if tmos.FileExists(genFile) {
				genDoc, _ := rdk_genutiltypes.GenesisDocFromFile(genFile)
				rdk_utils.SetPrefixes(sdkconfig, genDoc["bech32_prefix"].(string))
			} else {
				rdk_utils.SetPrefixes(sdkconfig, "ethm")
			}
			sdkconfig.Seal()
			return nil
		},
	}

	initRootCmd(rootCmd, encodingConfig)
	return rootCmd, encodingConfig
}

// initTendermintConfig helps to override default Tendermint Config values.
// return tmcfg.DefaultConfig if no custom configuration is required for the application.
func initTendermintConfig() *tmcfg.Config {
	cfg := tmcfg.DefaultConfig()

	// these values put a higher strain on node memory
	// cfg.P2P.MaxNumInboundPeers = 100
	// cfg.P2P.MaxNumOutboundPeers = 40

	return cfg
}

// initAppConfig helps to override default appConfig template and configs.
// return "", nil if no custom configuration is required for the application.
func initAppConfig() (string, interface{}) {
	customAppTemplate, customAppConfig := evmconfig.AppConfig("")

	srvCfg, ok := customAppConfig.(evmconfig.Config)
	if !ok {
		panic(fmt.Errorf("unknown app config type %T", customAppConfig))
	}

	//Default pruning for a rollapp, represent 2 weeks of states kept while pruning in intervals of 10 minutes
	srvCfg.Pruning = pruningtypes.PruningOptionCustom
	srvCfg.PruningInterval = "18000"
	srvCfg.PruningKeepRecent = "6048000"

	//Changing the default address to global instead of localhost
	srvCfg.JSONRPC.Address = "0.0.0.0:8545"
	srvCfg.JSONRPC.WsAddress = "0.0.0.0:8546"

	return customAppTemplate, srvCfg
}

func initRootCmd(
	rootCmd *cobra.Command,
	encodingConfig params.EncodingConfig,
) {
	ac := appCreator{
		encCfg: encodingConfig,
	}

	initCmd := genutilcli.InitCmd(app.ModuleBasics, app.DefaultNodeHome)
	initCmd.Flags().String(FlagBech32Prefix, "ethm", "set bech32 prefix for rollapp, if left blank default value is 'ethm'")

	initCmd.PostRunE = func(cmd *cobra.Command, args []string) error {
		prefix, _ := initCmd.Flags().GetString(FlagBech32Prefix)

		serverCtx := server.GetServerContextFromCmd(cmd)
		config := serverCtx.Config
		path := config.GenesisFile()

		genDoc, err := rdk_genutiltypes.GenesisDocFromFile(path)
		if err != nil {
			fmt.Println("Failed to read genesis doc from file", err)
		}

		genDoc["bech32_prefix"] = prefix

		genDocBytes, err := json.MarshalIndent(genDoc, "", "  ")
		if err != nil {
			return err
		}
		return tmos.WriteFile(path, genDocBytes, 0o644)

	}

	rootCmd.AddCommand(
		initCmd,
		rdk_genutilcli.CollectGenTxsCmd(banktypes.GenesisBalancesIterator{}, app.DefaultNodeHome),
		genutilcli.MigrateGenesisCmd(),
		genutilcli.GenTxCmd(app.ModuleBasics, encodingConfig.TxConfig, banktypes.GenesisBalancesIterator{}, app.DefaultNodeHome),

		genutilcli.ValidateGenesisCmd(app.ModuleBasics),
		AddGenesisAccountCmd(app.DefaultNodeHome),
		tmcli.NewCompletionCmd(rootCmd, true),
		debug.Cmd(),
		config.Cmd(),
	)

	rdkserver.AddRollappCommands(rootCmd, app.DefaultNodeHome, ac.newApp, ac.appExport, nil)
	rootCmd.AddCommand(StartCmd(ac.newApp, app.DefaultNodeHome))

	// add keybase, auxiliary RPC, query, and tx child commands
	rootCmd.AddCommand(
		rpc.StatusCommand(),
		queryCommand(),
		txCommand(),
		ethermintclient.KeyCommands(app.DefaultNodeHome),
	)

	rootCmd.AddCommand(evmserver.NewIndexTxCmd())
}

// queryCommand returns the sub-command to send queries to the app
func queryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "query",
		Aliases:                    []string{"q"},
		Short:                      "Querying subcommands",
		DisableFlagParsing:         false,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authcmd.GetAccountCmd(),
		rpc.ValidatorCommand(),
		rpc.BlockCommand(),
		authcmd.QueryTxsByEventsCmd(),
		authcmd.QueryTxCmd(),
	)

	app.ModuleBasics.AddQueryCommands(cmd)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

// txCommand returns the sub-command to send transactions to the app
func txCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "tx",
		Short:                      "Transactions subcommands",
		DisableFlagParsing:         false,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authcmd.GetSignCommand(),
		authcmd.GetSignBatchCommand(),
		authcmd.GetMultiSignCommand(),
		authcmd.GetMultiSignBatchCmd(),
		authcmd.GetValidateSignaturesCommand(),
		authcmd.GetBroadcastCommand(),
		authcmd.GetEncodeCommand(),
		authcmd.GetDecodeCommand(),
		authcmd.GetAuxToFeeCommand(),
	)

	app.ModuleBasics.AddTxCommands(cmd)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

type appCreator struct {
	encCfg params.EncodingConfig
}

func (ac appCreator) newApp(
	logger tmlog.Logger,
	db dbm.DB,
	traceStore io.Writer,
	appOpts servertypes.AppOptions,
) servertypes.Application {
	baseappOptions := server.DefaultBaseappOptions(appOpts)

	skipUpgradeHeights := make(map[int64]bool)
	for _, h := range cast.ToIntSlice(appOpts.Get(server.FlagUnsafeSkipUpgrades)) {
		skipUpgradeHeights[int64(h)] = true
	}

	return app.NewRollapp(
		logger,
		db,
		traceStore,
		true,
		skipUpgradeHeights,
		cast.ToString(appOpts.Get(flags.FlagHome)),
		cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod)),
		ac.encCfg,
		appOpts,
		baseappOptions...)
}

func (ac appCreator) appExport(
	logger tmlog.Logger,
	db dbm.DB,
	traceStore io.Writer,
	height int64,
	forZeroHeight bool,
	jailAllowedAddrs []string,
	appOpts rdk_servertypes.AppOptions,
) (rdk_servertypes.ExportedApp, error) {
	var rollapp *app.App
	homePath, ok := appOpts.Get(flags.FlagHome).(string)
	if !ok || homePath == "" {
		return rdk_servertypes.ExportedApp{}, errors.New("application home not set")
	}

	loadLatest := height == -1
	rollapp = app.NewRollapp(
		logger,
		db,
		traceStore,
		loadLatest,
		map[int64]bool{},
		homePath,
		cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod)),
		ac.encCfg,
		appOpts,
	)

	if height != -1 {
		if err := rollapp.LoadHeight(height); err != nil {
			return rdk_servertypes.ExportedApp{}, err
		}
	}

	return rollapp.ExportAppStateAndValidators(forZeroHeight, jailAllowedAddrs)
}
