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

func setup() (*b.Book, *b.Book) {
	myLib := l.NewLibrary()
	bookA := b.NewBook(1, "A")
	myLib.AddBook(bookA, 5)
	bookB := b.NewBook(2, "B")
	myLib.AddBook(bookB, 3)
	manager = NewManager(myLib)
	user = u.NewUser("Jack")
	return bookA, bookB
}

func TestHandleBorrowing_managerShouldHandleBorrowing(t *testing.T) {
	_, bookB := setup()
	assert.Equal(t, manager.library.ViewBooks(),
		"\n List of books\n "+
			"Id: 1, Name: A, AvailableCopies: 5, TotalCopies: 5\n "+
			"Id: 2, Name: B, AvailableCopies: 3, TotalCopies: 3")
	err := manager.HandleBorrow(&user, bookB.GetId())
	assert.Nil(t, err)
	assert.Equal(t, len(user.GetBooks()), 1)
	assert.Equal(t, user.GetBooks()[0].GetId(), bookB.GetId())
	assert.Equal(t, user.GetBooks()[0].GetName(), bookB.GetName())
	assert.Equal(t, manager.library.ViewBooks(),
		"\n List of books\n "+
			"Id: 1, Name: A, AvailableCopies: 5, TotalCopies: 5\n "+
			"Id: 2, Name: B, AvailableCopies: 2, TotalCopies: 3")
}
