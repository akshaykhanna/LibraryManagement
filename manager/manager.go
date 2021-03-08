package manager

import (
	l "LibraryManagement/library"
	u "LibraryManagement/user"
	"errors"
)

const (
	MAX_BOOKS_ALLOWED_PER_PERSON = 2
)

type Manager struct {
	library l.Library
}

func NewManager(library l.Library) Manager {
	return Manager{library: library}
}

func (m *Manager) HandleBorrow(user *u.User, bookId int) error {
	if m.CanBookBeBorrowed(user, bookId) {
		borrowedBook := m.library.BorrowBook(bookId)
		user.AddBook(borrowedBook)
		return nil
	}
	return errors.New("book can't be borrowed")
}

func (m *Manager) CanBookBeBorrowed(user *u.User, bookId int) bool {
	return m.library.CanBookBeBorrowed(bookId) && canUserBorrowBook(user, bookId)
}
func canUserBorrowBook(user *u.User, bookId int) bool {
	return !user.IsHavingBook(bookId) && len(user.GetBooks()) < MAX_BOOKS_ALLOWED_PER_PERSON
}
