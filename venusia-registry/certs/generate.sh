#!/bin/bash
mkcert -install
mkcert venusia.local localhost 127.0.0.1 ::1
cp venusia.local+3-key.pem venusia.local.key
cp venusia.local+3.pem venusia.local.crt
rm *.pem

