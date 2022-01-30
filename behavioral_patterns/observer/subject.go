package main

type Observable interface {
	registerObserver(obs Observer)
	unregisterObserver(obs Observer)
	notifyAll()
}

// Subject have a list of listeners and
// a data that gets changed, triggering them
type Subject struct {
	listeners []Listener
	data      string
}

// ChangeData will cause the Listeners to be called
func (s *Subject) ChangeData(data string) {
	s.data = data
	s.notifyAll()
}

// registerObserver adds an observer to the list
func (s *Subject) registerObserver(listener Listener) {
	s.listeners = append(s.listeners, listener)
}

// unregisterObserver removes an observer from the list
func (s *Subject) unregisterObserver(listener Listener) {
	var newListeners []Listener

	for _, l := range s.listeners {
		if listener.Name != l.Name {
			newListeners = append(newListeners, l)
		}
	}
	s.listeners = newListeners
}

// notifyAll calls all the listeners with the changed data
func (s *Subject) notifyAll() {
	for _, l := range s.listeners {
		l.onUpdate(s.data)
	}
}
