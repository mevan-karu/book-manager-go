package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

// Book represents a book in the bookstore
type Book struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

// BookStore manages the collection of books
type BookStore struct {
	books  map[int]Book
	nextID int
	mutex  sync.RWMutex
}

// NewBookStore creates a new BookStore instance
func NewBookStore() *BookStore {
	return &BookStore{
		books:  make(map[int]Book),
		nextID: 1,
	}
}

// AddBook adds a new book to the store
func (bs *BookStore) AddBook(name, author string) Book {
	bs.mutex.Lock()
	defer bs.mutex.Unlock()

	book := Book{
		ID:     bs.nextID,
		Name:   name,
		Author: author,
	}
	bs.books[bs.nextID] = book
	bs.nextID++
	return book
}

// GetBooks returns all books in the store
func (bs *BookStore) GetBooks() []Book {
	bs.mutex.RLock()
	defer bs.mutex.RUnlock()

	books := make([]Book, 0, len(bs.books))
	for _, book := range bs.books {
		books = append(books, book)
	}
	return books
}

// GetBookByID returns a book by its ID
func (bs *BookStore) GetBookByID(id int) (Book, bool) {
	bs.mutex.RLock()
	defer bs.mutex.RUnlock()

	book, exists := bs.books[id]
	return book, exists
}

var bookStore = NewBookStore()

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

	book := bookStore.AddBook(req.Name, req.Author)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

// getBooksHandler handles GET /books
func getBooksHandler(w http.ResponseWriter, r *http.Request) {
	books := bookStore.GetBooks()

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

	book, exists := bookStore.GetBookByID(id)
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func main() {
	r := mux.NewRouter()

	// API routes
	r.HandleFunc("/books", createBookHandler).Methods("POST")
	r.HandleFunc("/books", getBooksHandler).Methods("GET")
	r.HandleFunc("/books/{id}", getBookByIDHandler).Methods("GET")

	// Add some sample books for testing
	bookStore.AddBook("The Go Programming Language", "Alan Donovan")
	bookStore.AddBook("Clean Code", "Robert Martin")
	bookStore.AddBook("The Pragmatic Programmer", "David Thomas")

	port := ":8080"
	fmt.Printf("Server starting on port %s\n", port)
	fmt.Println("Available endpoints:")
	fmt.Println("  POST /books - Add a new book")
	fmt.Println("  GET /books - Get all books")
	fmt.Println("  GET /books/{id} - Get book by ID")
	
	log.Fatal(http.ListenAndServe(port, r))
}
