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
DOCKER_ENV_FILE=".env/.env.docker"

echo "🚀🚀🚀 Starting services with Docker Compose... 🚀🚀🚀"

if [ ! -f "$ENV_FILE" ]; then
  echo "❌ Environment file $ENV_FILE not found!"
  exit 1
fi

echo "📄 Loading environment from $ENV_FILE"

# Read env vars from file (handles values with spaces, =, and special chars)
POSTGRES_USER=""
POSTGRES_PASSWORD=""
POSTGRES_DB=""
while IFS='=' read -r key value; do
    # Strip carriage returns (Windows line endings)
    key=$(printf '%s' "$key" | tr -d '\r')
    value=$(printf '%s' "$value" | tr -d '\r')
    [ -z "$key" ] && continue
    case "$key" in \#*) continue ;; esac
    case "$key" in
        POSTGRES_USER)     POSTGRES_USER="$value" ;;
        POSTGRES_PASSWORD) POSTGRES_PASSWORD="$value" ;;
        POSTGRES_DB)       POSTGRES_DB="$value" ;;
    esac
done < "$ENV_FILE"

if [ -z "$POSTGRES_USER" ] || [ -z "$POSTGRES_PASSWORD" ] || [ -z "$POSTGRES_DB" ]; then
  echo "❌ Missing required POSTGRES_USER, POSTGRES_PASSWORD, or POSTGRES_DB in $ENV_FILE"
  exit 1
fi

echo "✅ Found DB credentials (user=$POSTGRES_USER, db=$POSTGRES_DB)"

# Generate a Docker-ready env file:
# - Copy all original vars (so both postgres and backend get identical values from the same parser)
# - Replace DATABASE_URL to use Docker compose service hostname (postgres-db) instead of localhost
cp "$ENV_FILE" "$DOCKER_ENV_FILE"
# Remove any existing DATABASE_URL line
sed -i'' -e '/^DATABASE_URL=/d' "$DOCKER_ENV_FILE"
# Append the Docker-specific DATABASE_URL
echo "DATABASE_URL=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@postgres-db:5432/${POSTGRES_DB}?sslmode=disable" >> "$DOCKER_ENV_FILE"

echo "📄 Generated $DOCKER_ENV_FILE with DATABASE_URL pointing to postgres-db"

# Start the containers in detached mode
# No compose-level ${} interpolation needed — everything comes from env_file
docker compose up -d

echo "✅ Services are running in the background."

if [ "$DETACH" = false ]; then
  echo "➡️ Tailing logs from the backend service. Press Ctrl+C to stop tailing."
  docker compose logs -f backend
fi