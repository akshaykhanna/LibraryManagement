package main

import (
	"LibraryManagement/book"
	"LibraryManagement/library"
	"LibraryManagement/manager"
	"LibraryManagement/user"
	"fmt"
)

func main() {
	book1 := book.NewBook(1, "Atomic Habits")
	book2 := book.NewBook(2, "Tool of Titans")
	myLib := library.NewLibrary()
	myLib.AddBook(book1, 5).AddBook(book2, 1)
	myManger := manager.NewManager(myLib)
	fmt.Println(myLib.ViewBooks())
	myUser := user.NewUser("Jack")

	borrowBook(myManger, myUser, myLib, book1)
	borrowBook(myManger, myUser, myLib, book2)
	borrowBook(myManger, myUser, myLib, book2)
	borrowBook(myManger, myUser, myLib, book1)

	returnBook(myManger, myUser, myLib, book2)
	returnBook(myManger, myUser, myLib, book2)
	returnBook(myManger, myUser, myLib, book1)

}

func borrowBook(myManger manager.Manager, user *user.User, myLib *library.Library, book *book.Book) {
	fmt.Println("---------------------------------")
	fmt.Printf("\nUser %s trying to borrow book %s (id : %d) from library. \n",
		user.GetName(), book.GetName(), book.GetId())
	err := myManger.HandleBorrow(user, book.GetId())
	if err != nil {
		fmt.Printf("Error: %s \n ", err.Error())
	} else {
		fmt.Println("Borrowing successfull !")
		fmt.Println(myLib.ViewBooks())
	}
}

func returnBook(myManger manager.Manager, user *user.User, myLib *library.Library, book *book.Book) {
	fmt.Println("---------------------------------")
	fmt.Printf("\nUser %s trying to return book %s (id : %d) to library. \n",
		user.GetName(), book.GetName(), book.GetId())
	err := myManger.HandleReturn(user, book.GetId())
	if err != nil {
		fmt.Printf("Error: %s \n", err.Error())
	} else {
		fmt.Println("Return successfull !")
		fmt.Println(myLib.ViewBooks())
	}
}
