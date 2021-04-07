package user

import b "LibraryManagement/book"

type User struct {
	name  string
	books b.Books
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

func (u *User) BorrowBook(book *b.Book) *User {
	u.books.AddBook(book)
	return u
}

func (u *User) ReturnBook(bookId int) *b.Book {
	return u.books.RemoveBook(bookId)
}

func (u *User) IsHavingBook(bookId int) bool {
	if u.books.GetBook(bookId) != nil {
		return true
	}
	return false
}

func (u *User) ViewBooks() string {
	return "User: " + u.books.GetBooksString()
}
