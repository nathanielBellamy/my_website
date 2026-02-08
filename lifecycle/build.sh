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
  if [ -f "./../../.env.$MODE" ]; then
      rm -f ./../../build/.env.* # remove old .env files
      cp "./../../.env.$MODE" ./../../build/ # copy in current
  else
      echo ".env.$MODE not found in root directory!"
      exit 1
  fi

  cd ../..
    cat << EOF

  📣  🏁  DONE:
GO SERVER BUILT

EOF
}

# Function for auth SPA build
build_auth_dev_spa() {
  cat << EOF

  📣  🏗️   BUILDING:
AUTH SPA

EOF
  SPA_ENV=$1
  cd auth/dev && npm run build-$SPA_ENV 
  cd ../..
  cat << EOF

  📣  🏁  DONE:
AUTH SPA BUILT

EOF

  # Perform the regex string replacement
  sed -i '' 's/\/assets/\.\/assets/g' build/auth/dev/index.html
  cat << EOF

  📣  🏁  DONE:
UPDATED ASSET PATHS IN AUTH SPA index.html

EOF
}

# Function for old-site SPA build
build_old_site_spa() {
  cat << EOF
  📣  🏗️   BUILDING:
old-site SPA
EOF

  SPA_ENV=$1
  cd old-site && npm run build-old-site-$SPA_ENV 
  cd ..
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
  cd marketing && npm run build-marketing-$SPA_ENV 
  cd ..
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
  cd admin && npm run build-admin-$SPA_ENV 
  cd ..
  cat << EOF

  📣  🏁  DONE:
ADMIN SPA BUILT

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
    build_auth_dev_spa "localhost"
    if [ "$incl_old_site" = true ]; then
      build_old_site_spa "localhost"
    fi
    build_marketing_spa "localhost"
    build_admin_spa "localhost"
    ;;
  remotedev)
    build_go_server "remotedev"
    build_auth_dev_spa "remotedev"
    build_old_site_spa "remotedev"
    build_marketing_spa "remotedev"
    build_admin_spa "remotedev"
    ;;
  prod)
    build_go_server "prod"
    build_auth_dev_spa "prod"
    build_old_site_spa "prod"
    build_marketing_spa "prod"
    build_admin_spa "prod"
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
