#!/bin/bash
mkcert -install
mkcert procyon-reverse.local localhost 127.0.0.1 ::1
cp procyon-reverse.local+3-key.pem procyon-reverse.local.key
cp procyon-reverse.local+3.pem procyon-reverse.local.crt
rm *.pem
