#!/bin/bash
# TODO: utils genesis transformer could be used to update the genesis file

set -x
tmp=$(mktemp)
CONFIG_DIRECTORY="$ROLLAPP_HOME_DIR/config"
GENESIS_FILE="$CONFIG_DIRECTORY/genesis.json"
DYMINT_CONFIG_FILE="$CONFIG_DIRECTORY/dymint.toml"
APP_CONFIG_FILE="$CONFIG_DIRECTORY/app.toml"

"$EXECUTABLE" keys add hub_genesis --keyring-backend test

jq '.consensus_params["block"]["max_gas"] = "400000000"' "$GENESIS_FILE" >"$tmp" && mv "$tmp" "$GENESIS_FILE"
jq '.consensus_params["block"]["max_bytes"] = "3145728"' "$GENESIS_FILE" >"$tmp" && mv "$tmp" "$GENESIS_FILE"

jq '.app_state.gov.voting_params.voting_period = "300s"' "$GENESIS_FILE" >"$tmp" && mv "$tmp" "$GENESIS_FILE"

jq '.app_state.bank.balances[0].coins[0].amount = "2000000000000000000000000000"' "$GENESIS_FILE" >"$tmp" && mv "$tmp" "$GENESIS_FILE"
jq '.app_state.bank.supply[0].amount = "2000000000000000000000000000"' "$GENESIS_FILE" >"$tmp" && mv "$tmp" "$GENESIS_FILE"

# ---------------------------- add genesis accounts for the hub ---------------------------- #
genesis_accounts=$(cat "$ROLLAPP_SETTLEMENT_INIT_DIR_PATH"/genesis_accounts.json)
jq --argjson genesis_accounts "$genesis_accounts" '.app_state.hubgenesis.state.genesis_accounts = $genesis_accounts' "$GENESIS_FILE" >"$tmp" && mv "$tmp" "$GENESIS_FILE"

# ---------------------------- add denom metadata ---------------------------- #
denom_metadata=$(cat "$ROLLAPP_SETTLEMENT_INIT_DIR_PATH"/denommetadata.json)
elevated_address=$("$EXECUTABLE" keys show "$KEY_NAME_ROLLAPP" --keyring-backend test --output json | jq -r .address)
jq --argjson denom_metadata "$denom_metadata" '.app_state.bank.denom_metadata = $denom_metadata' "$GENESIS_FILE" >"$tmp" && mv "$tmp" "$GENESIS_FILE"
jq --arg elevated_address "$elevated_address" '.app_state.denommetadata.params.allowed_addresses += [$elevated_address]' "$GENESIS_FILE" >"$tmp" && mv "$tmp" "$GENESIS_FILE"

# ----------------------------- update evm params ---------------------------- #

jq '.app_state.evm.params.extra_eips = [3855]' "$GENESIS_FILE" >"$tmp" && mv "$tmp" "$GENESIS_FILE"