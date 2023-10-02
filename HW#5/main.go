package main

import (
	"bookstore/book"
	"bookstore/database"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

var books []book.Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	books = database.GetBooks()
	w.Header().Set("Content-Type", "application/json")
	base := r.URL.Query().Get("filter")
	if base == "asc" {
		asc := database.FilterByPriceAsc()
		json.NewEncoder(w).Encode(asc)
	} else {
		json.NewEncoder(w).Encode(books)
	}

	fmt.Println("GetBooks Endpoint")

}


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books", getBooks).Methods("GET")
    fmt.Println("server is started")
	http.ListenAndServe(":3000", r)
}
