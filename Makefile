# Makefile for sec CLI

BINARY_NAME=cloak
VERSION=0.1.0

# Build for local system
build:
	go build -o $(BINARY_NAME)
	@echo "✅ Built $(BINARY_NAME)"

# Build for all major platforms
build-all:
	@mkdir -p dist
	GOOS=linux   GOARCH=amd64 go build -o dist/$(BINARY_NAME)-linux-amd64
	GOOS=darwin  GOARCH=amd64 go build -o dist/$(BINARY_NAME)-darwin-amd64
	GOOS=darwin  GOARCH=arm64 go build -o dist/$(BINARY_NAME)-darwin-arm64
	GOOS=windows GOARCH=amd64 go build -o dist/$(BINARY_NAME)-windows-amd64.exe
	@echo "✅ Built binaries for Linux, Mac (Intel+M1), Windows"

# Install locally (Mac/Linux)
install:
	sudo mv $(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)
	@echo "🚀 Installed $(BINARY_NAME) to /usr/local/bin"

# Clean build artifacts
clean:
	rm -f $(BINARY_NAME)
	rm -rf dist
	@echo "🧹 Cleaned build artifacts"

# Run the app directly
run:
	go run main.go

# Test everything
test:
	make build
	go test ./...

# Help
help:
	@echo "Available commands:"
	@echo "  make build       Build binary for local system"
	@echo "  make build-all   Build binaries for Linux, Mac (Intel+M1), Windows"
	@echo "  make install     Install binary to /usr/local/bin"
	@echo "  make clean       Clean binaries"
	@echo "  make run         Run locally with go run"
	@echo "  make test        Build and run all tests"
