package library

type Library struct {
	books []Book
}

func NewLibrary() Library {
	return Library{[]Book{}}
}
func (l *Library) ViewBooks() string {
	if len(l.books) == 0 {
		return "Library is empty"
	}
	return ""
}
func (l *Library) AddBooks(book Book) *Library {
	if !l.isBookPresent(book.GetId()) {
		l.books = append(l.books, book)
	}
	return l
}

func (l *Library) GetBook(bookId int) *Book {
	for _, book := range l.books {
		if book.GetId() == bookId {
			return &book
		}
	}
	return nil
}

func (l *Library) BorrowBook(bookId int) *Book {
	if l.isBookPresent(bookId) {
		book := l.GetBook(bookId)
		book.SetAvailableCopies(book.GetAvailableCopies() - 1)
		return book
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
