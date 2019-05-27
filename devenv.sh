#!/bin/bash -e

# NOTE: This file only demonstrates variables required to run an instance
# Use docker stack with proper environment insted

# Where to get/keep telegeram secrets for this instance?
export AGENT_SESSION_DIR="./session"

# Where is contract and CEO's (owner) private key?
export ETH_NODE_ADDRESS="https://rinkeby.infura.io/v3/55c9cb49e22641989af372edac7876fc"
export ETH_CONTRACT_ADDRESS="0x9d776b251eAB7C6e3AEF82B87ccb44A49df7197a"
export ETH_CONTRACT_CEO_KEYSTORE="./ceo.json"
export ETH_CONTRACT_CEO_PASSWORD="./ceo_password.txt"

# Where is tempoary Redis instance? (empty host to disable)
# NOTE: Redis is required for cost-effective production mode with multi-instance
export REDIS_HOST=""
export REDIS_PORT=""
export REDIS_DB=""
