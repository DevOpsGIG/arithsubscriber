#!/bin/bash
set -e
set -u

IMAGE=$1
CONTAINER=$2

if docker inspect -f {{.State.Running}} "${CONTAINER}"; then
  docker stop "${CONTAINER}" && \
  docker rm "${CONTAINER}" && \
  docker rmi "${IMAGE}"
fi
