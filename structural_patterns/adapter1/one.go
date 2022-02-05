package main

import "fmt"

type OneTV struct {
	channel int
	volume  int
	on      bool
}

func (tv *OneTV) getVolume() int {
	fmt.Println("OneTV volume is", tv.volume)
	return tv.volume
}

func (tv *OneTV) setVolume(vol int) {
	fmt.Println("Setting OneTV volume to", vol)
	tv.volume = vol
}

func (tv *OneTV) getChannel() int {
	fmt.Println("OneTV channel is", tv.channel)
	return tv.channel
}

func (tv *OneTV) setChannel(ch int) {
	fmt.Println("Setting OneTV channel to", ch)
	tv.channel = ch
}

func (tv *OneTV) setOnState(tvOn bool) {
	if tvOn == true {
		fmt.Println("OneTV is on")
	} else {
		fmt.Println("OneTV is off")
	}
	tv.on = tvOn
}
