#!/bin/bash -e

# Install Solidity compiler
brew list solidity || install solidity

# Install dependencies for tdlib
brew list gperf || brew install gperf
brew list cmake || brew install cmake
brew list openssl || brew install openssl

# Build tdlib

cd td
mkdir -p build
cd build

if [ -f .success ]; then
  echo ".success flag detected, skipping build tdlib"
  exit 0
fi

cmake -DCMAKE_BUILD_TYPE=Release -DOPENSSL_ROOT_DIR=/usr/local/opt/openssl/ ..
cmake --build . -- -j5

make install

touch .success
