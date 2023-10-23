#!/bin/bash

cd ..

docker-compose --env-file ./config/env.dev down

docker rm -f $(docker ps -a -q)

docker volume rm $(docker volume ls -q)

docker rmi $(docker images -a -q)

docker system prune -f

rm -rf ./database

rm -rf ./frontend/node_modules