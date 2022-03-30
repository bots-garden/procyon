#!/bin/bash


curl -v --request POST \
  --header 'Content-Type: application/json' \
  --data '{
      "executor": 2,
      "wasmFileName": "",
      "wasmFunctionHttpPort": 8085,
      "wasmRegistryUrl": "https://localhost:9999/wasm/download/hello.wasm"
    }
  ' http://localhost:9090/tasks

