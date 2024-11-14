package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/k33g/chronicles-of-aethelgard/ui"
	"github.com/k33g/chronicles-of-aethelgard/ui/colors"
)

type Position struct {
	X int
	Y int
}

type PlayerType string

const (
	Human PlayerType = "Human"
	Elf   PlayerType = "Elf"
	Dwarf PlayerType = "Dwarf"
)

type MonsterType string

const (
	Orc      MonsterType = "Orc"
	Goblin   MonsterType = "Goblin"
	Skeleton MonsterType = "Skeleton"
)

type NonPlayerCharacterType string

const (
	Merchant NonPlayerCharacterType = "Merchant"
	Guard    NonPlayerCharacterType = "Guard"
	Sorcerer NonPlayerCharacterType = "Sorcerer"
)

type Player struct {
	Name             string
	Position         Position
	PreviousPosition Position
	Kind             PlayerType
	HP               int
	MaxHP            int
	Attack           int
	Gold             int
	XP               int
}

type NonPlayerCharacter struct {
	Name     string
	Kind     NonPlayerCharacterType
	Placed   bool
	Messages []string // 🤔
}

type Monster struct {
	Name   string
	Kind   MonsterType
	Placed bool
	HP     int
	Attack int
}

type Room struct {
	ShortDescription     string
	Description          string
	Position             Position
	IsVisited            bool
	Monster              *Monster
	NonPlayerCharacter   *NonPlayerCharacter
	Gold                 int
	HealingPotion        int
	GoldGenerated          bool
	HealingPotionGenerated bool
}

var NB_MAX_NPC = 3
var NB_MAX_MONSTERS = 6

var START_POSITION = Position{X: 0, Y: 0}
var EXIT_POSITION = Position{X: 10, Y: 10}

func main() {
	// Start the Game
	dungeon := make(map[Position]*Room)

	// Start room
	room00 := Room{
		Description:        "This is the starting room.",
		Position:           START_POSITION,
		IsVisited:          false,
		Monster:            nil,
		NonPlayerCharacter: nil,
		Gold:               0,
		HealingPotion:      0,
	}

	// Exit room
	room1010 := Room{
		Description:        "This is the exit room.",
		Position:           EXIT_POSITION,
		IsVisited:          false,
		Monster:            nil,
		NonPlayerCharacter: nil,
		Gold:               0,
		HealingPotion:      0,
	}

	// Place the rooms start and exit in the dungeon
	dungeon[START_POSITION] = &room00
	dungeon[EXIT_POSITION] = &room1010

	// Initialize the NPC
	npcs := InitializeNPCs()

	// Initialize the monsters
	monsters := InitializeMonsters()

	// Initialize the player
	// TODO: create a function to choose the player type
	player := Player{
		Name:     "John Doe",
		Position: START_POSITION,
		Kind:     Human,
		HP:       100,
		MaxHP:    100,
		Attack:   10,
		Gold:     0,
		XP:       0,
	}

	// Start the game loop
	gameOver := false
	for !gameOver {

		// Display the current room
		currentRoom := dungeon[player.Position]

		// If the current room is nil (does not exist), create a new room
		if currentRoom == nil { // Create a new room

			ui.Println(colors.Magenta, "⏳ generating room...")

			// Gain 5 XP for discovering a new room
			player.XP += 5

			currentRoom = GenerateNewRoom(player.Position, npcs, NB_MAX_NPC, monsters, NB_MAX_MONSTERS)

			// Place the room in the dungeon
			dungeon[player.Position] = currentRoom
		}

		// Display the room
		currentRoom.IsVisited = true // 🤔
		DisplayRoom(currentRoom)

		// If there is a healing potion in the room
		if currentRoom.HealingPotion > 0 {
			ui.Println(colors.Magenta, "🧪 You find a healing potion!", currentRoom.HealingPotion)
			player.HP += currentRoom.HealingPotion
			if player.HP > player.MaxHP {
				player.HP = player.MaxHP
			}	
		}
		// If there is gold in the room
		if currentRoom.Gold > 0 {
			ui.Println(colors.Magenta, "💰 You find some gold!", currentRoom.Gold)
			player.Gold += currentRoom.Gold
		}

		// If there is a monster in the room
		if currentRoom.Monster != nil {
			// Fight or escape
			switch FightOrEscapeMenu(currentRoom, player) {
			case "i":
				// TODO: to be implemented
				// Fight
				ui.Println(colors.Red, "🔥 Fight!")

			case "s":
				// Escape
				ui.Println(colors.Red, "🚀 Escape! Return to the previous room!")
				player.Position = player.PreviousPosition
				player.PreviousPosition = currentRoom.Position

				// Return to the previous room
				currentRoom = dungeon[player.Position]
				// Display the room
				DisplayRoom(currentRoom)

			default:
				ui.Println(colors.Red, "🤔 I don't understand your choice")
			}

		}

		if currentRoom.NonPlayerCharacter != nil {
			// Chat with NPC ... or not

			switch ChatMenu(currentRoom, player) {
			case "h":
				// TODO: to be implemented
				ui.Println(colors.Purple, "🖐️🤖 HERE, CHAT WITH NPC")
				ui.Println(colors.Purple, "🖐️🤖 some information from the chat with NPC can  be used to gain XP 🤔")

			case "o":
				ui.Println(colors.Purple, "👋 Bye, thank you! 🙂")

			default:
				ui.Println(colors.Purple, "🤔 I don't understand your choice")
			}
		}

		// Main menu
		switch MainMenu(currentRoom, player) {
		case "n":
			player.PreviousPosition = player.Position
			player.Position = Position{X: player.Position.X, Y: player.Position.Y + 1}
		case "s":
			player.PreviousPosition = player.Position
			player.Position = Position{X: player.Position.X, Y: player.Position.Y - 1}
		case "w":
			player.PreviousPosition = player.Position
			player.Position = Position{X: player.Position.X - 1, Y: player.Position.Y}
		case "e":
			player.PreviousPosition = player.Position
			player.Position = Position{X: player.Position.X + 1, Y: player.Position.Y}
		case "x":
			gameOver = true
		default:
			ui.Println(colors.Red, "🤔 I don't understand your choice")
		}

	}

}

