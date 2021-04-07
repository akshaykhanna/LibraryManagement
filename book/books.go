package book

import "fmt"

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

func (b *Books) GetBooksString() string {
	if len(*b) == 0 {
		return "No book found"
	}
	booksString := fmt.Sprintf("%s", "List of books")
	for _, book := range *b {
		booksString = booksString + fmt.Sprintf("\n %s", book.View())
	}
	return booksString
}

func (b *Books) GetBook(bookId int) *Book {
	for _, book := range *b {
		if book.id == bookId {
			return book
		}
	}
	return nil
}
