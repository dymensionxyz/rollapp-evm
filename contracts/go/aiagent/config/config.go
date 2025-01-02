package config

import (
	"math/big"
	"time"
)

type Config struct {
	// OpenAI configuration
	OpenAIAPIKey         string        `json:"open_ai_api_key"`
	OpenAIBaseURL        string        `json:"open_ai_base_url"`
	PollRetryCount       int           `json:"poll_retry_count"`
	PollRetryWaitTime    time.Duration `json:"poll_retry_wait_time"`
	PollRetryMaxWaitTime time.Duration `json:"poll_retry_max_wait_time"`

	// Agent configuration
	HTTPServerAddress string `json:"http_server_address"`

	// DB configuration
	DBPath string `json:"db_path"`

	// Contract configuration
	ContractPollInterval time.Duration `json:"contract_poll_interval"`
	NodeURL              string        `json:"node_url"`
	Mnemonic             string        `json:"mnemonic"`
	ContractAddress      string        `json:"contract_address"`
	DerivationPath       string        `json:"derivation_path"`
	GasLimit             uint64        `json:"gas_limit"`
	GasFeeCap            *big.Int      `json:"gas_fee_cap"`
	GasTipCap            *big.Int      `json:"gas_tip_cap"`
}

// DefaultConfig returns a Config with default values
func DefaultConfig() Config {
	return Config{
		OpenAIAPIKey:         "",
		OpenAIBaseURL:        "https://api.openai.com",
		PollRetryCount:       10,
		PollRetryWaitTime:    10 * time.Millisecond,
		PollRetryMaxWaitTime: 4 * time.Second,
		HTTPServerAddress:    ":8080",
		ContractPollInterval: 10 * time.Second,
		NodeURL:              "http://127.0.0.1:8545",
		Mnemonic:             "",
		ContractAddress:      "",
		DerivationPath:       "",
		GasLimit:             1e8,
		GasFeeCap:            big.NewInt(3e16),
		GasTipCap:            big.NewInt(1e13),
	}
}
