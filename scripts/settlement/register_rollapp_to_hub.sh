#!/bin/bash

MAX_SEQUENCERS=1

# this account must be whitelisted on the hub for permissioned deployment setup
DEPLOYER=${HUB_PERMISSIONED_KEY-"$HUB_KEY_WITH_FUNDS"}

if [ "$HUB_RPC_URL" = "" ]; then
  echo "HUB_RPC_URL is not set, using 'http://localhost:36657'"
  HUB_RPC_URL="http://localhost:36657"
fi

if [ "$HUB_CHAIN_ID" = "" ]; then
  echo "HUB_CHAIN_ID is not set, using 'dymension_100-1'"
  HUB_CHAIN_ID="dymension_100-1"
fi

set -x
dymd tx rollapp create-rollapp "$ROLLAPP_CHAIN_ID" "$MAX_SEQUENCERS" "{\"Addresses\":[\"${SEQUENCER_ADDR}\"]}" \
	"$ROLLAPP_HOME_DIR"/init/denommetadata.json \
	--genesis-accounts-path "$ROLLAPP_HOME_DIR"/init/genesis_accounts.json \
	--from "$DEPLOYER" \
	--keyring-backend test --broadcast-mode block \
	--fees 1dym
set +x
