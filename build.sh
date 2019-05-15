#!/bin/bash -e

if [ ! -d td ]; then
  git clone git@github.com:tdlib/td.git
fi

cd td
mkdir -p build
cd build

if [ -f .success ]; then
  exit 0
fi

cmake -DCMAKE_BUILD_TYPE=Release -DOPENSSL_ROOT_DIR=/usr/local/opt/openssl/lib ..
cmake --build . -- -j5

make install

touch .success
