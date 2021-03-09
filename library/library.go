package library

import (
	b "LibraryManagement/book"
)

type Library struct {
	books libBooks
}

func NewLibrary() *Library {
	return &Library{[]*libBook{}}
}

func (l *Library) ViewBooks() string {
	return "Library: " + l.books.getBooksString()
}

func (l *Library) AddBook(book *b.Book, copies int) *Library {
	if !l.books.isBookPresent(book.GetId()) {
		l.books.addBook(NewLibBook(book.GetId(), book.GetName(), copies))
	}
	return l
}

func (l *Library) RemoveBook(bookId int) *b.Book {
	return l.books.removeBook(bookId).ToBook()
}

func (l *Library) GetBook(bookId int) *b.Book {
	return l.books.getLibBook(bookId).ToBook()
}

func (l *Library) CanBookBeBorrowed(bookId int) bool {
	return l.books.isBookPresent(bookId) && l.books.getLibBook(bookId).IsBookAvailable()
}

func (l *Library) BorrowBook(bookId int) *b.Book {
	if l.CanBookBeBorrowed(bookId) {
		book := l.books.getLibBook(bookId)
		if book.availableCopies == 1 {
			l.RemoveBook(bookId)
		} else {
			book.SetAvailableCopies(book.GetAvailableCopies() - 1)
		}
		return b.NewBook(book.GetId(), book.GetName())
	}
	return nil
}

func (l *Library) ReturnBook(book *b.Book) *Library {
	if l.books.isBookPresent(book.GetId()) {
		thatBook := l.books.getLibBook(book.GetId())
		thatBook.SetAvailableCopies(thatBook.GetAvailableCopies() + 1)
		return l
	} else {
		l.AddBook(book, 1)
	}
	return nil
}
