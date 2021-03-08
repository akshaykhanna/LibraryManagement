package user

import (
	b "LibraryManagement/book"
	"testing"
)

var user User

func setup() {
	user = NewUser("Jack")
}

func TestBorrowBook_whenBookBorrowedShouldAddItToUserBook(t *testing.T) {
	setup()
	user := user.BorrowBook(b.NewBook(1, "A"))
	expectedNoOfBooksWithUser := 1
	actualNoOfBooksWithUser := len(user.books)
	if actualNoOfBooksWithUser != expectedNoOfBooksWithUser {
		t.Errorf("BookBorrow failed, expected %d & got %d", expectedNoOfBooksWithUser, actualNoOfBooksWithUser)
	}
}
