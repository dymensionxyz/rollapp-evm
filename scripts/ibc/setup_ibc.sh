#!/bin/bash

if [ -z "$ROLLAPP_CHAIN_ID" ]; then
  echo "ROLLAPP_CHAIN_ID is not set"
  exit 1
fi

if [ -z "$HUB_CHAIN_ID" ]; then
  echo "HUB_CHAIN_ID is not set"
  exit 1
fi

if [ -z "$HUB_RPC_URL" ]; then
  echo "HUB_RPC_URL is not set"
  exit 1
fi

if [ -z "$BASE_DENOM" ]; then
  echo "BASE_DENOM is not set"
  exit 1
fi

if [ -z "$KEY_NAME_ROLLAPP" ]; then
  echo "KEY_NAME_ROLLAPP is not set"
  exit 1
fi

BASEDIR=$(dirname "$0")

IBC_PORT=transfer
IBC_VERSION=ics20-1

RELAYER_EXECUTABLE="rly"

# settlement config
SETTLEMENT_EXECUTABLE="dymd"
SETTLEMENT_CHAIN_ID=$HUB_CHAIN_ID
SETTLEMENT_KEY_NAME_GENESIS="local-user"

EXECUTABLE="rollapp-evm"

RELAYER_KEY_FOR_ROLLAP="relayer-rollapp-key"
RELAYER_KEY_FOR_HUB="relayer-hub-key"
RELAYER_PATH="hub-rollapp"
ROLLAPP_RPC_FOR_RELAYER="http://127.0.0.1:26657"
SETTLEMENT_RPC_FOR_RELAYER=$HUB_RPC_URL


if ! command -v $RELAYER_EXECUTABLE >/dev/null; then
  echo "$RELAYER_EXECUTABLE does not exist"
  echo "please run make install of github.com/dymensionxyz/dymension-relayer"
  exit 1
fi

# --------------------------------- rly init --------------------------------- #
RLY_PATH="$HOME/.relayer"
RLY_CONFIG_FILE="$RLY_PATH/config/config.yaml"
ROLLAPP_IBC_CONF_FILE="$BASEDIR/rollapp.json"
HUB_IBC_CONF_FILE="$BASEDIR/hub.json"

echo '# -------------------------- removing rly config ------------------------- #'
rm -rf "$RLY_PATH"

echo '# -------------------------- initializing rly config ------------------------- #'
rly config init

echo '# ------------------------- adding chains to rly config ------------------------- #'
tmp=$(mktemp)

jq --arg key "$RELAYER_KEY_FOR_ROLLAP" '.value.key = $key' $ROLLAPP_IBC_CONF_FILE > "$tmp" && mv "$tmp" $ROLLAPP_IBC_CONF_FILE
jq --arg chain "$ROLLAPP_CHAIN_ID" '.value."chain-id" = $chain' $ROLLAPP_IBC_CONF_FILE > "$tmp" && mv "$tmp" $ROLLAPP_IBC_CONF_FILE
jq --arg rpc "$ROLLAPP_RPC_FOR_RELAYER" '.value."rpc-addr" = $rpc' $ROLLAPP_IBC_CONF_FILE > "$tmp" && mv "$tmp" $ROLLAPP_IBC_CONF_FILE
jq --arg denom "0.0$BASE_DENOM" '.value."gas-prices" = $denom' $ROLLAPP_IBC_CONF_FILE > "$tmp" && mv "$tmp" $ROLLAPP_IBC_CONF_FILE

jq --arg key "$RELAYER_KEY_FOR_HUB" '.value.key = $key' $HUB_IBC_CONF_FILE > "$tmp" && mv "$tmp" $HUB_IBC_CONF_FILE
jq --arg chain "$SETTLEMENT_CHAIN_ID" '.value."chain-id" = $chain' $HUB_IBC_CONF_FILE > "$tmp" && mv "$tmp" $HUB_IBC_CONF_FILE
jq --arg rpc "$SETTLEMENT_RPC_FOR_RELAYER" '.value."rpc-addr" = $rpc' $HUB_IBC_CONF_FILE > "$tmp" && mv "$tmp" $HUB_IBC_CONF_FILE

rly chains add --file "$ROLLAPP_IBC_CONF_FILE" "$ROLLAPP_CHAIN_ID"
rly chains add --file "$HUB_IBC_CONF_FILE" "$SETTLEMENT_CHAIN_ID"

echo '# -------------------------------- creating keys ------------------------------- #'
rly keys add "$ROLLAPP_CHAIN_ID" "$RELAYER_KEY_FOR_ROLLAP" --coin-type 60
rly keys add "$SETTLEMENT_CHAIN_ID" "$RELAYER_KEY_FOR_HUB" --coin-type 60

RLY_HUB_ADDR=$(rly keys show "$SETTLEMENT_CHAIN_ID")
RLY_ROLLAPP_ADDR=$(rly keys show "$ROLLAPP_CHAIN_ID")

echo "# ------------------------------- balance of rly account on hub [$RLY_HUB_ADDR]------------------------------ #"
$SETTLEMENT_EXECUTABLE q bank balances "$(rly keys show "$SETTLEMENT_CHAIN_ID")" --node "$SETTLEMENT_RPC_FOR_RELAYER"
echo "From within the hub node: \"$SETTLEMENT_EXECUTABLE tx bank send $SETTLEMENT_KEY_NAME_GENESIS $RLY_HUB_ADDR 100dym --keyring-backend test --broadcast-mode block --fees 1dym --node "$SETTLEMENT_RPC_FOR_RELAYER" --chain-id dymension_1405-1 -y \""

echo "# ------------------------------- balance of rly account on rollapp [$RLY_ROLLAPP_ADDR] ------------------------------ #"
$EXECUTABLE q bank balances "$(rly keys show "$ROLLAPP_CHAIN_ID")" --node "$ROLLAPP_RPC_FOR_RELAYER"
echo "From within the rollapp node: \"$EXECUTABLE tx bank send $KEY_NAME_ROLLAPP $RLY_ROLLAPP_ADDR 10000000000000000000$BASE_DENOM --keyring-backend test --broadcast-mode block -y\""

echo "waiting to fund accounts. Press to continue..."
read -r answer

echo '# -------------------------------- creating IBC link ------------------------------- #'

rly paths new "$ROLLAPP_CHAIN_ID" "$SETTLEMENT_CHAIN_ID" "$RELAYER_PATH" --src-port "$IBC_PORT" --dst-port "$IBC_PORT" --version "$IBC_VERSION"


rly tx link "$RELAYER_PATH" --src-port "$IBC_PORT" --dst-port "$IBC_PORT" --version "$IBC_VERSION"
# Channel is currently not created in the tx link since we changed the relayer to support on demand blocks
# Which messed up with channel creation as part of tx link.
rly tx channel "$RELAYER_PATH"


echo '# -------------------------------- IBC channel established ------------------------------- #'
ROLLAPP_CHANNEL_ID=$(rly q channels "$ROLLAPP_CHAIN_ID" | jq -r 'select(.state == "STATE_OPEN") | .channel_id' | tail -n 1)
HUB_CHANNEL_ID=$(rly q channels "$SETTLEMENT_CHAIN_ID" | jq -r 'select(.state == "STATE_OPEN") | .channel_id' | tail -n 1)

echo "ROLLAPP_CHANNEL_ID: $ROLLAPP_CHANNEL_ID"
echo "HUB_CHANNEL_ID: $HUB_CHANNEL_ID"
