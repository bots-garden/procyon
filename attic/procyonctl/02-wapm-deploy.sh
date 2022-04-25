#!/bin/bash
#./procyonctl task deploy hello-v0.0.0.wasm hello rev1
#./procyonctl task deploy hello-v0.0.1.wasm hello rev2
#./procyonctl task deploy hey-v0.0.0.wasm hey rev1


procyonctl task deploy k33g/hello-world/1.0.1/hello-world.wasm hello-world rev1
procyonctl task deploy k33g/hello-world/1.0.2/hello-world.wasm hello-world rev2
procyonctl task deploy k33g/forty-two/1.0.0/forty-two.wasm forty-two rev1
