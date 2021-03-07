package library

import "LibraryManagement/book"

type Book struct {
	*book.Book
	totalCopies     int
	availableCopies int
}

func NewBook(id int, name string, totalCopies int) *Book {
	return &Book{book.NewBook(id, name), totalCopies, totalCopies}
}

func (b *Book) SetAvailableCopies(count int) {
	b.availableCopies = count
}

func (b *Book) GetAvailableCopies() int {
	return b.availableCopies
}
