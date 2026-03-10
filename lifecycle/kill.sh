#!/bin/bash

set -e

cd "$(dirname "$0")/.."

# Function to display usage information
usage() {
    echo "Usage: $0 [ -f | -d | -b | -l | -v ]"
    echo "  -f, --full      Teardown both the my_website_db and my_website_backend containers"
    echo "  -d, --db        Teardown only the my_website_db container"
    echo "  -b, --be        Teardown only the my_website_backend container"
    echo "  -l, --logs      Copy logs from the my_website_backend container without tearing down"
    echo "  -v, --volumes   Teardown entire stack AND erase all volumes and stored data (full reset)"
    exit 1
}

# --- Log copy function ---
copy_logs() {
    echo "Copying logs from my_website_backend container..."
    YEAR=$(date -u +"%Y")
    MONTH=$(date -u +"%m")
    mkdir -p "log/$YEAR/$MONTH"
    TIMESTAMP=$(date -u +"%Y-%m-%dT%H-%M-%SZ")
    LOG_FILE_NAME="${TIMESTAMP}-log.txt"
    docker logs my_website_backend > "log/$YEAR/$MONTH/$LOG_FILE_NAME"
    echo "Logs copied to log/$YEAR/$MONTH/$LOG_FILE_NAME"
}

# --- Teardown functions ---
teardown_my_website_backend() {
    echo "Tearing down my_website_backend container..."
    docker stop my_website_backend
    docker rm -f my_website_backend
    echo "my_website_backend container torn down."
}

teardown_db() {
    echo "Tearing down db container..."
    docker stop my_website_db
    docker rm -f my_website_db
    echo "DB container torn down."
}


# --- Main script logic ---
if [ $# -ne 1 ]; then
    usage
fi

case "$1" in
    -f|--full)
        copy_logs
        echo "Tearing down entire docker compose stack..."
        docker-compose down
        echo "Entire docker compose stack torn down."
        ;;
    -d|--db)
        teardown_db
        ;;
    -b|--be)
        copy_logs
        teardown_my_website_backend
        ;;
    -l|--logs)
        copy_logs
        ;;
    -v|--volumes)
        copy_logs
        echo "Tearing down entire docker compose stack and erasing all volumes..."
        docker-compose down -v
        echo "Entire docker compose stack torn down. All volumes erased."
        ;;
    *)
        usage
        ;;
esac

exit 0