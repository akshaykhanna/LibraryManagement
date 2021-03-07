package library

import "LibraryManagement/book"

type Book struct {
	book.Book
	noOfCopies int
}
