#!/bin/bash
echo "-----------------------------------------"

echo "forty-two default:"; ./procyonctl func get forty-two
echo "hello-world default:"; ./procyonctl func get hello-world

echo "-----------------------------------------"

echo "forty-two rev1:"; ./procyonctl func get-revision forty-two rev1
echo "hello-world rev1:"; ./procyonctl func get-revision hello-world rev1
echo "hello-world rev2:"; ./procyonctl func get-revision hello-world rev2

echo "-----------------------------------------"

./procyonctl func list

./procyonctl task set-default-revision forty-two rev1 on

#./procyonctl task set-default-revision hello-world rev2 on

./procyonctl task set-default-revision hello-world rev1 on
./procyonctl task set-default-revision hello-world rev2 off

./procyonctl task set-default-revision hello-world rev1 off
./procyonctl task set-default-revision hello-world rev2 on