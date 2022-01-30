package main

type oneAdapter struct {
	// tv field for the OneTV reference.
	tv *OneTV
}

func (s *oneAdapter) turnOn() {
	o.tv.setOnState(true)
}

func (s *oneAdapter) turnOff() {
	o.tv.setOnState(false)
}

func (s *oneAdapter) volumeUp() int {
	vol := o.tv.getVolume() + 1
	o.tv.setVolume(vol)

	return vol
}

func (s *oneAdapter) volumeDown() int {
	vol := o.tv.getVolume() - 1
	o.tv.setVolume(vol)

	return vol
}

func (s *oneAdapter) channelUp() int {
	ch := o.tv.getChannel() + 1
	o.tv.setChannel(ch)

	return ch
}

func (s *oneAdapter) channelDown() int {
	ch := o.tv.getChannel() - 1
	o.tv.setChannel(ch)

	return ch
}

func (s *oneAdapter) goToChannel(ch int) {
	o.tv.setChannel(ch)
}
