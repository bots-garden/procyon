#!/bin/bash


procyonctl task deploy hello-world.1.0.1.wasm hello-world rev1
procyonctl task deploy hello-world.1.0.2.wasm hello-world rev2
procyonctl task deploy forty-two.0.0.0.wasm forty-two rev1

