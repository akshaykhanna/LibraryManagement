package library

import (
	"fmt"
)

type libBooks []*libBook

func (lb *libBooks) addBook(book *libBook) *libBooks {
	*lb = append(*lb, book)
	return lb
}

func (lb *libBooks) removeBook(bookId int) *libBook {
	var newBooks libBooks
	var removedBook *libBook
	for _, book := range *lb {
		if book.GetId() != bookId {
			newBooks = append(newBooks, book)
		} else {
			removedBook = book
		}
	}
	*lb = newBooks
	return removedBook
}

func (lb *libBooks) getLibBook(bookId int) *libBook {
	for _, book := range *lb {
		if book.GetId() == bookId {
			return book
		}
	}
	return nil
}

func (lb *libBooks) isBookPresent(bookId int) bool {
	if lb.getLibBook(bookId) != nil {
		return true
	}
	return false
}

func (lb *libBooks) getBooksString() string {
	if len(*lb) == 0 {
		return "No book found"
	}
	booksString := fmt.Sprintf("%s", "List of books")
	for _, book := range *lb {
		booksString = booksString + fmt.Sprintf("\n %s", book.View())
	}
	return booksString
}
