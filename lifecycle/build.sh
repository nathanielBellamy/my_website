#!/bin/bash
#
cd "$(dirname "$0")/.."

# Load environment variables from config.env
if [ -f "config.env" ]; then
    export $(cat config.env | xargs)
else
    echo "config.env not found!"
    exit 1
fi

# Check if MODE is set
if [ -z "$MODE" ]; then
  echo "Please set the MODE environment variable. (localhost, remotedev, or prod)"
  exit 1
fi

# Function for Go server build
build_go_server() {
    cat << EOF
  📣  🏗️   BUILDING:
GO SERVER
EOF

  cd backend/go
  TARGET_ARCH="LINUX"
  if [ "$MODE" != "localhost" ]; then

    cat << EOF

  ⚡  GO TARGET ARCH:
LINUX

EOF

    GOOS=linux GOARCH=amd64 go build -o "./../../build" main.go
  else
    TARGET_ARCH="HOST_ARCHITECTURE"
    cat << EOF

  ⚡  GO TARGET ARCH:
HOST ARCHITECTURE

EOF
    go build -o "./../../build" main.go
  fi

  # copy .env file to build directory
  if [ -f "./../../.env/.env.$MODE" ]; then
      rm -f ./../../build/.env/.env.* # remove old .env files
      mkdir -p ./../../build/.env
      cp "./../../.env/.env.$MODE" ./../../build/.env/ # copy in current
  else
      echo ".env/.env.$MODE not found in .env directory!"
      exit 1
  fi

  cd ../..
    cat << EOF

  📣  🏁  DONE:
GO SERVER BUILT

EOF
}

# Function for auth SPA build
build_auth_admin_spa() {
  cat << EOF

  📣  🏗️   BUILDING:
AUTH SPA

EOF
  SPA_ENV=$1
  cd frontend/auth/admin && npm run build-$SPA_ENV
  cd ../../..
  cat << EOF

  📣  🏁  DONE:
AUTH SPA BUILT

EOF
}

# Function for old-site SPA build
build_old_site_spa() {
  cat << EOF
  📣  🏗️   BUILDING:
old-site SPA
EOF

  SPA_ENV=$1
  cd frontend/old-site && npm run build-old-site-$SPA_ENV
  cd ../..
  cat << EOF

  📣  🏁  DONE:
old-site SPA BUILT

EOF

  # Perform the regex string replacement
  sed -i '' -e 's/src="\/assets/src="\/old-site\/assets/g' -e 's/href="\/assets/href="\/old-site\/assets/g' build/old-site/index.html
  cat << EOF

  📣  🏁  DONE:
UPDATED ASSET PATHS IN old-site SPA index.html

EOF
}

# Function for marketing SPA build
build_marketing_spa() {
  cat << EOF

  📣  🏗️   BUILDING:
MARKETING SPA

EOF
  SPA_ENV=$1
  cd frontend/marketing && npm run build-marketing-$SPA_ENV
  cd ../..
  cat << EOF

  📣  🏁  DONE:
MARKETING SPA BUILT

EOF
}

# Function for admin SPA build
build_admin_spa() {
  cat << EOF

  📣  🏗️   BUILDING:
ADMIN SPA

EOF
  SPA_ENV=$1
  cd frontend/admin && npm run build-admin-$SPA_ENV
  cd ../..
  cat << EOF

  📣  🏁  DONE:
ADMIN SPA BUILT

EOF
}

