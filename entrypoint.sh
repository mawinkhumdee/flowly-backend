#!/bin/sh
echo "--- Runtime Debug Info ---"
echo "Working directory: $(pwd)"
echo "Listing contents of $(pwd):"
ls -la
echo "Listing contents of /app:"
ls -la /app
echo "Checking if /app/server exists:"
if [ -f "/app/server" ]; then
    echo "Binary found! Executing..."
    exec /app/server
else
    echo "ERROR: /app/server NOT FOUND"
    exit 1
fi
