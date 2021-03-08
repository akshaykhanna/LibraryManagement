package library

import (
	b "LibraryManagement/book"
	"fmt"
)

type libBook struct {
	*b.Book
	availableCopies int
}

func NewLibBook(id int, name string, availableCopies int) *libBook {
	return &libBook{b.NewBook(id, name), availableCopies}
}

func (b *libBook) SetAvailableCopies(count int) {
	b.availableCopies = count
}

func (b *libBook) GetAvailableCopies() int {
	return b.availableCopies
}
func (b *libBook) IsBookAvailable() bool {
	return b.availableCopies > 0
}

func (b *libBook) View() string {
	return fmt.Sprintf("%s, "+
		"AvailableCopies: %d", b.Book.View(), b.availableCopies)
}
