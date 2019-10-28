#!/bin/bash

# Formatting
go fmt $(go list ./... | grep -v /vendor/)

# Linters
golangci-lint run

# Tests
go test $(go list ./... | grep -v /vendor/)