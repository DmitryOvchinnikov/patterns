package main

type oneAdapter struct {
	// tv field for the OneTV reference.
	tv *OneTV
}

func (s *oneAdapter) turnOn() {
	s.tv.setOnState(true)
}

func (s *oneAdapter) turnOff() {
	s.tv.setOnState(false)
}

func (s *oneAdapter) volumeUp() int {
	vol := s.tv.getVolume() + 1
	s.tv.setVolume(vol)

	return vol
}

func (s *oneAdapter) volumeDown() int {
	vol := s.tv.getVolume() - 1
	s.tv.setVolume(vol)

	return vol
}

func (s *oneAdapter) channelUp() int {
	ch := s.tv.getChannel() + 1
	s.tv.setChannel(ch)

	return ch
}

func (s *oneAdapter) channelDown() int {
	ch := s.tv.getChannel() - 1
	s.tv.setChannel(ch)

	return ch
}

func (s *oneAdapter) goToChannel(ch int) {
	s.tv.setChannel(ch)
}
