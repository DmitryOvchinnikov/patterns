package main

type IterableCollection interface {
	CreateIterator() iterator
}

type iterator interface {
	HasNext() bool
	Next() *Book
}

type BookIterator struct {
	current int
	books   []Book
}

func (b *BookIterator) HasNext() bool {
	if b.current >= len(b.books) {
		return false
	}

	return true
}

func (b *BookIterator) Next() *Book {
	if !b.HasNext() {
		return nil
	}
	bk := b.books[b.current]
	b.current++

	return &bk
}
