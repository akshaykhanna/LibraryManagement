package library

import (
	b "LibraryManagement/book"
	"fmt"
)

type Library struct {
	books []*libBook
}

func NewLibrary() Library {
	return Library{[]*libBook{}}
}

func GetBooksString(l *Library) string {
	booksString := fmt.Sprintf("\n %s", "List of books")
	for _, book := range l.books {
		booksString = booksString + fmt.Sprintf("\n %s", book.View())
	}
	return booksString
}

func (l *Library) ViewBooks() string {
	if len(l.books) == 0 {
		return "Library is empty"
	}
	return GetBooksString(l)
}

func (l *Library) AddBook(book *b.Book, copies int) *Library {
	if !l.isBookPresent(book.GetId()) {
		l.books = append(l.books, NewLibBook(book.GetId(), book.GetName(), copies))
	}
	return l
}

func (l *Library) GetBook(bookId int) *b.Book {
	for _, book := range l.books {
		if book.GetId() == bookId {
			return b.NewBook(book.GetId(), book.GetName())
		}
	}
	return nil
}

func (l *Library) BorrowBook(bookId int) *b.Book {
	if l.CanBookBeBorrowed(bookId) {
		book := l.getLibBook(bookId)
		book.SetAvailableCopies(book.GetAvailableCopies() - 1)
		return b.NewBook(book.GetId(), book.GetName())
	}
	return nil
}

func (l *Library) CanBookBeBorrowed(bookId int) bool {
	return l.isBookPresent(bookId) && l.getLibBook(bookId).IsBookAvailable()
}

func (l *Library) getLibBook(bookId int) *libBook {
	for _, book := range l.books {
		if book.GetId() == bookId {
			return book
		}
	}
	return nil
}

func (l *Library) isBookPresent(bookId int) bool {
	for _, book := range l.books {
		if book.GetId() == bookId {
			return true
		}
	}
	return false
}
