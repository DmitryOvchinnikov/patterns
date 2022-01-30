package main

import "fmt"

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

func (l *Library) createIterator() iterator {
	return &BookIterator{
		books: l.Collection,
	}
}
