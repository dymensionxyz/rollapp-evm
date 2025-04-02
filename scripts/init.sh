#!/bin/bash
EXECUTABLE=$(which rollapp-evm)

if ! command -v "$EXECUTABLE" >/dev/null; then
  echo "$EXECUTABLE does not exist"
  echo "please run make install"
  exit 1
fi

if [ "$ROLLAPP_CHAIN_ID" = "" ]; then
  echo "ROLLAPP_CHAIN_ID is not set"
  exit 1
fi

if [ "$BASE_DENOM" = "" ]; then
  echo "BASE_DENOM is not set"
  exit 1
fi

# ---------------------------- initial parameters ---------------------------- #
CONFIG_DIRECTORY="$ROLLAPP_HOME_DIR/config"
APP_CONFIG_FILE="$CONFIG_DIRECTORY/app.toml"
GENESIS_FILE="$CONFIG_DIRECTORY/genesis.json"

# ---------------------------- check variables ---------------------------- #
if [ "$MONIKER" = "" ]; then
  MONIKER="${ROLLAPP_CHAIN_ID}-sequencer" # Default moniker value
fi

if [ "$KEY_NAME_ROLLAPP" = "" ]; then
  KEY_NAME_ROLLAPP="rol-user" # Default key name value
fi

# Default to 1,000,000,000 tokens
if [ "$TOTAL_SUPPLY" = "" ]; then
  TOTAL_SUPPLY="1000000000000000000000000000"
fi

if [ "$ROLLAPP_SETTLEMENT_INIT_DIR_PATH" = "" ]; then
  # ROLLAPP_SETTLEMENT_INIT_DIR_PATH is used as a target for generating the necessary
  # configuration files for RollApp initialization, such as denom metadata and genesis account
  # json files
  ROLLAPP_SETTLEMENT_INIT_DIR_PATH="${ROLLAPP_HOME_DIR}/init"
fi

if [ "$DA_CLIENT" = "" ]; then
  echo "DA_CLIENT type is not set"
  exit 1
fi

# FIXME: rename to DA_NETWORK
if [ "$CELESTIA_NETWORK" = "" ]; then
  echo "CELESTIA_NETWORK is not set"
  exit 1
fi

if [ "$CELESTIA_HOME_DIR" = "" ]; then
  echo "CELESTIA_HOME_DIR is not set"
  exit 1
fi

if [[ $CELESTIA_NETWORK == "mock" || $CELESTIA_NETWORK == "grpc" ]]; then
  mkdir -p "$CELESTIA_HOME_DIR"
fi

set_denom() {
  local denom=$1
  local success=true

  dasel put -f "$GENESIS_FILE" '.app_state.mint.params.mint_denom' -v "$denom" || success=false
  dasel put -f "$GENESIS_FILE" '.app_state.staking.params.bond_denom' -v "$denom" || success=false
  dasel put -t string -f "$GENESIS_FILE" '.app_state.gov.deposit_params.min_deposit.[0].denom' -v "$denom" || success=false
  dasel put -f "$GENESIS_FILE" '.app_state.evm.params.evm_denom' -v "$denom" || success=false
  dasel put -f "$GENESIS_FILE" '.app_state.evm.params.gas_denom' -v "$denom" || success=false
  dasel put -f "$GENESIS_FILE" '.app_state.claims.params.claims_denom' -v "$denom" || success=false

  if [ "$success" = false ]; then
    echo "An error occurred. Please refer to README.md"
    exit 1
  fi
}

