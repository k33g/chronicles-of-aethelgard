package config

import "rpg/models"

var Directions = map[string]models.Direction{
	"north": {DX: 0, DY: -1, Name: "north"},
	"south": {DX: 0, DY: 1, Name: "south"},
	"west":  {DX: -1, DY: 0, Name: "west"},
	"east":  {DX: 1, DY: 0, Name: "east"},
}

var ExitCell = models.Position{
	X: 10,
	Y: 10,
}


/* === MONSTER DESCRIPTIOM === */
var MonsterTypes = []models.Monster{
	{Name: "Skeleton", HP: 20, AttackPower: 5, Symbol: "💀"},
	{Name: "Troll", HP: 40, AttackPower: 8, Symbol: "👹"},
	{Name: "Dragon", HP: 100, AttackPower: 15, Symbol: "🐲"},
	{Name: "Werewolf", HP: 30, AttackPower: 7, Symbol: "🐺"},
	{Name: "Gobelin", HP: 50, AttackPower: 10, Symbol: "👺"},
}

/*
var NPCTypes = []models.NPC{
	{Type: models.Merchant, Symbol: "🤩"},
	{Type: models.Guard, Symbol: "🤠"},
	{Type: models.Sorcerer, Symbol: "😈"},
}
*/

var NPCTypesSymbols = map[models.NPCType]string{
	models.Merchant: "🤩",
	models.Guard:    "🤠",
	models.Sorcerer: "😈",
}

var StartingStats = map[models.Race]struct {
	HP     int
	Attack int
}{
	models.Human:  {100, 10},
	models.Elf:    {80, 20},
	models.Dwarf:  {120, 15},
	models.Wizard: {60, 25},
}
