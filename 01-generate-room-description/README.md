# Générer des descriptions de salles


<!-- TODO: à re écrire -->

```bash
docker compose up --watch
```



Il y a plusieurs façons de lancer le programme.

## Si Ollama et Go sont installés en local

dans le fichier `.env`, vous devez avoir:

```bash
OLLAMA_HOST=http://localhost:11434
LLM=qwen2.5:0.5b
```

Et vous lancez le programme avec:
```bash
go run main.go
```

## Si Ollama et Go sont installés dans des containers

dans le fichier `.env`, vous devez avoir:

```bash
OLLAMA_HOST=http://ollama-service:11434
LLM=qwen2.5:0.5b
```

Et vous lancez le programme avec:
```bash
docker compose up
```

Vous pouvez voir les résultats avec:
```bash
docker compose logs generate-description
```

## Si Ollama est installé en local et Go dans un container

Il est possible de "contacter" Ollama avec les paramètres suivants:

```bash
OLLAMA_HOST=http://host.docker.internal:11434
LLM=qwen2.5:0.5b
```

## Autre option : Devcontainer

<🚧 work in progress>


## Prompting

- Essayer d'autres prompts
- Jouer aussi avec les settings (temperature)