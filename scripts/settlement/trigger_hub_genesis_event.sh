#!/bin/bash

SETTLEMENT_EXECUTABLE="dymd"
HUB_PERMISSIONED_KEY=${HUB_PERMISSIONED_KEY-"$HUB_KEY_WITH_FUNDS"}

echo "Triggering genesis event on hub"
ROLLAPP_CHANNEL_ID=$(rly q channels "$ROLLAPP_CHAIN_ID" | jq -r 'select(.state == "STATE_OPEN") | .channel_id' | tail -n 1)
HUB_CHANNEL_ID=$(rollapp-evm q ibc channel end transfer ${ROLLAPP_CHANNEL_ID} | grep "channel_id:" | awk '{print $2}')

$SETTLEMENT_EXECUTABLE tx rollapp genesis-event ${ROLLAPP_CHAIN_ID} ${HUB_CHANNEL_ID} --from $HUB_PERMISSIONED_KEY --keyring-backend test --broadcast-mode block --gas auto --gas-adjustment 1.5 --fees 1dym