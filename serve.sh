#!/bin/bash

# Check if MODE is set
if [ -z "$MODE" ]; then
  echo "Please set the MODE environment variable. (localhost, remotedev, or prod)"
  exit 1
fi

# Function to start the Go server
serve() {
  echo "🚀🚀🚀 Now serving $MODE on :8080 🚀🚀🚀"
  ENV=$1
  cd build && MODE=ENV ./main && cd ..
  echo "🫡 Serve Is Out. Process completed.🫡"
}

# Handle different modes
case $MODE in
  localhost)
    serve "localhost"
    ;;
  remotedev)
    serve "remotedev"
    ;;
  prod)
    serve "prod"
    ;;
  *)
    echo "Invalid MODE. Choose between localhost, remotedev, or prod."
    exit 1
    ;;
esac


