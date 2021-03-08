package library

import b "LibraryManagement/book"

type libBook struct {
	*b.Book
	totalCopies     int
	availableCopies int
}

func NewLibBook(id int, name string, totalCopies int) *libBook {
	return &libBook{b.NewBook(id, name), totalCopies, totalCopies}
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
