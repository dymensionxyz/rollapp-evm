#!/bin/bash
# TODO: utils genesis transformer could be used to update the genesis file

set -x
CONFIG_DIRECTORY="$ROLLAPP_HOME_DIR/config"
GENESIS_FILE="$CONFIG_DIRECTORY/genesis.json"
DYMINT_CONFIG_FILE="$CONFIG_DIRECTORY/dymint.toml"
APP_CONFIG_FILE="$CONFIG_DIRECTORY/app.toml"

"$EXECUTABLE" keys add hub_genesis --keyring-backend test

dasel put string -f "$GENESIS_FILE" -p json 'consensus_params.block.max_gas' '400000000'
dasel put string -f "$GENESIS_FILE" -p json 'consensus_params.block.max_bytes' '3145728'

dasel put string -f "$GENESIS_FILE" -p json 'app_state.gov.voting_params.voting_period' '300s'

dasel put string -f "$GENESIS_FILE" -p json 'app_state.bank.balances.0.coins.0.amount' '2000000000000000000000000000'
dasel put string -f "$GENESIS_FILE" -p json 'app_state.bank.supply.0.amount' '2000000000000000000000000000'

# ---------------------------- add genesis accounts for the hub ---------------------------- #
genesis_accounts=$(cat "$ROLLAPP_SETTLEMENT_INIT_DIR_PATH"/genesis_accounts.json)
dasel put value -f "$GENESIS_FILE" -p json 'app_state.hubgenesis.state.genesis_accounts' -v "$genesis_accounts"

# ---------------------------- add elevated account ---------------------------- # TODO: can I remove it?
elevated_address=$("$EXECUTABLE" keys show "$KEY_NAME_ROLLAPP" --keyring-backend test --output json | dasel -r json -p json '.address')

# ---------------------------- add denom metadata ---------------------------- #
denom_metadata=$(cat "$ROLLAPP_SETTLEMENT_INIT_DIR_PATH"/denommetadata.json)
dasel put value -f "$GENESIS_FILE" -p json 'app_state.bank.denom_metadata' -v "$denom_metadata"
dasel put string -f "$GENESIS_FILE" -p json 'app_state.denommetadata.params.allowed_addresses' -v "$elevated_address" --append

# ----------------------------- update evm params ---------------------------- #
dasel put value -f "$GENESIS_FILE" -p json 'app_state.evm.params.extra_eips' -v '[3855]'
