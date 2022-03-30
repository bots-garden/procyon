#!/bin/bash

curl -v --request POST \
  --header 'Content-Type: application/json' \
  --data '{
      "executor": 1,
      "wasmFileName": "hello.wasm",
      "wasmFunctionHttpPort": 8082,
      "wasmRegistryUrl": "https://localhost:9999/wasm/download/hello.wasm"
    }
  ' http://localhost:9090/tasks


curl -v --request POST \
  --header 'Content-Type: application/json' \
  --data '{
      "executor": 1,
      "wasmFileName": "hey.wasm",
      "wasmFunctionHttpPort": 8083,
      "wasmRegistryUrl": "https://localhost:9999/wasm/download/hey.wasm"
    }
  ' http://localhost:9090/tasks



curl -v --request POST \
  --header 'Content-Type: application/json' \
  --data '{
      "executor": 2,
      "wasmFileName": "",
      "wasmFunctionHttpPort": 8084,
      "wasmRegistryUrl": "https://localhost:9999/wasm/download/hi.wasm"
    }
  ' http://localhost:9090/tasks


curl -v --request POST \
  --header 'Content-Type: application/json' \
  --data '{
      "executor": 1,
      "wasmFileName": "oh.wasm",
      "wasmFunctionHttpPort": 8085,
      "wasmRegistryUrl": "https://localhost:9999/wasm/download/oh.wasm"
    }
  ' http://localhost:9090/tasks
  