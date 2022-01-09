package main

import "fmt"

func main() {
	var bldr = newNotificationBuilder()

	bldr.SetTitle("New Notification")
	bldr.SetSubtitle("This is Subtitle")
	bldr.SetIcon("icon.png")
	bldr.SetImage("image.jpg")
	bldr.SetPriority(4)
	bldr.SetMessage("This is a basic Notification with some text in it")
	bldr.SetType("alert")

	notif, err := bldr.Build()
	if err != nil {
		fmt.Printf("Error creating the notification: %s", err)
	}
	fmt.Printf("Notification: %+v", notif)
}
