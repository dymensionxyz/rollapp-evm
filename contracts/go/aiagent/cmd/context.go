package cmd

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"agent/config"
)

type Context struct {
	Logger  *slog.Logger
	HomeDir string
	Config  config.Config
}

// InitContext initializes Context from config file and cmd flags
func InitContext() (Context, error) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	homeDir, err := GetHomeDir()
	if err != nil {
		return Context{}, fmt.Errorf("can't get home dir: %w", err)
	}

	configPath := filepath.Join(homeDir, "config", configFile)

	cfg, err := GetConfig(configPath)
	if err != nil {
		return Context{}, fmt.Errorf("can't get observer config: %w", err)
	}

	return Context{
		Logger:  logger,
		HomeDir: homeDir,
		Config:  cfg,
	}, nil
}

func GetHomeDir() (string, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(userHomeDir, DefaultAppDir), nil
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

func GetConfig(configPath string) (config.Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return config.Config{}, fmt.Errorf("read file: %w", err)
	}

	var cfg config.Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return config.Config{}, fmt.Errorf("unmarshal config: %w", err)
	}

	return cfg, nil
}

func StoreConfig(configPath string, cfg config.Config) error {
	jsonConfig, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal config: %w", err)
	}

	err = os.WriteFile(configPath, jsonConfig, 0o600)
	if err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	return nil
}
