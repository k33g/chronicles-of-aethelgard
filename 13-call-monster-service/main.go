package main

import (
	"13-call-monster-service/services"
	"fmt"
)

func main() {

	roomName, _ := services.GetRandomRoomName()
	roomDescription, _ := services.GetRoomDescription(roomName)
	fmt.Println("🏠 Room Description:\n", roomDescription)

	monsterName := "Giant Snake"
	monsterDescription, _ := services.GetMonsterDescription(monsterName)
	fmt.Println("🐍 Monster Description:\n", monsterDescription)

}
