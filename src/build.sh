#!/bin/bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

env GOOS=linux   GOARCH=386   go build -o $SCRIPT_DIR/../bin/linux-386/sbom *.go
env GOOS=linux   GOARCH=amd64 go build -o $SCRIPT_DIR/../bin/linux-amd64/sbom *.go
env GOOS=windows GOARCH=386   go build -o $SCRIPT_DIR/../bin/windows-386/sbom.exe *.go
env GOOS=windows GOARCH=amd64 go build -o $SCRIPT_DIR/../bin/windows-amd64/sbom.exe *.go