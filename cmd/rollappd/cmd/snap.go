package cmd

import (
	"fmt"

	"github.com/dymensionxyz/dymint/block"
	dymintconf "github.com/dymensionxyz/dymint/config"
	dymintconv "github.com/dymensionxyz/dymint/conv"
	localda "github.com/dymensionxyz/dymint/da/local"
	"github.com/dymensionxyz/dymint/settlement"
	"github.com/dymensionxyz/dymint/store"
	"github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/tendermint/tendermint/node"
	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/proxy"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/server/types"
	"github.com/spf13/cobra"

	"github.com/dymensionxyz/dymension-rdk/utils"
)

// SnapCmd rollbacks the app multistore to specific height and updates dymint state according to it
func SnapCmd(appCreator types.AppCreator) *cobra.Command {
	cmd := &cobra.Command{
		Use: "snapupdate",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := server.GetServerContextFromCmd(cmd)
			cfg := ctx.Config
			home := cfg.RootDir

			db, err := utils.OpenDB(home)
			if err != nil {
				return err
			}
			defer func() {
				err = db.Close()
			}()

			nodeConfig := dymintconf.DefaultConfig("", "")
			err = nodeConfig.GetViperConfig(cmd, ctx.Viper.GetString(flags.FlagHome))
			if err != nil {
				return err
			}

			app := appCreator(ctx.Logger, db, nil, ctx.Viper)

			ctx.Logger.Info("starting block manager with ABCI in-process")

			clientCreator := proxy.NewLocalClientCreator(app)

			newApp := proxy.NewAppConns(clientCreator)

			newApp.SetLogger(ctx.Logger)
			if err := newApp.Start(); err != nil {
				return fmt.Errorf("starting proxy app connections: %w", err)
			}

			info, err := newApp.Query().InfoSync(proxy.RequestInfo)
			if err != nil {
				return fmt.Errorf("querying info: %w", err)
			}

			fmt.Println(info)
			genDocProvider := node.DefaultGenesisDocProviderFunc(cfg)
			genesis, err := genDocProvider()
			if err != nil {
				return err
			}

			blockManager, err := liteBlockManager(cfg, nodeConfig, genesis, nil, clientCreator, ctx.Logger)
			if err != nil {
				return fmt.Errorf("start lite block manager: %w", err)
			}

			// rollback dymint state according to the app
			if err := blockManager.UpdateStateFromApp(); err != nil {
				return fmt.Errorf("updating dymint from app state: %w", err)
			}

			_, err = blockManager.Store.SaveState(blockManager.State, nil)
			if err != nil {
				return fmt.Errorf("save state: %w", err)
			}

			return err
		},
	}
	dymintconf.AddNodeFlags(cmd)
	return cmd
}

func liteBlockManager(cfg *config.Config, dymintConf *dymintconf.NodeConfig, genesis *tmtypes.GenesisDoc, slclient settlement.ClientI, clientCreator proxy.ClientCreator, logger log.Logger) (*block.Manager, error) {

	privValKey, err := p2p.LoadOrGenNodeKey(cfg.PrivValidatorKeyFile())
	if err != nil {
		return nil, err
	}
	signingKey, err := dymintconv.GetNodeKey(privValKey)
	if err != nil {
		return nil, err
	}

	err = dymintconv.GetNodeConfig(dymintConf, cfg)
	if err != nil {
		return nil, err
	}

	proxyApp := proxy.NewAppConns(clientCreator)
	if err := proxyApp.Start(); err != nil {
		return nil, fmt.Errorf("starting proxy app connections: %w", err)
	}

	var baseKV store.KV
	if dymintConf.RootDir == "" && dymintConf.DBPath == "" { // this is used for testing
		return nil, fmt.Errorf("wrong path")
	} else {
		baseKV = store.NewDefaultKVStore(dymintConf.RootDir, dymintConf.DBPath, "dymint")
	}
	mainKV := store.NewPrefixKV(baseKV, []byte{0})
	s := store.New(mainKV)

	dalc := &localda.DataAvailabilityLayerClient{}

	blockManager, err := block.NewManager(
		signingKey,
		dymintConf.BlockManagerConfig,
		genesis,
		s,
		nil,
		proxyApp,
		dalc,
		slclient,
		nil,
		nil,
		nil,
		nil,
		nil,
		logger,
	)
	if err != nil {
		return nil, fmt.Errorf("BlockManager initialization error: %w", err)
	}

	return blockManager, nil
}
