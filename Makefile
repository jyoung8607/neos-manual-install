# Makefile for the NEOS Manual Install Go application

# The name of the output binary
BINARY_NAME=neos-manual-install.exe

# Default target executed when you run `make`
all: build

# Build the Go application for Windows
build:
	@echo "Building for Windows..."
	GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME) .

# Clean up the build artifacts
clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)

# A phony target to avoid conflicts with a file named 'clean'
.PHONY: all build clean