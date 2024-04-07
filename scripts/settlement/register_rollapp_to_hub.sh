#!/bin/bash

MAX_SEQUENCERS=5

# this account must be whitelisted on the hub for permissioned deployment setup
DEPLOYER="local-user"

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
  --chain-id ${HUB_CHAIN_ID} -y
set +x