package game

import (
	"fmt"
	"math/rand"
	"rpg/config"
	"rpg/models"

	"rpg/ui"
	"rpg/ui/colors"
)

type Game struct {
	CurrentPos    models.Position
	PreviousPos   models.Position
	Rooms         map[models.Position]models.Room
	Player        models.Player
	LastDirection string
	PlacedNPCs    map[models.NPCType]bool // Pour suivre les PNJ déjà placés
}

func NewGame(playerRace models.Race) *Game {
	stats := config.StartingStats[playerRace]

	game := &Game{
		CurrentPos: models.Position{X: 0, Y: 0},
		Rooms:      make(map[models.Position]models.Room),
		Player: models.Player{
			Race:      playerRace,
			HP:        stats.HP,
			MaxHP:     stats.HP,
			Attack:    stats.Attack,
			Inventory: []models.Item{},
			Gold:      0,
			XP:        0,
		},
		PlacedNPCs: make(map[models.NPCType]bool),
	}

	game.Rooms[game.CurrentPos] = game.generateRoom()
	return game
}

func (g *Game) generateRoom() models.Room {
	/* === ROOM DESCRIPTIOM === */
	// BEGIN
	room := models.Room{
		Description: "You are in a dark room.",
		IsVisited:   false,
	}
	// END

	// Vérifier si tous les PNJ ont été placés
	allNPCsPlaced := len(g.PlacedNPCs) >= 3

	// 15% de chances d'avoir un PNJ si tous ne sont pas encore placés
	if !allNPCsPlaced && rand.Float32() < 0.15 {
		// Liste des types de PNJ non encore placés
		availableTypes := make([]models.NPCType, 0)
		for _, npcType := range []models.NPCType{models.Merchant, models.Guard, models.Sorcerer} {
			if !g.PlacedNPCs[npcType] {
				availableTypes = append(availableTypes, npcType)
			}
		}

		if len(availableTypes) > 0 {
			// Choisir aléatoirement parmi les types disponibles
			selectedType := availableTypes[rand.Intn(len(availableTypes))]
			room.NPC = &models.NPC{
				Type: selectedType,
				//Messages: g.getNPCMessages(selectedType),
			}
			g.PlacedNPCs[selectedType] = true
		}
	}

	// Génération des objets (20% de chances)
	if rand.Float32() < 0.2 {
		if rand.Float32() < 0.5 {
			room.Item = &models.Item{
				Name:        "Gold",
				Description: "bag of gold coins",
				Type:        "gold",
				Value:       rand.Intn(50) + 10,
			}
		} else {
			room.Item = &models.Item{
				Name:        "Potion",
				Description: "regeneration potion",
				Type:        "potion",
				Effect:      30,
			}
		}
	}

	// Génération des monstres (25% de chances)
	if rand.Float32() < 0.25 && room.NPC == nil {
		// choisir le type de monstre aléatoirement
		monster := config.MonsterTypes[rand.Intn(len(config.MonsterTypes))]
		monster.CurrentHP = monster.HP

		/* === MONSTER DESCRIPTIOM === */
		// BEGIN
		monster.Description = "You are facing a monster"
		// END

		room.Monster = &monster
	}

	return room
}

/*
func (g *Game) getNPCMessages(npcType models.NPCType) []string {
	switch npcType {
	case models.Merchant:
		return []string{
			"Bienvenue à ma boutique! J'ai des objets rares à vendre.",
			"Ces temps-ci, les potions de soin sont très demandées...",
			"Si vous trouvez de l'or, revenez me voir!",
			"Je peux vous faire un bon prix pour votre équipement.",
		}
	case models.Guard:
		return []string{
			"Halte! Cette zone est sous ma protection.",
			"Méfiez-vous des monstres qui rôdent dans les couloirs.",
			"J'ai entendu dire que le dragon garde un trésor...",
			"Je peux vous indiquer le chemin vers la sortie, si vous le souhaitez.",
		}
	case models.Sorcerer:
		return []string{
			"Je sens une grande magie en ces lieux...",
			"Voulez-vous que je vous enseigne quelques sorts?",
			"Les anciens secrets du château sont bien gardés.",
			"La magie peut vous aider à vaincre les monstres plus facilement.",
		}
	default:
		return []string{"..."}
	}
}
*/

func (g *Game) GetAvailableDirections() []string {
	// Toujours retourner toutes les directions possibles
	return []string{"north", "south", "west", "east"}
}

func (g *Game) Move(direction string) bool {
	dir := config.Directions[direction]
	newPos := models.Position{
		X: g.CurrentPos.X + dir.DX,
		Y: g.CurrentPos.Y + dir.DY,
	}

	// Générer une nouvelle pièce si nécessaire
	if _, exists := g.Rooms[newPos]; !exists {
		g.Rooms[newPos] = g.generateRoom()
	}

	g.PreviousPos = g.CurrentPos
	g.CurrentPos = newPos
	g.LastDirection = direction

	return g.ProcessRoom()
}

