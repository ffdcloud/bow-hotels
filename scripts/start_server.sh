#!/bin/bash

APP_DIR="/var/www/bow-hotels"
BUILD_FILENAME="build"
BUILD_OUTPUT="$APP_DIR/$BUILD_FILENAME"

cd $APP_DIR

# Build the Go binary
echo "Building the Go application..."
go build -o $BUILD_OUTPUT
chmod +x $BUILD_OUTPUT

# Stop the existing application (if running)
pkill $BUILD_FILENAME || true

# Start the new application
echo "Starting the application..."
nohup $BUILD_OUTPUT > $APP_DIR/app.log 2>&1 &

echo "Deployment complete!"
