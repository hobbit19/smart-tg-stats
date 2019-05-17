#!/bin/bash -e

# Install Solidity compiler
apt-get update
apt-get install -y software-properties-common
add-apt-repository -y ppa:ethereum/ethereum
apt-get update
apt-get install -y solc

# Install dependencies for tdlib
apt install -y libssl-dev g++ build-essential gperf zlib1g-dev cmake

# Build tdlib

cd td
mkdir -p build
cd build

if [ -f .success ]; then
  echo ".success flag detected, skipping build tdlib"
  exit 0
fi

cmake -DCMAKE_BUILD_TYPE=Release ..
cmake --build . -- -j5

make install

touch .success
