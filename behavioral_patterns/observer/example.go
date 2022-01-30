package main

func main() {
	// Construct two Listener observers and give each one a name
	listener1 := Listener{
		"Listener1",
	}
	listener2 := Listener{
		"Listener2",
	}

	// Create the Subject that the listeners will observe
	subj := &Subject{}
	// Register each listener with the Subject
	subj.registerObserver(listener1)
	subj.registerObserver(listener2)

	// Change the data in the Subject -- this will cause
	// the onUpdate method of each listener to be called
	subj.ChangeData("Monday!")
	subj.ChangeData("Sunday!")

	// Try to unregister one of the observers
	subj.unregisterObserver(listener2)
	subj.ChangeData("Friday !")
}
