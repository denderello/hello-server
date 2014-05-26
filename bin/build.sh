#!/bin/bash

rm -rf build/{src,pkg,bin}
mkdir -p build/{src/github.com/denderello/hello-server,pkg,bin}

cp *.go build/src/github.com/denderello/hello-server

(cd build ; ./build.sh)

./bin/go.sh install github.com/denderello/hello-server

docker build -t denderello/hello-server .
