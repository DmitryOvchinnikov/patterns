package main

import "fmt"

// magazine define the struct and embed the Publisher interface inside publication.
type magazine struct {
	publication
}

// String gives a string representation of the type.
func (m magazine) String() string {
	return fmt.Sprintf("This is a magazine named %s", m.name)
}

// createMagazine returns a new Magazine object.
func createMagazine(name string, pages int) Publisher {
	return &magazine{
		publication: publication{
			name:  name,
			pages: pages,
		},
	}
}
