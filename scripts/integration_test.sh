#!/bin/bash
set -e

[ $# -ne 3 ] && echo "Usage: $0 <Project ID> </path/to/credentials> <BOOKSHELF>" && exit 1
export PROJECT_ID=${1}
export GCLOUD_CRENTIAL_FILE_PATH=${2} 
export BOOKSHELF=${3}
go test ./... -v -coverprofile=coverage.out
go tool cover -func=coverage.out
echo "Run 'go tool cover -html=coverage.out' to see coverage in browser."
