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

			openAI := external.NewOpenAIClient(cmdCtx.Logger, cmdCtx.Config.External)

			aiOracle, err := contract.NewAIOracleClient(ctx, cmdCtx.Logger, cmdCtx.Config.Contract)
			if err != nil {
				return fmt.Errorf("new AI oracle client: %w", err)
			}

			levelDB, err := repository.NewLevelDB(cmdCtx.Config.DB.DBPath)
			if err != nil {
				return fmt.Errorf("new levelDB: %w", err)
			}

			aiAgent := agent.NewAgent(cmdCtx.Logger, cmdCtx.Config.Agent.HTTPServerAddress, aiOracle, openAI, levelDB)

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
