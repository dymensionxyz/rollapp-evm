package config

import (
	"math/big"
	"time"
)

// RandomnessServiceURL
type Config struct {
	External RNGConfig      `json:"external" mapstructure:"external"`
	Agent    AgentConfig    `json:"agent" mapstructure:"agent"`
	DB       LevelDBConfig  `json:"db" mapstructure:"db"`
	Contract ContractConfig `json:"contract" mapstructure:"contract"`
}

type RNGConfig struct {
	BaseURL              string        `json:"randomness_server_base_url" mapstructure:"base_url"`
	PollRetryCount       int           `json:"poll_retry_count" mapstructure:"poll_retry_count"`
	PollRetryWaitTime    time.Duration `json:"poll_retry_wait_time" mapstructure:"poll_retry_wait_time"`
	PollRetryMaxWaitTime time.Duration `json:"poll_retry_max_wait_time" mapstructure:"poll_retry_max_wait_time"`
}

func DefaultOpenAIConfig() RNGConfig {
	return RNGConfig{
		BaseURL:              "http://127.0.0.1:8081",
		PollRetryCount:       10,
		PollRetryWaitTime:    10 * time.Millisecond,
		PollRetryMaxWaitTime: 4 * time.Second,
	}
}

type AgentConfig struct {
	HTTPServerAddress string `json:"http_server_address" mapstructure:"http_server_address"`
}

func DefaultAgentConfig() AgentConfig {
	return AgentConfig{
		HTTPServerAddress: ":8080",
	}
}

type LevelDBConfig struct {
	DBPath string `json:"db_path" mapstructure:"db_path"`
}

func DefaultLevelDBConfig() LevelDBConfig {
	return LevelDBConfig{
		DBPath: "db",
	}
}

type ContractConfig struct {
	PollInterval    time.Duration `json:"poll_interval" mapstructure:"poll_interval"`
	NodeURL         string        `json:"node_url" mapstructure:"node_url"`
	Mnemonic        string        `json:"mnemonic" mapstructure:"mnemonic"`
	ContractAddress string        `json:"contract_address" mapstructure:"contract_address"`
	DerivationPath  string        `json:"derivation_path" mapstructure:"derivation_path"`
	GasLimit        uint64        `json:"gas_limit" mapstructure:"gas_limit"`
	GasFeeCap       *big.Int      `json:"gas_fee_cap" mapstructure:"gas_fee_cap"`
	GasTipCap       *big.Int      `json:"gas_tip_cap" mapstructure:"gas_tip_cap"`
}

func DefaultContractConfig() ContractConfig {
	return ContractConfig{
		PollInterval:    10 * time.Second,
		NodeURL:         "http://127.0.0.1:8545",
		Mnemonic:        "",
		ContractAddress: "",
		DerivationPath:  "",
		GasLimit:        1e8,
		GasFeeCap:       big.NewInt(3e16),
		GasTipCap:       big.NewInt(1e13),
	}
}

// DefaultConfig returns a Config with default values
func DefaultConfig() Config {
	return Config{
		External: DefaultOpenAIConfig(),
		Agent:    DefaultAgentConfig(),
		DB:       DefaultLevelDBConfig(),
		Contract: DefaultContractConfig(),
	}
}
