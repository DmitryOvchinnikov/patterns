package main

import "fmt"

// newspaper define the struct and embed the Publisher interface inside publication.
type newspaper struct {
	publication
}

// String gives a string representation of the type.
func (n newspaper) String() string {
	return fmt.Sprintf("This is a newspaper named %s", n.name)
}

// createNewspaper returns a new Newspaper object.
func createNewspaper(name string, pages int) Publisher {
	return &newspaper{
		publication: publication{
			name:  name,
			pages: pages,
		},
	}
}
