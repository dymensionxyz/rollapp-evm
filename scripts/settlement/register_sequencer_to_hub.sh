#!/bin/bash

if [ "$SETTLEMENT_EXECUTABLE" = "" ]; then
  DEFAULT_SETTLEMENT_EXECUTABLE=$(which dymd)
  echo "SETTLEMENT_EXECUTABLE is not set, using '${SETTLEMENT_EXECUTABLE}'"
  SETTLEMENT_EXECUTABLE=$DEFAULT_SETTLEMENT_EXECUTABLE

  if [ "$SETTLEMENT_EXECUTABLE" = "" ]; then
    echo "dymension binary not found in PATH. Exiting."
    exit 1
  fi
fi

if [ "$ROLLAPP_EXECUTABLE" = "" ]; then
  DEFAULT_ROLLAPP_EXECUTABLE=$(which rollapp-evm)
  echo "ROLLAPP_EXECUTABLE is not set, using '${DEFAULT_ROLLAPP_EXECUTABLE}'"
  ROLLAPP_EXECUTABLE=$DEFAULT_ROLLAPP_EXECUTABLE

  if [ "$ROLLAPP_EXECUTABLE" = "" ]; then
    echo "rollapp binary not found in PATH. Exiting."
    exit 1
  fi
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

#Register Sequencer
# DESCRIPTION="{\"Moniker\":\"${ROLLAPP_CHAIN_ID}-sequencer\",\"Identity\":\"\",\"Website\":\"\",\"SecurityContact\":\"\",\"Details\":\"\"}"
SEQ_PUB_KEY="$("$ROLLAPP_EXECUTABLE" dymint show-sequencer)"
BOND_AMOUNT="$("$SETTLEMENT_EXECUTABLE" q sequencer params -o json --node "$HUB_RPC_URL" | jq -r '.params.min_bond.amount')$("$SETTLEMENT_EXECUTABLE" q sequencer params -o json --node "$HUB_RPC_URL" | jq -r '.params.min_bond.denom')"

echo "$BOND_AMOUNT"

if [ "$METADATA_PATH" = "" ]; then
  DEFAULT_METADATA_PATH="${ROLLAPP_HOME_DIR}/init/sequencer-metadata.json"
  echo "METADATA_PATH is not set, using '$DEFAULT_METADATA_PATH"
  METADATA_PATH=$DEFAULT_METADATA_PATH

  if [ ! -f "$METADATA_PATH" ]; then
    echo "${METADATA_PATH} does not exist, would you like to use a dummy metadata file? (y/n)"
    read -r answer

    if [ "$answer" != "${answer#[Yy]}" ]; then
      cat <<EOF > "$METADATA_PATH"
{
  "moniker": "Sample Moniker",
  "details": "Some details about the sequencer",
  "p2p_seeds": [
    "seed1.example.com:26656",
    "seed2.example.com:26656"
  ],
  "rpcs": [
    "http://rpc1.example.com:26657",
    "http://rpc2.example.com:26657"
  ],
  "evm_rpcs": [
    "http://evm-rpc1.example.com:8545",
    "http://evm-rpc2.example.com:8545"
  ],
  "rest_api_urls": ["http://restapi.example.com"],
  "explorer_url": "http://explorer.example.com",
  "genesis_urls": [
    "http://genesis1.example.com",
    "http://genesis2.example.com"
  ],
  "contact_details": {
    "website": "http://website.example.com",
    "telegram": "https://t.me/example",
    "x": "https://twitter.com/example"
  },
  "extra_data": "RXh0cmEgZGF0YSBzYW1wbGU=",
  "snapshots": [
    {
      "snapshot_url": "http://snapshot1.example.com",
      "height": 123456,
      "checksum": "d41d8cd98f00b204e9800998ecf8427e"
    },
    {
      "snapshot_url": "http://snapshot2.example.com",
      "height": 789012,
      "checksum": "e2fc714c4727ee9395f324cd2e7f331f"
    }
  ],
  "gas_price": "1000000"
}
EOF
    else
      echo "You can't register a sequencer without sequencer metadata, please create the ${METADATA_PATH} and run the script again"
      exit 1
    fi
  fi
fi
set -x
"$SETTLEMENT_EXECUTABLE" tx sequencer create-sequencer "$SEQ_PUB_KEY" "$ROLLAPP_CHAIN_ID" "$BOND_AMOUNT" "$METADATA_PATH"\
  --from "$SEQUENCER_KEY_NAME" \
  --keyring-dir "$SEQUENCER_KEY_PATH" \
  --keyring-backend test \
  --fees 1dym \
  --gas auto --gas-adjustment 1.2

set +x
