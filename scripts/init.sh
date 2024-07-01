#!/bin/bash
tmp=$(mktemp)
EXECUTABLE=$(which rollapp-evm)

if ! command -v "$EXECUTABLE" >/dev/null; then
  echo "$EXECUTABLE does not exist"
  echo "please run make install"
  exit 1
fi

# ---------------------------- initial parameters ---------------------------- #
# 500,000  is staked
# set BASE_DENOM to the token denomination
STAKING_AMOUNT="500000000000000000000000$BASE_DENOM"

CONFIG_DIRECTORY="$ROLLAPP_HOME_DIR/config"

APP_CONFIG_FILE="$CONFIG_DIRECTORY/app.toml"
GENESIS_FILE="$CONFIG_DIRECTORY/genesis.json"
DENOM=$(echo "$BASE_DENOM" | sed 's/^.//')

# ---------------------------- check variables ---------------------------- #
if [ "$MONIKER" = "" ]; then
    MONIKER="${ROLLAPP_CHAIN_ID}-sequencer" # Default moniker value
fi

if [ "$KEY_NAME_ROLLAPP" = "" ]; then
    KEY_NAME_ROLLAPP="rol-user" # Default key name value
fi

# Default to 1,000,000,000 tokens
if [ "$TOTAL_SUPPLY" = "" ]; then
    TOTAL_SUPPLY="1000000000000000000000000"
fi

if [ "$ROLLAPP_SETTLEMENT_INIT_DIR_PATH" = "" ]; then
  # ROLLAPP_SETTLEMENT_INIT_DIR_PATH is used as a target for generating the necessary
  # configuration files for RollApp initialization, such as denom metadata and genesis account
  # json files
  ROLLAPP_SETTLEMENT_INIT_DIR_PATH="${ROLLAPP_HOME_DIR}/init"
fi

if [ "$ROLLAPP_CHAIN_ID" = "" ]; then
  echo "ROLLAPP_CHAIN_ID is not set"
  exit 1
fi

set_denom() {
  local denom=$1
  local success=true

  dasel put -f "$GENESIS_FILE" '.app_state.mint.params.mint_denom' -v "$denom" || success=false
  dasel put -f "$GENESIS_FILE" '.app_state.staking.params.bond_denom' -v "$denom" || success=false
  dasel put -t string -f "$GENESIS_FILE" '.app_state.gov.deposit_params.min_deposit.[0].denom' -v "$denom" || success=false
  dasel put -f "$GENESIS_FILE" '.app_state.evm.params.evm_denom' -v "$denom" || success=false
  dasel put -f "$GENESIS_FILE" '.app_state.claims.params.claims_denom' -v "$denom" || success=false

  if [ "$success" = false ]; then
    echo "An error occurred. Please refer to README.md"
    exit 1
  fi
}

update_genesis_params() {
  local success=true

  dasel put -f "$GENESIS_FILE" '.app_state.gov.voting_params.voting_period' -v "300s" || success=false
  dasel put -f "$GENESIS_FILE" '.app_state.bank.balances.[0].coins.[0].amount' -v "$TOTAL_SUPPLY" || success=false
  dasel put -f "$GENESIS_FILE" '.app_state.bank.supply.[0].amount' -v "$TOTAL_SUPPLY" || success=false

  if [ "$success" = false ]; then
    echo "An error occurred. Please refer to README.md"
    return 1
  fi
  echo "Successfully updated the genesis file"
}

add_genesis_accounts() {
  local success=true

  ALICE_MNEMONIC="mimic ten evoke card crowd upset tragic race borrow final vibrant gesture armed alley figure orange shock strike surge jaguar deposit hockey erosion taste"
  echo "$ALICE_MNEMONIC" |  dymd keys add genesis-wallet --keyring-backend test --keyring-dir "$ROLLAPP_HOME_DIR" --recover

  tee "$ROLLAPP_SETTLEMENT_INIT_DIR_PATH/genesis_accounts.json" >/dev/null <<EOF
[
  {"amount":
      {"amount":"50000000000000000000000","denom":"${BASE_DENOM}"},
      "address":"$(dymd keys show -a genesis-wallet --keyring-backend test --keyring-dir "${ROLLAPP_HOME_DIR}")"
    }
]
EOF
}

generate_denom_metadata() {
  tee "$ROLLAPP_SETTLEMENT_INIT_DIR_PATH/denommetadata.json" >/dev/null <<EOF
[
  {
    "description": "The native staking and governance token of the ${ROLLAPP_CHAIN_ID}",
    "denom_units": [
      {
        "denom": "${BASE_DENOM}",
        "exponent": 0
      },
      {
        "denom": "${DENOM}",
        "exponent": 18
      }
    ],
    "base": "${BASE_DENOM}",
    "display": "${DENOM}",
    "name": "${DENOM}",
    "symbol": "${DENOM}"
  }
]
EOF
}

