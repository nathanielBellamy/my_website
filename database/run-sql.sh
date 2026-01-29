#!/bin/bash

# This script runs the init.sql file to initialize the database.
# It should be run from the root of the project.

# Load environment variables from .env.localhost
if [ -f ".env.localhost" ]; then
    export $(cat .env.localhost | xargs)
else
    echo ".env.localhost not found!"
    exit 1
fi

psql -d postgres -f database/init.sql
