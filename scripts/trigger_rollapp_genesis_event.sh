#!/bin/bash

echo "Triggering genesis event on rollapp"
ROLLAPP_CHANNEL_ID=$(rly q channels "$ROLLAPP_CHAIN_ID" | jq -r 'select(.state == "STATE_OPEN") | .channel_id' | tail -n 1)
"$EXECUTABLE" tx hubgenesis genesis-event "$HUB_CHAIN_ID" "$ROLLAPP_CHANNEL_ID" --from "$KEY_NAME_ROLLAPP" --keyring-backend test --broadcast-mode block

