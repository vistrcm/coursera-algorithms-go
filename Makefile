.DEFAULT_GOAL := test_all

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: test_all
test_all: test benchmark

.PHONY: test
test:
	go test ./...

.PHONY: benchmark
benchmark:
	go test ./... -bench=.
