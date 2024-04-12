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

## Feature-full setup

For a more advanced setup that include `denom-metadata`, `genesis-accounts` and others, please refer to the [README.with-advanced-features.md](./README.with-advanced-features.md)

## Installing / Getting started

Build and install the ```rollapp-evm``` binary:

```shell
make install
```

### Initial configuration

export the following variables:

```shell
export EXECUTABLE="rollapp-evm"

export ROLLAPP_CHAIN_ID="rollappevm_1234-1"
export KEY_NAME_ROLLAPP="rol-user"
export BASE_DENOM="arax"
export DENOM=$(echo "$BASE_DENOM" | sed 's/^.//')
export MONIKER="$ROLLAPP_CHAIN_ID-sequencer"

export ROLLAPP_HOME_DIR="$HOME/.rollapp_evm"
export ROLLAPP_SETTLEMENT_INIT_DIR_PATH="${ROLLAPP_HOME_DIR}/init"

export HUB_KEY_WITH_FUNDS="hub-user" # This key should exist on the keyring-backend test
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

## Run a rollapp with local settlement node

### Run local dymension hub node

Follow the instructions on [Dymension Hub Docs](https://docs.dymension.xyz/validate/dymension-hub/join-network) to run local dymension hub node.

### Create sequencer keys

create sequencer key using `dymd`

```shell
dymd keys add sequencer --keyring-dir ~/.rollapp_evm/sequencer_keys --keyring-backend test
SEQUENCER_ADDR=`dymd keys show sequencer --address --keyring-backend test --keyring-dir ~/.rollapp_evm/sequencer_keys`
```

fund the sequencer account

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

### Register rollapp on settlement

```shell
sh scripts/settlement/register_rollapp_to_hub.sh
```

### Register sequencer for rollapp on settlement

```shell
sh scripts/settlement/register_sequencer_to_hub.sh
```

### Configure the rollapp

Modify `dymint.toml` in the chain directory (`~/.rollapp_evm/config`)
set:

```shell
ROLLAPP_HOME_DIR="$HOME/.rollapp_evm"
sed -i 's/settlement_layer.*/settlement_layer = "dymension"/' ${ROLLAPP_HOME_DIR}/config/dymint.toml
```

### Run rollapp locally

```shell
rollapp-evm start
```

## Setup IBC between rollapp and local dymension hub node

### Install dymension relayer

```shell
git clone https://github.com/dymensionxyz/go-relayer.git --branch v0.2.0-v2.3.1-relayer
cd go-relayer && make install
```

### Establish IBC channel

while the rollapp and the local dymension hub node running, run:

```shell
sh scripts/ibc/setup_ibc.sh
```

After successful run, the new established channels will be shown

### run the relayer

```shell
rly start hub-rollapp
```

## Developers guide

TODO
