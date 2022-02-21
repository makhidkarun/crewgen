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
	go build -o cmd/crewgen/crewgen  cmd/crewgen/main.go
	go build -o cmd/teamgen/teamgen  cmd/teamgen/main.go
.PHONY:build

test:	vet
	go test ./...
.PHONY:test

