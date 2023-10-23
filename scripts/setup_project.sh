#!/bin/bash

docker-compose --env-file ./config/env.dev up -d

chmod +x "$0"