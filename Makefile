# Makefile for Go project

# Variables
GOPATH := $(shell go env GOPATH)
GOCMD := go
GOBUILD := $(GOCMD) build
GORUN := $(GOCMD) run
GOTEST := $(GOCMD) test
GOGET := $(GOCMD) get
GOMOD := $(GOCMD) mod

# Main targets
.PHONY: all build run run-app test clean deps help

all: deps build

# Install dependencies
deps:
	@echo "Installing dependencies..."
	$(GOMOD) download

# Build the project
build:
	@echo "Building the project..."
	$(GOBUILD) ./...

# Run the main application
run:
	@echo "Running main application..."
	$(GORUN) main.go

dev:
	@echo "Running main application in dev mode..."
	$(GORUN) main.go -debug

# Run the app in cmd/app directory
run-app:
	@echo "Running app in cmd/app directory..."
	$(GORUN) cmd/app/main.go

# Run tests
test:
	@echo "Running tests..."
	$(GOTEST) ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -f main
	rm -f cmd/app/app

# Help target to show available commands
help:
	@echo "Available targets:"
	@echo "  all      - Install dependencies and build the project"
	@echo "  deps     - Install dependencies"
	@echo "  build    - Build the project"
	@echo "  run      - Run the main application"
	@echo "  run-app  - Run the app in cmd/app directory"
	@echo "  test     - Run tests"
	@echo "  clean    - Clean build artifacts"
	@echo "  help     - Show this help message"