#!/bin/sh

go build -o chat
./chat -addr=":3000"
