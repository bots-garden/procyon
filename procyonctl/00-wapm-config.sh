#!/bin/bash
procyon_config="$(cat <<-EOF
WASM_REGISTRY_URL="https://registry-cdn.wapm.io/contents"
PROCYON_URL="http://localhost:9090"
PROCYON_REVERSE_URL="http://localhost:8080"
EOF
)"
echo "${procyon_config}" > procyon.config
