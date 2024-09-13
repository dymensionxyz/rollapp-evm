#!/bin/bash
EXECUTABLE=$(which rollapp-evm)

if ! command -v "$EXECUTABLE" >/dev/null; then
  echo "$EXECUTABLE does not exist"
  echo "please run make install"
  exit 1
fi

if [ "$BECH32_PREFIX" = "" ]; then
  echo "BECH32_PREFIX is not set"
  exit 1
fi

if [ "$BASE_DENOM" = "" ]; then
  echo "BASE_DENOM is not set"
  exit 1
fi

if [ "$HUB_KEY_WITH_FUNDS" = "" ]; then
  echo "HUB_KEY_WITH_FUNDS is not set"
  exit 1
fi

if [ "$KEY_NAME_ROLLAPP" = "" ]; then
  echo "KEY_NAME_ROLLAPP is not set"
  exit 1
fi

if [ "$HUB_REST_URL" = "" ]; then
  echo "HUB_REST_URL is not set"
  exit 1
fi

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

RELAYER_KEY_FOR_ROLLAPP="relayer-rollapp-key"
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

dasel put -f "$ROLLAPP_IBC_CONF_FILE" '.value.key' -v "$RELAYER_KEY_FOR_ROLLAPP"
dasel put -f "$ROLLAPP_IBC_CONF_FILE" '.value.chain-id' -v "$ROLLAPP_CHAIN_ID"
dasel put -f "$ROLLAPP_IBC_CONF_FILE" '.value.account-prefix' -v "$BECH32_PREFIX"
dasel put -f "$ROLLAPP_IBC_CONF_FILE" '.value.rpc-addr' -v "$ROLLAPP_RPC_FOR_RELAYER"
dasel put -f "$ROLLAPP_IBC_CONF_FILE" '.value.gas-prices' -v "1000000000$BASE_DENOM"

dasel put -f "$HUB_IBC_CONF_FILE"  '.value.key' -v "$RELAYER_KEY_FOR_HUB"
dasel put -f "$HUB_IBC_CONF_FILE"  '.value.chain-id' -v "$SETTLEMENT_CHAIN_ID"
dasel put -f "$HUB_IBC_CONF_FILE"  '.value.rpc-addr' -v "$SETTLEMENT_RPC_FOR_RELAYER"

rly chains add --file "$ROLLAPP_IBC_CONF_FILE" "$ROLLAPP_CHAIN_ID"
rly chains add --file "$HUB_IBC_CONF_FILE" "$SETTLEMENT_CHAIN_ID"

echo -e '--------------------------------- Setting min-loop-duration to 100ms in rly config... --------------------------------'
sed -i.bak '/min-loop-duration:/s/.*/            min-loop-duration: 100ms/' "$RLY_CONFIG_FILE"

echo -e '--------------------------------- Creating keys for rly... --------------------------------'

rly keys add "$ROLLAPP_CHAIN_ID" "$RELAYER_KEY_FOR_ROLLAPP" --coin-type 60
rly keys add "$SETTLEMENT_CHAIN_ID" "$RELAYER_KEY_FOR_HUB" --coin-type 60

RLY_HUB_ADDR=$(rly keys show "$SETTLEMENT_CHAIN_ID")

echo '--------------------------------- Funding rly account on hub ['"$RLY_HUB_ADDR"']... --------------------------------'
DYM_BALANCE=$("$SETTLEMENT_EXECUTABLE" q bank balances "$RLY_HUB_ADDR" -o json | jq -r '.balances[0].amount')

if [ "$(echo "$DYM_BALANCE >= 100000000000000000000" | bc)" -eq 1 ]; then
  echo "${RLY_HUB_ADDR} already funded"
else
  "$SETTLEMENT_EXECUTABLE" tx bank send "$SETTLEMENT_KEY_NAME_GENESIS" "$RLY_HUB_ADDR" 100dym --keyring-backend test --fees 1dym --node "$SETTLEMENT_RPC_FOR_RELAYER" -y || exit 1
fi

echo '--------------------------------- Creating IBC path... --------------------------------'

rly paths new "$SETTLEMENT_CHAIN_ID" "$ROLLAPP_CHAIN_ID" "$RELAYER_PATH" --src-port "$IBC_PORT" --dst-port "$IBC_PORT" --version "$IBC_VERSION"

dasel put -r yaml -f "$RLY_CONFIG_FILE" "chains.$SETTLEMENT_CHAIN_ID.value.http-addr" -v "$HUB_REST_URL";
dasel put -r yaml -f "$RLY_CONFIG_FILE" "chains.$SETTLEMENT_CHAIN_ID.value.is-dym-hub" -v true -t bool;
dasel put -r yaml -f "$RLY_CONFIG_FILE" "chains.$ROLLAPP_CHAIN_ID.value.is-dym-rollapp" -v true -t bool;
dasel put -r yaml -f "$RLY_CONFIG_FILE" "chains.$ROLLAPP_CHAIN_ID.value.trust-period" -v "240h"; # 10 days

rly tx link "$RELAYER_PATH" --src-port "$IBC_PORT" --dst-port "$IBC_PORT" --version "$IBC_VERSION" --max-clock-drift 70m

echo '# -------------------------------- IBC channel established ------------------------------- #'
echo "Channel Information:"

channel_info=$(rly q channels "$ROLLAPP_CHAIN_ID" | jq '{ "rollapp-channel": .channel_id, "hub-channel": .counterparty.channel_id }')
rollapp_channel=$(echo "$channel_info" | jq -r '.["rollapp-channel"]')
hub_channel=$(echo "$channel_info" | jq -r '.["hub-channel"]')

echo "$channel_info"

echo -e '--------------------------------- Set channel-filter --------------------------------'

if [ "$rollapp_channel" = "" ] || [ "$hub_channel" = "" ]; then
  echo "Both channels must be provided. Something is wrong. Exiting."
  exit 1
fi

sed -i.bak '/rule:/s/.*/            rule: "allowlist"/' "$RLY_CONFIG_FILE"
sed -i.bak '/channel-list:/s/.*/            channel-list: ["'"$rollapp_channel"'","'"$hub_channel"'"]/' "$RLY_CONFIG_FILE"
echo "Config file updated successfully."
