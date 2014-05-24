#!/bin/bash

docker ps | grep 'denderello/hello-server:latest' | awk '{ print $1 }' | xargs docker stop
docker ps -a | grep 'denderello/hello-server:latest' | awk '{ print $1 }' | xargs docker rm
