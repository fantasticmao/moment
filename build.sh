#!/bin/bash

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./build/macos/moment .
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/linux/moment .
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./build/windows/moment.exe .