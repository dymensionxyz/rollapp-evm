#!/bin/bash

EXECUTABLE="rollapp-evm"
KEYRING_PATH="$HOME/.rollapp_evm/sequencer_keys"
KEY_NAME_SEQUENCER="sequencer"

#Register Sequencer
DESCRIPTION="{\"Moniker\":\"${ROLLAPP_CHAIN_ID}-sequencer\",\"Identity\":\"\",\"Website\":\"\",\"SecurityContact\":\"\",\"Details\":\"\"}"
SEQ_PUB_KEY="$($EXECUTABLE dymint show-sequencer)" # sequencing key: for signing blocks on the rollapp chain
BOND_AMOUNT="100000dym"

SEQUENCER_ADDRESS=$(dymd keys show $KEY_NAME_SEQUENCER --address --keyring-backend test --keyring-dir $KEYRING_PATH)
dymd tx bank send $KEY_NAME_ROLLAPP --keyring-backend test $SEQUENCER_ADDRESS 100000000000000000000000adym --node $HUB_RPC_URL --chain-id $HUB_CHAIN_ID

set -x
dymd tx sequencer create-sequencer "$SEQ_PUB_KEY" "$ROLLAPP_CHAIN_ID" "$DESCRIPTION" "$BOND_AMOUNT" \
  --from "$KEY_NAME_SEQUENCER" \
  --keyring-dir "$KEYRING_PATH" \
  --keyring-backend test \
  --broadcast-mode block \
  --fees 1dym \
  --node "$HUB_RPC_URL" \
  --chain-id "$HUB_CHAIN_ID" -y
set +x
