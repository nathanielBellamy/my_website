#!/bin/bash

# Check if MODE is set
if [ -z "$MODE" ]; then
  echo "Please set the MODE environment variable. (localhost, remotedev, or prod)"
  exit 1
fi

# Function for Go server build
build_go_server() {
  echo "Building Go server..."
  cd backend && cd go && go build -o "./../../build" main.go
  cd .. && cd ..
  echo "Go server built successfully."
}

# Function for auth SPA build
build_auth_dev_spa() {
  echo "Building Auth SPA..."
  SPA_ENV=$1
  cd auth && cd dev && npm run build-$SPA_ENV 
  cd .. && cd ..
  echo "Auth SPA built successfully."
}

# Function for main SPA build
build_main_spa() {
  echo "Building Main SPA..."
  SPA_ENV=$1
  cd frontend && npm run build-frontend-$SPA_ENV 
  cd ..
  echo "Main SPA built successfully."
}

# Handle different modes
case $MODE in
  localhost)
    build_go_server "localhost"
    build_auth_dev_spa "localhost"
    build_main_spa "localhost"
    ;;
  remotedev)
    build_go_server "remotedev"
    build_auth_dev_spa "remotedev"
    build_main_spa "remotedev"
    ;;
  prod)
    build_go_server "prod"
    build_auth_dev_spa "prod"
    build_main_spa "prod"
    ;;
  *)
    echo "Invalid MODE. Choose between localhost, remotedev, or prod."
    exit 1
    ;;
esac

echo "Build process completed."
