#!/bin/bash

set -e

IMAGE=zvonimir/wurbs
VERSION=$(cat ./version)
SUFFIX=""
TAG="$VERSION$SUFFIX"

docker build --platform linux/amd64 --tag $IMAGE:$TAG .
docker tag $IMAGE:$TAG $IMAGE:latest

docker push $IMAGE:$TAG
docker push $IMAGE:latest
