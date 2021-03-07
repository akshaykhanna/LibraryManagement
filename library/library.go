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
	if !l.isBookPresent(book.Id()) {
		l.books = append(l.books, book)
	}
	return l
}

func (l *Library) isBookPresent(id int) bool {
	for _, book := range l.books {
		if book.Id() == id {
			return true
		}
	}
	return false
}
