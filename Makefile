NAME=moment
BUILD_DIR=build
GO_BUILD=go build

all: darwin-amd64 linux-amd64 windows-amd64
	@echo "build moment success!"

.PHONY: all

darwin-amd64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GO_BUILD) -o ./$(BUILD_DIR)/macos/$(NAME)

linux-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO_BUILD) -o ./$(BUILD_DIR)/linux/$(NAME)

windows-amd64:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GO_BUILD) -o ./$(BUILD_DIR)/windows/$(NAME).exe

test:
	go test ./app -v

.PHONY: test

clean:
	rm -rf $(BUILD_DIR)

.PHONY: clean