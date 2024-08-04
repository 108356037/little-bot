#!/bin/bash
docker build -t 108356037/tigris-bot:alpha-v2 .

docker rm -f tigris-bot-short
docker rm -f tigris-bot-long

docker run --restart=always --name tigris-bot-short --env-file .env.short -tid 108356037/tigris-bot:alpha-v2
docker run --restart=always --name tigris-bot-long --env-file .env.long -tid 108356037/tigris-bot:alpha-v2