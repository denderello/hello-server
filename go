#!/bin/bash

docker run \
  --rm \
  -i \
  -t \
  -v $(pwd)/code:/opt/hello-server \
  -e "GOPATH=/opt/hello-server"\
  denderello/hello-server-dev "go $*"
