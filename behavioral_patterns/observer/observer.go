package main

import "fmt"

type Observer interface {
	updated(data string)
}

type subscriber struct {
	name string
}

// updated has called when subscriber got data change.
func (s *subscriber) updated(data string) {
	fmt.Println("Subscriber:", s.name, "got data change:", data)
}
