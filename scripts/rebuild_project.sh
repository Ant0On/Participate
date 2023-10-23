#!/bin/bash

docker-compose --env-file ./config/env.dev up -d --build

chmod +x "$0"