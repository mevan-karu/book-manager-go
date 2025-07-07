# Book Manager Go API

A simple REST API for managing books in a bookstore, built with Go.

## Features

- Add new books to the bookstore
- Retrieve all books
- Get a specific book by ID
- Thread-safe operations using mutex
- JSON request/response format

## API Endpoints

### POST /api/books
Add a new book to the bookstore.

**Request Body:**
```json
{
  "name": "Book Name",
  "author": "Author Name"
}
```

**Response:**
```json
{
  "id": 1,
  "name": "Book Name",
  "author": "Author Name"
}
```

### GET /api/books
Get all books from the bookstore.

**Response:**
```json
[
  {
    "id": 1,
    "name": "Book Name",
    "author": "Author Name"
  }
]
```

### GET /api/books/{id}
Get a specific book by ID.

**Response:**
```json
{
  "id": 1,
  "name": "Book Name",
  "author": "Author Name"
}
```

## Setup and Running

1. Initialize the Go module and download dependencies:
   ```bash
   go mod tidy
   ```

2. Run the application:
   ```bash
   go run main.go
   ```

3. The server will start on port 8080

## Sample Usage

### Add a book:
```bash
curl -X POST http://localhost:8080/api/books \
  -H "Content-Type: application/json" \
  -d '{"name":"The Go Programming Language","author":"Alan Donovan"}'
```

### Get all books:
```bash
curl http://localhost:8080/api/books
```

### Get a specific book:
```bash
curl http://localhost:8080/api/books/1
```

## Project Structure

- `main.go` - Main application file containing the REST API implementation
- `go.mod` - Go module file with dependencies
- `README.md` - This documentation file

The application includes sample books for testing purposes when it starts.

## Deploying to Choreo

This project is configured to be deployed as a service component in Choreo. The following files support the Choreo deployment:

- `.choreo/component.yaml` - Defines the service configuration for Choreo
- `openapi.yaml` - OpenAPI specification for the REST API

To deploy to Choreo:

1. Ensure your code is pushed to a GitHub repository
2. In Choreo, create a new service component 
3. Select the GitHub repository and branch
4. Choose Go as the buildpack
5. Complete the deployment configuration
6. Deploy the service

After deployment, Choreo will provide you with endpoints to access your API through their managed gateway.
