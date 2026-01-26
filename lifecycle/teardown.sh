#!/bin/bash

# Ensure we are in the project root
cd "$(dirname "$0")/.."

# Default to stop, but use down if --down is passed
DOWN_COMMAND="stop"
if [ "$1" == "--down" ]; then
    DOWN_COMMAND="down"
fi

# Copy logs before stopping/destroying the container
if [ ! -z "$(docker ps -q -f name=my_website_backend)" ]; then
    echo "Copying logs from backend container..."
    LOG_DIR="log/$(date -u +%Y)/$(date -u +%m)"
    mkdir -p "$LOG_DIR"
    TIMESTAMP=$(date -u +%Y-%m-%dT%H-%M-%SZ)
    LOG_FILE="$LOG_DIR/${TIMESTAMP}-log.txt"
    docker cp my_website_backend:/app/log.txt "$LOG_FILE"
    echo "Logs saved to $LOG_FILE"
else
    echo "Backend container not running or does not exist. Skipping log copy."
fi

# Stop or tear down the containers
echo "Running 'docker-compose $DOWN_COMMAND'..."
docker-compose $DOWN_COMMAND

echo "Teardown complete."
