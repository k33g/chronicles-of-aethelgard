# Room service

## Objectif

Faire un service qui permet d'obtenir une description de pièce à partir d'un nom.

## Que fait ce programme ?

```mermaid
graph TD
    Start[Start Program] --> Init[Initialize Context & Environment]
    Init --> Client[Create Ollama Client]
    Client --> Server[Setup HTTP Server]
    
    subgraph endpoints[HTTP Endpoints]
        Server --> EP1[GET /api/room/generate/name]
        Server --> EP2[POST /api/room/generate/description]
        
        EP1 --> F1[Setup Stream Response]
        F1 --> G1[Create Name Generation Prompt]
        G1 --> H1[Configure LLM Parameters<br/>temperature: 2.0]
        H1 --> I1[Stream Response<br/>Random Room Name]
        
        EP2 --> F2[Setup Stream Response]
        F2 --> G2[Parse JSON Request<br/>room_name]
        G2 --> H2[Create Description Prompt]
        H2 --> I2[Configure LLM Parameters<br/>temperature: 0.8]
        I2 --> J2[Stream Response<br/>Room Description]
    end
    
    Server --> Listen[Listen on Port]
    Listen --> End[Handle Server Errors]
    
    style Start fill:#f9f
    style End fill:#f99
    style endpoints fill:#eef
    style Listen fill:#9f9
    
```

## Allons voir le code

[Le code](main.go)

## Que font le 🐳 compose file & le Dockerfile ?

- [Le 🐳 compose file](compose.yml) 
- [Dockerfile](Dockerfile)

## Lancer l'application

```bash
docker compose up --watch
```

## 🚧 Travaillez un peu

- Ré-utilisez le code de `01-generate-room-description` pour définir le prompt (instructions, questions, ...) qui seront utilisés par le end-point `POST /api/room/generate/description`
  - `systemInstructions`
  - `generationInstructions`


## Testez les services

### Avec curl

- `query-room-name.sh`
- `query-room-description.sh`


### Si vous n'avez pas curl

```bash
docker run --rm --network host curlimages/curl:8.6.0 \
    --silent --no-buffer "http://localhost:5050/api/room/generate/name" 

docker run --rm --network host curlimages/curl:8.6.0 \
    --silent --no-buffer "http://localhost:5050/api/room/generate/description" \
    -H "Content-Type: application/json" \
    -d '{"room_name":"Minion Lair"}'
```

## Questions ?

## Quittez Docker Compose

[README](../README.md)


