NAME=moment
BIN_DIR=bin
GO_BUILD=go build

PLATFORM_LIST=darwin-amd64 darwin-arm64 linux-amd64 linux-armv5 linux-armv6 linux-armv7 linux-armv8 \
	windows-amd64 windows-arm64

default: build

.PHONY: default

build: darwin-amd64 linux-amd64 windows-amd64 # Most used

build-all: $(PLATFORM_LIST)

darwin-amd64:
	GOARCH=amd64 GOOS=darwin $(GO_BUILD) -o $(BIN_DIR)/$(NAME)-$@

darwin-arm64:
	GOARCH=arm64 GOOS=darwin $(GO_BUILD) -o $(BIN_DIR)/$(NAME)-$@

linux-amd64:
	GOARCH=amd64 GOOS=linux $(GO_BUILD) -o ./$(BIN_DIR)/$(NAME)-$@

linux-armv5:
	GOARCH=arm GOOS=linux GOARM=5 $(GO_BUILD) -o $(BIN_DIR)/$(NAME)-$@

linux-armv6:
	GOARCH=arm GOOS=linux GOARM=6 $(GO_BUILD) -o $(BIN_DIR)/$(NAME)-$@

linux-armv7:
	GOARCH=arm GOOS=linux GOARM=7 $(GO_BUILD) -o $(BIN_DIR)/$(NAME)-$@

linux-armv8:
	GOARCH=arm64 GOOS=linux $(GO_BUILD) -o $(BIN_DIR)/$(NAME)-$@

windows-amd64:
	GOARCH=amd64 GOOS=windows $(GO_BUILD) -o $(BIN_DIR)/$(NAME)-$@.exe

windows-arm64:
	GOARCH=arm64 GOOS=windows $(GO_BUILD) -o $(BIN_DIR)/$(NAME)-$@.exe

test:
	go test ./... -race -coverprofile=coverage.txt -covermode=atomic -v

.PHONY: test

clean:
	rm $(BIN_DIR)/*

.PHONY: clean

help:
	@echo 'make clean: clean project'
	@echo 'make test: compile and test project'
	@echo 'make [build]: compile and build project'
	@echo 'make build-all: compile and build project for all platform'

.PHONY: help
