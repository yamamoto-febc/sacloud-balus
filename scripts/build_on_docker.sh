#!/bin/bash

set -e

DOCKER_IMAGE_NAME="$1-build"
DOCKER_CONTAINER_NAME="$1-build-container"

if [[ $(docker ps -a | grep $DOCKER_CONTAINER_NAME) != "" ]]; then
  docker rm -f $DOCKER_CONTAINER_NAME 2>/dev/null
fi

docker build -t $DOCKER_IMAGE_NAME -f scripts/Dockerfile.build .

docker run --name $DOCKER_CONTAINER_NAME \
       -e SAKURACLOUD_ACCESS_TOKEN \
       -e SAKURACLOUD_ACCESS_TOKEN_SECRET \
       -e SAKURACLOUD_ZONES \
       -e SAKURACLOUD_TRACE_MODE \
       -e SACLOUD_BALUS_RECEIVE_PATH \
       -e SACLOUD_BALUS_RECEIVE_PORT \
       -e SACLOUD_BALUS_RECEIVE_SECRET \
       -e SACLOUD_BALUS_SEND_TOKEN \
       -e SACLOUD_BALUS_SEND_SECRET \
       -e SACLOUD_BALUS_SEND_MODULE_ID \
       -e SACLOUD_BALUS_DEBUG \
       -e TESTARGS \
       $DOCKER_IMAGE_NAME make "$@"
if [[ "$@" == *"build"* ]]; then
  docker cp $DOCKER_CONTAINER_NAME:`docker inspect -f "{{ .Config.WorkingDir  }}" $DOCKER_CONTAINER_NAME`/bin ./
fi
docker rm -f $DOCKER_CONTAINER_NAME 2>/dev/null
