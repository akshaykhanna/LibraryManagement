package book

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var books Books

func setuBopks() (*Book, *Book) {
	books = Books{}
	book1 := NewBook(1, "A")
	book2 := NewBook(2, "B")
	books.AddBook(book1).AddBook(book2)
	return book1, book2
}

func TestBooks_ShouldAddBooks(t *testing.T) {
	setuBopks()
	book3 := NewBook(3, "Zero To One")
	books.AddBook(book3)
	assert.Equal(t, 3, len(books))
}

func TestBooks_ShouldRemoveBook(t *testing.T) {
	setuBopks()
	books.RemoveBook(1)
	assert.Equal(t, 1, len(books))
}

func TestBooks_ShouldGetBooksString(t *testing.T) {
	setuBopks()
	expectedString := "List of books\n Id: 1, Name: A\n Id: 2, Name: B"
	assert.Equal(t, expectedString, books.GetBooksString())
}

func TestBooks_ShouldGetABook(t *testing.T) {
	book1, _ := setuBopks()
	expectedBook := book1
	assert.Equal(t, expectedBook, books.GetBook(book1.id))
}

func TestBooks_ShouldReturnNilIfBookNotPresent(t *testing.T) {
	setuBopks()
	assert.Nil(t, books.GetBook(3))
}
