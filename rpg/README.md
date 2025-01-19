# RPG

## Présentation

```raw
--------------------------------------------------------
⬛️⬛️⬛️⬛️⬛️⬛️
⬛️👺⬛️⬛️⬛️⬛️
⬛️⬜️🤠⬜️🤩⬛️
⬛️👹⬛️⬛️🙂⬛️
⬛️⬛️⬛️⬛️⬛️⬛️
--------------------------------------------------------
You: 🙂
NPC: 🤠 guard 😈 sorcerer 🤩 merchant 
Monsters: 💀 Skeleton 👹 Troll 🐲 Dragon 🐺 Werewolf 👺 Gobelin 
```

## Structure du jeu

- Configuration: [`config/game_config.go`](config/game_config.go)

## 🚧 Travaillez un peu

### Adapter l'utilisation des bots

- Dans le code [`services/bots.go`](services/bots.go) *(nommage, variables d'environnement, ...)*
- Dans le Compose file [`compose.yml`](compose.yml) *(nommage, variables d'environnement, ...)*

### Ajouter le descriptif des salles du Donjon

- Cherchez: `/* === ROOM DESCRIPTIOM === */` dans [`game/game.go`](game/game.go)
- Ajoutez le code

### Ajouter la description des monstres

- Cherchez: `/* === MONSTER DESCRIPTIOM === */` dans [`game/game.go`](game/game.go)
- Ajoutez le code
 
### Ajouter la possibilité de discuter avec les PNJ (NPC)

- Cherchez: `/* === NPC CHAT === */` dans [`game/game.go`](game/game.go)
- Ajoutez le code

### Ajouter la possibilité de discuter avec le Boss 
> pour q'il vous laisse partir

- Cherchez: `/* === END LEVEL BOSS === */` dans [`game/game.go`](game/game.go)
- Ajoutez le code

## Lancez le jeu

Dans un terminal:
```bash
docker compose up --watch
```

### Jouez 

Dans un autre terminal:
```bash
docker exec -it rpg-game-1 go run main.go
```

## Questions ?

## Quittez Docker Compose

[README](../README.md)