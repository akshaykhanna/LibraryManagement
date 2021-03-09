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
func getMockLib(noOfCopiesForBookA int, noOfCopiesForBookB int) *l.Library {
	bookA, bookB := getMockBooks()
	myLib := l.NewLibrary()
	if noOfCopiesForBookA > 0 {
		myLib.AddBook(bookA, noOfCopiesForBookA)
	}
	if noOfCopiesForBookB > 0 {
		myLib.AddBook(bookB, noOfCopiesForBookB)
	}
	return myLib
}

func setupManger(library *l.Library) {
	manager = NewManager(library)
}

func setupUser(name string) {
	user = *u.NewUser(name)
}

func setupBorrowData() {
	setupManger(getMockLib(5, 3))
	setupUser("Jack")
}

func setupReturnData() {
	setupManger(getMockLib(5, 0))
	setupUser("Jack")

}

func TestHandleBorrowing_managerShouldHandleBorrowing(t *testing.T) {
	setupBorrowData()
	_, bookB := getMockBooks()
	assert.Equal(t,
		"Library: List of books\n "+
			"Id: 1, Name: A, AvailableCopies: 5\n "+
			"Id: 2, Name: B, AvailableCopies: 3", manager.library.ViewBooks())
	err := manager.HandleBorrow(&user, bookB.GetId())
	assert.Nil(t, err)
	assert.Equal(t, len(user.GetBooks()), 1)
	assert.Equal(t, user.GetBooks()[0].GetId(), bookB.GetId())
	assert.Equal(t, user.GetBooks()[0].GetName(), bookB.GetName())
	assert.Equal(t,
		"Library: List of books\n "+
			"Id: 1, Name: A, AvailableCopies: 5\n "+
			"Id: 2, Name: B, AvailableCopies: 2", manager.library.ViewBooks())
}

func TestHandleReturn_managerShouldHandleReturnOfBookWhichNotPresentInLibrary(t *testing.T) {
	setupReturnData()
	_, bookB := getMockBooks()
	user.BorrowBook(bookB)
	assert.Equal(t, manager.library.ViewBooks(),
		"Library: List of books\n "+
			"Id: 1, Name: A, AvailableCopies: 5")
	err := manager.HandleReturn(&user, bookB.GetId())
	assert.Nil(t, err)
	assert.Equal(t, len(user.GetBooks()), 0)
	assert.Equal(t, manager.library.ViewBooks(),
		"Library: List of books\n "+
			"Id: 1, Name: A, AvailableCopies: 5\n "+
			"Id: 2, Name: B, AvailableCopies: 1")
}

func TestHandleReturn_managerShouldHandleReturnOfBookWhichIsAlreadyPresentInLibrary(t *testing.T) {
	setupReturnData()
	bookA, _ := getMockBooks()
	user.BorrowBook(bookA)
	assert.Equal(t, manager.library.ViewBooks(),
		"Library: List of books\n "+
			"Id: 1, Name: A, AvailableCopies: 5")
	err := manager.HandleReturn(&user, bookA.GetId())
	assert.Nil(t, err)
	assert.Equal(t, len(user.GetBooks()), 0)
	assert.Equal(t, manager.library.ViewBooks(),
		"Library: List of books\n "+
			"Id: 1, Name: A, AvailableCopies: 6")
}
