APP_NAME=secrethor-cli
VERSION=$(shell git describe --tags --always)
BUILD_DIR=dist

all: clean build

build:
	mkdir -p $(BUILD_DIR)
	GOOS=linux   GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)-linux   ./main.go
	GOOS=darwin  GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)-darwin  ./main.go
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME).exe     ./main.go

clean:
	rm -rf $(BUILD_DIR)

package:
	cd $(BUILD_DIR) && \
	tar -czf $(APP_NAME)-linux.tar.gz $(APP_NAME)-linux && \
	tar -czf $(APP_NAME)-darwin.tar.gz $(APP_NAME)-darwin && \
	zip $(APP_NAME)-windows.zip $(APP_NAME).exe