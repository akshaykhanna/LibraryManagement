package library

import (
	"github.com/stretchr/testify/assert"
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

func TestViewLibBook_shouldReturnLibBookString(t *testing.T) {
	book := NewLibBook(1, "Atomic Habits", 5)
	assert.Equal(t, book.View(), "Id: 1, Name: Atomic Habits, AvailableCopies: 5, TotalCopies: 5")
}
