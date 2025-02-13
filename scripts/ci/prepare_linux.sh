#!/usr/bin/env bash

#
# IMPORTANT: This file should only include production
# dependencies. Dependencies for integration tests should be in the
# ./prepare_linux_integration_test_setup_only.sh file in this
# directory.
#

set -ex

# Set up Node.js, jq
curl -fsSL https://deb.nodesource.com/setup_16.x | sudo -E bash -
sudo apt-get update -y
sudo apt-get install -y nodejs cmake

# Set up Go
./scripts/ci/prepare_go.sh
#
# IMPORTANT: Only run in Github CI.
#
if [[ "$1" == "--integration-tests" ]]; then
    ./scripts/ci/prepare_linux_integration_test_setup_only.sh
fi

# Set up project
sudo npm install --global yarn
yarn

# Download  unixODBC source code and compile
curl -LO http://www.unixodbc.org/unixODBC-2.3.11.tar.gz
tar -xvf unixODBC-2.3.11.tar.gz
cd unixODBC-2.3.11
./configure
make
sudo make install
cd .. && sudo rm -rf unixODBC-2.3.11

# Compile Go plugin
cd runner/plugins/odbc && go build -trimpath -buildmode=plugin -mod=readonly -modcacherw -ldflags="-s -w" -o ./odbc.build ./odbc.go && cd ../../..
