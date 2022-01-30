package main

// Define the interface for the observable type
type observable interface {
	registerObserver(obs observer)
	unregisterObserver(obs observer)
	notifyAll()
}

// The DataSubject will have a list of listeners
// and a field that gets changed, triggering them
type DataSubject struct {
	observers []DataListener
	field     string
}

// ChangeItem will cause the Listeners to be called
func (s *DataSubject) ChangeItem(data string) {
	s.field = data
}

// registerObserver adds an observer to the list
func (s *DataSubject) registerObserver(l DataListener) {
	s.observers = append(s.observers, l)
}

// unregisterObserver removes an observer from the list
func (s *DataSubject) unregisterObserver(l DataListener) {
	var newObs []DataListener

	for _, obs := range s.observers {
		if l.Name != obs.Name {
			newObs = append(newObs, obs)
		}
	}
	s.observers = newObs
}

// notifyAll calls all the listeners with the changed data
func (s *DataSubject) notifyAll() {
	for _, obs := range s.observers {
		obs.onUpdate(s.field)
	}
}
