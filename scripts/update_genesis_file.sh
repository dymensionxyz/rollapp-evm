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

# this is a static module account for the hubgenesis module
# retrieved using  'rollapp-evm q auth module-accounts' command
module_account_address=$(rollapp-evm debug addr F54EBEEF798CA51615C02D13888768F9960863F2 | grep "Bech32 Acc" | awk '{print $3}')

# Construct the JSON object with the obtained address
module_account=$(jq -n \
  --arg address "$module_account_address" \
  '[{
      "@type": "/cosmos.auth.v1beta1.ModuleAccount",
      "base_account": {
          "account_number": "0",
          "address": $address,
          "pub_key": null,
          "sequence": "0"
      },
      "name": "hubgenesis",
      "permissions": []
  }]')

jq --argjson module_account "$module_account" '.app_state.auth.accounts += $module_account' "$GENESIS_FILE" >"$tmp" && mv "$tmp" "$GENESIS_FILE"

module_account_balance=$(
  jq -n \
    --arg address "$module_account_address" \
    --arg denom "$BASE_DENOM" \
    '[{
    "address": $address,
    "coins": [
      {
        "denom": $denom,
        "amount": "60000000000000000000000"
      }
    ]
  }]'
)

hubgenesis_tokens=$(
jq -n \
  --arg denom "$BASE_DENOM" \
  --arg amount "60000000000000000000000" \
  '[{
    "denom": $denom,
    "amount": $amount
  }]'
)

jq --argjson hubgenesis_tokens "$hubgenesis_tokens" '.app_state.hubgenesis.state.genesis_tokens = $hubgenesis_tokens' "$GENESIS_FILE" >"$tmp" && mv "$tmp" "$GENESIS_FILE"

jq '.app_state.bank.balances[0].coins[0].amount = "2000000000000000000000000000"' "$GENESIS_FILE" >"$tmp" && mv "$tmp" "$GENESIS_FILE"
jq --argjson module_account_balance "$module_account_balance" '.app_state.bank.balances += $module_account_balance' "$GENESIS_FILE" >"$tmp" && mv "$tmp" "$GENESIS_FILE"

jq '.app_state.bank.supply[0].amount = "2000060000000000000000000000"' "$GENESIS_FILE" >"$tmp" && mv "$tmp" "$GENESIS_FILE"

# ---------------------------- add elevated account ---------------------------- #
elevated_address=$("$EXECUTABLE" keys show "$KEY_NAME_ROLLAPP" --keyring-backend test --output json | jq -r .address)
elevated_address_json=$(jq -n \
  --arg address "$elevated_address" \
  '[{
        "address": $address
    }]')
jq --argjson elevated_address_json "$elevated_address_json" '.app_state.hubgenesis.params.genesis_triggerer_allowlist += $elevated_address_json' "$GENESIS_FILE" >"$tmp" && mv "$tmp" "$GENESIS_FILE"

# ---------------------------- add denom metadata ---------------------------- #
denom_metadata=$(cat "$ROLLAPP_SETTLEMENT_INIT_DIR_PATH"/denommetadata.json)
jq --argjson denom_metadata "$denom_metadata" '.app_state.bank.denom_metadata = $denom_metadata' "$GENESIS_FILE" >"$tmp" && mv "$tmp" "$GENESIS_FILE"
jq --arg elevated_address "$elevated_address" '.app_state.denommetadata.params.allowed_addresses += [$elevated_address]' "$GENESIS_FILE" >"$tmp" && mv "$tmp" "$GENESIS_FILE"

# ----------------------------- update evm params ---------------------------- #

jq '.app_state.evm.params.extra_eips = [3855]' "$GENESIS_FILE" >"$tmp" && mv "$tmp" "$GENESIS_FILE"