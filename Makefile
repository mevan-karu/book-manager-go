.PHONY: build run test clean deps db-up db-down

# Build the application
build:
	go build -o book-manager .

# Run the application
run:
	go run main.go

# Install dependencies
deps:
	go mod tidy
	go mod download

# Start PostgreSQL with docker-compose
db-up:
	docker-compose up -d postgres

# Stop PostgreSQL
db-down:
	docker-compose down

# Run with local database (starts db and runs app)
dev: db-up
	@echo "Waiting for database to be ready..."
	@sleep 5
	@export DB_HOST=localhost && \
	export DB_PORT=5432 && \
	export DB_USER=postgres && \
	export DB_PASSWORD=password && \
	export DB_NAME=bookstore && \
	export DB_SSLMODE=disable && \
	export PORT=8080 && \
	go run main.go

# Test the API endpoints
test-api:
	@echo "Testing API endpoints..."
	@echo "1. Creating a book:"
	@curl -X POST http://localhost:8080/api/books \
		-H "Content-Type: application/json" \
		-d '{"name":"Test Book","author":"Test Author"}' && echo
	@echo "2. Getting all books:"
	@curl http://localhost:8080/api/books && echo
	@echo "3. Getting book by ID:"
	@curl http://localhost:8080/api/books/1 && echo

# Clean build artifacts
clean:
	rm -f book-manager

# Show help
help:
	@echo "Available targets:"
	@echo "  build    - Build the application"
	@echo "  run      - Run the application"
	@echo "  deps     - Install dependencies"
	@echo "  db-up    - Start PostgreSQL with docker-compose"
	@echo "  db-down  - Stop PostgreSQL"
	@echo "  dev      - Start database and run application"
	@echo "  test-api - Test the API endpoints"
	@echo "  clean    - Clean build artifacts"
	@echo "  help     - Show this help message"
