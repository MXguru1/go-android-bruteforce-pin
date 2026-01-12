#!/bin/bash
echo "Building go-android-bruteforce-pin..."

export CGO_ENABLED=1
go build -ldflags="-s -w" -o go-android-bruteforce-pin cmd/sony-xperia-z1-compact/main.go

echo "Done."
