package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Position struct {
	x, y int
}

type Direction struct {
	dx, dy int
	name   string
}

type Item struct {
	name        string
	description string
	effect      int // Pour les potions : points de vie restaurés
}

type Monster struct {
	name          string
	hp            int
	attackPower   int
	currentHP     int
}

type NPC struct {
	name     string
	messages []string
}

type Room struct {
	description    string
	availableDirs  []string
	npc           *NPC
	item          *Item
	monster       *Monster
	isAccessible  bool
}

type Player struct {
	hp        int
	maxHP     int
	inventory []Item
	attack    int
}

type Game struct {
	currentPos    Position
	previousPos   Position
	visitedCells  map[Position]bool
	rooms         map[Position]Room
	directions    map[string]Direction
	lastDirection string
	player        Player
	npcs          map[string]NPC
	monsters      []Monster
}

func NewGame() *Game {
	directions := map[string]Direction{
		"haut":    {0, -1, "haut"},
		"bas":     {0, 1, "bas"},
		"gauche":  {-1, 0, "gauche"},
		"droite":  {1, 0, "droite"},
	}

	// Création des NPCs
	npcs := map[string]NPC{
		"A": {name: "Sage", messages: []string{
			"Bienvenue aventurier!",
			"Méfie-toi des monstres qui rôdent...",
			"As-tu trouvé des trésors?",
		}},
		"B": {name: "Marchand", messages: []string{
			"J'ai des objets rares à vendre!",
			"Revenez me voir plus tard...",
			"Les potions sont très utiles par ici.",
		}},
		"C": {name: "Guerrier", messages: []string{
			"Je peux t'apprendre à combattre!",
			"Les gobelins sont faibles mais nombreux.",
			"Garde toujours une potion avec toi.",
		}},
	}

	// Création des types de monstres
	monsters := []Monster{
		{name: "Gobelin", hp: 20, attackPower: 5},
		{name: "Troll", hp: 40, attackPower: 8},
		{name: "Squelette", hp: 25, attackPower: 6},
		{name: "Loup", hp: 30, attackPower: 7},
		{name: "Bandit", hp: 35, attackPower: 6},
	}

	return &Game{
		currentPos:   Position{0, 0},
		visitedCells: make(map[Position]bool),
		rooms:        make(map[Position]Room),
		directions:   directions,
		player: Player{
			hp:        100,
			maxHP:     100,
			inventory: []Item{},
			attack:    10,
		},
		npcs:     npcs,
		monsters: monsters,
	}
}

func (g *Game) generateRoom(pos Position) Room {
	descriptions := []string{
		"Une pièce sombre avec des toiles d'araignées",
		"Une salle illuminée par des torches",
		"Un couloir humide aux murs de pierre",
		"Une ancienne bibliothèque poussiéreuse",
		"Une salle au sol couvert de mousse",
	}

	items := []Item{
		{name: "Or", description: "Un sac de pièces d'or", effect: 0},
		{name: "Potion", description: "Une potion de régénération", effect: 30},
	}

	room := Room{
		description:   descriptions[rand.Intn(len(descriptions))],
		isAccessible:  rand.Float32() > 0.2, // 20% de chances d'être inaccessible
	}

	if room.isAccessible {
		// Génération aléatoire du contenu de la pièce
		if rand.Float32() < 0.3 { // 30% de chances d'avoir un NPC
			npcKeys := []string{"A", "B", "C"}
			npc := g.npcs[npcKeys[rand.Intn(len(npcKeys))]]
			room.npc = &npc
		}

		if rand.Float32() < 0.2 { // 20% de chances d'avoir un item
			item := items[rand.Intn(len(items))]
			room.item = &item
		}

		if rand.Float32() < 0.25 { // 25% de chances d'avoir un monstre
			monster := g.monsters[rand.Intn(len(g.monsters))]
			monster.currentHP = monster.hp // Réinitialiser les HP du monstre
			room.monster = &monster
		}
	}

	return room
}

