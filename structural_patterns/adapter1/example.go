package main

import "fmt"

func main() {
	// Create instances of the two TV types with some default valueo.
	tv1 := &OneTV{
		channel: 13,
		volume:  35,
		on:      true,
	}
	tv2 := &TwoTV{
		channel: 20,
		volume:  9,
		on:      true,
	}

	// TwoTV implements the "television" interface, ...
	tv2.turnOn()
	tv2.volumeUp()
	tv2.volumeDown()
	tv2.channelUp()
	tv2.channelDown()
	tv2.turnOff()

	fmt.Printf("----------------\n")

	// oneAdapter for the OneTV type because it has
	// an interface that's different from the one we want to have.
	adapter := &oneAdapter{
		tv: tv1,
	}
	adapter.turnOn()
	adapter.volumeUp()
	adapter.volumeDown()
	adapter.channelUp()
	adapter.channelDown()
	adapter.goToChannel(68)
	adapter.turnOff()
}
