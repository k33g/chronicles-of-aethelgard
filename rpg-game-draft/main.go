package main

import (
	"fmt"
	"rpg-game/game"
	"rpg-game/models"
	"rpg-game/screen"
	"strings"

	"github.com/k33g/chronicles-of-aethelgard/ui"
	"github.com/k33g/chronicles-of-aethelgard/ui/colors"
)

func chooseRace() models.Race {
	ui.Println(colors.Purple, "Choose a character!")

	ui.Println(colors.Purple, "1. Human (HP: 100, Attack: 10)")
	ui.Println(colors.Purple, "2. Elf (HP: 80, Attack: 12)")
	ui.Println(colors.Purple, "3. Dwarf (HP: 120, Attack: 8)")
	ui.Println(colors.Purple, "4. Magician (HP: 60, Attack: 15)")

	for {
		choice, _ := ui.Input(colors.Purple, "Your choice (1-4)? ")

		switch strings.TrimSpace(choice) {
		case "1":
			return models.Human
		case "2":
			return models.Elf
		case "3":
			return models.Dwarf
		case "4":
			return models.Wizard
		case "0":
			return models.Nothing
		default:
			ui.Println(colors.Red, "Bad choice! 😡 Please, try again!")
		}
	}
}

/*
rpg-game/
├── main.go
├── models/
│   └── types.go
├── config/
│   └── game_config.go
├── game/
│   └── game.go
└── ui/

	└── display.go
*/
func main() {
	//rand.Seed(time.Now().UnixNano())

	ui.Println(colors.Purple, "-------------------------------")
	ui.Println(colors.Purple, "  🏰 Chronicles of Aethelgard")
	ui.Println(colors.Purple, "-------------------------------")

	race := chooseRace()

	if race == models.Nothing {
		ui.Println(colors.Blue, "👋 Bye, thank you! 🙂")
		return
	}

	ui.Println(colors.Magenta, fmt.Sprintf("You choose %s!🔥", race))

	gameInstance := game.NewGame(race)
	//screen.DisplayLegend()
	//screen.DisplayCommands()

	gameOver := false
	for !gameOver {
		//screen.DisplayMap(gameInstance.CurrentPos, gameInstance.Rooms)

		input, _ := ui.Input(colors.Green, fmt.Sprintf("🤖 ->[%d,%d] [(n)orth/(s)outh/(w)est/(e)ast] | [(m)ap/s(t)atus/(r)ules/(q)uit] ? ", gameInstance.CurrentPos.X, gameInstance.CurrentPos.Y))

		input = strings.ToLower(strings.TrimSpace(input))
		switch input {
		case "m":
			screen.DisplayMap(gameInstance.CurrentPos, gameInstance.Rooms)
		case "t":
			screen.DisplayStatus(gameInstance.Player)
		case "r":
			// TODO:
			/*
				🖐️🤖 HERE, CHAT WITH GM
			*/
		case "q":
			ui.Println(colors.Blue, "👋 Bye, thank you! 🙂")
			return
		case "n":
			gameOver = gameInstance.Move("north")
		case "s":
			gameOver = gameInstance.Move("south")
		case "w":
			gameOver = gameInstance.Move("west")
		case "e":
			gameOver = gameInstance.Move("east")
		default:
			ui.Println(colors.Red, "Bad command!")
		}

		if gameOver {
			if gameInstance.Player.HP <= 0 {
				ui.Println(colors.Red,"Game Over! You're dead! 💀...")
			} else {
				ui.Println(colors.Blue,"🎉 Congratulations! You found the exit! 👏")
				//screen.DisplayStatus(gameInstance.Player)
			}
		}
	}
}
