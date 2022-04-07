#!/bin/bash

./procyonctl task set-default-revision forty-two rev1 on

#./procyonctl task set-default-revision hello-world rev2 on

./procyonctl task set-default-revision hello-world rev1 on
./procyonctl task set-default-revision hello-world rev2 off

./procyonctl task set-default-revision hello-world rev1 off
./procyonctl task set-default-revision hello-world rev2 on