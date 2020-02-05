.DEFAULT_GOAL := build

install: # Install dependencies
	export GO111MODULE=on 
	go mod download
	go mod vendor

build:
	go build

test:
	go test ./...