package main

import "fmt"

func main() {
	mag1, _ := newPublication("magazine", "Time", 50)
	mag2, _ := newPublication("magazine", "Life", 40)
	news1, _ := newPublication("newspaper", "The Herald", 60)
	news2, _ := newPublication("newspaper", "The Standard", 30)

	pubDetails(mag1)
	pubDetails(mag2)
	pubDetails(news1)
	pubDetails(news2)
}

func pubDetails(pub Publisher) {
	fmt.Printf("----------------\n")
	fmt.Printf("%v\n", pub)
	fmt.Printf("Type: %T\n", pub)
	fmt.Printf("Name: %s\n", pub.getName())
	fmt.Printf("Pages: %d\n", pub.getPages())
	fmt.Printf("----------------\n")
}
