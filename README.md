# Book Manager Go API

A simple REST API for managing books in a bookstore, built with Go.

## Features

- Add new books to the bookstore
- Retrieve all books
- Get a specific book by ID
- Thread-safe operations using mutex
- JSON request/response format

## API Endpoints

### POST /books
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

### GET /books
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

### GET /books/{id}
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
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{"name":"The Go Programming Language","author":"Alan Donovan"}'
```

### Get all books:
```bash
curl http://localhost:8080/books
```

### Get a specific book:
```bash
curl http://localhost:8080/books/1
```

## Project Structure

- `main.go` - Main application file containing the REST API implementation
- `go.mod` - Go module file with dependencies
- `README.md` - This documentation file

The application includes sample books for testing purposes when it starts.
