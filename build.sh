#!/bin/bash
#

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
  cd backend && cd go
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
  if [ -f ".env.$MODE" ]; then
      rm ./../../build/.env.* # remove old .env files
      cp ".env.$MODE" ./../../build/ # copy in current
  else
      echo ".env.$MODE not found!"
      exit 1
  fi

  cd .. && cd ..
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
  cd auth && cd dev && npm run build-$SPA_ENV 
  cd .. && cd ..
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

# Function for old-site main SPA build
build_main_spa() {
  cat << EOF

  📣  🏗️   BUILDING:
FRONTEND SPA

EOF
  SPA_ENV=$1
  cd frontend && npm run build-frontend-$SPA_ENV 
  cd ..
  cat << EOF

  📣  🏁  DONE:
FRONTEND SPA BUILT

EOF

  # Perform the regex string replacement
  sed -i '' -e 's/src="\/assets/src="\/old-site\/assets/g' -e 's/href="\/assets/href="\/old-site\/assets/g' build/frontend/index.html
  cat << EOF

  📣  🏁  DONE:
UPDATED ASSET PATHS IN FRONTEND SPA index.html

EOF
}

# Function for marketing SPA build
build_marketing_spa() {
  cat << EOF

  📣  🏗️   BUILDING:
MARKETING SPA

EOF
  SPA_ENV=$1
  cd marketing && npm run build-frontend-$SPA_ENV 
  cd ..
  cat << EOF

  📣  🏁  DONE:
MARKETING SPA BUILT

EOF

  # Perform the regex string replacement
  sed -i '' 's/\/assets/\.\/assets/g' build/marketing/index.html
  cat << EOF

  📣  🏁  DONE:
UPDATED ASSET PATHS IN MARKETING SPA index.html

EOF
}

cat << EOF

  📣  🏗️   BUILDING WEBSITE
  ⚡  MODE:
${MODE}

EOF

# Check if we only want to build the Go server
if [ "$1" == "--server-only" ]; then
  build_go_server $MODE
  exit 0
fi

# Handle different modes
case $MODE in
  localhost)
    build_go_server "localhost"
    build_auth_dev_spa "localhost"
    build_main_spa "localhost"
    build_marketing_spa "localhost"
    ;;
  remotedev)
    build_go_server "remotedev"
    build_auth_dev_spa "remotedev"
    build_main_spa "remotedev"
    build_marketing_spa "remotedev"
    ;;
  prod)
    build_go_server "prod"
    build_auth_dev_spa "prod"
    build_main_spa "prod"
    build_marketing_spa "prod"
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
