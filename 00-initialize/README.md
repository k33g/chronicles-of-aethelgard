# Initialisation

## Volume partagé

```bash
docker volume create ollama_shared_data
```
Ce volume sera utilisé par tous les autres projets Docker Compose

## Téléchargement des modèles

```bash
docker compose up
```

## Essayer Ollama

```bash
docker ps
```

```bash
CONTAINER ID   IMAGE                                                 COMMAND                  CREATED         STATUS         PORTS                       NAMES
40ccf4427d4f   ollama/ollama:0.5.1                                   "/bin/ollama serve"      2 minutes ago   Up 2 minutes   0.0.0.0:11434->11434/tcp    00-initialize-ollama-service-1
```


```bash
docker exec -it 00-initialize-ollama-service-1 bash
# Then
ollama run qwen2.5:0.5b
describe a room in a dungeon
```
> Ctrl + D pour quitter Ollama

```bash
docker exec -it 00-initialize-ollama-service-1 bash
# Then
ollama run qwen2.5:1.5b
describe a room in a dungeon
```

> Ctrl + D puis `exit` pour sortir du conteneur


```bash
docker compose logs golang-version
```

Quitter Docker Compose