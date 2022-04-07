#!/bin/bash

curl -v \
  -F "file=@./samples/hello/hello.wasm" \
	-H "Content-Type: multipart/form-data" \
	-X POST https://localhost:9999/wasm/upload

curl -v \
  -F "file=@./samples/hey/hey.wasm" \
	-H "Content-Type: multipart/form-data" \
	-X POST https://localhost:9999/wasm/upload

curl -v \
  -F "file=@./samples/hi/hi.wasm" \
	-H "Content-Type: multipart/form-data" \
	-X POST https://localhost:9999/wasm/upload

curl -v \
  -F "file=@./samples/yo/yo.wasm" \
	-H "Content-Type: multipart/form-data" \
	-X POST https://localhost:9999/wasm/upload


