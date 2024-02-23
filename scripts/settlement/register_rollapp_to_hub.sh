#!/bin/bash

KEYRING_PATH="$HOME/.rollapp_evm/sequencer_keys"
KEY_NAME_SEQUENCER="sequencer"

MAX_SEQUENCERS=5

#Register rollapp 
dymd tx rollapp create-rollapp "$ROLLAPP_CHAIN_ID" "$MAX_SEQUENCERS" '{"Addresses":[]}' \
  --from "$KEY_NAME_SEQUENCER" \
  --keyring-backend test \
  --keyring-dir "$KEYRING_PATH" \
  --broadcast-mode block \
  --gas-prices=100000000adym
