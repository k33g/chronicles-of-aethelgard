package models

type Position struct {
	X, Y int
}

type Direction struct {
	DX, DY int
	Name   string
}

type Race string

const (
	Human   Race = "human"
	Elf     Race = "elf"
	Dwarf   Race = "dwarf"
	Wizard  Race = "magician"
	Nothing Race = "nothing"
)

type Item struct {
	Name        string
	Description string
	Effect      int    // Pour les potions : points de vie restaurés
	Type        string // "or" ou "potion"
	Value       int    // Pour l'or : valeur en pièces
}

type Monster struct {
	Name        string
	HP          int
	AttackPower int
	CurrentHP   int
}

type NPCType string

const (
	Merchant NPCType = "merchant"
	Guard    NPCType = "guard"
	Sorcerer NPCType = "sorcerer"
)

type NPC struct {
	Type     NPCType
	Messages []string
}

type Room struct {
	Description  string
	NPC          *NPC
	Item         *Item
	Monster      *Monster
	IsVisited    bool
	PreviousRoom *Position
}

type Player struct {
	Race      Race
	HP        int
	MaxHP     int
	Inventory []Item
	Attack    int
	Gold      int
	XP        int // Ajout des points d'expérience
}
