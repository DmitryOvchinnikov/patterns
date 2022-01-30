package main

import "fmt"

func main() {
	var bldr = newBlueprintBuilder()

	bldr.SetPropertyOne("New Blueprint")
	bldr.SetPropertyTwo("This is second property")
	bldr.SetPropertyThree("This is third property")

	notif, err := bldr.Build()
	if err != nil {
		fmt.Printf("Error creating the Blueprint: %s", err)
	}
	fmt.Printf("Blueprint: %+v", notif)
}
