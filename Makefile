# Build the application
all: build

build:
		@echo "Building..."
		@go build -o main cmd/api/main.go

# Run the application
run-server:
		@go run cmd/api/main.go

.PHONY: all build run-server