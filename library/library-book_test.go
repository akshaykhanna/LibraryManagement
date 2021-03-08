package library

import (
	"testing"
)

func TestBook_IsBookAvailable_shouldReturnTrueIfAvailable(t *testing.T) {
	newBook := NewLibBook(1, "A", 5)
	if !newBook.IsBookAvailable() {
		t.Errorf("IsBookAvailable failed, expected %v & got %v", true, newBook.IsBookAvailable())
	}
}

func TestBook_IsBookAvailable_shouldReturnFalseIfNotAvailable(t *testing.T) {
	newBook := NewLibBook(1, "A", 5)
	newBook.SetAvailableCopies(0)
	if newBook.IsBookAvailable() {
		t.Errorf("IsBookAvailable failed, expected %v & got %v", false, newBook.IsBookAvailable())
	}
}
