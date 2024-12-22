package config

import "rpg-game/models"

var Directions = map[string]models.Direction{
	"north": {DX: 0, DY: -1, Name: "north"},
	"south": {DX: 0, DY: 1, Name: "south"},
	"west":  {DX: -1, DY: 0, Name: "west"},
	"east":  {DX: 1, DY: 0, Name: "east"},
}

var RoomDescriptions = []string{
	"Une pièce sombre avec des toiles d'araignées",
	"Une salle illuminée par des torches",
	"Un couloir humide aux murs de pierre",
	"Une ancienne bibliothèque poussiéreuse",
	"Une salle au sol couvert de mousse",
	"Une crypte aux murs gravés de runes",
	"Une salle du trône abandonnée",
	"Une cuisine en ruines",
}

var NPCMessages = map[models.NPCType][]string{
	models.Merchant: {
		"Voulez-vous voir mes marchandises ?",
		"J'ai des objets rares à vendre !",
		"Les prix sont négociables...",
	},
	models.Guard: {
		"Halte ! Cette zone est surveillée.",
		"Faites attention aux monstres qui rôdent...",
		"Je peux vous indiquer le chemin si besoin.",
	},
	models.Sorcerer: {
		"Je sens une grande magie en ces lieux...",
		"Voulez-vous apprendre quelques sorts ?",
		"Les anciens secrets reposent ici.",
	},
}

var MonsterTypes = []models.Monster{
	{Name: "Gobelin", HP: 20, AttackPower: 5},
	{Name: "Troll", HP: 40, AttackPower: 8},
	{Name: "Dragon", HP: 100, AttackPower: 15},
	{Name: "Wolf", HP: 30, AttackPower: 7},
	{Name: "Bear", HP: 50, AttackPower: 10},
}

var StartingStats = map[models.Race]struct {
	HP     int
	Attack int
}{
	models.Human:  {100, 10},
	models.Elf:    {80, 12},
	models.Dwarf:  {120, 8},
	models.Wizard: {60, 15},
}
