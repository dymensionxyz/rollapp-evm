#!/bin/bash

BASEDIR=$(dirname "$0")

IBC_PORT=transfer
IBC_VERSION=ics20-1

RELAYER_EXECUTABLE="rly"

# settlement config
SETTLEMENT_EXECUTABLE="dymd"
SETTLEMENT_CHAIN_ID=$("$SETTLEMENT_EXECUTABLE" config | jq -r '."chain-id"')
SETTLEMENT_RPC_FOR_RELAYER=$("$SETTLEMENT_EXECUTABLE" config | jq -r '."node"')

SETTLEMENT_KEY_NAME_GENESIS="$HUB_KEY_WITH_FUNDS"

# rollapp config
ROLLAPP_CHAIN_ID=$("$EXECUTABLE" config | jq -r '."chain-id"')
ROLLAPP_RPC_FOR_RELAYER=$("$EXECUTABLE" config | jq -r '."node"')

RELAYER_KEY_FOR_ROLLAP="relayer-rollapp-key"
RELAYER_KEY_FOR_HUB="relayer-hub-key"
RELAYER_PATH="hub-rollapp"

if ! command -v "$RELAYER_EXECUTABLE" >/dev/null; then
  echo "$RELAYER_EXECUTABLE does not exist"
  echo "please run make install of github.com/dymensionxyz/dymension-relayer"
  exit 1
fi

# --------------------------------- rly init --------------------------------- #
RLY_PATH="$HOME/.relayer"
RLY_CONFIG_FILE="$RLY_PATH/config/config.yaml"
ROLLAPP_IBC_CONF_FILE="$BASEDIR/rollapp.json"
HUB_IBC_CONF_FILE="$BASEDIR/hub.json"

if [ -f "$RLY_CONFIG_FILE" ]; then
  printf "======================================================================================================\n"
  echo "A rly config file already exists. Overwrite? (y/N)"
  printf "======================================================================================================\n"
  read -r answer
  if [[ "$answer" == "Y" || "$answer" == "y" ]]; then
    rm -rf "$RLY_PATH"
  fi
fi

echo '--------------------------------- Initializing rly config... --------------------------------'
rly config init

echo '--------------------------------- Adding chains to rly config.. --------------------------------'
tmp=$(mktemp)

jq --arg key "$RELAYER_KEY_FOR_ROLLAP" '.value.key = $key' "$ROLLAPP_IBC_CONF_FILE" >"$tmp" && mv "$tmp" "$ROLLAPP_IBC_CONF_FILE"
jq --arg chain "$ROLLAPP_CHAIN_ID" '.value."chain-id" = $chain' "$ROLLAPP_IBC_CONF_FILE" >"$tmp" && mv "$tmp" "$ROLLAPP_IBC_CONF_FILE"
jq --arg bech "$BECH32_PREFIX" '.value["account-prefix"] = $bech' "$ROLLAPP_IBC_CONF_FILE" >"$tmp" && mv "$tmp" "$ROLLAPP_IBC_CONF_FILE"
jq --arg rpc "$ROLLAPP_RPC_FOR_RELAYER" '.value."rpc-addr" = $rpc' "$ROLLAPP_IBC_CONF_FILE" >"$tmp" && mv "$tmp" "$ROLLAPP_IBC_CONF_FILE"
jq --arg denom "0.0$BASE_DENOM" '.value."gas-prices" = $denom' "$ROLLAPP_IBC_CONF_FILE" >"$tmp" && mv "$tmp" "$ROLLAPP_IBC_CONF_FILE"

jq --arg key "$RELAYER_KEY_FOR_HUB" '.value.key = $key' "$HUB_IBC_CONF_FILE" >"$tmp" && mv "$tmp" "$HUB_IBC_CONF_FILE"
jq --arg chain "$SETTLEMENT_CHAIN_ID" '.value."chain-id" = $chain' "$HUB_IBC_CONF_FILE" >"$tmp" && mv "$tmp" "$HUB_IBC_CONF_FILE"
jq --arg rpc "$SETTLEMENT_RPC_FOR_RELAYER" '.value."rpc-addr" = $rpc' "$HUB_IBC_CONF_FILE" >"$tmp" && mv "$tmp" "$HUB_IBC_CONF_FILE"

rly chains add --file "$ROLLAPP_IBC_CONF_FILE" "$ROLLAPP_CHAIN_ID"
rly chains add --file "$HUB_IBC_CONF_FILE" "$SETTLEMENT_CHAIN_ID"

echo -e '--------------------------------- Creating keys for rly... --------------------------------'

rly keys add "$ROLLAPP_CHAIN_ID" "$RELAYER_KEY_FOR_ROLLAP" --coin-type 60
rly keys add "$SETTLEMENT_CHAIN_ID" "$RELAYER_KEY_FOR_HUB" --coin-type 60

RLY_HUB_ADDR=$(rly keys show "$SETTLEMENT_CHAIN_ID")
RLY_ROLLAPP_ADDR=$(rly keys show "$ROLLAPP_CHAIN_ID")

echo '--------------------------------- Funding rly account on hub ['"$RLY_HUB_ADDR"']... --------------------------------'
DYM_BALANCE=$(${SETTLEMENT_EXECUTABLE} q bank balances ${RLY_HUB_ADDR} -o json | jq -r '.balances[0].amount')

if [ "$(echo "$DYM_BALANCE >= 100000000000000000000" | bc)" -eq 1 ]; then
  echo "${RLY_HUB_ADDR} already funded"
else
  "$SETTLEMENT_EXECUTABLE" tx bank send "$SETTLEMENT_KEY_NAME_GENESIS" "$RLY_HUB_ADDR" 100dym --keyring-backend test --broadcast-mode block --fees 1dym --node "$SETTLEMENT_RPC_FOR_RELAYER" -y
fi

echo '--------------------------------- Funding rly account on rollapp ['"$RLY_ROLLAPP_ADDR"'].. --------------------------------'

RA_BALANCE=$(${EXECUTABLE} q bank balances ${RLY_ROLLAPP_ADDR} -o json | jq -r '.balances[0].amount')

if [ "$(echo "$RA_BALANCE >= 100000000000000000000" | bc)" -eq 1 ]; then
  echo "${RLY_ROLLAPP_ADDR} already funded"
else
  "$EXECUTABLE" tx bank send "$KEY_NAME_ROLLAPP" "$RLY_ROLLAPP_ADDR" 100000000000000000000"$BASE_DENOM" --keyring-backend test --broadcast-mode block -y --fees 4000000000$BASE_DENOM
fi

echo '--------------------------------- Creating IBC path... --------------------------------'

rly paths new "$ROLLAPP_CHAIN_ID" "$SETTLEMENT_CHAIN_ID" "$RELAYER_PATH" --src-port "$IBC_PORT" --dst-port "$IBC_PORT" --version "$IBC_VERSION"

rly tx link "$RELAYER_PATH" --src-port "$IBC_PORT" --dst-port "$IBC_PORT" --version "$IBC_VERSION"
# Channel is currently not created in the tx link since we changed the relayer to support on demand blocks
# Which messed up with channel creation as part of tx link.
rly tx channel "$RELAYER_PATH"

echo '# -------------------------------- IBC channel established ------------------------------- #'
echo "Channel Information:"
echo "$(rly q channels "$ROLLAPP_CHAIN_ID" | jq '{ "rollapp-channel": .channel_id, "hub-channel": .counterparty.channel_id }')"
