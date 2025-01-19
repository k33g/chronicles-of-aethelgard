# Utiliser les services des PNJ

## Helpers

J'ai développé quelques utilitaires: [`services/rooms.go`](services/bots.go)

Voici comment les utiliser: [`main.go`](main.go)

## Ce qui est important

Comprendre le fichier compose pour lancer les services: [Le 🐳 compose file](compose.yml) (utilisation des DNS Docker Compose)

## Lancer l'application

```bash
docker compose up --watch
```
> Et attendez un peu ⏳

## 🚧 Travaillez un peu

👋 Vous devez:
- Adapter les services [`compose.yml`](compose.yml) de bots PNJ en fonction du nom de vos bots
  - **Rappelez vous, ils utilisent les fichiers `character-sheet-*.md`**
  - Attention aux dépendances entre services
- Adapter le nom des fonctions d'appel à vos bots, variables d'environnement, ...  dans [`services/rooms.go`](services/bots.go) et [`main.go`](main.go)

## Questions ?

## Quittez Docker Compose

[README](../README.md)