package screen

import (
	"fmt"
	"rpg/config"
	"rpg/models"
	"strings"

	"rpg/ui"
	"rpg/ui/colors"
)


func getMonsterSymbol(monster *models.Monster) string {
	return monster.Symbol
}

func getNPCSymbol(npc *models.NPC) string {
	return config.NPCTypesSymbols[npc.Type]
}


func DisplayLegend() {

	ui.Println(colors.Black, "--------------------------------------------------------")

	ui.Println(colors.Black, "You: 🙂")

	ui.Print(colors.Black, "NPC: ")
	for kind, symbol := range config.NPCTypesSymbols {
		ui.Print(colors.Black, fmt.Sprintf("%s %s ", symbol, kind))
	}
	ui.Println(colors.Black, "")

	ui.Print(colors.Black, "Monsters: ")
	for _, monster := range config.MonsterTypes {
		ui.Print(colors.Black, fmt.Sprintf("%s %s ", monster.Symbol, monster.Name))
	}
	ui.Println(colors.Black, "")

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
					//sb.WriteString(room.Monster.Symbol)
				} else if room.NPC != nil {
					//sb.WriteString(getNPCSymbol(room.NPC) + " ")
					sb.WriteString(getNPCSymbol(room.NPC))
					//sb.WriteString(room.NPC.Symbol)
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
