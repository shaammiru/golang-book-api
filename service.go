package main

type BookService struct {
	bookStorage BookStorage
}

func NewBookService(bookStorage BookStorage) *BookService {
	return &BookService{bookStorage}
}

func (bs *BookService) CreateBook(b *BookSchema) (*BookSchema, error) {
	book, err := bs.bookStorage.Create(b)
	if err != nil {
		return nil, err
	}
	return book, nil
}
