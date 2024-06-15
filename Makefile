.PHONY: build build-sender build-receiver install install-sender install-receiver up down logs clean test-sender test-receiver fmt vet lint restart init status

# Linux install command
build: build-sender build-receiver

build-sender:
	@echo "Build sender application"
	cd apps/sender/src && go build -o sender cmd/main.go

build-receiver:
	@echo "Build receiver application"
	cd apps/receiver/src && go build -o receiver cmd/main.go

install: install-sender install-receiver

install-sender:
	@echo "Install sender application to /usr/local/bin"
	mv apps/sender/src/sender /usr/local/bin/sender

install-receiver:
	@echo "Install receiver application to /usr/local/bin"
	mv apps/receiver/src/receiver /usr/local/bin/receiver

# Docker Compose commands
up:
	@echo "Starting the application with Docker Compose..."
	cd apps && docker-compose up -d --build

down:
	@echo "Stopping the application..."
	cd apps && docker-compose down

logs:
	@echo "Displaying logs..."
	cd apps && docker-compose logs --follow --timestamps

# Clean up dangling images and containers
clean:
	@echo "Cleaning up dangling images and containers..."
	docker system prune -f

# Run Go tests for sender
test-sender:
	@echo "Running tests for sender..."
	cd apps/sender/src && docker-compose run --rm sender go test ./...

# Run Go tests for receiver
test-receiver:
	@echo "Running tests for receiver..."
	cd apps/receiver/src && docker-compose run --rm receiver go test ./...

# Format Go code
fmt:
	@echo "Formatting Go code..."
	cd apps/sender/src && docker-compose run --rm sender go fmt ./...
	cd apps/receiver/src && docker-compose run --rm receiver go fmt ./...

# Run Go vet
vet:
	@echo "Running Go vet..."
	cd apps/sender/src && docker-compose run --rm sender go vet ./...
	cd apps/receiver/src && docker-compose run --rm receiver go vet ./...

# Lint Go code
lint:
	@echo "Linting Go code..."
	cd apps/sender/src && docker-compose run --rm sender golangci-lint run
	cd apps/receiver/src && docker-compose run --rm receiver golangci-lint run

# Rebuild and start the application
restart: down up
	@echo "Restarted the application..."

# Initialize the environment (example: pulling images)
init:
	@echo "Initializing the environment..."
	cd apps && docker-compose pull

# Display the status of services
status:
	@echo "Displaying the status of services..."
	cd apps && docker-compose ps -a