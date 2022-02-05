package main

import "fmt"

type TwoTV struct {
	channel int
	volume  int
	on      bool
}

func (tv *TwoTV) turnOn() {
	fmt.Println("TwoTV is now on")
	tv.on = true
}

func (tv *TwoTV) turnOff() {
	fmt.Println("TwoTV is now off")
	tv.on = false
}

func (tv *TwoTV) volumeUp() int {
	tv.volume++
	fmt.Println("Increasing TwoTV volume to", tv.volume)
	return tv.volume
}

func (tv *TwoTV) volumeDown() int {
	tv.volume--
	fmt.Println("Decreasing TwoTV volume to", tv.volume)
	return tv.volume
}

func (tv *TwoTV) channelUp() int {
	tv.channel++
	fmt.Println("Increasing TwoTV channel to", tv.channel)
	return tv.channel
}

func (tv *TwoTV) channelDown() int {
	tv.channel--
	fmt.Println("Decreasing TwoTV channel to", tv.channel)
	return tv.channel
}
