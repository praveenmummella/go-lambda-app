#!/bin/bash
set -e  # Exit on error

# Clean previous builds
rm -f bootstrap function-*.zip

echo "Building Go binary for Lambda..."
# Key change: Added CGO_ENABLED=0 for static compilation
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bootstrap main.go

echo "Creating deployment package..."
VERSION=$(date +%Y%m%d%H%M%S)
zip -0 "function-$VERSION.zip" bootstrap

# Create a consistent 'function.zip' for local testing
cp "function-$VERSION.zip" function.zip

echo "Build complete:"
ls -lh function-*.zip

# Verify the binary is statically linked (optional)
echo "Verifying binary compatibility..."
file bootstrap
ldd bootstrap 2>&1 | grep -q "not a dynamic executable" && echo "Binary is statically linked" || echo "Warning: Binary is dynamically linked"
