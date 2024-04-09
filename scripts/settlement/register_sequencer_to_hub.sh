#!/bin/bash

EXECUTABLE="rollapp-evm"
KEYRING_PATH="$ROLLAPP_HOME_DIR/sequencer_keys"
KEY_NAME_SEQUENCER="sequencer"
FUND_AMOUNT="1100dym"
BOND_AMOUNT="1000dym"

#Register Sequencer
DESCRIPTION="{\"Moniker\":\"${ROLLAPP_CHAIN_ID}-sequencer\",\"Identity\":\"\",\"Website\":\"\",\"SecurityContact\":\"\",\"Details\":\"\"}"
SEQ_PUB_KEY="$($EXECUTABLE dymint show-sequencer)" # sequencing key: for signing blocks on the rollapp chain

echo "Add the sequencer key to the ~/.rollapp_evm/sequencer_keys directory"
dymd keys add $KEY_NAME_SEQUENCER --keyring-dir $KEYRING_PATH --keyring-backend test
SEQUENCER_ADDRESS=$(dymd keys show $KEY_NAME_SEQUENCER --address --keyring-backend test --keyring-dir $KEYRING_PATH)
echo "Fund the sequencer from the local-user account"
dymd tx bank send local-user --keyring-backend test $SEQUENCER_ADDRESS $FUND_AMOUNT --node $HUB_RPC_URL --chain-id $HUB_CHAIN_ID --fees 1dym -b block -y

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