add_denom_metadata() {
  local success=true

  denom_metadata=$(cat "$ROLLAPP_SETTLEMENT_INIT_DIR_PATH"/denommetadata.json)
  elevated_address=$("$EXECUTABLE" keys show "$KEY_NAME_ROLLAPP" --keyring-backend test -a)

  dasel put -f "$GENESIS_FILE" '.app_state.bank.denom_metadata' -v "$denom_metadata" || success=false
  dasel put -t json -f "$GENESIS_FILE" '.app_state.denommetadata.params.allowed_addresses.' -v "$elevated_address" || success=false

  if [ "$success" = false ]; then
    echo "An error occurred. Please refer to README.md"
    return 1
  fi
}

set_consensus_params() {
  local success=true

  BLOCK_SIZE="500000"

  dasel put -f "$GENESIS_FILE" '.consensus_params.block.max_gas' -v "400000000" || success=false
  dasel put -f "$GENESIS_FILE" '.consensus_params.block.max_bytes' -v "$BLOCK_SIZE" || success=false
  dasel put -f "$GENESIS_FILE" '.consensus_params.evidence.max_bytes' -v "$BLOCK_SIZE" || success=false

  if [ "$success" = false ]; then
    echo "An error occurred. Please refer to README.md"
    return 1
  fi
}

SKIP_BASE_FEE=${SKIP_EVM_BASE_FEE-false}

set_EVM_params() {
  local success=true

  SKIP_BASE_FEE_LOWER=$(echo "$SKIP_BASE_FEE" | tr '[:upper:]' '[:lower:]')

  dasel put -t bool -f "$GENESIS_FILE" 'app_state.feemarket.params.no_base_fee' -v "$SKIP_BASE_FEE_LOWER" || success=false
  dasel put -f "$GENESIS_FILE" '.app_state.evm.params.extra_eips.[]' -v '3855' || success=false
  dasel put -f "$GENESIS_FILE" '.app_state.feemarket.params.min_gas_price' -v "10000000.0" || success=false

  if [ "$success" = false ]; then
    echo "An error occurred. Please refer to README.md"
    return 1
  fi
}

# --------------------------------- run init --------------------------------- #

# Verify that a genesis file doesn't exists for the dymension chain
if [ -f "$GENESIS_FILE" ]; then
  printf "\n======================================================================================================\n"
  echo "A genesis file already exists at $GENESIS_FILE."
  echo "Building the chain will delete all previous chain data. Continue? (y/n)"
  printf "\n======================================================================================================\n"
  read -r answer
  if [ "$answer" != "${answer#[Yy]}" ]; then
    rm -rf "$ROLLAPP_HOME_DIR"
  else
    exit 1
  fi
fi

# ------------------------------- init rollapp ------------------------------- #
"$EXECUTABLE" init "$MONIKER" --chain-id "$ROLLAPP_CHAIN_ID"

if [ ! -d "$ROLLAPP_SETTLEMENT_INIT_DIR_PATH" ]; then
  mkdir -p "$ROLLAPP_SETTLEMENT_INIT_DIR_PATH"
  echo "creating init directory : $ROLLAPP_SETTLEMENT_INIT_DIR_PATH"
else
  echo "init directory : $ROLLAPP_SETTLEMENT_INIT_DIR_PATH already exists"
fi

# ------------------------------- client config ------------------------------ #

"$EXECUTABLE" config chain-id "$ROLLAPP_CHAIN_ID"

# -------------------------------- app config -------------------------------- #
# Modify app.toml minimum-gas-prices using sed command based on the OS
dasel put -t string -f "$APP_CONFIG_FILE" 'minimum-gas-prices' -v "0$BASE_DENOM" || success=false

set_denom "$BASE_DENOM"
set_consensus_params
set_EVM_params
add_genesis_accounts
generate_denom_metadata

# --------------------- adding keys and genesis accounts --------------------- #
# Local genesis account
"$EXECUTABLE" keys add "$KEY_NAME_ROLLAPP" --keyring-backend test
"$EXECUTABLE" add-genesis-account "$KEY_NAME_ROLLAPP" "$TOTAL_SUPPLY$BASE_DENOM" --keyring-backend test

# Set sequencer's operator address
operator_address=$("$EXECUTABLE" keys show "$KEY_NAME_ROLLAPP" -a --keyring-backend test --bech val)
dasel put -f "$GENESIS_FILE" '.app_state.sequencers.genesis_operator_address' -v "$operator_address"
"$EXECUTABLE" validate-genesis

# Ask if to include a governor on genesis
echo "Do you want to include a governor on genesis? (Y/n) "
read -r answer
if [ ! "$answer" != "${answer#[Nn]}" ] ;then
  "$EXECUTABLE" gentx "$KEY_NAME_ROLLAPP" "$STAKING_AMOUNT" --chain-id "$ROLLAPP_CHAIN_ID" --keyring-backend test --home "$ROLLAPP_HOME_DIR" --fees 4000000000000"$BASE_DENOM"
  "$EXECUTABLE" collect-gentxs --home "$ROLLAPP_HOME_DIR"
fi

update_genesis_params
"$EXECUTABLE" validate-genesis
