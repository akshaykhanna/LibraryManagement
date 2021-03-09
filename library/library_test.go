package library

import (
	b "LibraryManagement/book"
	"github.com/stretchr/testify/assert"
	"testing"
)

var library Library

func setup() {
	library = *NewLibrary()
	library.AddBook(b.NewBook(1, "A"), 5).AddBook(b.NewBook(2, "B"), 3)
}

func TestViewBooks_shouldReturnEmptyWhenNoBooks(t *testing.T) {
	library := NewLibrary()
	actualBooksString := library.ViewBooks()
	const expectedBooksString = "No book found"
	if actualBooksString != expectedBooksString {
		t.Errorf("ViewBooks failed, expected %v & got %v", expectedBooksString, actualBooksString)
	}
}
func TestViewBooks_shouldReturnBooksStringWhenBooksArePresent(t *testing.T) {
	setup()
	actualBooksString := library.ViewBooks()
	const expectedBooksString = `
 List of books
 Id: 1, Name: A, AvailableCopies: 5
 Id: 2, Name: B, AvailableCopies: 3`
	if actualBooksString != expectedBooksString {
		t.Errorf("ViewBooks failed, expected %v & got %v", expectedBooksString, actualBooksString)
	}
}

func TestViewBooks_shouldNotReturnEmptyWhenBooksAreThere(t *testing.T) {
	library := NewLibrary()
	library.books = append(library.books, NewLibBook(1, "A", 5))
	actualBooksString := library.ViewBooks()
	const expectedBooksString = "Library is empty"
	if actualBooksString == expectedBooksString {
		t.Errorf("ViewBooks failed, expected %v & got %v", expectedBooksString, actualBooksString)
	}
}

func TestAddBooks_shouldAddBooksToLibrary(t *testing.T) {
	library := NewLibrary()
	library.AddBook(b.NewBook(1, "A"), 5)
	if len(library.books) == 0 {
		t.Errorf("AddBook failed, expected %v & got %v", "more than 0 books", "0")
	}
}

func TestAddBooks_shouldOnlyAddBooksToLibraryIfBookWithThatIdIsNotPresent(t *testing.T) {
	library := NewLibrary()
	library.AddBook(b.NewBook(1, "A"), 5).AddBook(b.NewBook(1, "B"), 3)
	if len(library.books) != 1 {
		t.Errorf("AddBook failed, expected %v & got %v", "1 books", 2)
	}
}

func TestGetBook_shouldReturnBookWithPassedId(t *testing.T) {
	setup()
	book := library.GetBook(2)
	if book.GetName() != "B" {
		t.Errorf("GetBooks failed, expected %v & got %v", "libBook with name B", "libBook not found")
	}
}

func TestBorrowBook_whenBookBorrowedShouldReduceItAvailableCount(t *testing.T) {
	setup()
	borrowedBookId := 1
	libBorrowedBook := library.books.getLibBook(borrowedBookId)
	assert.Equal(t, libBorrowedBook.availableCopies, 5)
	borrowedBook := library.BorrowBook(borrowedBookId)
	assert.Equal(t, libBorrowedBook.availableCopies, 4)
	assert.Equal(t, borrowedBook.GetId(), borrowedBookId)
	assert.Equal(t, borrowedBook.GetName(), "A")
}

func TestBorrowBook_ShouldRemoveBookFromLibIfOnlyOneCopyPresentBeforeBorrowing(t *testing.T) {
	setup()
	borrowedBookId := 3
	library.AddBook(b.NewBook(borrowedBookId, "C"), 1)
	libBorrowedBook := library.books.getLibBook(borrowedBookId)
	assert.Equal(t, libBorrowedBook.availableCopies, 1)
	borrowedBook := library.BorrowBook(borrowedBookId)
	assert.False(t, library.books.isBookPresent(borrowedBookId))
	assert.Equal(t, borrowedBook.GetId(), borrowedBookId)
	assert.Equal(t, borrowedBook.GetName(), "C")
}

func TestReturnBook_shouldUpdateBookQuantityIfBookAlreadyThere(t *testing.T) {
	setup()
	returnedBook := b.NewBook(2, "B")
	libReturnedBook := library.books.getLibBook(returnedBook.GetId())
	assert.Equal(t, libReturnedBook.GetAvailableCopies(), 3)
	library.ReturnBook(returnedBook)
	assert.Equal(t, libReturnedBook.GetAvailableCopies(), 4)
}

func TestReturnBook_shouldAddNewBookQuantityIfBookNotPresent(t *testing.T) {
	setup()
	returnedBook := b.NewBook(3, "Monk who sold his ferrari")
	assert.False(t, library.books.isBookPresent(returnedBook.GetId()))
	library.ReturnBook(returnedBook)
	assert.True(t, library.books.isBookPresent(returnedBook.GetId()))
	libReturnedBook := library.books.getLibBook(returnedBook.GetId())
	assert.Equal(t, libReturnedBook.GetAvailableCopies(), 1)
}

func TestCanBookBeBorrow_shouldReturnFalseIfBookIsNotPresent(t *testing.T) {
	setup()
	flag := library.CanBookBeBorrowed(3)
	if flag {
		t.Errorf("CanBookBeBorrowed failed, expected %v & got %v",
			"book not to be borrowed when not present", "but it is borrowed")
	}

}

func TestCanBookBeBorrow_shouldReturnFalseIfBookIsNotAvailable(t *testing.T) {
	setup()
	borrowBookId := 2
	library.books.getLibBook(borrowBookId).SetAvailableCopies(0)
	flag := library.CanBookBeBorrowed(borrowBookId)
	if flag {
		t.Errorf("canBookBeBorrowed failed, expected %v & got %v",
			"book not to be borrowed when not present", "it can be borrowed")
	}

}

func TestCanBookBeBorrow_shouldReturnTrueIfBookIsPresentAndAvailable(t *testing.T) {
	setup()
	borrowBookId := 2
	flag := library.CanBookBeBorrowed(borrowBookId)
	if !flag {
		t.Errorf("canBookBeBorrowed failed, expected %v & got %v",
			"book can be borrowed when present & available", " it can't be borrowed")
	}

}
