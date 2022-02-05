package main

import (
	"fmt"
)

func main() {
	// iterate via callback function.
	lib.IterateBooks(bookCallback)

	// iterate via anonymous function.
	//lib.IterateBooks(func(b Book) error {
	//	fmt.Println("Book author:", b.Author)
	//	return nil
	//})

	// iterate via a BookIterator.
	//i := lib.CreateIterator()
	//for i.HasNext() {
	//	book := i.Next()
	//	fmt.Printf("Book %+v\n", book)
	//}
}

func bookCallback(b Book) error {
	fmt.Println("Book title:", b.Name)
	return nil
}
