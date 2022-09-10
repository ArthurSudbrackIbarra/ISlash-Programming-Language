@echo off

:: docker-compose file path.
set DOCKER_COMPOSE_FILE_PATH="%~dp0..\docker-compose.yaml"

:: Stops the container.
docker stop islash-playground-container
:: Removes the container.
docker rm islash-playground-container
:: Removes the image.
docker rmi islash-programming-language_islash-playground
:: Starts the container.
docker compose -f %DOCKER_COMPOSE_FILE_PATH% up -d
:: Enters inside the container.
docker exec -it islash-playground-container sh
