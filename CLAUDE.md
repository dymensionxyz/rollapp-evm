# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

This is the d-rollapp-evm repository - a template implementation of a Dymension RollApp with EVM execution layer using Evmos. It serves as a reference implementation for building EVM-compatible rollup chains on Dymension.

## Core Architecture

### Module Structure

The application is built on Cosmos SDK with the following key components:

- **EVM Module**: Evmos EVM implementation for Ethereum compatibility
- **ERC20 Module**: Token bridging and conversion between Cosmos native and ERC20
- **IBC Module**: Inter-blockchain communication for cross-chain transfers
- **Hub Genesis Module**: State synchronization with Dymension Hub
- **Sequencer Module**: Block production and sequencing logic
- **Mint Module**: Token minting and inflation logic from dymension-rdk
- **Epochs Module**: Time-based triggers for periodic operations

### Key Directories

- `app/`: Main application logic, keepers, and module wiring
- `app/ante/`: Custom ante handlers for transaction validation
- `app/upgrades/`: Chain upgrade handlers (DRS-2 through DRS-7)
- `cmd/rollappd/`: CLI commands and application entry point
- `scripts/`: Setup, initialization, and deployment scripts
- `ra_evm_be_rpc/`: EVM RPC backend implementation

### Upgrade System

The app uses Dymension Rollapp Standard (DRS) versioning:

- Current version: DRS-7
- Upgrades are defined in `app/upgrades/drs-*/`
- Each upgrade has its own handler and migration logic

## Build and Development Commands

### Build and Install

```bash
# Set required environment variable
export BECH32_PREFIX=ethm  # Required for all builds

# Build the binary
make build

# Install to $GOPATH/bin
make install BECH32_PREFIX=$BECH32_PREFIX

# The installed binary will be named 'rollapp-evm'
```

### Testing

```bash
# Run all tests with coverage
go test -race -coverprofile=coverage.txt ./...

# Run tests with go-acc (used in CI)
go install github.com/ory/go-acc@v0.2.6
go-acc -o coverage.txt ./... -- -v --race
```

### Linting

```bash
# Run golangci-lint (version 1.62.2)
golangci-lint run --timeout 5m

# Format Go code
gofumpt -w .
```

### Protocol Buffers

```bash
# Generate protobuf files (uses Docker)
make proto-gen

# Clean proto generation containers
make proto-clean
```

### Genesis Generation

```bash
# Generate genesis template for mainnet
make generate-genesis env=mainnet

# Generate genesis template for testnet  
make generate-genesis env=testnet

# Optional: specify DRS version
make generate-genesis env=mainnet DRS_VERSION=7
```

## Local Development Setup

### Initial Configuration

```bash
# Core configuration
export EXECUTABLE="rollapp-evm"
export DA_CLIENT="mock"  # Options: mock, celestia, loadnetwork, sui, aptos, walrus
export ROLLAPP_CHAIN_ID="rollappevm_1234-1"
export KEY_NAME_ROLLAPP="rol-user"
export BASE_DENOM="arax"
export DENOM=$(echo "$BASE_DENOM" | sed 's/^.//')
export MONIKER="$ROLLAPP_CHAIN_ID-sequencer"
export ROLLAPP_HOME_DIR="$HOME/.rollapp_evm"
export ROLLAPP_SETTLEMENT_INIT_DIR_PATH="${ROLLAPP_HOME_DIR}/init"
export SKIP_EVM_BASE_FEE=true  # Optional: disable fees

# Initialize
sh scripts/init.sh

# Start the rollapp
rollapp-evm start
```

### Settlement Layer Integration

```bash
# Hub configuration
export SETTLEMENT_LAYER="dymension"  # or "mock" for local testing
export HUB_KEY_WITH_FUNDS="hub-user"
export HUB_RPC_URL="http://localhost:36657"
export HUB_REST_URL="http://localhost:1318"
export HUB_CHAIN_ID="dymension_100-1"

# Register rollapp on hub
sh scripts/settlement/register_rollapp_to_hub.sh

# Register sequencer
sh scripts/settlement/register_sequencer_to_hub.sh
```

### IBC Setup

```bash
# Setup IBC channel (requires relayer)
sh scripts/ibc/setup_ibc.sh
```

## Important Configurations

### Dymint Configuration (`~/.rollapp_evm/config/dymint.toml`)

Key parameters to configure:

- `settlement_layer`: "dymension" or "mock"
- `node_address`: Hub RPC endpoint
- `rollapp_id`: Chain ID
- `max_idle_time`: Empty block timeout
- `batch_submit_time`: State update frequency

### Application Configuration (`~/.rollapp_evm/config/app.toml`)

- `minimum-gas-prices`: Set token denomination for fees

## Dependencies

- **Go**: 1.23.1+ required
- **Tools**: dasel, jq (for scripts)
- **Docker**: For protobuf generation
- **Relayer**: dymensionxyz/go-relayer v0.3.4-v2.5.2-relayer-canon-6 (for IBC)

## Key Imports and Forks

The project uses several Dymension-specific dependencies:

- `github.com/dymensionxyz/dymension-rdk`: RollApp development kit
- `github.com/dymensionxyz/dymint`: Block sequencing (replaces Tendermint)
- `github.com/evmos/evmos/v12`: EVM execution layer
- `github.com/cosmos/cosmos-sdk v0.46.16`: Base framework
- `github.com/cosmos/ibc-go/v6`: IBC implementation

## CI/CD Workflows

- **test.yml**: Runs unit tests with coverage
- **golangci_lint.yml**: Runs linting checks
- **build.yml**: Builds the binary
- **e2e-tests.yml**: End-to-end integration tests
- **generate_genesis_template.yml**: Generates genesis templates

## Data Availability Layers

Supports multiple DA options via `DA_CLIENT` environment variable:

- `mock`: Local database (development)
- `celestia`: Celestia network
- `sui`: Sui network (requires `SUI_MNEMONIC`)
- `aptos`: Aptos network (requires `APT_PRIVATE_KEY`)
- `walrus`: Walrus network
- `loadnetwork`: Load testing network
