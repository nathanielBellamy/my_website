#!/usr/bin/env bash
cd "$(dirname "$0")/.."

DETACH=false

while [[ "$#" -gt 0 ]]; do
    case $1 in
        -d|--detach) DETACH=true ;;
        *) echo "Unknown parameter: $1"; exit 1 ;;
    esac
    shift
done

ENV_FILE=".env/.env.production"

echo "🚀🚀🚀 Starting services with Docker Compose... 🚀🚀🚀"

if [ ! -f "$ENV_FILE" ]; then
  echo "❌ Environment file $ENV_FILE not found!"
  exit 1
fi

echo "📄 Using environment from $ENV_FILE"

# Start the containers in detached mode
# --env-file provides variables for compose-file interpolation (e.g. ${POSTGRES_USER})
# The compose file's env_file directive delivers variables into the containers
docker compose --env-file "$ENV_FILE" up -d

echo "✅ Services are running in the background."

if [ "$DETACH" = false ]; then
  echo "➡️ Tailing logs from the backend service. Press Ctrl+C to stop tailing."
  docker compose logs -f backend
fi