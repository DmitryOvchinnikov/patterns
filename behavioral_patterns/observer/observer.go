package main

import "fmt"

// Define the interface for the observer type
type observer interface {
	onUpdate(data string)
}

// DataListener observer will have a name
type DataListener struct {
	Name string
}

// onUpdate to conform to the interface
func (l *DataListener) onUpdate(data string) {
	fmt.Println("Listener:", l.Name, "got data change:", data)
}
