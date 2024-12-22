package screen

import (
	"fmt"
	"rpg-game/models"
	"strings"

	"github.com/k33g/chronicles-of-aethelgard/ui"
	"github.com/k33g/chronicles-of-aethelgard/ui/colors"
)

func getMonsterSymbol(monster *models.Monster) string {
	switch monster.Name {
	case "Gobelin":
		return "👺"
	case "Troll":
		return "👹"
	case "Dragon":
		return "🐲"
	case "Wolf":
		return "🐺"
	case "Bear":
		return "🐻"
	default:
		return "M"
	}
}



func getNPCSymbol(npc *models.NPC) string {
	switch npc.Type {
	case models.Merchant:
		return "🤗"
	case models.Guard:
		return "🤠"
	case models.Sorcerer:
		return "🎃"
	default:
		return "n"
	}
}



func DisplayLegend() {

	ui.Println(colors.Black, "--------------------------------------------------------")

	ui.Println(colors.Black,"You: 🙂")
	ui.Println(colors.Black,"NPC: 🤗 merchant 🤠 guard 🎃 sorcerer")
	ui.Println(colors.Black,"Monsters: 👺 Gobelin 👹 Troll 🐲 Dragon 🐺 Wolf 🐻 Bear")

	ui.Println(colors.Black, "--------------------------------------------------------")


}

func DisplayMap(currentPos models.Position, rooms map[models.Position]models.Room) {
	ui.Println(colors.Black, "--------------------------------------------------------")

	minX, maxX, minY, maxY := 0, 0, 0, 0
	for pos := range rooms {
		if pos.X < minX {
			minX = pos.X
		}
		if pos.X > maxX {
			maxX = pos.X
		}
		if pos.Y < minY {
			minY = pos.Y
		}
		if pos.Y > maxY {
			maxY = pos.Y
		}
	}

	minX--
	maxX++
	minY--
	maxY++

	var sb strings.Builder

	// Display map
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			pos := models.Position{X: x, Y: y}
			if pos == currentPos {
				//sb.WriteString("P ")
				sb.WriteString("🙂")
			} else if room, exists := rooms[pos]; exists {
				if room.Monster != nil {
					//sb.WriteString(getMonsterSymbol(room.Monster) + " ")
					sb.WriteString(getMonsterSymbol(room.Monster))
				} else if room.NPC != nil {
					//sb.WriteString(getNPCSymbol(room.NPC) + " ")
					sb.WriteString(getNPCSymbol(room.NPC))
				} else if room.IsVisited {
					//sb.WriteString("# ")
					sb.WriteString("⬜️")
				} else {
					//sb.WriteString("? ") // Case découverte mais non visitée
					sb.WriteString("⬜️")
				}
			} else {
				//sb.WriteString(". ") // Case non découverte
				sb.WriteString("⬛️")
			}
		}
		sb.WriteString("\n")
	}
	fmt.Print(sb.String())

	DisplayLegend()
}

func DisplayStatus(player models.Player) {

	ui.Println(colors.Purple, "--------------------------------------------------------")
	ui.Println(colors.Purple, fmt.Sprintf("Player's status (%s):", player.Race))
	ui.Println(colors.Purple, "--------------------------------------------------------")

	ui.Println(colors.Purple, fmt.Sprintf("💚 Health points: %d/%d", player.HP, player.MaxHP))
	ui.Println(colors.Purple, fmt.Sprintf("💪 Attack strength: %d", player.Attack))
	ui.Println(colors.Purple, fmt.Sprintf("🤓 Experience: %d", player.XP))
	ui.Println(colors.Purple, fmt.Sprintf("⭐️ Gold: %d coins", player.Gold))

}


