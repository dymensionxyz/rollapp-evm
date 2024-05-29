#!/bin/bash

if [ ! -d "$ROLLAPP_SETTLEMENT_INIT_DIR_PATH" ]; then
  mkdir -p "$ROLLAPP_SETTLEMENT_INIT_DIR_PATH"
  echo "Creating the ROLLAPP_SETTLEMENT_INIT_DIR_PATH: $ROLLAPP_SETTLEMENT_INIT_DIR_PATH"
else
  echo "ROLLAPP_SETTLEMENT_INIT_DIR_PATH already exists: $ROLLAPP_SETTLEMENT_INIT_DIR_PATH"
fi

# use concrete mnemonics so the addresses are always the same, just for convenience and recognisability
ALICE_MNEMONIC="mimic ten evoke card crowd upset tragic race borrow final vibrant gesture armed alley figure orange shock strike surge jaguar deposit hockey erosion taste"
BOB_MNEMONIC="matrix venture pair label proud ignore manual crunch brand board welcome suspect purity steak melt atom stadium vanish bullet hill angry bulk visa analyst"
echo $ALICE_MNEMONIC |  dymd keys add alice-genesis --keyring-backend test --keyring-dir ${ROLLAPP_HOME_DIR} --recover
echo $BOB_MNEMONIC |  keys add bob-genesis --keyring-backend test --keyring-dir ${ROLLAPP_HOME_DIR} --recover

tee "$ROLLAPP_SETTLEMENT_INIT_DIR_PATH/genesis_accounts.json" >/dev/null <<EOF
[
  {"amount":
      {"amount":"10000000000000000000000","denom":"${BASE_DENOM}"},
      "address": "$(dymd keys show -a bob-genesis --keyring-backend test --keyring-dir ${ROLLAPP_HOME_DIR})"
  },
  {"amount":
      {"amount":"50000000000000000000000","denom":"${BASE_DENOM}"},
      "address":"$(dymd keys show -a alice-genesis --keyring-backend test --keyring-dir ${ROLLAPP_HOME_DIR})"
    }
]
EOF
