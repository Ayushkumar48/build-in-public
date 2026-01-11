.PHONY: dev run build clean test test-coverage install-deps install-air fmt lint \
        swagger api-client api-gen help

# ----------------------------
# Development
# ----------------------------

# Run with hot reload using Air
dev:
	air

# Run without hot reload
run:
	go run cmd/server/main.go

# Build the binary
build:
	@echo "Building application..."
	go build -o bin/server cmd/server/main.go
	@echo "Build complete! Binary at bin/server"

# ----------------------------
# Cleanup
# ----------------------------

# Clean temporary files and binaries
clean:
	@echo "Cleaning..."
	rm -rf tmp/ bin/ build-errors.log docs/
	@echo "Clean complete!"

# ----------------------------
# Testing
# ----------------------------

# Run tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -v -cover ./...

# ----------------------------
# Dependencies
# ----------------------------

# Install dependencies
install-deps:
	@echo "Installing dependencies..."
	go mod download
	go mod tidy
	@echo "Dependencies installed!"

# Install Air for hot reloading
install-air:
	@echo "Installing Air..."
	go install github.com/air-verse/air@latest
	@echo "Air installed!"

# ----------------------------
# Code quality
# ----------------------------

# Format code
fmt:
	go fmt ./...

# Run linter
lint:
	golangci-lint run

# ----------------------------
# Swagger / OpenAPI
# ----------------------------

# Generate Swagger (OpenAPI) docs
swagger:
	@echo "Generating Swagger docs..."
	swag init \
		--dir cmd/server,internal \
		--generalInfo main.go \
		--output docs
	@echo "Swagger docs generated at docs/"

# ----------------------------
# Frontend API client
# ----------------------------

# Generate typed frontend API client (uses local swagger.json)
api-client:
	@echo "Generating frontend API client..."
	cd frontend && \
	bunx openapi-typescript-codegen \
		--input ../docs/swagger.json \
		--output src/lib/api \
		--client fetch
	@echo "Frontend API client generated!"

# Combined target (most important)
api-gen: swagger api-client

# ----------------------------
# Help
# ----------------------------

help:
	@echo "Available commands:"
	@echo "  make dev            - Run with hot reload (Air)"
	@echo "  make run            - Run without hot reload"
	@echo "  make build          - Build the application"
	@echo "  make clean          - Clean temporary files"
	@echo "  make test           - Run tests"
	@echo "  make test-coverage  - Run tests with coverage"
	@echo "  make install-deps   - Install Go dependencies"
	@echo "  make install-air    - Install Air for hot reloading"
	@echo "  make fmt            - Format code"
	@echo "  make lint           - Run linter"
	@echo "  make swagger        - Generate Swagger docs"
	@echo "  make api-client     - Generate frontend typed API client"
	@echo "  make api-gen        - Generate Swagger + frontend client"
	@echo "  make help           - Show this help message"
