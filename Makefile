.PHONY: build run test cover coverage configure dep

all: fmt configure build cover run test

build:
	go build

run: build
	./gopherbadger -md="README.md"

test:
	go test -v

cover:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o=coverage.html

coverage: cover

configure:
	go mod download -x
	go mod verify
	go mod tidy -v

dep: configure

fmt:
	go fmt ./...
	go mod edit -fmt
