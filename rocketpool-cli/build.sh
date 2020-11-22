#!/bin/bash

GOOS=linux GOARCH=arm64 go build -o rocketpool-cli-linux-arm64 rocketpool-cli.go
GOOS=linux GOARCH=amd64 go build -o rocketpool-cli-linux-amd64 rocketpool-cli.go
GOOS=darwin GOARCH=amd64 go build -o rocketpool-cli-darwin-amd64 rocketpool-cli.go
GOOS=windows GOARCH=amd64 go build -o rocketpool-cli-windows-amd64.exe rocketpool-cli.go
