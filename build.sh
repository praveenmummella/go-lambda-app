#!/bin/bash
set -e  # Exit on error

# Clean previous builds
rm -f bootstrap function-*.zip

echo "Building Go binary for Lambda..."
GOOS=linux GOARCH=amd64 go build -o bootstrap main.go

echo "Creating deployment package..."
VERSION=$(date +%Y%m%d%H%M%S)
zip -0 "function-$VERSION.zip" bootstrap

# Create a consistent 'function.zip' for local testing
cp "function-$VERSION.zip" function.zip

echo "Build complete:"
ls -lh function-*.zip
