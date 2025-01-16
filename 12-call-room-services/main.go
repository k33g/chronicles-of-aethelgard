package main

import (
	"12-call-room-services/services"
	"fmt"
)

func main() {
	// Get random room name ...
	roomName, _ := services.GetRandomRoomName()
	roomDescription, _ := services.GetRoomDescription(roomName)
	fmt.Println("🏠 Room Description:\n", roomDescription)

}
