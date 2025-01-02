package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"agent/agent"
	"agent/config"
	"agent/contract"
	"agent/external"
	"agent/repository"
	"github.com/spf13/cobra"
)

const (
	configFileName = "config"
	configFileExt  = "json"
	configFile     = configFileName + "." + configFileExt
)

var (
	DefaultAppDir = ".rollapp-agent"
)

// RootCmd builds commands for the CLI
func RootCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{}
	cmd.AddCommand(
		InitCmd(),
		StartCmd(),
	)
	return cmd, nil
}

func InitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Init agent",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

			homeDir, err := GetHomeDir()
			if err != nil {
				return fmt.Errorf("get home dir: %w", err)
			}

			configPath := filepath.Join(homeDir, "config")

			err = InitConfig(logger, configPath, configFile)
			if err != nil {
				return fmt.Errorf("can't init observer config: %w", err)
			}

			return nil
		},
	}
	return cmd
}

func StartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start agent",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmdCtx, err := InitContext()
			if err != nil {
				return fmt.Errorf("init command context: %w", err)
			}

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			openAI := external.NewOpenAIClient(
				cmdCtx.Config.OpenAIAPIKey,
				cmdCtx.Config.OpenAIBaseURL,
				cmdCtx.Config.PollRetryCount,
				cmdCtx.Config.PollRetryWaitTime,
				cmdCtx.Config.PollRetryMaxWaitTime,
			)

			aiOracle, err := contract.NewAIOracleClient(
				ctx,
				cmdCtx.Logger,
				contract.Config{
					NodeURL:         cmdCtx.Config.NodeURL,
					Mnemonic:        cmdCtx.Config.Mnemonic,
					ContractAddress: cmdCtx.Config.ContractAddress,
					DerivationPath:  cmdCtx.Config.DerivationPath,
					GasLimit:        cmdCtx.Config.GasLimit,
					GasFeeCap:       cmdCtx.Config.GasFeeCap,
					GasTipCap:       cmdCtx.Config.GasTipCap,
					PollInterval:    cmdCtx.Config.ContractPollInterval,
				},
			)
			if err != nil {
				return fmt.Errorf("new AI oracle client: %w", err)
			}

			levelDB, err := repository.NewLevelDB(cmdCtx.Config.DBPath)
			if err != nil {
				return fmt.Errorf("new levelDB: %w", err)
			}

			aiAgent := agent.NewAgent(
				cmdCtx.Logger,
				cmdCtx.Config.HTTPServerAddress,
				aiOracle,
				openAI,
				levelDB,
			)

			go aiAgent.Run(ctx)

			cmdCtx.Logger.Info("Agent started")

			stop := make(chan os.Signal, 1)
			signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
			<-stop

			cancel()
			_ = aiAgent.Close()

			cmdCtx.Logger.Info("Agent stopped")

			return nil
		},
	}
	return cmd
}

func InitConfig(logger *slog.Logger, configFolder string, configFile string) error {
	configPath := filepath.Join(configFolder, configFile)
	stat, _ := os.Stat(configPath)
	if stat != nil {
		logger.Info("Skipping creating config file: file already exists", "configPath", configPath)
		return nil
	}

	err := os.MkdirAll(configFolder, os.ModePerm)
	if err != nil {
		return fmt.Errorf("mkdir: %w", err)
	}

	err = StoreConfig(configPath, config.DefaultConfig())
	if err != nil {
		return fmt.Errorf("store config: %w", err)
	}

	logger.Info("Created config file", "filePath", configPath)

	return nil
}
