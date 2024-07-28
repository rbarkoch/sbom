#!/bin/bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

env GOOS=linux   GOARCH=386   go build -o $SCRIPT_DIR/../bin/sbom-linux-386 *.go
env GOOS=linux   GOARCH=amd64 go build -o $SCRIPT_DIR/../bin/sbom-linux-amd64 *.go
env GOOS=windows GOARCH=386   go build -o $SCRIPT_DIR/../bin/sbom-windows-386.exe *.go
env GOOS=windows GOARCH=amd64 go build -o $SCRIPT_DIR/../bin/sbom-windows-amd64.exe *.go