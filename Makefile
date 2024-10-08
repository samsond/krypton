BINARY_NAME=kptn
BUILD_DIR=./cmd/krypton
VERSION=$(shell git describe --tags --always --dirty)

.PHONY: build test tidy

# Build the kptn binary
build:
	@echo "Building $(BINARY_NAME)..."
	go build -ldflags "-X main.version=$(VERSION)" -o $(BINARY_NAME) $(BUILD_DIR)

# Run tests
test:
	@echo "Running tests..."
	go test ./...

# Tidy up dependencies
tidy:
	go mod tidy