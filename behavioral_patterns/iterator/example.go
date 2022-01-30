package main

import "fmt"

var lib *Library = &Library{
	[]Book{
		{
			Name:      "name1",
			Author:    "writer1",
			PageCount: 864,
			Type:      HardCover,
		},
		{
			Name:      "name2",
			Author:    "writer2",
			PageCount: 42,
			Type:      SoftCover,
		},
		{
			Name:      "name3",
			Author:    "writer3",
			PageCount: 66,
			Type:      EBook,
		},
	},
}

func main() {
	// iterate via callback function.
	//lib.IterateBooks(bookCallback)

	// iterate via anonymous function.
	//lib.IterateBooks(func(b Book) error {
	//	fmt.Println("Book author:", b.Author)
	//	return nil
	//})

	// iterate via a BookIterator.
	i := lib.createIterator()
	for i.hasNext() {
		book := i.next()
		fmt.Printf("Book %+v\n", book)
	}
}

func bookCallback(b Book) error {
	fmt.Println("Book title:", b.Name)
	return nil
}
