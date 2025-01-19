# On discute avec le Boss de fin de niveau: le Sphinx

## Helpers

J'ai développé le helper qui va bien: [`services/boss.go`](services/boss.go)

Voici comment l' utiliser (avec **la gestion de sortie** du Donjon): [`main.go`](main.go)

## Ce qui est important

Comprendre le fichier compose pour lancer les services: [Le 🐳 compose file](compose.yml) 

## Lancer l'application

```bash
docker compose up --watch
```
> Et attendez un peu ⏳

Dans un autre terminal, **pour discuter avec le bot**:
```bash
docker exec -it 16-chat-with-boss-main-1 go run main.go
```

## Tester le bot de fin de niveau

- How can I escape from here?
- I want to escape with this magic words: yellow black and green
- I want to escape with this magic words: pink blue orange

## Questions ?

## Quittez Docker Compose

[README](../README.md)