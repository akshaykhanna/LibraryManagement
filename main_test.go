package main

import "testing"

func TestViewBooks(t *testing.T) {
	actualBooksString := viewBooks()
	const expectedBooksString = "Library is empty"
	if actualBooksString != expectedBooksString {
		t.Errorf("ViewBooks failed, expected %v & got %v", expectedBooksString, actualBooksString)
	}
}
