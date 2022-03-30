#!/bin/bash
#tinygo build -o main.wasm -target wasm ./main.go

version="0.0.1"

tinygo build -o compute_${version}.wasm -target wasi ./main.go

ls -lh *.wasm

