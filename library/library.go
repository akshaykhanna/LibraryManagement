package library

type Library struct {
	books []Book
}

func NewLibrary() Library {
	return Library{[]Book{}}
}
func (l Library) ViewBooks() string {
	if len(l.books) == 0 {
		return "Library is empty"
	}
	return ""
}
