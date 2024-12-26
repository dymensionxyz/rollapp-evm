#!/bin/bash
set -x

# Get environment parameter
if [ $# -lt 2 ]; then
    echo "Error: Both Environment (mainnet/testnet) and DRS parameters are required"
    exit 1
fi

ENVIRONMENT=$1
DRS=$2

if [ "$ENVIRONMENT" != "mainnet" ] && [ "$ENVIRONMENT" != "testnet" ]; then
    echo "Error: Environment must be either 'mainnet' or 'testnet'"
    exit 1
fi

if [ "$GENESIS_FILE" = "" ]; then
  DEFAULT_GENESIS_FILE_PATH="${HOME}/.rollapp_evm/config/genesis.json"
  echo "GENESIS_FILE is not set, using default: ${DEFAULT_GENESIS_FILE_PATH}"
  export GENESIS_FILE=$DEFAULT_GENESIS_FILE_PATH
fi

update_params() {
  local success=true
  
  # Create temp directory and copy genesis file
  TEMP_DIR=$(mktemp -d)
  TEMP_GENESIS="${TEMP_DIR}/genesis_temp.json"
  cp "$GENESIS_FILE" "$TEMP_GENESIS"

  BLOCK_SIZE="500000"

  # Update all dasel commands to use TEMP_GENESIS instead of GENESIS_FILE
  dasel put -f "$TEMP_GENESIS" '.consensus_params.block.max_gas' -v "400000000" || success=false
  dasel put -f "$TEMP_GENESIS" '.consensus_params.block.max_bytes' -v "$BLOCK_SIZE" || success=false
  dasel put -f "$TEMP_GENESIS" '.consensus_params.evidence.max_bytes' -v "$BLOCK_SIZE" || success=false
  dasel put -f "$TEMP_GENESIS" 'app_state.distribution.params.base_proposer_reward' -v '0.8' || success=false
  dasel put -f "$TEMP_GENESIS" 'app_state.distribution.params.community_tax' -v "0.00002" || success=false
  dasel put -t bool -f "$GENESIS_FILE" 'app_state.feemarket.params.no_base_fee' -v false || success=false
  dasel put -f "$GENESIS_FILE" '.app_state.feemarket.params.min_gas_price' -v "1000000000.0" || success=false
  dasel put -f "$GENESIS_FILE" '.app_state.auth.accounts' -t json -r json -v '[]' || success=false 
  dasel put -f "$GENESIS_FILE" '.app_state.bank.balances' -t json -r json -v '[]' || success=false 
  dasel put -f "$TEMP_GENESIS" 'app_state.rollappparams.params.drs_version' -v "$DRS" -t int || success=false
  dasel put -f "$TEMP_GENESIS" 'app_state.rollappparams.params.da' -v "celestia" || success=false
  

# Update jq command to use temp file
  if ! jq '.app_state.evm.params.extra_eips = ["3855"]' "$TEMP_GENESIS" > "${TEMP_DIR}/temp.json"; then
    echo "Error updating JSON file"
    success=false
  else
    mv "${TEMP_DIR}/temp.json" "$TEMP_GENESIS"
  fi


  # Update jq command to use temp file
  if ! jq '.app_state.gov.deposit_params.min_deposit[0].amount = "1000000000000000000000"' "$TEMP_GENESIS" > "${TEMP_DIR}/temp.json"; then
    echo "Error updating JSON file"
    success=false
  else
    mv "${TEMP_DIR}/temp.json" "$TEMP_GENESIS"
  fi

  # these vary depending on environment
  if [ "$ENVIRONMENT" = "mainnet" ]; then
    UNBONDING_TIME="1814400s" # 2 weeks
    VOTING_PERIOD="432000s" # 5 days
  else
    UNBONDING_TIME="1309600s" # ~2 weeks + 1 day
    VOTING_PERIOD="300s"
  fi
  
  dasel put -f "$TEMP_GENESIS" '.app_state.sequencers.params.unbonding_time' -v "$UNBONDING_TIME" || success=false
  dasel put -f "$TEMP_GENESIS" '.app_state.staking.params.unbonding_time' -v "$UNBONDING_TIME" || success=false
  dasel put -f "$TEMP_GENESIS" 'app_state.gov.voting_params.voting_period' -v "$VOTING_PERIOD" || success=false

  if [ "$success" = false ]; then
    echo "An error occurred."
    rm -rf "$TEMP_DIR"
    return 1
  fi
  
  # Create templates directory with DRS subdirectory if it doesn't exist
  mkdir -p "./genesis-templates/DRS/${DRS}"
  
  # Copy the modified genesis file to the DRS-specific template location
  cp "$TEMP_GENESIS" "./genesis-templates/DRS/${DRS}/genesis-${ENVIRONMENT}.json"
  
  # Cleanup temp directory
  rm -rf "$TEMP_DIR"
  
  set +x
}

update_params












  

  