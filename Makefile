.PHONY: build
build:
	go build -v ./cmd/concurency

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: run
run:
	go run -v ./cmd/concurency

.DEFAULT_GOAL := build