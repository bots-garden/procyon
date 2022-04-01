#!/bin/bash
wget -qO- https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/utils/install.sh | bash -s -- -v 0.9.0
source /home/gitpod/.wasmedge/env 
go get github.com/second-state/WasmEdge-go/wasmedge@v0.9.0
