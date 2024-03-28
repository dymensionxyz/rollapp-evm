#!/bin/bash

if [ ! -d "$ROLLAPP_SETTLEMENT_INIT_DIR_PATH" ]; then
  mkdir -p "$ROLLAPP_SETTLEMENT_INIT_DIR_PATH"
  echo "Creating the ROLLAPP_SETTLEMENT_INIT_DIR_PATH: $ROLLAPP_SETTLEMENT_INIT_DIR_PATH"
else
  echo "ROLLAPP_SETTLEMENT_INIT_DIR_PATH already exists: $ROLLAPP_SETTLEMENT_INIT_DIR_PATH"
fi

dymd keys add alice-genesis --keyring-backend test
dymd keys add bob-genesis --keyring-backend test

tee "$ROLLAPP_SETTLEMENT_INIT_DIR_PATH/genesis_accounts.json" >/dev/null <<EOF
[
  {"amount":
      {"amount":"10000000000000000000000","denom":"a${DENOM}"},
      "address": "$(dymd keys show bob-genesis --keyring-backend test --output json | jq -r .address)"
  },
  {"amount":
      {"amount":"50000000000000000000000","denom":"a${DENOM}"},
      "address":"$(dymd keys show alice-genesis --keyring-backend test --output json | jq -r .address)"
    }
]
EOF
