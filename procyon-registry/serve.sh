#!/bin/bash
REGISTRY_CRT="certs/procyon-registry.local.crt" \
REGISTRY_KEY="certs/procyon-registry.local.key" \
REGISTRY_HTTP=7070 \
REGISTRY_HTTPS=7070 \
REGISTRY_FUNCTIONS_PATH="./functions" \
node index.js 



