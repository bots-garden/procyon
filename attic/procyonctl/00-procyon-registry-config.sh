#!/bin/bash
procyon_config="$(cat <<-EOF
WASM_REGISTRY_URL="https://localhost:7070/get"
PROCYON_URL="http://localhost:9090"
PROCYON_REVERSE_URL="http://localhost:8080"
EOF
)"
echo "${procyon_config}" > procyon.config
