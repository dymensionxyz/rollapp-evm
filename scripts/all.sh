#!/bin/bash

export ROLLAPP_CHAIN_ID="rollex_1443-1"
export KEY_NAME_ROLLAPP="rol-user"
export SETTLEMENT_KEY_NAME="local-user"
export BASE_DENOM="alex"
export DENOM=$(echo "$BASE_DENOM" | sed 's/^.//')
export MONIKER="$ROLLAPP_CHAIN_ID-sequencer"
export BOND_AMOUNT="1000dym"
export ROLLAPP_HOME_DIR="$HOME/.rollapp_evm"
export HUB_HOME_DIR="$HOME/.dymension"
export ROLLAPP_SETTLEMENT_INIT_DIR_PATH="$HOME/.rollapp_evm/init"

export HUB_RPC_ENDPOINT="https://rpc.hwpd.noisnemyd.xyz"
export HUB_RPC_PORT="443"
export HUB_RPC_URL="https://rpc.hwpd.noisnemyd.xyz:443"
export HUB_CHAIN_ID="dymension_1405-1"

rm -rf $ROLLAPP_HOME_DIR
rm -rf $HUB_HOME_DIR

sh ./init.sh

dymd keys add $SETTLEMENT_KEY_NAME --recover --keyring-backend test
dymd keys add sequencer --keyring-dir ~/.rollapp_evm/sequencer_keys --keyring-backend test
SEQUENCER_ADDR=`dymd keys show sequencer --address --keyring-backend test --keyring-dir ~/.rollapp_evm/sequencer_keys`
dymd tx bank send $KEY_NAME_ROLLAPP $SEQUENCER_ADDR ${BOND_AMOUNT} --keyring-backend test --broadcast-mode block --fees 1dym -y --node ${HUB_RPC_URL} --chain-id $HUB_CHAIN_ID

sh settlement/generate_denom_metadata.sh
sh settlement/add_genesis_accounts.sh
sh settlement/register_rollapp_to_hub.sh
sh settlement/register_sequencer_to_hub.sh


sed -i '' 's/settlement_layer.*/settlement_layer = "dymension"/' ${ROLLAPP_HOME_DIR}/config/dymint.toml
sed -i '' 's/node_address.*/node_address = "https:\/\/rpc.hwpd.noisnemyd.xyz:443"/' ${ROLLAPP_HOME_DIR}/config/dymint.toml

sh update_genesis_file.sh

rollapp-evm start