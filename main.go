package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	books = append(books, Book{ID: "1", Isbn: "987564", Title: "first book", Author: &Author{Firstname: "behrooz", Lastname: "razzaghi"}})
	books = append(books, Book{ID: "2", Isbn: "345645", Title: "second book", Author: &Author{Firstname: "ali", Lastname: "vali"}})

	r := mux.NewRouter()
	r.HandleFunc("/books", getbooks).Methods("GET")
	r.HandleFunc("/book/{id}", getbook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/book/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
