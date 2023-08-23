#!/bin/bash

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

# Check if PW is set
if [ -z "$PW" ]; then
  echo "Please set the PW environment variable. (Your site's password - localhost and remotedev only.)"
  exit 1
fi

# Function to start the Go server
serve() {
  echo "🚀🚀🚀 Now serving $MODE on :8080 🚀🚀🚀"
  cd build 
  export MODE=$MODE PW=$PW
  ./main
  cd ..
  echo "🫡 Server Is Out. Process completed.🫡"
}

# Handle different modes
# safeguard against erroneous mode var
case $MODE in
  localhost)
    serve
    ;;
  remotedev)
    serve
    ;;
  prod)
    serve
    ;;
  *)
    echo "Invalid MODE. Choose between localhost, remotedev, or prod."
    exit 1
    ;;
esac


