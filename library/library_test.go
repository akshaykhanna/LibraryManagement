package library

import (
	b "LibraryManagement/book"
	"testing"
)

var library Library

func setup() {
	library = NewLibrary()
	library.AddBook(b.NewBook(1, "A"), 5).AddBook(b.NewBook(2, "B"), 3)
}

func TestViewBooks_shouldReturnEmptyWhenNoBooks(t *testing.T) {
	library := NewLibrary()
	actualBooksString := library.ViewBooks()
	const expectedBooksString = "Library is empty"
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
	borrowedBook := library.BorrowBook(1)
	libBorrowedBook := library.getLibBook(1)
	expectedAvailableCopies := 4
	if expectedAvailableCopies != libBorrowedBook.availableCopies {
		t.Errorf("BorrowBooks failed, expected %v & got %d", expectedAvailableCopies, libBorrowedBook.availableCopies)
	}
	if borrowedBook.GetName() != "A" && borrowedBook.GetId() != 1 {
		t.Errorf("BorrowBooks failed, expected %v & got %v", b.NewBook(1, "A"), borrowedBook)
	}
}
