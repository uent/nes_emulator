# Running the Project Locally

This document explains how to run this Go project locally.

## Prerequisites

- Go installed (latest stable version recommended, version 1.22 or compatible with the go.mod file)
- Git (to clone the repository)

## Steps to Run

### 1. Install Dependencies

```bash
go mod download
```

Or use the devfile command:
```bash
# If using the devfile with a compatible tool
devfile exec install
```

### 2. Build the Project

```bash
go build ./...
```

Or use the devfile command:
```bash
# If using the devfile with a compatible tool
devfile exec build
```

### 3. Run the Project

This project has two entry points:

#### Main application in root directory

```bash
# Build and run the main application
go run main.go
```
This will output: "Hello, Go!"

#### Application in cmd/app directory

```bash
# Build and run the app in cmd/app
go run cmd/app/main.go
```
This will output: "Hello from the app!"

### 4. Running Tests

```bash
go test ./...
```

Or use the devfile command:
```bash
# If using the devfile with a compatible tool
devfile exec test
```

## Development with Devfile

This project includes a devfile.yaml which can be used with compatible tools to standardize the development environment. The devfile specifies commands for installing dependencies, building, and testing.