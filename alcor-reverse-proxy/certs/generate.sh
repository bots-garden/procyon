#!/bin/bash
mkcert -install
mkcert alcor.local localhost 127.0.0.1 ::1
cp alcor.local+3-key.pem alcor.local.key
cp alcor.local+3.pem alcor.local.crt
rm *.pem

