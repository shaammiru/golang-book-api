package main

import (
	"log"
	"net/http"
)

func main() {
	bookStorage := NewBookStorage()
	bookService := NewBookService(*bookStorage)
	bookHandler := NewBookHandler(*bookService)

	http.HandleFunc("/books", bookHandler.CreateBook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