update_genesis_params() {
  local success=true

  dasel put -f "$GENESIS_FILE" '.app_state.gov.voting_params.voting_period' -v "300s" || success=false
  dasel put -f "$GENESIS_FILE" '.app_state.gov.tally_params.threshold' -v "0.490000000000000000" || success=false
  dasel put -f "$GENESIS_FILE" '.app_state.sequencers.params.unbonding_time' -v "1814400s" || success=false # 2 weeks
  dasel put -f "$GENESIS_FILE" '.app_state.staking.params.unbonding_time' -v "1814400s" || success=false    # 2 weeks

  if [ "$success" = false ]; then
    echo "An error occurred. Please refer to README.md"
    return 1
  fi
  echo "Successfully updated the genesis file"
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

add_denom_metadata_to_genesis() {
  local success=true

  #denom_metadata=$(cat "$ROLLAPP_SETTLEMENT_INIT_DIR_PATH"/denommetadata.json)
  #dasel put -f "$GENESIS_FILE" '.app_state.bank.denom_metadata' -v "$denom_metadata" || success=false
  jq --argjson metadata "$(cat "$ROLLAPP_SETTLEMENT_INIT_DIR_PATH"/denommetadata.json)" '.app_state.bank.denom_metadata = $metadata' "$GENESIS_FILE" >temp.json && mv temp.json "$GENESIS_FILE"

  if [ "$success" = false ]; then
    echo "An error occurred. Please refer to README.md"
    return 1
  fi
}

set_consensus_params() {
  local success=true

  BLOCK_SIZE="500000"
  COMMIT=$(git log -1 --format='%H')

  case $DA_CLIENT in
  "celestia")
    case $CELESTIA_NETWORK in
    "celestia" | "mocha")
      DA="celestia"
      ;;
    "mock" | *)
      DA="mock"
      ;;
    esac
    ;;
  "weavevm")
    DA="weavevm"
    ;;
  "grpc")
    DA="grpc"
    ;;
  "avail")
    DA="avail"
    ;;
  *)
    DA="mock"
    ;;
  esac

  VERSION=$(rollapp-evm version --long | grep DRS-)
  DRS_VERSION="${VERSION#*-}"

  dasel put -f "$GENESIS_FILE" '.consensus_params.block.max_gas' -v "400000000" || success=false
  dasel put -f "$GENESIS_FILE" '.consensus_params.block.max_bytes' -v "$BLOCK_SIZE" || success=false
  dasel put -f "$GENESIS_FILE" '.consensus_params.evidence.max_bytes' -v "$BLOCK_SIZE" || success=false
  dasel put -f "$GENESIS_FILE" '.app_state.rollappparams.params.da' -v "$DA" || success=false
  dasel put -f "$GENESIS_FILE" '.app_state.rollappparams.params.drs_version' -v $DRS_VERSION -t int || success=false

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

update_configuration_weavevm_da() {
  if [[ "$OSTYPE" == "darwin"* ]]; then
    # macOS-specific sed
    sed -i '' "s|da_layer =.*|da_layer = [\"weavevm\"]|" "${CONFIG_DIRECTORY}/dymint.toml"
    sed -i '' "s|da_config =.*|da_config = [\"{\\\\\"endpoint\\\\\":\\\\\"https:\/\/testnet-rpc.wvm.dev\\\\\",\\\\\"chain_id\\\\\":9496,\\\\\"timeout\\\\\":60000000000,\\\\\"private_key_hex\\\\\":\\\\\"${WVM_PRIV_KEY}\\\\\"}\"]|" "${CONFIG_DIRECTORY}/dymint.toml"
  else
    # Linux/Other OS-specific sed
    sed -i "s|da_layer =.*|da_layer = \"weavevm\"|" "${CONFIG_DIRECTORY}/dymint.toml"
    sed -i "s|da_config =.*|da_config = [\"{\\\\\"endpoint\\\\\":\\\\\"https:\/\/testnet-rpc.wvm.dev\\\\\",\\\\\"chain_id\\\\\":9496,\\\\\"timeout\\\\\":60000000000,\\\\\"private_key_hex\\\\\":\\\\\"${WVM_PRIV_KEY}\\\\\"}\"]|" "${CONFIG_DIRECTORY}/dymint.toml"
  fi
}

update_configuration_avail_da() {
  if [[ "$OSTYPE" == "darwin"* ]]; then
    # macOS-specific sed
    sed -i '' "s|da_layer =.*|da_layer = [\"avail\"]|" "${CONFIG_DIRECTORY}/dymint.toml"
    sed -i '' "s|da_config =.*|da_config = [\"{\\\\\"endpoint\\\\\":\\\\\"https:\/\/turing-rpc.avail.so\/rpc\\\\\",\\\\\"app_id\\\\\":1,\\\\\"seed\\\\\":\\\\\"${MNEMONIC}\\\\\"}\"]|" "${CONFIG_DIRECTORY}/dymint.toml"
  else
    # Linux/Other OS-specific sed
    sed -i "s|da_layer =.*|da_layer = [\"avail\"]|" "${CONFIG_DIRECTORY}/dymint.toml"
    sed -i "s|da_config =.*|da_config = [\"{\\\\\"endpoint\\\\\":\\\\\"https:\/\/turing-rpc.avail.so\/rpc\\\\\",\\\\\"app_id\\\\\":1,\\\\\"seed\\\\\":\\\\\"${MNEMONIC}\\\\\"}\"]|" "${CONFIG_DIRECTORY}/dymint.toml"
  fi
}

