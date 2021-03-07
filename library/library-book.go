package library

import "LibraryManagement/book"

type Book struct {
	*book.Book
	noOfCopies int
}

func NewBook(id int, name string, noOfCopies int) *Book {
	return &Book{book.NewBook(id, name), noOfCopies}
}
