# On discute avec les bots "pour de vrai"

## Helpers

J'ai développé un utilitaire: [`tools/input.go`](services/input.go)

Voici comment l' utiliser: [`main.go`](main.go)

## Ce qui est important

Comprendre le fichier compose pour lancer les services: [Le 🐳 compose file](compose.yml) (utilisation des DNS Docker Compose)

## Lancer l'application

```bash
docker compose up --watch
```
> Et attendez un peu ⏳

Dans un autre terminal, **pour discuter avec le bot**:
```bash
docker exec -it 15-chat-with-bot-services-main-1 go run main.go
```

## 🚧 Travaillez un peu

- 👋 Comme pour l'exercice précédent il faudra adapter le code, les variables, ... dans dans [`services/rooms.go`](services/bots.go) et [`main.go`](main.go) et [`compose.yml`](compose.yml).
- Faites ceci pour un seul bot PNJ
- Allez modifier la fiche du PNJ pour qu'il puisse vous donner un indice de couleurs pour répondre au Sphinx.

## Questions ?

## Quittez Docker Compose

[README](../README.md)