update_configuration_celestia_da() {
  if [[ $CELESTIA_NETWORK != "mock" && $CELESTIA_NETWORK != "grpc" ]]; then
    celestia_namespace_id=$(openssl rand -hex 10)
    if [ ! -d "$CELESTIA_HOME_DIR" ]; then
      echo "Celestia light client is expected to be initialized in this directory: $CELESTIA_HOME_DIR"
      echo "but it does not exist, please initialize the light client or update the 'CELESTIA_HOME_DIR'"
      echo "environment variable"
      exit 1
    fi

    celestia_token=$(celestia light auth admin --p2p.network "$CELESTIA_NETWORK" --node.store "$CELESTIA_HOME_DIR")

    if [[ "$OSTYPE" == "darwin"* ]]; then
      sed -i '' "s/da_layer =.*/da_layer = [\"celestia\"]/" "${CONFIG_DIRECTORY}/dymint.toml"
      sed -i '' "s/da_config .*/da_config =[\"{\\\\\"base_url\\\\\": \\\\\"http:\/\/localhost:26658\\\\\", \\\\\"timeout\\\\\": 60000000000, \\\\\"gas_prices\\\\\":1.0, \\\\\"gas_adjustment\\\\\": 1.3, \\\\\"namespace_id\\\\\": \\\\\"${celestia_namespace_id}\\\\\", \\\\\"auth_token\\\\\":\\\\\"${celestia_token}\\\\\"}\"]/" "${CONFIG_DIRECTORY}/dymint.toml"
    else
      sed -i "s/da_layer =.*/da_layer = [\"celestia\"]/" "${CONFIG_DIRECTORY}/dymint.toml"
      sed -i "s/da_config .*/da_config = [\"{\\\\\"base_url\\\\\": \\\\\"http:\/\/localhost:26658\\\\\", \\\\\"timeout\\\\\": 60000000000, \\\\\"gas_prices\\\\\":1.0, \\\\\"gas_adjustment\\\\\": 1.3, \\\\\"namespace_id\\\\\": \\\\\"${celestia_namespace_id}\\\\\", \\\\\"auth_token\\\\\":\\\\\"${celestia_token}\\\\\"}\"]/" "${CONFIG_DIRECTORY}/dymint.toml"
    fi
  fi
}

update_configuration_sui_da() {
  if [[ "$OSTYPE" == "darwin"* ]]; then
    # macOS-specific sed
    sed -i '' "s|da_layer =.*|da_layer = [\"sui\"]|" "${CONFIG_DIRECTORY}/dymint.toml"
    sed -i '' "s|da_config =.*|da_config = [\"{\\\\\"chain_id\\\\\":9496,\\\\\"rpc_url\\\\\":\\\\\"https:\/\/fullnode.testnet.sui.io:443\\\\\",\\\\\"noop_contract_address\\\\\":\\\\\"0xcf119583badb169bfc9a031ec16fb6a79a5151ff7aa0d229f2a35b798ddcd9d6\\\\\",\\\\\"gas_budget\\\\\":\\\\\"10000000\\\\\",\\\\\"timeout\\\\\":5000000000,\\\\\"mnemonic_env\\\\\":\\\\\"SUI_MNEMONIC\\\\\"}\"]|" "${CONFIG_DIRECTORY}/dymint.toml"
  else
    # Linux/Other OS-specific sed
    sed -i "s|da_layer =.*|da_layer = [\"sui\"]|" "${CONFIG_DIRECTORY}/dymint.toml"
    sed -i "s|da_config =.*|da_config = [\"{\\\\\"chain_id\\\\\":9496,\\\\\"rpc_url\\\\\":\\\\\"https:\/\/fullnode.testnet.sui.io:443\\\\\",\\\\\"noop_contract_address\\\\\":\\\\\"0xcf119583badb169bfc9a031ec16fb6a79a5151ff7aa0d229f2a35b798ddcd9d6\\\\\",\\\\\"gas_budget\\\\\":\\\\\"10000000\\\\\",\\\\\"timeout\\\\\":5000000000,\\\\\"mnemonic_env\\\\\":\\\\\"SUI_MNEMONIC\\\\\"}\"]|" "${CONFIG_DIRECTORY}/dymint.toml"
  fi
}

update_configuration_aptos_da() {
  if [[ "$OSTYPE" == "darwin"* ]]; then
    # macOS-specific sed
    sed -i '' "s|da_layer =.*|da_layer = [\"aptos\"]|" "${CONFIG_DIRECTORY}/dymint.toml"
    sed -i '' "s|da_config =.*|da_config = [\"{\\\\\"network\\\\\":\\\\\"testnet\\\\\",\\\\\"pri_key_env\\\\\":\\\\\"APT_PRIVATE_KEY\\\\\"}\"]|" "${CONFIG_DIRECTORY}/dymint.toml"
  else
    # Linux/Other OS-specific sed
    sed -i "s|da_layer =.*|da_layer = [\"aptos\"]|" "${CONFIG_DIRECTORY}/dymint.toml"
    sed -i "s|da_config =.*|da_config = [\"{\\\\\"network\\\\\":\\\\\"testnet\\\\\",\\\\\"pri_key_env\\\\\":\\\\\"APT_PRIVATE_KEY\\\\\"}\"]|" "${CONFIG_DIRECTORY}/dymint.toml"
  fi
}

