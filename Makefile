.DEFAULT_GOAL := test

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: test
test:
	go test ./... -bench=.

.PHONY: test_only
test_only:
	go test ./...
