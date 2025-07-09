# Makefile for the NEOS Manual Install Go application

# The name of the output binary
BINARY_NAME_WINDOWS=c2-neos-alt-fix-install.exe
BINARY_NAME_MACOS=c2-neos-alt-fix-install-darwin

# Default target executed when you run `make`
all: build-windows build-macos

# Build the Go application for Windows
build-windows:
	@echo "Building for Windows..."
	GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME_WINDOWS) ./cmd/installer

# Build the Go application for macOS
build-macos:
	@echo "Building for macOS..."
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY_NAME_MACOS) ./cmd/installer

# Clean up the build artifacts
clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME_WINDOWS) $(BINARY_NAME_MACOS)

# A phony target to avoid conflicts with a file named 'clean'
.PHONY: all build clean