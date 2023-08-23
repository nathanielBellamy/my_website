#!/bin/bash

# Call the original build.sh to create the build
./build.sh

# Either create or empty out a top-level directory named 'dist'
if [ -d "dist" ]; then
    rm -rf dist/*
else
    mkdir dist
fi

# Copy the serve.sh file, config.env file, and the entire build directory into the 'dist' directory
cp serve.sh dist/
cp config.env dist/
cp -r build dist/

echo "Distribution built and ready in 'dist' directory!"

