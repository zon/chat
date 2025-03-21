#!/bin/bash

set -e

IMAGE=zvonimir/wurbs
VERSION=$(cat ./version)
SUFFIX=""
TAG="$VERSION$SUFFIX"

podman build --tag $IMAGE:$TAG .
podman tag $IMAGE:$TAG $IMAGE:latest

podman push $IMAGE:$TAG
podman push $IMAGE:latest