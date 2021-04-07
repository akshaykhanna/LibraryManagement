package user

import (
	b "LibraryManagement/book"
	"github.com/stretchr/testify/assert"
	"testing"
)

var user User

func setup() {
	user = *NewUser("Jack")
}

func TestBorrowBook_whenBookBorrowedShouldAddItToUserBook(t *testing.T) {
	setup()
	user := user.BorrowBook(b.NewBook(1, "A"))
	expectedNoOfBooksWithUser := 1
	actualNoOfBooksWithUser := len(user.books)
	if actualNoOfBooksWithUser != expectedNoOfBooksWithUser {
		t.Errorf("AddBook failed, expected %d & got %d", expectedNoOfBooksWithUser, actualNoOfBooksWithUser)
	}
}
func TestReturnBook_ShouldRemoveBookFromUserBooks(t *testing.T) {
	setup()
	book := b.NewBook(1, "A")
	user := user.BorrowBook(book)
	assert.Equal(t, len(user.books), 1)
	user.ReturnBook(book.GetId())
	assert.Equal(t, len(user.books), 0)
}

func TestUser_IsHavingBook_ShouldReturnTrueIfUserHasThatBook(t *testing.T) {
	setup()
	book := b.NewBook(1, "A")
	user.BorrowBook(book)
	flag := user.IsHavingBook(book.GetId())
	if !flag {
		t.Errorf("IsHavingBook failed, expected %v & got %v", true, flag)
	}
}

func TestUser_IsHavingBook_ShouldReturnFalseIfUserDoNotHaveThatBook(t *testing.T) {
	setup()
	book := b.NewBook(1, "A")
	flag := user.IsHavingBook(book.GetId())
	if flag {
		t.Errorf("IsHavingBook failed, expected %v & got %v", false, flag)
	}
}
