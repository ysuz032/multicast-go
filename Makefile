.PHONY: up down logs build clean test-sender test-receiver fmt vet lint restart init status

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

build:
	@echo "Building the Docker images..."
	cd apps && docker-compose build

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