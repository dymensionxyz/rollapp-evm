#!/bin/bash

set -e

# check if ROLLAPP_CHAIN_ID env var is set
if [ -z "$ROLLAPP_CHAIN_ID" ]; then
  export ROLLAPP_CHAIN_ID="rollex_1235-1"
fi
echo "Using $ROLLAPP_CHAIN_ID as ROLLAPP_CHAIN_ID"

if [ -z "$HUB_CHAIN_ID" ]; then
  export HUB_CHAIN_ID="dymension_100-1"
fi
echo "Using $HUB_CHAIN_ID as HUB_CHAIN_ID"

if [ -z "$HUB_RPC_URL" ]; then
  export HUB_RPC_URL="http://127.0.0.1:36657"
fi
echo "Using $HUB_RPC_URL as HUB_RPC_URL"

MNEMONIC="pink draw film undo drama horror eternal hill team spin dolphin crane essay boost couple cereal jungle crime visa record bean knock giggle recycle"

echo "Set the environment variables"
# export ROLLAPP_CHAIN_ID="rollex_1450-1"
export KEY_NAME_ROLLAPP="rol-user"
export SETTLEMENT_KEY_NAME="local-user"
export BASE_DENOM="alxx"
export DENOM=$(echo "$BASE_DENOM" | sed 's/^.//')
export MONIKER="$ROLLAPP_CHAIN_ID-sequencer"
export ROLLAPP_HOME_DIR="$HOME/.rollapp_evm"
export HUB_HOME_DIR="$HOME/.dymension"
export ROLLAPP_SETTLEMENT_INIT_DIR_PATH="$HOME/.rollapp_evm/init"

echo "Remove the existing directories"
rm -rf $ROLLAPP_HOME_DIR
rm -rf $HUB_HOME_DIR

echo "Run the init.sh script"
sh ./init.sh

unset -e

sed -i '' 's/settlement_layer.*/settlement_layer = "dymension"/' ${ROLLAPP_HOME_DIR}/config/dymint.toml
sed -i '' "s/node_address.*/node_address = \"$(echo "$HUB_RPC_URL" | sed 's/\//\\\//g')\"/" ${ROLLAPP_HOME_DIR}/config/dymint.toml

sed -i 's/settlement_layer.*/settlement_layer = "dymension"/' ${ROLLAPP_HOME_DIR}/config/dymint.toml
sed -i "s/node_address.*/node_address = \"$(echo "$HUB_RPC_URL" | sed 's/\//\\\//g')\"/" ${ROLLAPP_HOME_DIR}/config/dymint.toml

set -e

echo "Import the local-user key to the dymd keyring"
echo "$MNEMONIC" | dymd keys add "$SETTLEMENT_KEY_NAME" --recover --keyring-backend test

echo "Generate denom metadata"
sh settlement/generate_denom_metadata.sh
echo "Add genesis accounts"
sh settlement/add_genesis_accounts.sh
echo "Register rollapp to the hub"
sh settlement/register_rollapp_to_hub.sh
echo "Register sequencer to the hub"
sh settlement/register_sequencer_to_hub.sh

echo "Update the genesis file"
sh update_genesis_file.sh

echo "Start the rollapp-evm"
# rollapp-evm start