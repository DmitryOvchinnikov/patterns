package main

type Publisher interface {
	setName(name string)
	setPages(pages int)
	getName() string
	getPages() int
}

type publication struct {
	name      string
	pages     int
	publisher string
}

func (p *publication) setName(name string) {
	p.name = name
}

func (p *publication) setPages(pages int) {
	p.pages = pages
}

func (p *publication) getName() string {
	return p.name
}

func (p *publication) getPages() int {
	return p.pages
}
