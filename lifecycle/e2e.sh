#!/bin/bash
cd "$(dirname "$0")/.."

HEADLESS=true
ENV="localhost"

while [[ "$#" -gt 0 ]]; do
    case $1 in
        -h|--head) HEADLESS=false ;;
        -e|--env) ENV="$2"; shift ;;
        *) echo "Unknown parameter passed: $1"; exit 1 ;;
    esac
    shift
done

echo "🧪 Running E2E tests..."
echo "📍 Environment: $ENV"
echo "🖥️ Headless: $HEADLESS"

# Install dependencies if node_modules doesn't exist
if [ ! -d "e2e/node_modules" ]; then
    echo "📦 Installing E2E dependencies..."
    cd e2e && npm install && cd ..
fi

CYPRESS_CMD="npx cypress run"
if [ "$HEADLESS" = false ]; then
    CYPRESS_CMD="npx cypress run --headed --browser chrome"
fi

# Set base URL based on environment
case $ENV in
    "localhost")
        BASE_URL="http://localhost:8080"
        ;;
    "remotedev")
        BASE_URL="https://dev.nateschieber.dev" # Assuming this is the remote dev URL
        ;;
    *)
        echo "Unknown environment: $ENV"
        exit 1
        ;;
esac

echo "🌐 Base URL: $BASE_URL"

cd e2e && $CYPRESS_CMD --env mode=$ENV,baseUrl=$BASE_URL
