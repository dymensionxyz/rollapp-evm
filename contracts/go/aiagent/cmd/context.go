package cmd

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"agent/config"
	"github.com/spf13/viper"
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

	configPath := filepath.Join(homeDir, "config")

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

func GetConfig(configPath string) (config.Config, error) {
	v := viper.New()
	v.AddConfigPath(configPath)
	v.SetConfigName(configFileName)
	v.SetConfigType(configFileExt)

	if err := v.ReadInConfig(); err != nil {
		return config.Config{}, err
	}

	var conf config.Config
	if err := v.Unmarshal(&conf); err != nil {
		return config.Config{}, err
	}

	return conf, nil
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
