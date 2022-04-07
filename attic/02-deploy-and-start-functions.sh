#!/bin/bash

curl -v --request POST \
  --header 'Content-Type: application/json' \
  --data '{
      "executor": 1,
      "wasmFileName": "hello.wasm",
      "wasmRegistryUrl": "https://localhost:9999/wasm/download/hello.wasm",
      "functionName": "hello",
      "functionRevision": "first",
      "defaultRevision": true
    }
  ' http://localhost:9090/tasks


curl -v --request POST \
  --header 'Content-Type: application/json' \
  --data '{
      "executor": 1,
      "wasmFileName": "hey.wasm",
      "wasmRegistryUrl": "https://localhost:9999/wasm/download/hey.wasm",
      "functionName": "hey",
      "functionRevision": "first",
      "defaultRevision": true
    }
  ' http://localhost:9090/tasks



curl -v --request POST \
  --header 'Content-Type: application/json' \
  --data '{
      "executor": 1,
      "wasmFileName": "hi.wasm",
      "wasmRegistryUrl": "https://localhost:9999/wasm/download/hi.wasm",
      "functionName": "hi",
      "functionRevision": "first",
      "defaultRevision": true
    }
  ' http://localhost:9090/tasks


curl -v --request POST \
  --header 'Content-Type: application/json' \
  --data '{
      "executor": 1,
      "wasmFileName": "yo.wasm",
      "wasmRegistryUrl": "https://localhost:9999/wasm/download/yo.wasm",
      "functionName": "yo",
      "functionRevision": "first",
      "defaultRevision": true
    }
  ' http://localhost:9090/tasks


curl -v --request POST \
  --header 'Content-Type: application/json' \
  --data '{
      "executor": 1,
      "wasmFileName": "hello.wasm",
      "wasmRegistryUrl": "https://localhost:9999/wasm/download/hello.wasm",
      "functionName": "hello",
      "functionRevision": "orange",
      "defaultRevision": false
    }
  ' http://localhost:9090/tasks

curl -v --request POST \
  --header 'Content-Type: application/json' \
  --data '{
      "executor": 1,
      "wasmFileName": "hello.wasm",
      "wasmRegistryUrl": "https://localhost:9999/wasm/download/hello.wasm",
      "functionName": "hello",
      "functionRevision": "red",
      "defaultRevision": false
    }
  ' http://localhost:9090/tasks
