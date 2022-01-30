package main

import "fmt"

type Observer interface {
	onUpdate(data string)
}

// Listener observer will have a name.
type Listener struct {
	Name string
}

// onUpdate to conform to the Observer interface
func (l *Listener) onUpdate(data string) {
	fmt.Println("Listener:", l.Name, "got data change:", data)
}
