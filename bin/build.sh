#!/bin/bash

# create the needed GOPATH structure so that gotools plays nicely
mkdir -p build/{src/github.com/denderello/hello-server,pkg,bin}

# link all needed projects files to intermediate GOPATH
# build the final container with the compiled binary
(cd build/src/github.com/denderello/hello-server ; ln -f ../../../../../*.go .)
(cd build ; ln -f ../Makefile)

# build the development Docker container
(cd build ; ./build.sh)

# build and install the project to the intermediate GOPATH
./bin/go.sh install github.com/denderello/hello-server

# build the final container with the compiled binary
docker build -t denderello/hello-server .
