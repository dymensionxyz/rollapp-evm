# Dymension EVM Rollapp

## Rollapp-evm - A template EVM RollApp chain

This repository hosts `rollapp-evm`, a template implementation of a dymension rollapp with `EVM` execution layer.

`rollapp-evm` is an example of a working RollApp using `dymension-RDK` and `dymint`.

It uses Cosmos-SDK's [simapp](https://github.com/cosmos/cosmos-sdk/tree/main/simapp) as a reference, but with the following changes:

- minimal app setup
- wired with EVM and ERC20 modules by [Evmos](https://github.com/evmos/evmos)
- wired IBC for [ICS 20 Fungible Token Transfers](https://github.com/cosmos/ibc/tree/main/spec/app/ics-020-fungible-token-transfer)
- Uses `dymint` for block sequencing and replacing `tendermint`
- Uses modules from `dymension-RDK` to sync with `dymint` and provide RollApp custom logic

## Overview

**Note**: Requires [Go 1.22.1](https://go.dev/dl/)

## Installing / Getting started

Build and install the ```rollapp-evm``` binary:

```shell
export BECH32_PREFIX=ethm
make install BECH32_PREFIX=$BECH32_PREFIX
```

### Initial configuration

export the following variables:

```shell
export EXECUTABLE="rollapp-evm"
export BECH32_PREFIX="ethm"
export ROLLAPP_CHAIN_ID="rollappevm_1234-1"
export KEY_NAME_ROLLAPP="rol-user"
export BASE_DENOM="arax"
export DENOM=$(echo "$BASE_DENOM" | sed 's/^.//')
export MONIKER="$ROLLAPP_CHAIN_ID-sequencer"

export ROLLAPP_HOME_DIR="$HOME/.rollapp_evm"
export ROLLAPP_SETTLEMENT_INIT_DIR_PATH="${ROLLAPP_HOME_DIR}/init"
export SKIP_EVM_BASE_FEE=true # optional, removes fees on the rollapp
```

And initialize the rollapp:

```shell
sh scripts/init.sh
```

### Run rollapp

```shell
rollapp-evm start
```

You should have a running local rollapp!

## Run a rollapp with a settlement node

### Run local dymension hub node

Follow the instructions on [Dymension Hub docs](https://docs.dymension.xyz/develop/get-started/run-base-layers) to run local dymension hub node

all scripts are adjusted to use local hub node that's hosted on the default port `localhost:36657`.

configuration with a remote hub node is also supported, the following variables must be set:

```shell
export HUB_RPC_ENDPOINT="localhost"
export HUB_RPC_PORT="36657" # default: 36657

export HUB_RPC_URL="http://${HUB_RPC_ENDPOINT}:${HUB_RPC_PORT}"
export HUB_CHAIN_ID="dymension_100-1"

dymd config chain-id ${HUB_CHAIN_ID}
dymd config node ${HUB_RPC_URL}

export HUB_KEY_WITH_FUNDS="hub-user" # This key should exist on the keyring-backend test
```

### Create sequencer keys

create sequencer key using `dymd`

```shell
dymd keys add sequencer --keyring-dir ~/.rollapp_evm/sequencer_keys --keyring-backend test
SEQUENCER_ADDR=`dymd keys show sequencer --address --keyring-backend test --keyring-dir ~/.rollapp_evm/sequencer_keys`
```

fund the sequencer account (if you're using a remote hub node, you must fund the sequencer account or you must have an account with enough funds in your keyring)

```shell
# retrieve the minimal bond amount from hub sequencer params
# you have to account for gas fees so it should the final value should be increased
BOND_AMOUNT="$(dymd q sequencer params -o json | jq -r '.params.min_bond.amount')$(dymd q sequencer params -o json | jq -r '.params.min_bond.denom')"

# Extract the numeric part
NUMERIC_PART=$(echo $BOND_AMOUNT | sed 's/adym//')

# Add 100000000000000000000 for fees
NEW_NUMERIC_PART=$(echo "$NUMERIC_PART + 100000000000000000000" | bc)

# Append 'adym' back
TRANSFER_AMOUNT="${NEW_NUMERIC_PART}adym"

dymd tx bank send $HUB_KEY_WITH_FUNDS $SEQUENCER_ADDR ${TRANSFER_AMOUNT} --keyring-backend test --broadcast-mode block --fees 1dym -y --node ${HUB_RPC_URL}
```

### Generate denommetadata

```shell
sh scripts/settlement/generate_denom_metadata.sh
```

### Add genesis accounts

```shell
sh scripts/settlement/add_genesis_accounts.sh
```

### Register rollapp on settlement

```shell
# for permissioned deployment setup, you must have access to an account whitelisted for rollapp
# registration, assuming you want to import an existing account, you can do:
export HUB_PERMISSIONED_KEY="hub-rollapp-permission" # This key is for creating rollapp in permissioned setup
echo <memonic> | dymd keys add $HUB_PERMISSIONED_KEY --recover
# input mnemonic from the account that has the permission to register rollapp

sh scripts/settlement/register_rollapp_to_hub.sh
```

### Register sequencer for rollapp on settlement

```shell
sh scripts/settlement/register_sequencer_to_hub.sh
```

### Configure the rollapp

Modify `dymint.toml` in the chain directory (`~/.rollapp_evm/config`)

linux:

```shell
sed -i 's/settlement_layer.*/settlement_layer = "dymension"/' ${ROLLAPP_HOME_DIR}/config/dymint.toml
sed -i '/node_address =/c\node_address = '\"$HUB_RPC_URL\" "${ROLLAPP_HOME_DIR}/config/dymint.toml"
sed -i '/rollapp_id =/c\rollapp_id = '\"$ROLLAPP_CHAIN_ID\" "${ROLLAPP_HOME_DIR}/config/dymint.toml"
```

mac:

```shell
sed -i '' 's/settlement_layer.*/settlement_layer = "dymension"/' ${ROLLAPP_HOME_DIR}/config/dymint.toml
sed -i '' 's|node_address =.*|node_address = '\"$HUB_RPC_URL\"'|' "${ROLLAPP_HOME_DIR}/config/dymint.toml"
sed -i '' 's|rollapp_id =.*|rollapp_id = '\"$ROLLAPP_CHAIN_ID\"'|' "${ROLLAPP_HOME_DIR}/config/dymint.toml"
```

### Update the Genesis file to include the denommetadata, genesis accounts, module account and elevated accounts

```shell
sh scripts/update_genesis_file.sh
```

Validate genesis file:

```shell
rollapp-evm validate-genesis
```

```shell
# this script automatically adds 2 vesting accounts, adjust the timestampts to your liking or skip this step
sh scripts/add_vesting_accounts_to_genesis_file.sh
```

### Change to 3s block time for ibc connection initialization

Linux:

```shell
sed -i 's/empty_blocks_max_time = "1h0m0s"/empty_blocks_max_time = "3s"/' ${ROLLAPP_HOME_DIR}/config/dymint.toml
```

Mac:

```shell
sed -i '' 's/empty_blocks_max_time = "1h0m0s"/empty_blocks_max_time = "3s"/' ${ROLLAPP_HOME_DIR}/config/dymint.toml
```

### Run rollapp locally

```shell
rollapp-evm start
```

## Setup IBC between rollapp and local dymension hub node

### Install dymension relayer

```shell
git clone https://github.com/dymensionxyz/go-relayer.git --branch v0.3.0-v2.5.2-relayer
cd go-relayer && make install
```

### Establish IBC channel

Verify you have all the environment variables defined earlier set.
while the rollapp and the local dymension hub node running, run:

```shell
sh scripts/ibc/setup_ibc.sh
```

After successful run, the new established channels will be shown

### Configure empty block time to 1hr

Stop the rollapp:

```shell
kill $(pgrep rollapp-evm)
```

Linux:

```shell
sed -i 's/empty_blocks_max_time = "3s"/empty_blocks_max_time = "3600s"/' ${ROLLAPP_HOME_DIR}/config/dymint.toml
```

Mac:

```shell
sed -i '' 's/empty_blocks_max_time = "3s"/empty_blocks_max_time = "3600s"/' ${ROLLAPP_HOME_DIR}/config/dymint.toml
```

Start the rollapp:

```shell
rollapp-evm start
```

### run the relayer

```shell
rly start hub-rollapp
```

### Trigger genesis events

Trigger the genesis events on the rollapp

```shell
sh ./scripts/trigger_rollapp_genesis_event.sh
```

Trigger the genesis events on the hub

```shell
sh ./scripts/settlement/trigger_hub_genesis_event.sh
```

## Developers guide

TODO
