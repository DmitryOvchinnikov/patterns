package main

func main() {

	subscriber1 := subscriber{
		"subscriber1",
	}
	subscriber2 := subscriber{
		"subscriber2",
	}

	// Create the Object that the subscribers will observe.
	obj := &Object{}
	// Register each Subscriber with the Object
	obj.registerObserver(subscriber1)
	obj.registerObserver(subscriber2)

	// Change the data in the Object -- this will cause
	// the Updated method of each Subscriber to be called
	obj.changeTo("Monday!")
	obj.changeTo("Sunday!")

	// Try to unregister one of the observers
	obj.unregisterObserver(subscriber2)
	obj.changeTo("Friday!")
}
