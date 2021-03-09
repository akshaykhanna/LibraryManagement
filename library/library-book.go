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

func (lb *libBook) SetAvailableCopies(count int) {
	lb.availableCopies = count
}

func (lb *libBook) GetAvailableCopies() int {
	return lb.availableCopies
}
func (lb *libBook) IsBookAvailable() bool {
	return lb.availableCopies > 0
}

func (lb *libBook) ToBook() *b.Book {
	if lb != nil {
		return b.NewBook(lb.GetId(), lb.GetName())
	}
	return nil
}

func (lb *libBook) View() string {
	return fmt.Sprintf("%s, "+
		"AvailableCopies: %d", lb.Book.View(), lb.availableCopies)
}
