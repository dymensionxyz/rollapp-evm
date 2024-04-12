#!/bin/bash

if [ ! -d "$ROLLAPP_SETTLEMENT_INIT_DIR_PATH" ]; then
  mkdir -p "$ROLLAPP_SETTLEMENT_INIT_DIR_PATH"
  echo "Creating the ROLLAPP_SETTLEMENT_INIT_DIR_PATH: $ROLLAPP_SETTLEMENT_INIT_DIR_PATH"
else
  echo "ROLLAPP_SETTLEMENT_INIT_DIR_PATH already exists: $ROLLAPP_SETTLEMENT_INIT_DIR_PATH"
fi

dymd keys add alice-genesis
dymd keys add bob-genesis

tee "$ROLLAPP_SETTLEMENT_INIT_DIR_PATH/genesis_accounts.json" >/dev/null <<EOF
[
  {"amount":
      {"amount":"10000000000000000000000","denom":"${BASE_DENOM}"},
      "address": "$(dymd keys show -a bob-genesis)"
  },
  {"amount":
      {"amount":"50000000000000000000000","denom":"${BASE_DENOM}"},
      "address":"$(dymd keys show -a alice-genesis)"
    }
]
EOF
