#!/bin/bash

docker-compose down --volumes

docker-compose --env-file ./config/env.dev -f docker-compose-dev.yml up --build -d

chmod +x "$0"