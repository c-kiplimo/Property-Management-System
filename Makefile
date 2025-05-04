# Makefile

# Variables
GO=go
SWAG=swag
FLYWAY=flyway

# Set Go binary path
BIN_DIR=./bin

# Initialize the database
initdb:
	$(FLYWAY) migrate

# Generate Swagger Docs
swagger:
	$(SWAG) init

# Build the Go app
build:
	$(GO) build -o $(BIN_DIR)/main .

# Run the Go app
run:
	$(GO) run main.go

# Start the app with migration
migrate-and-run:
	$(MAKE) initdb && $(MAKE) run

.PHONY: initdb swagger build run migrate-and-run
