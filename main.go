package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var err error

func main() {

	// set database connection
	dsbn := "host=localhost user=postgres password= dbname=go-bookstore-db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsbn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// migrate the schema
	db.AutoMigrate(&Book{})
	Seed(db)

	// create a new router
	r := mux.NewRouter()

	// define routes
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		getBooksHandler(w, db)

	}).Methods("GET")

	// enable CORS middleware
	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			next.ServeHTTP(w, r)
		})
	}

	// attach cors middleware to router
	r.Use(corsMiddleware)

	// listen for requests in 8080
	log.Fatal(http.ListenAndServe(":8080", r))
}

// test api
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the bookstore"))
}

// get all the books from the database
func getBooksHandler(w http.ResponseWriter, db *gorm.DB) {
	var books []Book

	// make sure db is connected]
	if db == nil {
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	} else {
		db.Find(&books)
		json.NewEncoder(w).Encode(books)
	}
}
