#!/bin/bash
  set -x
if [ "$GENESIS_FILE" = "" ]; then
  DEFAULT_GENESIS_FILE_PATH="${HOME}/.rollapp_evm/config/genesis.json"
  echo "GENESIS_FILE is not set, using default: ${DEFAULT_GENESIS_FILE_PATH}"
  export GENESIS_FILE=$DEFAULT_GENESIS_FILE_PATH
fi

update_params() {
  local success=true

  BLOCK_SIZE="500000"

  dasel put -f "$GENESIS_FILE" '.consensus_params.block.max_gas' -v "400000000" || success=false
  dasel put -f "$GENESIS_FILE" '.consensus_params.block.max_bytes' -v "$BLOCK_SIZE" || success=false
  dasel put -f "$GENESIS_FILE" '.consensus_params.evidence.max_bytes' -v "$BLOCK_SIZE" || success=false
  dasel put -t bool -f "$GENESIS_FILE" 'app_state.feemarket.params.no_base_fee' -v false || success=false
  dasel put -f "$GENESIS_FILE" '.app_state.evm.params.extra_eips.[]' -v '3855' || success=false
  dasel put -f "$GENESIS_FILE" '.app_state.feemarket.params.min_gas_price' -v "10000000.0" || success=false
  dasel put -f "$GENESIS_FILE" 'app_state.distribution.params.base_proposer_reward' -v '0.8' || success=false
  dasel put -f "$GENESIS_FILE" 'app_state.distribution.params.community_tax' -v "0.00002" || success=false
  # these vary depending on environment
  dasel put -f "$GENESIS_FILE" 'app_state.gov.voting_params.voting_period' -v "300s" || success=false
  dasel put -f "$GENESIS_FILE" '.app_state.sequencers.params.unbonding_time' -v "1209600s" || success=false # 2 weeks
  dasel put -f "$GENESIS_FILE" '.app_state.staking.params.unbonding_time' -v "1209600s" || success=false # 2 weeks

  if [ "$success" = false ]; then
    echo "An error occurred."
    return 1
  fi
  set +x
}

update_params
