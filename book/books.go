package book

type Books []*Book

func (b *Books) AddBook(book *Book) *Books {
	*b = append(*b, book)
	return b
}

func (b *Books) RemoveBook(bookId int) *Book {
	var newBooks Books
	var removedBook *Book
	for _, book := range *b {
		if book.GetId() != bookId {
			newBooks = append(newBooks, book)
		} else {
			removedBook = book
		}
	}
	*b = newBooks
	return removedBook
}
