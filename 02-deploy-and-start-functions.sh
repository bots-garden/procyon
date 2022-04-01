#!/bin/bash

curl -v --request POST \
  --header 'Content-Type: application/json' \
  --data '{
      "executor": 2,
      "wasmFileName": "hello.wasm",
      "wasmRegistryUrl": "https://localhost:9999/wasm/download/hello.wasm",
      "functionName": "hello",
      "functionRevision": "",
      "defaultRevision": true
    }
  ' http://localhost:9090/tasks


curl -v --request POST \
  --header 'Content-Type: application/json' \
  --data '{
      "executor": 2,
      "wasmFileName": "hey.wasm",
      "wasmRegistryUrl": "https://localhost:9999/wasm/download/hey.wasm",
      "functionName": "hey",
      "functionRevision": "",
      "defaultRevision": true
    }
  ' http://localhost:9090/tasks



curl -v --request POST \
  --header 'Content-Type: application/json' \
  --data '{
      "executor": 2,
      "wasmFileName": "hi.wasm",
      "wasmRegistryUrl": "https://localhost:9999/wasm/download/hi.wasm",
      "functionName": "hi",
      "functionRevision": "",
      "defaultRevision": true
    }
  ' http://localhost:9090/tasks


curl -v --request POST \
  --header 'Content-Type: application/json' \
  --data '{
      "executor": 2,
      "wasmFileName": "yo.wasm",
      "wasmRegistryUrl": "https://localhost:9999/wasm/download/yo.wasm",
      "functionName": "yo",
      "functionRevision": "",
      "defaultRevision": true
    }
  ' http://localhost:9090/tasks
