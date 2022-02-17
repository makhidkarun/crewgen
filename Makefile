.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint:	fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

build: vet
	go build -o cmd/web/crewgen  cmd/web/main.go
.PHONY:build

test:	vet
	go test ./...
.PHONY:test

