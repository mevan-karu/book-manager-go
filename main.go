package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Book represents a book in the bookstore
type Book struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Name   string `json:"name" gorm:"not null"`
	Author string `json:"author" gorm:"not null"`
}

// Database instance
var db *gorm.DB

// initDatabase initializes the database connection and creates tables
func initDatabase() {
	var err error

	// Get database configuration from environment variables
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "")
	dbname := getEnv("DB_NAME", "bookstore")
	sslmode := getEnv("DB_SSLMODE", "disable")

	// Construct connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode)

	// Open database connection
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate the schema
	err = db.AutoMigrate(&Book{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database connected and tables created successfully")
}

// getEnv gets an environment variable with a fallback default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// BookService provides database operations for books
type BookService struct{}

// AddBook adds a new book to the database
func (bs *BookService) AddBook(name, author string) (*Book, error) {
	book := Book{
		Name:   name,
		Author: author,
	}

	result := db.Create(&book)
	if result.Error != nil {
		return nil, result.Error
	}

	return &book, nil
}

// GetBooks returns all books from the database
func (bs *BookService) GetBooks() ([]Book, error) {
	var books []Book
	result := db.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

// GetBookByID returns a book by its ID from the database
func (bs *BookService) GetBookByID(id uint) (*Book, error) {
	var book Book
	result := db.First(&book, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

var bookService = &BookService{}

// createBookHandler handles POST /books
func createBookHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name   string `json:"name"`
		Author string `json:"author"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Name == "" || req.Author == "" {
		http.Error(w, "Name and author are required", http.StatusBadRequest)
		return
	}

	book, err := bookService.AddBook(req.Name, req.Author)
	if err != nil {
		http.Error(w, "Failed to create book: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

// getBooksHandler handles GET /books
func getBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := bookService.GetBooks()
	if err != nil {
		http.Error(w, "Failed to get books: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// getBookByIDHandler handles GET /books/{id}
func getBookByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	book, err := bookService.GetBookByID(uint(id))
	if err != nil {
		if err.Error() == "record not found" {
			http.Error(w, "Book not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to get book: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func main() {
	// Initialize database
	initDatabase()

	r := mux.NewRouter()
	apiRouter := r.PathPrefix("/api").Subrouter()

	// API routes
	apiRouter.HandleFunc("/books", createBookHandler).Methods("POST")
	apiRouter.HandleFunc("/books", getBooksHandler).Methods("GET")
	apiRouter.HandleFunc("/books/{id}", getBookByIDHandler).Methods("GET")

	// Add some sample books for testing (only if database is empty)
	var count int64
	db.Model(&Book{}).Count(&count)
	if count == 0 {
		sampleBooks := []Book{
			{Name: "The Go Programming Language", Author: "Alan Donovan"},
			{Name: "Clean Code", Author: "Robert Martin"},
			{Name: "The Pragmatic Programmer", Author: "David Thomas"},
		}
		for _, book := range sampleBooks {
			bookService.AddBook(book.Name, book.Author)
		}
		log.Println("Sample books added to database")
	}

	port := getEnv("PORT", "8080")
	if port[0] != ':' {
		port = ":" + port
	}

	fmt.Printf("Server starting on port %s\n", port)
	fmt.Println("Available endpoints:")
	fmt.Println("  POST /api/books - Add a new book")
	fmt.Println("  GET /api/books - Get all books")
	fmt.Println("  GET /api/books/{id} - Get book by ID")
	fmt.Printf("Database connected to: %s\n", getEnv("DB_HOST", "localhost"))

	log.Fatal(http.ListenAndServe(port, r))
}
