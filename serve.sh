#!/bin/bash

# Function to start the Go server
serve() {
  echo "🚀🚀🚀 Now serving $MODE on :8080 🚀🚀🚀"
  cd build 
  export MODE=$MODE
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


