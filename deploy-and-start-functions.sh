#!/bin/bash

curl -v --request POST \
  --header 'Content-Type: application/json' \
  --data '{
      "executor": 2,
      "wasmFileName": "hello.wasm",
      "wasmFunctionHttpPort": 8081,
      "wasmRegistryUrl": "https://localhost:9999/wasm/download/hello.wasm"
    }
  ' http://localhost:9090/tasks


curl -v --request POST \
  --header 'Content-Type: application/json' \
  --data '{
      "executor": 2,
      "wasmFileName": "hey.wasm",
      "wasmFunctionHttpPort": 8082,
      "wasmRegistryUrl": "https://localhost:9999/wasm/download/hey.wasm"
    }
  ' http://localhost:9090/tasks



curl -v --request POST \
  --header 'Content-Type: application/json' \
  --data '{
      "executor": 2,
      "wasmFileName": "hi.wasm",
      "wasmFunctionHttpPort": 8083,
      "wasmRegistryUrl": "https://localhost:9999/wasm/download/hi.wasm"
    }
  ' http://localhost:9090/tasks


curl -v --request POST \
  --header 'Content-Type: application/json' \
  --data '{
      "executor": 2,
      "wasmFileName": "yo.wasm",
      "wasmFunctionHttpPort": 8084,
      "wasmRegistryUrl": "https://localhost:9999/wasm/download/yo.wasm"
    }
  ' http://localhost:9090/tasks
