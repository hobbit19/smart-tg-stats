#!/bin/bash -e

# Where to get/keep telegeram secrets?
export AGENT_SESSION_DIR="./session"

# Where is the contract and CEO private file?
export ETH_NODE_ADDRESS="https://rinkeby.infura.io/v3/55c9cb49e22641989af372edac7876fc"
export ETH_CONTRACT_ADDRESS="0x9d776b251eAB7C6e3AEF82B87ccb44A49df7197a"
export ETH_CONTRACT_CEO_KEYSTORE="./ceo.json"
export ETH_CONTRACT_CEO_PASSWORD="./ceo_password.txt"
