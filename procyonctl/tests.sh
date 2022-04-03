#!/bin/bash
./procyonctl task deploy hello-v0.0.0.wasm hello rev1
./procyonctl task deploy hello-v0.0.1.wasm hello rev2
./procyonctl task deploy hey-v0.0.0.wasm hey rev1

