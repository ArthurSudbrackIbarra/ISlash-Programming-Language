#!/bin/bash

# docker-compose file path.
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
DOCKER_COMPOSE_FILE_PATH="$SCRIPT_DIR/../docker-compose.yaml"

# Stops the container.
docker stop islash-container
# Removes the container.
docker rm islash-container
# Removes the image.
docker rmi islash/islash-programming-language:v1
# Starts the container.
docker compose up -d
# Enters inside the container.
docker exec -it islash-container /bin/bash
