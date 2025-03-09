# NES Emulator Guidelines

## Commands
- Build: `make build` or `go build ./...`
- Run: `make run` or `go run main.go`
- Run app: `make run-app` or `go run cmd/app/main.go`
- Test all: `make test` or `go test ./...`
- Test single package: `go test ./pkg/memory`
- Install deps: `make deps` or `go mod download`
- Clean: `make clean`

## Code Style
- Use standard Go conventions and idioms
- Follow package structure: cmd/ (entrypoints), pkg/ (public), internal/ (private)
- Format: Use gofmt/goimports, tabs for indentation
- Naming: PascalCase for exported items, camelCase for unexported
- Imports: Group standard library, external, and project imports
- Error handling: Return errors with context using fmt.Errorf
- Comments: Document all exported items and packages
- Testing: Include unit tests for all functionality
- Memory maps are 64KB following NES architecture

This file serves as guidance for agents working in this repo.