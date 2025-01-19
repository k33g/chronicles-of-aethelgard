# Utiliser les "room services"

## Helpers

J'ai développé quelques utilitaires: [`services/rooms.go`](services/rooms.go)

Voici comment les utiliser: [`main.go`](main.go)

## Ce qui est important

Comprendre le fichier compose pour lancer les services: [Le 🐳 compose file](compose.yml) (utilisation des DNS Docker Compose)

```mermaid
graph TD
    subgraph docker-compose
        O[ollama-service<br/>image: k33g/chronicles-of-aethelgard:0.0.2<br/>port: 11434]
        
        subgraph room[room-services]
            R[room-services<br/>Custom Dockerfile]
            RE[Environment:<br/>OLLAMA_HOST=http://ollama-service:11434<br/>LLM=qwen2.5:1.5b]
            RV[Volume:<br/>/05-room-services/:/app]
        end
        
        subgraph main[main-service]
            M[main<br/>Custom Dockerfile]
            ME[Environment:<br/>ROOM_SERVICES_HOST=http://room-services:8080]
            MW[Watch:<br/>./main.go]
            MV[Volume:<br/>./:/app]
        end
        
        O --> R
        R --> M
        
        style O fill:#f9d,stroke:#333
        style R fill:#9df,stroke:#333
        style M fill:#df9,stroke:#333
    end
```


## Lancer l'application

```bash
docker compose up --watch
```
> Et attendez un peu ⏳


## Questions ?

## Quittez Docker Compose

[README](../README.md)