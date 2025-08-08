.PHONY: build run test clean migrate

# Build the application
build:
	go build -o bin/tenant-management cmd/api/main.go

# Run the application
run:
	go run cmd/api/main.go

# Run tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# Clean build artifacts
clean:
	rm -rf bin/

# Install dependencies
deps:
	go mod download
	go mod tidy

# Run database migrations (placeholder)
migrate:
	@echo "Running database migrations..."
	# Add actual migration commands here

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run

# Generate docs
docs:
	@echo "Generating API documentation..."
	# Add swagger/doc generation commands here

# Docker build
docker-build:
	docker build -t tenant-management:latest .

# Docker run
docker-run:
	docker run -p 8080:8080 --env-file .env tenant-management:latest
