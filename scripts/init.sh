#!/bin/bash
tmp=$(mktemp)

EXECUTABLE="rollapp-evm"
ROLLAPP_CHAIN_DIR="$HOME/.rollapp_evm"

set_denom() {
  denom=$1
  jq --arg denom $denom '.app_state.mint.params.mint_denom = $denom' "$GENESIS_FILE" > "$tmp" && mv "$tmp" "$GENESIS_FILE"
  jq --arg denom $denom '.app_state.staking.params.bond_denom = $denom' "$GENESIS_FILE" > "$tmp" && mv "$tmp" "$GENESIS_FILE"
  jq --arg denom $denom '.app_state.gov.deposit_params.min_deposit[0].denom = $denom' "$GENESIS_FILE" > "$tmp" && mv "$tmp" "$GENESIS_FILE"
  
  jq --arg denom $denom '.app_state.evm.params.evm_denom = $denom' "$GENESIS_FILE" > "$tmp" && mv "$tmp" "$GENESIS_FILE"
  jq --arg denom $denom '.app_state.claims.params.claims_denom = $denom' "$GENESIS_FILE" > "$tmp" && mv "$tmp" "$GENESIS_FILE"
}

set_EVM_params() {
  jq '.consensus_params["block"]["max_gas"] = "40000000"' "$GENESIS_FILE" > "$tmp" && mv "$tmp" "$GENESIS_FILE"
  jq '.app_state["feemarket"]["params"]["no_base_fee"] = true' "$GENESIS_FILE" > "$tmp" && mv "$tmp" "$GENESIS_FILE"
  jq '.app_state["feemarket"]["params"]["min_gas_price"] = "0.0"' "$GENESIS_FILE" > "$tmp" && mv "$tmp" "$GENESIS_FILE"
}

# ---------------------------- initial parameters ---------------------------- #
# Assuming 1,000,000 tokens
#half is staked
MODULE_AMOUNT=1000000000000000000000000
TOTAL_SUPPLY=2000000000000000000000000
TOKEN_AMOUNT="1000000000000000000000000$DENOM"
STAKING_AMOUNT="500000000000000000000000$DENOM"


CONFIG_DIRECTORY="$ROLLAPP_CHAIN_DIR/config"
GENESIS_FILE="$CONFIG_DIRECTORY/genesis.json"
DYMINT_CONFIG_FILE="$CONFIG_DIRECTORY/dymint.toml"
APP_CONFIG_FILE="$CONFIG_DIRECTORY/app.toml"

# --------------------------------- run init --------------------------------- #
if ! command -v $EXECUTABLE >/dev/null; then
  echo "$EXECUTABLE does not exist"
  echo "please run make install"
  exit 1
fi

if [ -z "$ROLLAPP_CHAIN_ID" ]; then
  echo "ROLLAPP_CHAIN_ID is not set"
  exit 1
fi

# Verify that a genesis file doesn't exists for the dymension chain
if [ -f "$GENESIS_FILE" ]; then
  printf "\n======================================================================================================\n"
  echo "A genesis file already exists [$GENESIS_FILE]. building the chain will delete all previous chain data. continue? (y/n)"
  printf "\n======================================================================================================\n"
  read -r answer
  if [ "$answer" != "${answer#[Yy]}" ]; then
    rm -rf "$ROLLAPP_CHAIN_DIR"
  else
    exit 1
  fi
fi

# ------------------------------- init rollapp ------------------------------- #
$EXECUTABLE init "$MONIKER" --chain-id "$ROLLAPP_CHAIN_ID"

# ------------------------------- client config ------------------------------ #
$EXECUTABLE config chain-id "$ROLLAPP_CHAIN_ID"

# -------------------------------- app config -------------------------------- #
sed -i'' -e "s/^minimum-gas-prices *= .*/minimum-gas-prices = \"0$DENOM\"/" "$APP_CONFIG_FILE"
set_denom "$DENOM"
set_EVM_params

# --------------------- adding keys and genesis accounts --------------------- #
#local genesis account
$EXECUTABLE keys add "$KEY_NAME_ROLLAPP" --keyring-backend test
$EXECUTABLE add-genesis-account "$KEY_NAME_ROLLAPP" "$TOKEN_AMOUNT" --keyring-backend test


# set sequencer's operator address
operator_address=$($EXECUTABLE keys show "$KEY_NAME_ROLLAPP" -a --keyring-backend test --bech val)
jq --arg addr $operator_address '.app_state["sequencers"]["genesis_operator_address"] = $addr' "$GENESIS_FILE" > "$tmp" && mv "$tmp" "$GENESIS_FILE"

whitelisted_address=$($EXECUTABLE keys show "$KEY_NAME_ROLLAPP" -a --keyring-backend test)

jq --arg address $whitelisted_address '.app_state.hubgenesis.params.genesis_triggerer_whitelist += [{ "address": $address }]' "$GENESIS_FILE" > "$tmp" && mv "$tmp" "$GENESIS_FILE"

module_name="hubgenesis"
module_address=$($EXECUTABLE q make-address $module_name)

jq --arg module_address $module_address --arg module_name $module_name '.app_state.auth.accounts += [
    {
     "@type": "/cosmos.auth.v1beta1.ModuleAccount",
     "base_account": {
         "account_number": "0",
         "address": $module_address,
         "pub_key": null,
         "sequence": "0"
     },
     "name": $module_name,
     "permissions": []
    }
]' "$GENESIS_FILE" > "$tmp" && mv "$tmp" "$GENESIS_FILE"

jq --arg module_address $module_address --arg amount $MODULE_AMOUNT '.app_state.bank.balances += [
    {
      "address": $module_address,
      "coins": [
        {
          "denom": "urax",
          "amount": $amount
        }
      ]
    }
]' "$GENESIS_FILE" > "$tmp" && mv "$tmp" "$GENESIS_FILE"

jq --arg total_supply $TOTAL_SUPPLY '.app_state.bank.supply = [
    {
      "denom": "urax",
      "amount": $total_supply
    }
]' "$GENESIS_FILE" > "$tmp" && mv "$tmp" "$GENESIS_FILE"

echo "Do you want to include staker on genesis? (Y/n) "
read -r answer
if [ ! "$answer" != "${answer#[Nn]}" ] ;then
  $EXECUTABLE gentx "$KEY_NAME_ROLLAPP" "$STAKING_AMOUNT" --chain-id "$ROLLAPP_CHAIN_ID" --keyring-backend test --home "$ROLLAPP_CHAIN_DIR"
  $EXECUTABLE collect-gentxs --home "$ROLLAPP_CHAIN_DIR"
fi


$EXECUTABLE validate-genesis
