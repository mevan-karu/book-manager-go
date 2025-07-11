# Book Manager Go API

A REST API for managing books in a bookstore, built with Go and PostgreSQL using GORM.

## Features

- Add new books to the bookstore
- Retrieve all books
- Get a specific book by ID
- PostgreSQL database integration with GORM
- Automatic database table creation
- Environment variable configuration
- JSON request/response format

## Database

This application uses PostgreSQL as the database backend with GORM for ORM functionality. The database connection and tables are automatically configured using environment variables.

### Database Schema

The application automatically creates the following table structure:

```sql
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    author VARCHAR NOT NULL
);
```

## Environment Variables

The following environment variables need to be configured:

- `DB_HOST` - PostgreSQL host (default: localhost)
- `DB_PORT` - PostgreSQL port (default: 5432)
- `DB_USER` - Database username (default: postgres)
- `DB_PASSWORD` - Database password (required)
- `DB_NAME` - Database name (default: bookstore)
- `DB_SSLMODE` - SSL mode (default: disable)
- `PORT` - Server port (default: 8080)

See `.env.example` for a template of required environment variables.

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

### Prerequisites

1. PostgreSQL database server
2. Go 1.21 or later

### Local Development

1. Initialize the Go module and download dependencies:
   ```bash
   go mod tidy
   ```

2. Set up your PostgreSQL database and create a database named `bookstore` (or your preferred name)

3. Set the required environment variables:
   ```bash
   export DB_HOST=localhost
   export DB_PORT=5432
   export DB_USER=postgres
   export DB_PASSWORD=your_password
   export DB_NAME=bookstore
   export DB_SSLMODE=disable
   export PORT=8080
   ```

   Or copy `.env.example` to `.env` and modify the values:
   ```bash
   cp .env.example .env
   # Edit .env with your database credentials
   ```

4. Run the application:
   ```bash
   go run main.go
   ```

5. The server will start on the configured port and automatically create the necessary database tables

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

- `main.go` - Main application file containing the REST API implementation with PostgreSQL integration
- `go.mod` - Go module file with dependencies (includes GORM and PostgreSQL driver)
- `go.sum` - Go module checksums
- `.env.example` - Example environment variables configuration
- `README.md` - This documentation file
- `openapi.yaml` - OpenAPI specification for the REST API

The application includes sample books that are automatically added to the database on first run if the database is empty.

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
