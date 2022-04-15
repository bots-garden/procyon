#!/bin/bash
mkcert -install
mkcert procyon-registry.local localhost 127.0.0.1 ::1
cp procyon-registry.local+3-key.pem procyon-registry.local.key
cp procyon-registry.local+3.pem procyon-registry.local.crt
rm *.pem

