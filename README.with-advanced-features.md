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

**Note**: Requires [Go 1.20](https://go.dev/)

## Quick guide

Get started with [building RollApps](https://docs.dymension.xyz/develop/get-started/setup)

## Installing / Getting started

Build and install the ```rollapp-evm``` binary:

```shell
make install
```

### Initial configuration

export the following variables:

```shell
export ROLLAPP_CHAIN_ID="rollappevm_1234-1"
export KEY_NAME_ROLLAPP="rol-user"
export BASE_DENOM="arax"
export DENOM=$(echo "$BASE_DENOM" | sed 's/^.//')
export MONIKER="$ROLLAPP_CHAIN_ID-sequencer"
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
export HUB_RPC_ENDPOINT="http://localhost"
export HUB_RPC_PORT="36657" # default: 36657
export HUB_RPC_URL="http://3.71.160.88:36657"
export HUB_CHAIN_ID="dymension_100-1"
```

### Create sequencer keys

create sequencer key using `dymd`

```shell
dymd keys add sequencer --keyring-dir ~/.rollapp_evm/sequencer_keys --keyring-backend test
SEQUENCER_ADDR=`dymd keys show sequencer --address --keyring-backend test --keyring-dir ~/.rollapp_evm/sequencer_keys`
```

fund the sequencer account (if you're using a remote hub node, you must fund the sequencer account or you must have an account with enough funds in your keyring)

```shell
BOND_AMOUNT="100000dym"
dymd tx bank send local-user $SEQUENCER_ADDR ${BOND_AMOUNT} --keyring-backend test --broadcast-mode block --fees 1dym -y --node ${HUB_RPC_URL}
```

### Generate denommetadata

```shell
export ROLLAPP_SETTLEMENT_INIT_DIR_PATH="$HOME/.rollapp_evm/init"

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
dymd keys add local-user --recover
# input mnemonic from the account that has the permission to register rollapp

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

### Update the Genesis file to include the denommetadata, genesis accounts, module account and elevated accounts

```shell
sh scripts/update_genesis_file.sh
```

```shell
# this script automatically adds 2 vesting accounts, adjust the timestampts to your liking or skip this step
sh scripts/add_vesting_accounts_to_genesis_file.sh
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