// ----------------------------------
// NPCs
// ----------------------------------
func InitializeNPCs() []*NonPlayerCharacter {
	// TODO: change the name with the AI generated names

	npc1 := NonPlayerCharacter{
		Name:   "Bob",
		Kind:   Merchant,
		Placed: false,
	}

	npc2 := NonPlayerCharacter{
		Name:   "Alice",
		Kind:   Guard,
		Placed: false,
	}

	npc3 := NonPlayerCharacter{
		Name:   "Eve",
		Kind:   Sorcerer,
		Placed: false,
	}

	return []*NonPlayerCharacter{&npc1, &npc2, &npc3}
}

func CountPlacedNPC(npcs []*NonPlayerCharacter) int {
	count := 0
	for _, npc := range npcs {
		if npc.Placed {
			count++
		}
	}
	return count
}

func FindFirstUnplacedNPC(npcs []*NonPlayerCharacter) (*NonPlayerCharacter, bool) {
	for _, npc := range npcs {
		if !npc.Placed {
			return npc, true
		}
	}
	return nil, false // Return nil and false if none found
}

func ChatMenu(currentRoom *Room, player Player) string {
	choice, _ := ui.Input(
		colors.Purple,
		fmt.Sprintf("> Room: (%d,%d)- H:%d/%d A:%d XP:%d $:%d - 👋 You meet %s [%s]! c(h)at or n(o)t?",
			currentRoom.Position.X,
			currentRoom.Position.Y,
			player.HP,
			player.MaxHP,
			player.Attack,
			player.XP,
			player.Gold,
			currentRoom.NonPlayerCharacter.Name,
			currentRoom.NonPlayerCharacter.Kind,
		))
	return strings.TrimSpace(choice)
}

// ----------------------------------
// Monsters
// ----------------------------------
func InitializeMonsters() []*Monster {
	// TODO: change the name with the AI generated names

	monster1 := Monster{
		Name:   "Blouch",
		Kind:   Orc,
		Placed: false,
		HP:     20,
		Attack: 5,
		//Position: Position{},
	}

	monster2 := Monster{
		Name:   "Hertuiop",
		Kind:   Goblin,
		Placed: false,
		HP:     10,
		Attack: 3,
		//Position: Position{},
	}

	monster3 := Monster{
		Name:   "Necro",
		Kind:   Skeleton,
		Placed: false,
		HP:     15,
		Attack: 4,
		//Position: Position{},
	}

	monster4 := Monster{
		Name:   "Warf",
		Kind:   Orc,
		Placed: false,
		HP:     20,
		Attack: 5,
		//Position: Position{},
	}

	monster5 := Monster{
		Name:   "Zouille",
		Kind:   Goblin,
		Placed: false,
		HP:     10,
		Attack: 3,
		//Position: Position{},
	}

	monster6 := Monster{
		Name:   "Thorg",
		Kind:   Skeleton,
		Placed: false,
		HP:     15,
		Attack: 4,
		//Position: Position{},
	}

	return []*Monster{&monster1, &monster2, &monster3, &monster4, &monster5, &monster6}

}

