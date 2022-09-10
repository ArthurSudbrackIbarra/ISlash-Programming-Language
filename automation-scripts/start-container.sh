#!/bin/bash

# docker-compose file path.
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
DOCKER_COMPOSE_FILE_PATH="$SCRIPT_DIR/../docker-compose.yaml"

# Stops the container.
docker stop islash-playground-container
# Removes the container.
docker rm islash-playground-container
# Removes the image.
docker rmi islash-programming-language_islash-playground
# Starts the container.
docker compose up -d
# Enters inside the container.
docker exec -it islash-playground-container sh