func (g *Game) ProcessRoom() bool {
	room := g.Rooms[g.CurrentPos]

	// You found the exit 🎉
	if g.CurrentPos.X == config.ExitCell.X && g.CurrentPos.Y == config.ExitCell.Y {
		ui.Println(colors.Orange, "You are almost out of the castle! 🏰")
		ui.Println(colors.Orange, "But first, you must give the 🦁 Sphinx all three colors to escape...")

		/* === END LEVEL BOSS === */

		// BEGIN
		ui.Println(colors.Orange, "🦁 Sphinx: Welcome to the final challenge, brave adventurer!")
		// END

		ui.Println(colors.Orange, "--------------------------------------------------------")
		ui.Println(colors.Orange, "Congratulations! You have reached the castle exit!")
		ui.Println(colors.Orange, "Here is your final report:")
		ui.Println(colors.Orange, fmt.Sprintf("Health points: %d/%d", g.Player.HP, g.Player.MaxHP))
		ui.Println(colors.Orange, fmt.Sprintf("Experience gained: %d", g.Player.XP))
		ui.Println(colors.Orange, fmt.Sprintf("Gold amassed: %d coins", g.Player.Gold))

		ui.Println(colors.Orange, "NPCs encountered:")

		for npcType := range g.PlacedNPCs {
			ui.Println(colors.Orange, fmt.Sprintf("- %s", string(npcType)))
		}
		ui.Println(colors.Orange, "--------------------------------------------------------")

		return true
	}

	ui.Println(colors.Pink, "--------------------------------------------------------")
	ui.Println(colors.Pink, fmt.Sprintf("You are in the room (%d, %d)", g.CurrentPos.X, g.CurrentPos.Y))
	ui.Println(colors.Orange, "--------------------------------------------------------")
	ui.Println(colors.Orange, "Room description:")
	ui.Println(colors.Orange, room.Description)
	ui.Println(colors.Orange, "--------------------------------------------------------")

	// Update the room state
	if !room.IsVisited {
		room.IsVisited = true
		g.Rooms[g.CurrentPos] = room // Sauvegarder l'état mis à jour
		g.Player.XP += 10            // Gain d'XP pour l'exploration
	}

	// Non Player Character (NPC)
	if room.NPC != nil {
		ui.Println(colors.Blue, fmt.Sprintf("You meet a %s", room.NPC.Type))
		input, _ := ui.Input(colors.Blue, "Do you want to chat? (y/n)")

		if input == "y" {
			/* === NPC CHAT === */
			// BEGIN
			ui.Println(colors.Blue, "👋 Hello World 🌍\n\n")
			// END

			g.Player.XP += 5 // XP bonus to chat with NPC
		}
		ui.Println(colors.Blue, "--------------------------------------------------------")
	}

	if room.Item != nil {
		ui.Println(colors.Orange, "--------------------------------------------------------")

		ui.Println(colors.Orange, fmt.Sprintf("You find a %s", room.Item.Description))

		if room.Item.Type == "potion" {
			ui.Println(colors.Orange, fmt.Sprintf("You use the potion and recover %d health points", room.Item.Effect))
			g.Player.HP = min(g.Player.HP+room.Item.Effect, g.Player.MaxHP)
		} else {
			g.Player.Gold += room.Item.Value
			ui.Println(colors.Orange, fmt.Sprintf("You pick up %d gold coins", room.Item.Value))
		}

		ui.Println(colors.Orange, "--------------------------------------------------------")

		g.Player.XP += 15 // XP bonus for finding an item
		room.Item = nil
		g.Rooms[g.CurrentPos] = room
	}

	if room.Monster != nil {

		ui.Println(colors.Red, fmt.Sprintf("🙀 You meet a %s", room.Monster.Name))

		ui.Println(colors.Red, room.Monster.Description)

		input, _ := ui.Input(colors.Red, "Do you want to (f)ight or (e)scape? 👀")

		// Escape
		if input == "e" {
			g.CurrentPos = g.PreviousPos
			ui.Println(colors.Red, "--------------------------------------------------------")
			return false
		}

		if !g.Combat(room.Monster) {
			ui.Println(colors.Red, "--------------------------------------------------------")
			return true
		}
		room.Monster = nil
		g.Rooms[g.CurrentPos] = room
	}
	return false
}

func (g *Game) Combat(monster *models.Monster) bool {

	ui.Println(colors.Red, "--------------------------------------------------------")
	ui.Println(colors.Red, fmt.Sprintf("Combat against %s (HP: %d, Attack: %d)", monster.Name, monster.CurrentHP, monster.AttackPower))

	for {
		// Player's turn
		damage := g.Player.Attack + rand.Intn(5) - 2
		monster.CurrentHP -= damage

		ui.Println(colors.Red, fmt.Sprintf("You inflict %d damage to the %s", damage, monster.Name))

		if monster.CurrentHP <= 0 {

			ui.Println(colors.Red, fmt.Sprintf("You have defeated the %s!", monster.Name))

			xpGained := rand.Intn(20) + 30 // Entre 30 et 49 XP
			goldReward := rand.Intn(30) + 10
			g.Player.XP += xpGained
			g.Player.Gold += goldReward

			ui.Println(colors.Red, fmt.Sprintf("You win %d XP and %d gold coins!", xpGained, goldReward))
			ui.Println(colors.Red, "--------------------------------------------------------")

			return true
		}

		// Monster's turn
		damage = monster.AttackPower + rand.Intn(3) - 1
		g.Player.HP -= damage

		ui.Println(colors.Red, fmt.Sprintf("The %s inflicts %d damage to you", monster.Name, damage))

		ui.Println(colors.Red, fmt.Sprintf("Your health points: %d/%d", g.Player.HP, g.Player.MaxHP))

		if g.Player.HP <= 0 {
			ui.Println(colors.Red, "--------------------------------------------------------")
			return false
		}

		ui.Println(colors.Red, "Press Enter to continue the fight...")
		fmt.Scanln()
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
