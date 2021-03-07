package book

type Book struct {
	id   int
	name string
}

func NewBook(id int, name string) *Book {
	return &Book{id, name}
}
func (b *Book) Id() int {
	return b.id
}

func (b *Book) Name() string {
	return b.name
}