func (g *Game) getAvailableDirections() []string {
	if !g.rooms[g.currentPos].isAccessible {
		return []string{}
	}

	// Si la pièce a déjà été visitée, retourner les directions déjà définies
	if room, exists := g.rooms[g.currentPos]; exists && len(room.availableDirs) > 0 {
		return room.availableDirs
	}

	available := []string{}
	
	// Direction de retour
	var returnDir string
	if g.lastDirection != "" {
		switch g.lastDirection {
		case "haut":
			returnDir = "bas"
		case "bas":
			returnDir = "haut"
		case "gauche":
			returnDir = "droite"
		case "droite":
			returnDir = "gauche"
		}
		available = append(available, returnDir)
	}

	// Pour la position initiale (0,0), toutes les directions sont disponibles
	if g.currentPos.x == 0 && g.currentPos.y == 0 && g.lastDirection == "" {
		return []string{"haut", "bas", "gauche", "droite"}
	}

	// Liste des directions possibles sauf le retour
	remainingDirs := []string{}
	for _, dir := range []string{"haut", "bas", "gauche", "droite"} {
		if dir != returnDir {
			remainingDirs = append(remainingDirs, dir)
		}
	}

	// Mélanger les directions restantes
	rand.Shuffle(len(remainingDirs), func(i, j int) {
		remainingDirs[i], remainingDirs[j] = remainingDirs[j], remainingDirs[i]
	})

	// Ajouter au moins une direction supplémentaire, et possiblement plus
	numExtra := rand.Intn(2) + 1 // 1 ou 2 directions supplémentaires
	for i := 0; i < numExtra && i < len(remainingDirs); i++ {
		available = append(available, remainingDirs[i])
	}

	// S'assurer qu'il y a toujours au moins une direction en plus du retour
	if len(available) < 2 && len(remainingDirs) > 0 {
		available = append(available, remainingDirs[0])
	}

	// Sauvegarder les directions disponibles pour cette pièce
	if room, exists := g.rooms[g.currentPos]; exists {
		room.availableDirs = available
		g.rooms[g.currentPos] = room
	}

	return available
}

func (g *Game) combat(monster *Monster) bool {
	fmt.Printf("\nCombat contre %s (HP: %d, Attaque: %d)\n", monster.name, monster.currentHP, monster.attackPower)
	
	for {
		// Tour du joueur
		damage := g.player.attack + rand.Intn(5) - 2 // Variation de ±2
		monster.currentHP -= damage
		fmt.Printf("Vous infligez %d dégâts au %s\n", damage, monster.name)
		
		if monster.currentHP <= 0 {
			fmt.Printf("Vous avez vaincu le %s!\n", monster.name)
			return true
		}
		
		// Tour du monstre
		damage = monster.attackPower + rand.Intn(3) - 1 // Variation de ±1
		g.player.hp -= damage
		fmt.Printf("Le %s vous inflige %d dégâts\n", monster.name, damage)
		fmt.Printf("Vos points de vie: %d/%d\n", g.player.hp, g.player.maxHP)
		
		if g.player.hp <= 0 {
			fmt.Println("Vous êtes mort!")
			return false
		}

		fmt.Print("Appuyez sur Entrée pour continuer le combat...")
		fmt.Scanln()
	}
}

