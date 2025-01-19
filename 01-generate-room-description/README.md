# Générer des descriptions de salles

## Que fais le code ?

```mermaid
graph TD
    A[Start] --> B[Set Context & Get Env Variables]
    B --> C[Initialize Ollama Client]
    
    subgraph prompt[Prompt Construction]
        D1[Define System Instructions<br/>You are a Generator] 
        D2[Define Generation Instructions<br/>Room Description Rules]
        D3[Set User Content<br/>QUESTION]
        D1 & D2 & D3 --> E[Create Messages Array]
    end
    
    subgraph config[Chat Configuration]
        F1[Set Stream Mode: true]
        F2[Set Parameters:<br/>temperature: 1.8<br/>repeat_penalty: 2.2<br/>etc...]
        F1 & F2 --> G[Create Chat Request]
    end
    
    C --> prompt
    prompt --> config
    G --> H[Start Chat Completion]
    H --> I[Stream Response to Console]
    I --> J[End Program]
    
    style A fill:#f9f,stroke:#333
    style J fill:#f9f,stroke:#333
    style prompt fill:#eef,stroke:#333
    style config fill:#efe,stroke:#333
```

## Allons voir le code

[Le code](main.go)

## Que font le 🐳 compose file & le Dockerfile ?

- [Le 🐳 compose file](compose.yml) ... C'est quoi ce `watch` ?
- [Dockerfile](Dockerfile)

## Lancer l'application

```bash
docker compose up --watch
```
> Et attendez un peu ⏳ ... qur tous les services soient démarrés.

## 🚧 Travaillez un peu

- Essayez avec d'autre noms de pièces : `userContent`
- Vous pouvez modifier les instructions : `systemInstructions` & `generationInstructions`
- Jouer aussi avec les settings (en fait uniquement la `temperature`) ... 🤔 mais pourquoi ? *(Explications à donner)*

## Questions ?

## Quittez Docker Compose

[README](../README.md)