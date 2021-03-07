package library

import (
	"testing"
)

func TestViewBooks(t *testing.T) {
	library := NewLibrary()
	actualBooksString := library.ViewBooks()
	const expectedBooksString = "Library is empty"
	if actualBooksString != expectedBooksString {
		t.Errorf("ViewBooks failed, expected %v & got %v", expectedBooksString, actualBooksString)
	}
}
