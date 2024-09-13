#!/bin/bash

# this account must be whitelisted on the hub for permissioned deployment setup
DEPLOYER=${HUB_PERMISSIONED_KEY-"$HUB_KEY_WITH_FUNDS"}

if [ "$EXECUTABLE" = "" ]; then
  DEFAULT_EXECUTABLE=$(which dymd)

  if [ "$DEFAULT_EXECUTABLE" = "" ]; then
    echo "dymd not found in PATH. Exiting."
    exit 1
  fi
  echo "EXECUTABLE is not set, using '${DEFAULT_EXECUTABLE}'"
  EXECUTABLE=$DEFAULT_SEQUENCER_KEY_PATH
fi

if [ "$SEQUENCER_KEY_PATH" = "" ]; then
  DEFAULT_SEQUENCER_KEY_PATH="${ROLLAPP_HOME_DIR}/sequencer_keys"
  echo "SEQUENCER_KEY_PATH is not set, using '${DEFAULT_SEQUENCER_KEY_PATH}'"
  SEQUENCER_KEY_PATH=$DEFAULT_SEQUENCER_KEY_PATH
fi

if [ "$SEQUENCER_KEY_NAME" = "" ]; then
  DEFAULT_SEQUENCER_KEY_NAME="sequencer"
  echo "SEQUENCER_KEY_PATH is not set, using '${DEFAULT_SEQUENCER_KEY_PATH}'"
  SEQUENCER_KEY_NAME=$DEFAULT_SEQUENCER_KEY_NAME
fi

if [ "$HUB_RPC_URL" = "" ]; then
  echo "HUB_RPC_URL is not set, using 'http://localhost:36657'"
  HUB_RPC_URL="http://localhost:36657"
fi

if [ "$HUB_CHAIN_ID" = "" ]; then
  echo "HUB_CHAIN_ID is not set, using 'dymension_100-1'"
  HUB_CHAIN_ID="dymension_100-1"
fi

if [ "$ROLLAPP_ALIAS" = "" ]; then
  DEFAULT_ALIAS="${ROLLAPP_CHAIN_ID%%_*}"
  echo "ROLLAPP_ALIAS is not set, using '$DEFAULT_ALIAS'"
  ROLLAPP_ALIAS=$DEFAULT_ALIAS
fi

if [ "$ROLLAPP_HOME_DIR" = "" ]; then
  DEFAULT_ROLLAPP_HOME_DIR=${HOME}/.rollapp_evm
  echo "ROLLAPP_ALIAS is not set, using '$DEFAULT_ROLLAPP_HOME_DIR'"
  ROLLAPP_HOME_DIR=$DEFAULT_ROLLAPP_HOME_DIR
fi

if [ "$BECH32_PREFIX" = "" ]; then
  echo "BECH32_PREFIX is not set, exiting "
  exit 1
fi

if [ "$METADATA_PATH" = "" ]; then
  DEFAULT_METADATA_PATH="${ROLLAPP_HOME_DIR}/init/rollapp-metadata.json"
  echo "METADATA_PATH is not set, using '$DEFAULT_METADATA_PATH"
  METADATA_PATH=$DEFAULT_METADATA_PATH

  if [ ! -f "$METADATA_PATH" ]; then
    echo "${METADATA_PATH} does not exist, would you like to use a dummy metadata file? (y/n)"
    read -r answer

    if [ "$answer" != "${answer#[Yy]}" ]; then
      cat <<EOF > "$METADATA_PATH"
{
  "website": "https://dymension.xyz/",
  "description": "This is a description of the Rollapp.",
  "logo_data_uri": "data:image/jpeg;base64,/000",
  "token_logo_uri": "data:image/jpeg;base64,/000",
  "telegram": "https://t.me/example",
  "x": "https://x.com/dymension"
}
EOF
    else
      echo "You can't register a rollapp without rollapp metadata, please create the ${METADATA_PATH} and run the script again"
      exit 1
    fi
  fi

fi

if [ "$NATIVE_DENOM_PATH" = "" ]; then
  DEFAULT_NATIVE_DENOM_PATH="${ROLLAPP_HOME_DIR}/init/rollapp-native-denom.json"
  echo "NATIVE_DENOM_PATH is not set, using '$DEFAULT_NATIVE_DENOM_PATH"
  NATIVE_DENOM_PATH=$DEFAULT_NATIVE_DENOM_PATH

  if [ ! -f "$NATIVE_DENOM_PATH" ]; then
    echo "${NATIVE_DENOM_PATH} does not exist, would you like to use a dummy native-denom file? (y/n)"
    read -r answer

    if [ "$answer" != "${answer#[Yy]}" ]; then
      cat <<EOF > "$NATIVE_DENOM_PATH"
{
  "display": "DEN",
  "base": "aden",
  "exponent": 18
}
EOF
    else
      echo "You can't register a rollapp without a native denom, please create the ${NATIVE_DENOM_PATH} and run the script again"
      exit 1
    fi
  fi

fi

GENESIS_PATH="${ROLLAPP_HOME_DIR}/config/genesis.json"
GENESIS_HASH=$(sha256sum "$GENESIS_PATH" | awk '{print $1}' | sed 's/[[:space:]]*$//')
SEQUENCER_ADDR=$(dymd keys show "$SEQUENCER_KEY_NAME" --address --keyring-backend test --keyring-dir "$SEQUENCER_KEY_PATH")

echo "deployer" $DEPLOYER;

set -x
"$SETTLEMENT_EXECUTABLE" tx rollapp create-rollapp "$ROLLAPP_CHAIN_ID" "$ROLLAPP_ALIAS" EVM \
  --bech32-prefix "$BECH32_PREFIX" \
  --init-sequencer "$SEQUENCER_ADDR" \
  --genesis-checksum "$GENESIS_HASH" \
  --metadata "$METADATA_PATH" \
  --native-denom "$NATIVE_DENOM_PATH" \
  --initial-supply 1 \
	--from "$DEPLOYER" \
	--keyring-backend test \
  --gas auto --gas-adjustment 1.2 \
	--fees 1dym
set +x
