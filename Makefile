.DEFAULT_GOAL := all

.PHONY: all
all: fmt lint test

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: test
test:
	go test ./... -bench=.

.PHONY: test_only
test_only:
	go test ./...
