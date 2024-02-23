#!/bin/bash

EXECUTABLE="rollapp-evm"
KEYRING_PATH="$HOME/.rollapp_evm/sequencer_keys"
KEY_NAME_SEQUENCER="sequencer"


#Register Sequencer
DESCRIPTION="{\"Moniker\":\"myrollapp-sequencer\",\"Identity\":\"\",\"Website\":\"\",\"SecurityContact\":\"\",\"Details\":\"\"}";
SEQ_PUB_KEY="$($EXECUTABLE dymint show-sequencer)"

dymd tx sequencer create-sequencer "$SEQ_PUB_KEY" "$ROLLAPP_CHAIN_ID" "$DESCRIPTION" \
  --from "$KEY_NAME_SEQUENCER" \
  --keyring-dir "$KEYRING_PATH" \
  --keyring-backend test \
  --broadcast-mode block \
  --gas-prices=1000000000adym
