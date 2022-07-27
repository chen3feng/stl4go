#!/bin/bash

set -e

go build

go test ./... -coverprofile=coverage.out

golint

go tool cover -html=coverage.out
