package user

import b "LibraryManagement/book"

type User struct {
	name  string
	books []b.Book
}

func NewUser(name string) User {
	return User{name, []b.Book{}}
}

func (u *User) BorrowBook(book b.Book) *User {
	u.books = append(u.books, book)
	return u
}
