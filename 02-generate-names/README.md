# Générer des noms de personnages

## Objectif

- Générer des noms de personnages avec Ollama,
- Obtenir une format de sortie en JSON


## Aller lire le programme


## Lancer le programme

Et vous lancez le programme avec:
```bash
docker compose up
```

Vous pouvez voir les résultats avec:
```bash
docker compose logs generate-names
```

> ✋ pour les autres méthodes de lancement, voir le README de l'étape précédente.

## Il existe maintenant une autre méthode pour générer du JSON


## Remarques

le format JSON ne va pas forcément fonctionner avec tous les modèles (structured output)

- https://ollama.com/blog/structured-outputs
- https://k33g.hashnode.dev/generating-json-with-an-llm-the-old-method-and-the-new-method 

## Jouez avec les paramètres

- 👋 jouer avec la température
- génerer plusieurs personnages de plusieurs races
