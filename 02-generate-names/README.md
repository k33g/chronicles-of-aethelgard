# Générer des noms de personnages

## Que voulons nous faire ?

- Générer des noms de personnages de JdR "dans un style médiéval"
- Obtenir un format de sortie en JSON
- Sauver le résultat

## Allons voir le code

> - 👋 il ressemble beaucoup au code précédent
> - mais il faut que l'on explique au LLM de façon précise que l'on veut du JSON, structuré d'une manière spécifique

[Le code](main.go)

## Que font le 🐳 compose file & le Dockerfile ?

- [Le 🐳 compose file](compose.yml)
- [Dockerfile](Dockerfile)

## Lancer l'application

```bash
docker compose up --watch
```

Vous pouvez voir les résultats avec:
```bash
docker compose logs generate-names
```

## 🚧 Travaillez un peu

- 👋 jouez avec la température
- génerez plusieurs personnages de plusieurs races

## Autre méthode pour générer du JSON

La méthode que je viens de vous présenter n'est pas fiable à 100%, mais il existe depuis peu une autre méthode pour générer du JSON avec un LLM.

On parle de **Structured Output**. 
> ✋ cela ne fonctionnera pas avec tous les modèles

- Blog post officiel: https://ollama.com/blog/structured-outputs
- Blog post écrit par votre serviteur: https://k33g.hashnode.dev/generating-json-with-an-llm-the-old-method-and-the-new-method 

## Questions ?

## Quittez Docker Compose

[README](../README.md)