func (g *Game) displayMap() string {
	minX, maxX, minY, maxY := 0, 0, 0, 0
	for pos := range g.rooms {
		if pos.x < minX {
			minX = pos.x
		}
		if pos.x > maxX {
			maxX = pos.x
		}
		if pos.y < minY {
			minY = pos.y
		}
		if pos.y > maxY {
			maxY = pos.y
		}
	}

	minX--
	maxX++
	minY--
	maxY++

	var sb strings.Builder
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			pos := Position{x, y}
			if pos == g.currentPos {
				sb.WriteString("P ")
			} else if room, exists := g.rooms[pos]; exists {
				if !room.isAccessible {
					sb.WriteString("X ")
				} else if room.npc != nil {
					sb.WriteString("N ")
				} else if room.monster != nil {
					sb.WriteString("M ")
				} else {
					sb.WriteString("# ")
				}
			} else {
				sb.WriteString(". ")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (g *Game) displayStatus() {
	fmt.Printf("\nPoints de vie: %d/%d\n", g.player.hp, g.player.maxHP)
	fmt.Println("\nInventaire:")
	if len(g.player.inventory) == 0 {
		fmt.Println("Vide")
	} else {
		for _, item := range g.player.inventory {
			fmt.Printf("- %s (%s)\n", item.name, item.description)
		}
	}
}

func (g *Game) move(direction string) bool {
	dir := g.directions[direction]
	newPos := Position{
		x: g.currentPos.x + dir.dx,
		y: g.currentPos.y + dir.dy,
	}

	// Générer une nouvelle pièce si nécessaire
	if _, exists := g.rooms[newPos]; !exists {
		g.rooms[newPos] = g.generateRoom(newPos)
	}

	// Vérifier si la pièce est accessible
	if !g.rooms[newPos].isAccessible {
		fmt.Println("Cette direction est bloquée!")
		return false
	}

	g.previousPos = g.currentPos
	g.currentPos = newPos
	g.lastDirection = direction

	room := g.rooms[g.currentPos]
	
	// Afficher la description de la pièce
	fmt.Printf("\nVous êtes dans la pièce (%d, %d)\n", g.currentPos.x, g.currentPos.y)
	fmt.Println(room.description)

	// Gérer la rencontre avec un NPC
	if room.npc != nil {
		fmt.Printf("\nVous rencontrez %s!\n", room.npc.name)
		fmt.Print("Voulez-vous discuter? (o/n) ")
		var input string
		fmt.Scanln(&input)
		if input == "o" {
			fmt.Println(room.npc.messages[rand.Intn(len(room.npc.messages))])
		}
	}

	// Gérer la découverte d'un item
	if room.item != nil {
		fmt.Printf("\nVous trouvez %s!\n", room.item.description)
		g.player.inventory = append(g.player.inventory, *room.item)
		if room.item.name == "Potion" {
			fmt.Printf("Vous utilisez la potion et récupérez %d points de vie\n", room.item.effect)
			g.player.hp = min(g.player.hp+room.item.effect, g.player.maxHP)
		}
		room.item = nil // L'item a été ramassé
	}

	// Gérer la rencontre avec un monstre
	if room.monster != nil {
		fmt.Printf("\nVous rencontrez un %s!\n", room.monster.name)
		if !g.combat(room.monster) {
			return true // Fin du jeu si le joueur meurt
		}
		room.monster = nil // Le monstre a été vaincu
	}

	// Marquer la case comme visitée
	if !g.visitedCells[newPos] {
		g.visitedCells[newPos] = true
	}

	return len(g.visitedCells) >= 10
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	rand.Seed(time.Now().UnixNano())
	game := NewGame()
	gameOver := false

	fmt.Println("Bienvenue dans le jeu d'exploration!")
	fmt.Println("Commandes disponibles:")
	fmt.Println("- map : afficher la carte")
	fmt.Println("- status : afficher votre état")
	fmt.Println("- une direction (haut/bas/gauche/droite)")
	fmt.Println("\nLégende de la carte:")
	fmt.Println("P: Votre position")
	fmt.Println("N: PNJ")
	fmt.Println("M: Monstre")
	fmt.Println("#: Case visitée")
	fmt.Println("X: Case inaccessible")
	fmt.Println(".: Case inexplorée")

	game.rooms[Position{0, 0}] = game.generateRoom(Position{0, 0})
	fmt.Print(game.displayMap())

	for !gameOver && game.player.hp > 0 {
		availableDirections := game.getAvailableDirections()
		fmt.Printf("\nDirections disponibles: %v\n", availableDirections)
		fmt.Print("Que voulez-vous faire ? ")

		var input string
		fmt.Scanln(&input)

		switch input {
		case "map":
			fmt.Print(game.displayMap())
		case "status":
			game.displayStatus()
		default:
			validMove := false
			for _, dir := range availableDirections {
				if input == dir {
					validMove = true
					gameOver = game.move(dir)
					break
				}
			}
			if !validMove {
				fmt.Println("Action invalide!")
				continue
			}
		}

		if gameOver {
			if game.player.hp <= 0 {
				fmt.Println("\nGame Over! Vous êtes mort...")
			} else {
				fmt.Println("\nFélicitations! Vous avez exploré 10 salles différentes!")
				game.displayStatus()
			}
		}
	}
}