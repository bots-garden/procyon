#!/bin/bash

# publish the wasm function to the registry

wasm_file="compute_0.0.1.wasm"
wasm_registry="https://localhost:9999"

curl -v \
  -F "file=@${wasm_file}" \
	-H "Content-Type: multipart/form-data" \
	-X POST ${wasm_registry}/upload
