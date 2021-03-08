package book

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var book *Book

func setup() {
	book = NewBook(1, "Rich Dad Poor Dad")
}

func TestViewBook_ShouldReturnBookString(t *testing.T) {
	setup()
	assert.Equal(t, book.View(), "Id: 1, Name: Rich Dad Poor Dad")
}
