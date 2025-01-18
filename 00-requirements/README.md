# Initialisation: Vérification des pré-requis

> Vous avez besoin de Docker sur votre machine, mais normalement si vous avez lu les [pré-requis](https://github.com/k33g/chronicles-of-aethelgard-at-snowcamp-2025) vous devez déjà l'avoir installé et vous avec déjà chargé les images nécessaires.

## Nous allons utiliser Docker Compose

- Expliquer Docker Compose
- Montrer le fichier `compose.yml`

## Téléchargement des modèles et démarrage d'ollama

Lancez:
```bash
docker compose up
```

## Essayer Ollama

Puis dans un autre terminal:
```bash
docker ps
```

Vous devriez obtenir la liste des containers "tournant" sur votre machine:
```bash
CONTAINER ID   IMAGE                                                 COMMAND                  CREATED         STATUS         PORTS                       NAMES
40ccf4427d4f   ollama/ollama:0.5.1                                   "/bin/ollama serve"      2 minutes ago   Up 2 minutes   0.0.0.0:11434->11434/tcp    00-requirements-ollama-service-1
```

Théoriquement le nom du container est `00-requirements-ollama-service-1`, donc "entrez" en mode interactif dans le container avec la commande suivante (n'arrêtez pas Docker Compose):
```bash
docker exec -it 00-requirements-ollama-service-1 bash
```

Ensuite (dans le container), lancez `ollama` en utilisant le modèle `qwen2.5:0.5b`
```bash
ollama run qwen2.5:0.5b
```

Une fois qu'Ollama est lancé, posez votre question (en anglais):
```text
Explain the role playing game
```
> 👋 bien sûr vous pouvez aussi lui demander quelle est la meilleur pizza au monde


> - Ctrl + D pour quitter Ollama
> - `exit` pour sortir du conteneur

## Vérifions que l'image Go a bien été chargée

```bash
docker compose logs golang-version
```

## Quitter Docker Compose

```bash
docker compose down
```
> ou Ctrl+C, mais c'est moins propre


[README](../README.md)