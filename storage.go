package main

type BookSchema struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

var books = []*BookSchema{}

type BookStorage struct{}

func NewBookStorage() *BookStorage {
	return &BookStorage{}
}

func (s *BookStorage) Create(b *BookSchema) error {
	books = append(books, b)
	return nil
}
