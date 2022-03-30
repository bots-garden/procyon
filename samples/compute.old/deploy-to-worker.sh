#!/bin/bash

wasm_file="compute_0.0.1.wasm"
wasm_function_port="8087"
wasm_registry="https://localhost:9999"

worker_url="http://localhost:9090"

curl -v --request POST \
  --header 'Content-Type: application/json' \
  --data '{
      "executor": 1,
      "wasmFileName": "'"${wasm_file}"'",
      "wasmFunctionHttpPort": '"${wasm_function_port}"',
      "wasmRegistryUrl": "'"${wasm_registry}"'/functions/'"${wasm_file}"'"
    }
  ' ${worker_url}/tasks

