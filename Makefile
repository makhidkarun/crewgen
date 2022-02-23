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
	cp -Rp pkg/datamine/data bin
	cp -Rp cmd/crewgen/web bin
	go build -o bin/crewgen  cmd/crewgen/main.go
	go build -o bin/teamgen  cmd/teamgen/main.go
.PHONY:build

test:	vet
	go test ./...
.PHONY:test

