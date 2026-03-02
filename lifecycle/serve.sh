#!/bin/bash
cd "$(dirname "$0")/.."

echo "🚀🚀🚀 Starting services with Docker Compose... 🚀🚀🚀"

# Build and start the containers in detached mode
docker compose up --build -d

echo "✅ Services are running in the background."
echo "➡️ Tailing logs from the backend service. Press Ctrl+C to stop tailing."

# Tail the logs of the backend service
docker compose logs -f backend