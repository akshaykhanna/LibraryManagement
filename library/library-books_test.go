package library

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var books libBooks

func setupLibBooks() (*libBook, *libBook) {
	books = libBooks{}
	book1 := NewLibBook(1, "A", 3)
	book2 := NewLibBook(2, "B", 1)
	books.addBook(book1).addBook(book2)
	return book1, book2
}

func TestLibBooks_ShouldAddBooks(t *testing.T) {
	setupLibBooks()
	book3 := NewLibBook(3, "Zero To One", 2)
	books.addBook(book3)
	assert.Equal(t, 3, len(books))
}

func TestLibBooks_ShouldRemoveBook(t *testing.T) {
	book1, _ := setupLibBooks()
	assert.Equal(t, 2, len(books))
	books.removeBook(book1.GetId())
	assert.Equal(t, 1, len(books))
}

func TestLibBooks_ShouldGetALibBook(t *testing.T) {
	book1, _ := setupLibBooks()
	expectedBook := book1
	assert.Equal(t, expectedBook, books.getLibBook(book1.GetId()))
}

func TestLibBooks_ShouldReturnNilIfBookNotPresent(t *testing.T) {
	setupLibBooks()
	assert.Nil(t, books.getLibBook(3))
}

func TestLibBooks_ShouldReturnTrueIfBookIsPresent(t *testing.T) {
	setupLibBooks()
	assert.True(t, books.isBookPresent(2))
}

func TestLibBooks_ShouldReturnFalseIfBookIsNotPresent(t *testing.T) {
	setupLibBooks()
	assert.False(t, books.isBookPresent(3))
}

func TestLibBooks_ShouldGetLibBooksString(t *testing.T) {
	setupLibBooks()
	expectedString := "List of books\n Id: 1, Name: A, AvailableCopies: 3\n Id: 2, Name: B, AvailableCopies: 1"
	assert.Equal(t, expectedString, books.getBooksString())
}
