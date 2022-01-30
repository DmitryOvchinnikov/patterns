package main

import "fmt"

func newPublication(pubType string, name string, pages int) (Publisher, error) {
	if pubType == "newspaper" {
		return createNewspaper(name, pages), nil
	}
	if pubType == "magazine" {
		return createMagazine(name, pages), nil
	}

	return nil, fmt.Errorf("newPublication(): no such publication type")
}
