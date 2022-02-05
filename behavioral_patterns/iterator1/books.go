package main

import (
	"fmt"
)

type BookType int

const (
	HardCover BookType = iota
	SoftCover
	EBook
)

type Book struct {
	Name      string
	Author    string
	PageCount int
	Type      BookType
}

type Library struct {
	Collection []Book
}

func (l *Library) IterateBooks(f func(book Book) error) {
	var err error

	for _, b := range l.Collection {
		err = f(b)
		if err != nil {
			fmt.Println("Error encountered:", err)
			break
		}
	}
}

func (l *Library) CreateIterator() iterator {
	return &BookIterator{
		books: l.Collection,
	}
}

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