func CountPlacedMonsters(monsters []*Monster) int {
	count := 0
	for _, monster := range monsters {
		if monster.Placed {
			count++
		}
	}
	return count
}

func FindFirstUnplacedMonster(monsters []*Monster) (*Monster, bool) {
	for _, monster := range monsters {
		if !monster.Placed {
			return monster, true
		}
	}
	return nil, false // Return nil and false if none found
}

func FightOrEscapeMenu(currentRoom *Room, player Player) string {
	choice, _ := ui.Input(
		colors.Red,
		fmt.Sprintf("> Room: (%d,%d)- H:%d/%d A:%d XP:%d $:%d - 🙀 You meet %s [%s]! f(i)ght or e(s)cape?",
			currentRoom.Position.X,
			currentRoom.Position.Y,
			player.HP,
			player.MaxHP,
			player.Attack,
			player.XP,
			player.Gold,
			currentRoom.Monster.Name,
			currentRoom.Monster.Kind,
		))
	return strings.TrimSpace(choice)
}

// ----------------------------------
// Rooms
// ----------------------------------
func DisplayRoom(room *Room) {
	ui.Println(colors.Blue, "Room Description:", room.Description)
	ui.Println(colors.Red, "Monster:", room.Monster)
	ui.Println(colors.Orange, "NPC:", room.NonPlayerCharacter)
	ui.Println(colors.Yellow, "Gold:", room.Gold)
	ui.Println(colors.Green, "Healing Potion:", room.HealingPotion)
}

func GetShortRoomDescription() string {
	RoomDescriptions := []string{
		"A dark room with spider webs",
		"A room lit by torches",
		"A damp stone-walled corridor",
		"An old dusty library",
		"A room with moss-covered floor",
		"A crypt with rune-carved walls",
		"An abandoned throne room",
		"A ruined kitchen",
	}
	return RoomDescriptions[rand.Intn(len(RoomDescriptions))]
}

func GenerateNewRoom(position Position, npcs []*NonPlayerCharacter, NB_MAX_NPC int, monsters []*Monster, NB_MAX_MONSTERS int) *Room {
	room := &Room{
		Description:        GetShortRoomDescription(),
		Position:           position,
		IsVisited:          false,
		Monster:            nil,
		NonPlayerCharacter: nil,
		Gold:               0,
		HealingPotion:      0,
	}
	// TODO: implement gold and healing potion
	// 20% chance to place a gold
	if rand.Float32() < 0.20 && !room.GoldGenerated {
		room.Gold = rand.Intn(50) + 10
		room.GoldGenerated = true
	}

	// 30% chance to place a healing potion
	if rand.Float32() < 0.30 && !room.HealingPotionGenerated {
		room.HealingPotion = rand.Intn(20) + 5
		room.HealingPotionGenerated = true
	}

	// 20% chance to place a NPC
	if CountPlacedNPC(npcs) < NB_MAX_NPC && rand.Float32() < 0.20 {
		npc, found := FindFirstUnplacedNPC(npcs)
		if found {
			npc.Placed = true
			room.NonPlayerCharacter = npc
		}
	}
	// No NPC placed then we can place a monster
	if room.NonPlayerCharacter == nil {
		// 30% chance to place a monster
		if CountPlacedMonsters(monsters) < NB_MAX_MONSTERS && rand.Float32() < 0.30 {
			monster, found := FindFirstUnplacedMonster(monsters)
			if found {
				monster.Placed = true
				room.Monster = monster
			}
		}
	}

	return room
}

func MainMenu(currentRoom *Room, player Player) string {
	choice, _ := ui.Input(
		colors.Green,
		fmt.Sprintf("> Room: (%d,%d)- H:%d/%d A:%d XP:%d $:%d - move (n\u2191,s\u2193,w\u2190,e\u2192) e(x)it? ",
			currentRoom.Position.X,
			currentRoom.Position.Y,
			player.HP,
			player.MaxHP,
			player.Attack,
			player.XP,
			player.Gold,
		))
	return strings.TrimSpace(choice)
}
