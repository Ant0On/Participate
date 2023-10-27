#!/bin/bash

docker-compose down --volumes

docker-compose --env-file ./config/env.dev -f docker-compose-debug.yml up --build -d

chmod +x "$0"