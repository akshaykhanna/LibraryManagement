package book

import "fmt"

type Book struct {
	id   int
	name string
}

func NewBook(id int, name string) *Book {
	return &Book{id, name}
}
func (b *Book) GetId() int {
	return b.id
}

func (b *Book) GetName() string {
	return b.name
}

func (b *Book) View() string {
	return fmt.Sprintf("Id: %d, Name: %s", b.id, b.name)
}