# Function to inject New Relic browser agent config into built SPA index.html files.
# Values are read from the mode's .env file (e.g. .env/.env.localhost), mirroring the
# secret substitution done in .github/workflows/deploy.yml for remotedev/prod.
inject_new_relic_config() {
  ENV_MODE=$1
  ENV_FILE="./.env/.env.$ENV_MODE"

  if [ ! -f "$ENV_FILE" ] || ! grep -q "^NEW_RELIC_BROWSER_LICENSE_KEY=" "$ENV_FILE"; then
    cat << EOF

  ⚠️   SKIPPING New Relic browser config injection
  (no NEW_RELIC_BROWSER_LICENSE_KEY found in ${ENV_FILE})

EOF
    return
  fi

  cat << EOF

  📣  💉  INJECTING:
NEW RELIC BROWSER CONFIG (${ENV_MODE})

EOF

  # Load only the NEW_RELIC_* vars from the mode's env file (scoped to this function)
  while IFS='=' read -r key value; do
    export "$key"="$value"
  done < <(grep '^NEW_RELIC_' "$ENV_FILE")

  inject_file() {
    FILE=$1
    AGENT_ID_VALUE=$2
    APP_ID_VALUE=$3
    AGENT_PLACEHOLDER=$4
    APP_PLACEHOLDER=$5

    if [ ! -f "$FILE" ]; then
      return
    fi

    sed -i '' \
      -e "s|__NEW_RELIC_ACCOUNT_ID__|${NEW_RELIC_ACCOUNT_ID}|g" \
      -e "s|__NEW_RELIC_TRUST_KEY__|${NEW_RELIC_TRUST_KEY}|g" \
      -e "s|__NEW_RELIC_BROWSER_LICENSE_KEY__|${NEW_RELIC_BROWSER_LICENSE_KEY}|g" \
      -e "s|${AGENT_PLACEHOLDER}|${AGENT_ID_VALUE}|g" \
      -e "s|${APP_PLACEHOLDER}|${APP_ID_VALUE}|g" \
      "$FILE"

    echo "  ✅  Injected New Relic config into $FILE"
  }

  inject_file "build/marketing/browser/index.html" \
    "$NEW_RELIC_AGENT_ID_MARKETING" "$NEW_RELIC_APP_ID_MARKETING" \
    "__NEW_RELIC_AGENT_ID_MARKETING__" "__NEW_RELIC_APP_ID_MARKETING__"

  inject_file "build/admin/browser/index.html" \
    "$NEW_RELIC_AGENT_ID_ADMIN" "$NEW_RELIC_APP_ID_ADMIN" \
    "__NEW_RELIC_AGENT_ID_ADMIN__" "__NEW_RELIC_APP_ID_ADMIN__"

  inject_file "build/auth/admin/browser/index.html" \
    "$NEW_RELIC_AGENT_ID_AUTH" "$NEW_RELIC_APP_ID_AUTH" \
    "__NEW_RELIC_AGENT_ID_AUTH__" "__NEW_RELIC_APP_ID_AUTH__"

  inject_file "build/old-site/index.html" \
    "$NEW_RELIC_AGENT_ID_OLDSITE" "$NEW_RELIC_APP_ID_OLDSITE" \
    "__NEW_RELIC_AGENT_ID_OLDSITE__" "__NEW_RELIC_APP_ID_OLDSITE__"

  cat << EOF

  📣  🏁  DONE:
NEW RELIC BROWSER CONFIG INJECTED

EOF
}

######

cat << EOF

  📣  🏗️   BUILDING WEBSITE
  ⚡  MODE:
${MODE}

EOF

SERVER_ONLY=false
incl_old_site=false

# Parse command-line arguments
while [[ "$#" -gt 0 ]]; do
  case "$1" in
    --server-only)
      SERVER_ONLY=true
      ;;
    --incl-old)
      incl_old_site=true
      ;;
    *)
      echo "Unknown parameter passed: $1"
      exit 1
      ;;
  esac
  shift
done

# Check if we only want to build the Go server
if [ "$SERVER_ONLY" = true ]; then
  build_go_server $MODE
  exit 0
fi

# Handle different modes
case $MODE in
  localhost)
    build_go_server "localhost"
    build_auth_admin_spa "localhost"
    if [ "$incl_old_site" = true ]; then
      build_old_site_spa "localhost"
    fi
    build_marketing_spa "localhost"
    build_admin_spa "localhost"
    inject_new_relic_config "localhost"
    ;;
  remotedev)
    build_go_server "remotedev"
    build_auth_admin_spa "remotedev"
    build_old_site_spa "remotedev"
    build_marketing_spa "remotedev"
    build_admin_spa "remotedev"
    inject_new_relic_config "remotedev"
    ;;
  prod)
    build_go_server "prod"
    build_auth_admin_spa "prod"
    build_old_site_spa "prod"
    build_marketing_spa "prod"
    build_admin_spa "prod"
    inject_new_relic_config "prod"
    ;;
  *)
    echo "Invalid MODE. Choose between localhost, remotedev, or prod."
    exit 1
    ;;
esac

cat << EOF

  📣  🏁  DONE:
BUILD COMPLETE
CHECK ABOVE OUTPUT FOR WARNINGS

  ⚡  VERIFY
  ⚡  VERIFY
  ⚡  VERIFY

  ⚡  GO TARGET ARCH:
${TARGET_ARCH}

  ⚡  MODE:
${MODE}

  🚀🚀🚀  
  🚀🚀🚀  Happy
  🚀🚀🚀  Coding

EOF
