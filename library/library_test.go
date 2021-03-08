package library

import (
	"testing"
)

var library Library

func setup() {
	library = NewLibrary()
	library.AddBooks(NewBook(1, "A", 5)).AddBooks(NewBook(2, "B", 3))
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
	library.books = append(library.books, NewBook(1, "A", 5))
	actualBooksString := library.ViewBooks()
	const expectedBooksString = "Library is empty"
	if actualBooksString == expectedBooksString {
		t.Errorf("ViewBooks failed, expected %v & got %v", expectedBooksString, actualBooksString)
	}
}

func TestAddBooks_shouldAddBooksToLibrary(t *testing.T) {
	library := NewLibrary()
	library.AddBooks(NewBook(1, "A", 5))
	if len(library.books) == 0 {
		t.Errorf("AddBooks failed, expected %v & got %v", "more than 0 books", "0")
	}
}

func TestAddBooks_shouldOnlyAddBooksToLibraryIfBookWithThatIdIsNotPresent(t *testing.T) {
	library := NewLibrary()
	library.AddBooks(NewBook(1, "A", 5)).AddBooks(NewBook(1, "B", 3))
	if len(library.books) != 1 {
		t.Errorf("AddBooks failed, expected %v & got %v", "1 books", 2)
	}
}

func TestGetBook_shouldReturnBookWithPassedId(t *testing.T) {
	setup()
	book := library.GetBook(2)
	if book.GetName() != "B" {
		t.Errorf("GetBooks failed, expected %v & got %v", "book with name B", "book not found")
	}
}

func TestBorrowBook_whenBookBorrowedShouldReduceItAvailableCount(t *testing.T) {
	setup()
	book := library.BorrowBook(1)
	expectedAvailableCopies := 4
	if expectedAvailableCopies != book.GetAvailableCopies() {
		t.Errorf("BorrowBooks failed, expected %v & got %d", "availableCopies", book.GetAvailableCopies())
	}
}
