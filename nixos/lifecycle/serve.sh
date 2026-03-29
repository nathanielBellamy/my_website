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
# --env-file provides vars for compose-level ${} interpolation in the YAML
# (env_file: in the YAML only injects vars INTO containers, not for YAML substitution)
docker compose --env-file "$DOCKER_ENV_FILE" up -d

echo "✅ Services are running in the background."

# --- Verify Grafana health ---
echo ""
echo "🔍 Verifying Grafana container health..."
GRAFANA_OK=false
for i in 1 2 3 4 5 6; do
  sleep 5
  HEALTH=$(docker exec my_website_grafana wget -qO- http://localhost:3000/api/health 2>/dev/null)
  if echo "$HEALTH" | grep -q '"database": "ok"'; then
    echo "✅ Grafana is healthy: $HEALTH"
    GRAFANA_OK=true
    break
  fi
  echo "   Attempt $i/6: Grafana not ready yet..."
done

if [ "$GRAFANA_OK" = false ]; then
  echo "⚠️  Grafana may not be healthy. Dumping diagnostics:"
  echo "--- Container status ---"
  docker inspect --format='{{.State.Status}} (health: {{.State.Health.Status}})' my_website_grafana 2>/dev/null || echo "Container not found"
  echo "--- Provisioning files ---"
  docker exec my_website_grafana ls -la /etc/grafana/provisioning/dashboards/ 2>/dev/null || echo "Cannot list provisioning dir"
  echo "--- Grafana config ---"
  docker exec my_website_grafana cat /etc/grafana/grafana.ini 2>/dev/null | head -20
  echo "--- Last 50 Grafana logs ---"
  docker logs my_website_grafana 2>&1 | tail -50
  echo "--- Home dashboard test ---"
  docker exec my_website_grafana wget -qO- http://localhost:3000/api/dashboards/home 2>&1 | head -5
  echo ""
  echo "⚠️  Review the above output to diagnose Grafana startup issues."
fi

if [ "$DETACH" = false ]; then
  echo "➡️ Tailing logs from the backend service. Press Ctrl+C to stop tailing."
  docker compose --env-file "$DOCKER_ENV_FILE" logs -f backend
fi