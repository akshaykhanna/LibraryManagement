package manager

import (
	b "LibraryManagement/book"
	l "LibraryManagement/library"
	u "LibraryManagement/user"
	"github.com/stretchr/testify/assert"
	"testing"
)

var manager Manager
var user u.User

func getMockBooks() (*b.Book, *b.Book) {
	bookA := b.NewBook(1, "A")
	bookB := b.NewBook(2, "B")
	return bookA, bookB
}

func setupMangerAndUser(library l.Library) {
	manager = NewManager(library)
	user = u.NewUser("Jack")
}

func setupBorrowData() (*b.Book, *b.Book) {
	bookA, bookB := getMockBooks()
	myLib := l.NewLibrary()
	myLib.AddBook(bookA, 5)
	myLib.AddBook(bookB, 3)
	setupMangerAndUser(myLib)
	return bookA, bookB
}

func TestHandleBorrowing_managerShouldHandleBorrowing(t *testing.T) {
	_, bookB := setupBorrowData()
	assert.Equal(t, manager.library.ViewBooks(),
		"\n List of books\n "+
			"Id: 1, Name: A, AvailableCopies: 5\n "+
			"Id: 2, Name: B, AvailableCopies: 3")
	err := manager.HandleBorrow(&user, bookB.GetId())
	assert.Nil(t, err)
	assert.Equal(t, len(user.GetBooks()), 1)
	assert.Equal(t, user.GetBooks()[0].GetId(), bookB.GetId())
	assert.Equal(t, user.GetBooks()[0].GetName(), bookB.GetName())
	assert.Equal(t, manager.library.ViewBooks(),
		"\n List of books\n "+
			"Id: 1, Name: A, AvailableCopies: 5\n "+
			"Id: 2, Name: B, AvailableCopies: 2")
}
