#!/bin/bash

MAX_SEQUENCERS=5

# this account must be whitelisted on the hub for permissioned deployment setup
DEPLOYER=${HUB_PERMISSIONED_KEY-"$HUB_KEY_WITH_FUNDS"}

if [ -z "$HUB_RPC_URL" ]; then
  echo "HUB_RPC_URL is not set, using 'http://localhost:36657'"
  HUB_RPC_URL="http://localhost:36657"
fi

if [ -z "$HUB_CHAIN_ID" ]; then
  echo "HUB_CHAIN_ID is not set, using 'dymension_100-1'"
  HUB_CHAIN_ID="dymension_100-1"
fi

set -x
#Register rollapp
dymd tx rollapp create-rollapp "$ROLLAPP_CHAIN_ID" "$MAX_SEQUENCERS" '{"Addresses":[]}' \
  --from "$DEPLOYER" \
  --keyring-backend test \
  --broadcast-mode block \
  --fees 1dym \
  -y

set +x
