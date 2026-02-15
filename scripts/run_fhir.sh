#!/bin/bash

# Build the CLI tool
echo "Building zh-fhir CLI..."
go build -o zh-fhir ./cmd/zh-fhir

# Start the terminology server in the background
echo "Starting Terminology Server on port 8080..."
./zh-fhir -term-server -port 8080 &

# Keep the script running
wait
