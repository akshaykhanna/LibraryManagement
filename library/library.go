package library

import b "LibraryManagement/book"

type Library struct {
	books []*libBook
}

func NewLibrary() Library {
	return Library{[]*libBook{}}
}
func (l *Library) ViewBooks() string {
	if len(l.books) == 0 {
		return "Library is empty"
	}
	return ""
}
func (l *Library) AddBook(book *b.Book, noOfCopies int) *Library {
	if !l.isBookPresent(book.GetId()) {
		l.books = append(l.books, NewLibBook(book.GetId(), book.GetName(), noOfCopies))
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

func (l *Library) CanBookBeBorrowed(bookId int) bool {
	return l.isBookPresent(bookId) && l.getLibBook(bookId).IsBookAvailable()
}
