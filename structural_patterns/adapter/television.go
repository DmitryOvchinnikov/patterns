package main

// television interface to use with both TV types.
type television interface {
	volumeUp() int
	volumeDown() int
	channelUp() int
	channelDown() int
	turnOn()
	turnOff()
	goToChannel(ch int)
}