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
	user := user.AddBook(b.NewBook(1, "A"))
	expectedNoOfBooksWithUser := 1
	actualNoOfBooksWithUser := len(user.books)
	if actualNoOfBooksWithUser != expectedNoOfBooksWithUser {
		t.Errorf("BookBorrow failed, expected %d & got %d", expectedNoOfBooksWithUser, actualNoOfBooksWithUser)
	}
}

func TestUser_IsHavingBook_ShouldReturnTrueIfUserHasThatBook(t *testing.T) {
	setup()
	book := b.NewBook(1, "A")
	user.AddBook(book)
	flag := user.IsHavingBook(book.GetId())
	if !flag {
		t.Errorf("IsHavingBook failed, expected %v & got %v", true, flag)
	}
}

func TestUser_IsHavingBook_ShouldReturnFalsefUserDoNotHaveThatBook(t *testing.T) {
	setup()
	book := b.NewBook(1, "A")
	flag := user.IsHavingBook(book.GetId())
	if flag {
		t.Errorf("IsHavingBook failed, expected %v & got %v", false, flag)
	}
}
