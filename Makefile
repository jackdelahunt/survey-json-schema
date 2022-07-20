.DEFAULT_GOAL := binary
SHELL = bash

 
BUILDFLAGS :=

ifdef DEBUG
BUILDFLAGS := -gcflags "all=-N -l" $(BUILDFLAGS)
endif

# The details of the application:
binary:=askjschema

 
# Requires golangci-lint to be installed @ $(go env GOPATH)/bin/golangci-lint
# https://golangci-lint.run/usage/install/
lint: lint
	golangci-lint run cmd/... pkg/... internal/...

# Build binaries
# NOTE it may be necessary to use CGO_ENABLED=0 for backwards compatibility with centos7 if not using centos7
binary: ## Compile the rhoas binary into the local project directory
	go build -o ${binary} ./cmd/askjschema
.PHONY: binary

install: ## Compile and install rhoas and add it to the PAth 
	go install ./cmd/askjschema
.PHONY: install

test: ## Run unit tests
	go test ./pkg/...
.PHONY: test

format: ## Clean up code and dependencies
	@go mod tidy

	@gofmt -w `find . -type f -name '*.go'`
.PHONY: format

# Check http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
.PHONY: help

