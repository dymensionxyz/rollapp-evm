#!/bin/bash

MAX_SEQUENCERS=5

# this account must be whitelisted on the hub for permissioned deployment setup
DEPLOYER="local-user"

if [ -z "$HUB_RPC_URL" ]; then
  echo "HUB_RPC_URL is not set, using 'http://localhost:36657'"
  HUB_RPC_URL="http://localhost:36657"
fi

if [ -z "$HUB_CHAIN_ID" ]; then
  echo "HUB_CHAIN_ID is not set, using 'dymension_100-1'"
  HUB_CHAIN_ID="dymension_100-1"
fi

# this file is generated using the scripts/settlement/generate_denom_metadata.sh
DENOM_METADATA_PATH="$HOME/.rollapp_evm/init/denommetadata.json"
# this file is generated using the scripts/settlement/add_genesis_accounts.sh
GENESIS_ACCOUNTS_PATH="$HOME/.rollapp_evm/init/genesis_accounts.json"

set -x
#Register rollapp
dymd tx rollapp create-rollapp "$ROLLAPP_CHAIN_ID" "$MAX_SEQUENCERS" '{"Addresses":[]}' \
  "$DENOM_METADATA_PATH" \
  --genesis-accounts-path "$GENESIS_ACCOUNTS_PATH" \
  --from "$DEPLOYER" \
  --keyring-backend test \
  --broadcast-mode block \
  --fees 1dym \
  --node ${HUB_RPC_URL} \
  --chain-id ${HUB_CHAIN_ID}
set +x
