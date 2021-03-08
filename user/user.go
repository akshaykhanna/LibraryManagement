package user

import b "LibraryManagement/book"

type User struct {
	name  string
	books []*b.Book
}

func NewUser(name string) *User {
	return &User{name, []*b.Book{}}
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetBooks() []*b.Book {
	return u.books
}
func (u *User) AddBook(book *b.Book) *User {
	u.books = append(u.books, book)
	return u
}

func (u *User) RemoveBook(bookId int) *b.Book {
	var newBooks []*b.Book
	var returnedBook *b.Book
	for _, book := range u.books {
		if book.GetId() != bookId {
			newBooks = append(newBooks, book)
		} else {
			returnedBook = book
		}
	}
	u.books = newBooks
	return returnedBook
}

func (u *User) IsHavingBook(bookId int) bool {
	for _, book := range u.books {
		if book.GetId() == bookId {
			return true
		}
	}
	return false
}
