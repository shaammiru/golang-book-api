package main

import (
	"encoding/json"
	"net/http"
)

type BookHandler struct {
	bookService BookService
}

func NewBookHandler(bookService BookService) *BookHandler {
	return &BookHandler{bookService}
}

func (bh *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book BookSchema
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdBook, err := bh.bookService.CreateBook(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdBook)
}
