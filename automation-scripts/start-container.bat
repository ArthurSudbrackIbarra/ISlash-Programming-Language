@echo off

:: docker-compose file path.
set DOCKER_COMPOSE_FILE_PATH="%~dp0..\docker-compose.yaml"

:: Stops the container.
docker stop islash-container
:: Removes the container.
docker rm islash-container
:: Removes the image.
docker rmi islash/islash-programming-language:v1
:: Starts the container.
docker compose -f %DOCKER_COMPOSE_FILE_PATH% up -d
:: Enters inside the container.
docker exec -it islash-container /bin/bash
