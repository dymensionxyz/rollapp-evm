package config

import "time"

type Config struct {
	// OpenAI configuration
	OpenAIAPIKey  string `json:"open_ai_api_key"`
	OpenAIBaseURL string `json:"open_ai_base_url"`

	// Internal configuration
	HTTPServerAddress    string        `json:"http_server_address"`
	EventPollingInterval time.Duration `json:"event_polling_interval"`

	// Rollapp configuration
	RollappNodeURL     string `json:"rollapp_node_url"`
	KeyringPath        string `json:"keyring_path"` // ??
	HexContractAddress string `json:"hex_contract_address"`

	// Gas configuration
	GasLimit  uint64 `json:"gas_limit"`
	GasFeeCap string `json:"gas_fee_cap"`
	GasTipCap string `json:"gas_tip_cap"`
}
