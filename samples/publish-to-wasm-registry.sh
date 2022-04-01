#!/bin/bash

curl -v \
  -F "file=@hello/hello.wasm" \
	-H "Content-Type: multipart/form-data" \
	-X POST https://localhost:9999/wasm/upload

curl -v \
  -F "file=@hey/hey.wasm" \
	-H "Content-Type: multipart/form-data" \
	-X POST https://localhost:9999/wasm/upload

curl -v \
  -F "file=@hi/hi.wasm" \
	-H "Content-Type: multipart/form-data" \
	-X POST https://localhost:9999/wasm/upload

curl -v \
  -F "file=@yo/yo.wasm" \
	-H "Content-Type: multipart/form-data" \
	-X POST https://localhost:9999/wasm/upload

