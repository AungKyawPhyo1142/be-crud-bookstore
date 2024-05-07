package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/google/uuid"
)

type Book struct {
	ID		string `json:"id"`
	Title	string `json:"title"`
	Author	string `json:"author"`
}

// replace this with actual database
// ! this is only for testing http requests at the moment
var books = map[string]Book{
	// "1": {ID: generateID(), Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"},
	// "2": {ID: generateID(), Title: "To Kill a Mockingbird", Author: "Harper Lee"},
	// "3": {ID: generateID(), Title: "1984", Author: "George Orwell"},
}

func generateID() string {
	return uuid.New().String()
}

// get all the books
func getAllBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(books)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// get a book by id
func getBook(w http.ResponseWriter, r *http.Request) {
	// get the id from the url query string
	id := r.URL.Path[len("/books/"):]
	
	book, ok := books[id]

	if !ok {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(book)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// create a book
func createBook(w http.ResponseWriter, r *http.Request) {
	var book Book

	fmt.Println(r.Body)

	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book.ID = generateID()
	books[book.ID] = book

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(book)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/books", getAllBooks)
	http.HandleFunc("/books/", getBook)
	http.HandleFunc("/books/create", createBook)
	http.ListenAndServe(":8000", nil)
}