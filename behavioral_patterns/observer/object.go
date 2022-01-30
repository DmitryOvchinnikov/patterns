package main

type Observable interface {
	registerObserver(obs Observer)
	unregisterObserver(obs Observer)
	notify()
}

// Object have a list of subscribers and
// a data that gets changed, triggering them.
type Object struct {
	subscribers []subscriber
	data        string
}

// changeTo will cause the subscribers to be called.
func (o *Object) changeTo(data string) {
	o.data = data
	o.notify()
}

// registerObserver adds an observer to the list.
func (o *Object) registerObserver(s subscriber) {
	o.subscribers = append(o.subscribers, s)
}

// unregisterObserver removes an observer from the list.
func (o *Object) unregisterObserver(s subscriber) {
	var newSubs []subscriber

	for _, sub := range o.subscribers {
		if s.name != sub.name {
			newSubs = append(newSubs, sub)
		}
	}
	o.subscribers = newSubs
}

// notify calls the subscribers with the changed data.
func (o *Object) notify() {
	for _, sub := range o.subscribers {
		sub.updated(o.data)
	}
}
