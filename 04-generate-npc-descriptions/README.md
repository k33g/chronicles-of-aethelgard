# Générer des descriptions de PNJ

## Que fait ce programme ?

```mermaid
graph TD
    A[Début] --> B[Initialisation<br/>Ollama Client]
    B --> C[Lecture des fichiers<br/>instructions.md & steps.md]
    C --> D[Lecture du<br/>personnage<br/>fichier JSON]
    D --> E[Construction du<br/>prompt]
    E --> F[Envoi requête à<br/>Ollama API]
    F --> G[Réception réponse<br/>en streaming]
    G --> H[Sauvegarde du<br/>character sheet]
    
```

## Allons voir le code

[Le code](main.go)

## Que font le 🐳 compose file & le Dockerfile ?

- [Le 🐳 compose file](compose.yml) ... ✋👀 **le LLM a changé**
- [Dockerfile](Dockerfile)

## Lancer l'application

```bash
docker compose up --watch
```

## 🚧 Travaillez un peu

### Créer les instructions pour aider le LLM

Allez modifier le fichier [`steps.md`](steps.md) pour expliquer au LLM ce dont vous avez besoin (en anglais 🇬🇧).

- Expliquez au LLM qu'il doit générer une fiche descriptive pour un personnage donné
- Donnez les informations dont vous avez besoin (Kind, Name, Age, Family, Occupation, Physical Appearance, Background Story, Quote, Personality, ...)
- Expliquez le format de sortie que vous souhaitez (markdown)

😈 **Soyez créatifs**


### Générer une fiche pour les 3 personnages

Modifiez 3 fois le code pour pouvoir générer 3 feuilles de personnage (il faut modifier le code Go: juste le nom du fichier JSON à charger)

> Vous pouvez jouer avec la température si les résultat ne vous conviennent pas.

## Questions ?

## Quittez Docker Compose

[README](../README.md)



- Essayer d'écrire les instructions pour générer des descriptions de PNJ.
Guider pour le faire (une sorte de template)
- Essayer avec différents modèles.
- Changer la temperature