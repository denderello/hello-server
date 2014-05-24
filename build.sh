#!/bin/bash

(cd code ; ./build.sh)
./go install github.com/denderello/hello-server

docker build -t denderello/hello-server .