update_configuration_mock_da() {
  if [[ "$OSTYPE" == "darwin"* ]]; then
    sed -i '' "s/da_layer =.*/da_layer = [\"mock\"]/" "${CONFIG_DIRECTORY}/dymint.toml"
    sed -i '' "s/da_config .*/da_config =[\"\"]/" "${CONFIG_DIRECTORY}/dymint.toml"
  else
    sed -i "s/da_layer =.*/da_layer = [\"mock\"]/" "${CONFIG_DIRECTORY}/dymint.toml"
    sed -i "s/da_config .*/da_config =[\"\"]/" "${CONFIG_DIRECTORY}/dymint.toml"
  fi
}

update_configuration() {
  case $DA_CLIENT in
  "weavevm")
    update_configuration_weavevm_da
    ;;
  "celestia")
    update_configuration_celestia_da
    ;;
  "avail")
    update_configuration_avail_da
    ;;
  "sui")
    update_configuration_sui_da
    ;;
  "aptos")
    update_configuration_aptos_da
    ;;
  "mock")
    update_configuration_mock_da
    ;;
  esac

  if [[ ! $SETTLEMENT_LAYER == "mock" ]]; then
    if [[ "$OSTYPE" == "darwin"* ]]; then
      sed -i '' 's/settlement_layer.*/settlement_layer = "dymension"/' "${CONFIG_DIRECTORY}/dymint.toml"
      sed -i '' -e "/settlement_node_address =/s/.*/settlement_node_address = \"${HUB_RPC_URL//\//\\/}\"/" "${CONFIG_DIRECTORY}/dymint.toml"
      sed -i '' -e "/rollapp_id =/s/.*/rollapp_id = \"$ROLLAPP_CHAIN_ID\"/" "${CONFIG_DIRECTORY}/dymint.toml"
      sed -i '' -e "/minimum-gas-prices =/s/.*/minimum-gas-prices = \"1000000000${BASE_DENOM}\"/" "${CONFIG_DIRECTORY}/app.toml"
    else
      sed -i 's/settlement_layer.*/settlement_layer = "dymension"/' "${CONFIG_DIRECTORY}/dymint.toml"
      sed -i '/settlement_node_address =/c\settlement_node_address = '\""$HUB_RPC_URL"\" "${CONFIG_DIRECTORY}/dymint.toml"
    fi
  fi

  if [[ "$OSTYPE" == "darwin"* ]]; then
    sed -i '' -e "/rollapp_id =/s/.*/rollapp_id = \"$ROLLAPP_CHAIN_ID\"/" "${CONFIG_DIRECTORY}/dymint.toml"
    sed -i '' -e "/minimum-gas-prices =/s/.*/minimum-gas-prices = \"1000000000${BASE_DENOM}\"/" "${CONFIG_DIRECTORY}/app.toml"
  else
    sed -i '/rollapp_id =/c\rollapp_id = '\""$ROLLAPP_CHAIN_ID"\" "${CONFIG_DIRECTORY}/dymint.toml"
    sed -i '/minimum-gas-prices =/c\minimum-gas-prices = '\"1000000000"$BASE_DENOM"\" "${CONFIG_DIRECTORY}/app.toml"
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
generate_denom_metadata
add_denom_metadata_to_genesis
update_configuration

# --------------------- adding keys and genesis accounts --------------------- #
# Local genesis account
"$EXECUTABLE" keys add "$KEY_NAME_ROLLAPP" --keyring-backend test
"$EXECUTABLE" add-genesis-account "$KEY_NAME_ROLLAPP" "$TOTAL_SUPPLY$BASE_DENOM" --keyring-backend test

# Ask if to include a governor on genesis
echo "Do you want to include a governor on genesis? (Y/n) "
read -r answer
if [ ! "$answer" != "${answer#[Nn]}" ]; then
  STAKING_AMOUNT="500000000000000000000000$BASE_DENOM"
  "$EXECUTABLE" gentx "$KEY_NAME_ROLLAPP" "$STAKING_AMOUNT" --chain-id "$ROLLAPP_CHAIN_ID" --keyring-backend test --home "$ROLLAPP_HOME_DIR" --fees 4000000000000"$BASE_DENOM"
  "$EXECUTABLE" collect-gentxs --home "$ROLLAPP_HOME_DIR"
fi

update_genesis_params
"$EXECUTABLE" validate-genesis
