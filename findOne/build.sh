#!/bin/bash
echo "Building the binary"
GOOS=linux GOARCH=amd64 go build -o main main.go

echo "Create Zip file"
zip deployment.zip main

echo "Cleaning up"
rm main