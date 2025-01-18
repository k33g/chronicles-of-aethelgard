# Bot NPC + RAG
> Retrieval Augmented Generation

Cet exemple est à titre éducatif. Nous ne l'utiliserons pas pour le "jeu" définitif.

Mais cette technique peut être utile avec des petits LLM qui ne savent pas utiliser correctement des "gros contextes".

Cette fois nous allons travailler avec 2 modèles:

- `qwen2.5:0.5b`
- `snowflake-arctic-embed:33m`

Le 2ème LLM sert à générer des embeddings 😮🤔

- [Embeddings ?](docs/01-embeddings.md)
- [RAG ?](docs/02-rag.drawio)
- [Calcul de distance ?](docs/03-distance.md)

Nous avons donc 2 programmes:

## `chunking_test.go` pour générer les chunks

```mermaid
graph TD
    A[Start Test] --> B[Initialize Environment]
    B --> C[Create Ollama Client]
    C --> D[Load Character Info]
    D --> E[Read Character Sheet]
    
    subgraph embedding[Embedding Generation]
        E --> F[Split Content into Chunks]
        F --> G[Initialize Vector Store]
        G --> H[Process Each Chunk]
        H --> |For each chunk| I[Generate Embedding]
        I --> J[Create Vector Record]
        J --> K[Add to Vector Store]
        K --> |Next Chunk| H
    end
    
    subgraph storage[Store Persistence]
        K --> L[Convert Store to JSON]
        L --> M[Generate Store Filename]
        M --> N[Write JSON to File]
    end
    
    N --> O[End Test]
    
    style A fill:#f9f
    style O fill:#f9f
    style embedding fill:#ddf
    style storage fill:#dfd
```

**[Le code](chunking_test.go)**

## `main.go` pour la complétion avec recherche de similarité

```mermaid
graph TD
    A[Start] --> B[Initialize Environment]
    B --> C[Create Ollama Client]
    C --> D[Load Character]
    D --> E[Load Vector Store from JSON]
    
    subgraph server[HTTP Server]
        E --> F[Initialize Memory Array]
        F --> G[Setup POST /api/chat Endpoint]
    end
    
    subgraph request[Request Processing]
        G --> H[Receive Question]
        H --> I[Generate Question Embedding]
        
        subgraph similarity[Similarity Search]
            I --> J[Calculate Cosine Similarities]
            J --> K[Sort Similarities]
            K --> L[Select Top 5 Similar Chunks]
            L --> M[Create New Context]
        end
        
        subgraph chat[Chat Generation]
            M --> N[Build Messages Array]
            N --> O[Add System Instructions]
            O --> P[Add Memory]
            P --> Q[Add Current Question]
            Q --> R[Configure LLM]
            R --> S[Stream Response]
        end
        
        S --> T[Save to Memory]
        T --> U[Ready for Next Request]
    end
    
    style A fill:#f9f
    style server fill:#dfd
    style similarity fill:#fdf
    style chat fill:#ddf
```

**[Le code](main.go)**

## Que font le 🐳 compose file & le Dockerfile ?

- [Le 🐳 compose file](compose.yml) 
- [Dockerfile](Dockerfile)

## Lancer l'application

```bash
docker compose up --watch
```


## 🚧 Travaillez un peu

- On va juste se contenter d'exécuter les codes
- Jetez un coup d'oeil à [`store-grym.json`](store-grym.json)



## Testez les services (au moins un des services)

### Avec curl

- `query-1.sh`, ensuite, essayez `docker compose logs bot-with-memory`
- `query-2.sh`, ensuite, essayez `docker compose logs bot-with-memory`
- `query-3.sh`, ensuite, essayez `docker compose logs bot-with-memory`
- `query-4.sh`, ensuite, essayez `docker compose logs bot-with-memory`

> Bien sûr, adaptez les requête (numéro de port HTTP par exemple)

### Si vous n'avez pas curl

```bash
docker run --rm --network host curlimages/curl:8.6.0 \
    --silent --no-buffer "http://localhost:5051/api/chat" \
    -H "Content-Type: application/json" \
    -d '{"question":"What is your name?"}'
```

etc ...

## Questions ?

## Quittez Docker Compose

[README](../README.md)