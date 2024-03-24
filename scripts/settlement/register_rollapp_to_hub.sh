#!/bin/bash

MAX_SEQUENCERS=5
DEPLOYER="local-user"

#Register rollapp 
dymd tx rollapp create-rollapp "$ROLLAPP_CHAIN_ID" "$MAX_SEQUENCERS" '{"Addresses":[]}' \
  --from "$DEPLOYER" \
  --keyring-backend test \
  --broadcast-mode block \
  --fees 1dym

