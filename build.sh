#!/bin/bash
set -e  # Exit on error

echo "Building Go binary..."
GOOS=linux GOARCH=amd64 go build -o bootstrap main.go

echo "Creating ZIP package..."
zip function.zip bootstrap

echo "Done! Your Lambda package:"
ls -lh function.zip

# Cleanup (optional)
rm bootstrap

echo "Lambda package ready: function.zip"
