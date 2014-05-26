#!/bin/bash

docker run \
  --rm \
  -i \
  -t \
  -v $(pwd)/build:/opt/hello-server \
  -w "/opt/hello-server" \
  -e "GOPATH=/opt/hello-server"\
  denderello/hello-server-dev "make $*"
