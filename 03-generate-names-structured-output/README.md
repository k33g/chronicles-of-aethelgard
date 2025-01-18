# Générer des noms de personnages / Version alternative

## Principe

Avec cette nouvelle méthode, vous pouvez définir le format précis de ce que vous attendez:

```golang
schema := map[string]any{
    "type": "object",
    "properties": map[string]any{
        "name": map[string]any{
            "type": "string",
        },
        "kind": map[string]any{
            "type": "string",
        },
    },
    "required": []string{"name", "kind"},
}
```
> pour une sortie de ce type: `{"name": "John Doe","kind": "Human"}`

Ensuite cette structure/format sera fournie au LLM de la façon suivante :

```golang
req := &api.ChatRequest{
    Model:    model,
    Messages: messages,
    Options: map[string]interface{}{
        "temperature":    0.0, // 🤔
        "repeat_last_n":  2,
        "repeat_penalty": 1.8,
        "top_k":          10,
        "top_p":          0.5,
    },
    Format:    json.RawMessage(jsonModel), // ✋✋✋
    Stream:    &noStream,
}
```


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

## 🚧 Travaillez un peu

Il nous faut 3 personnages
- un(e) Elf -> Elf
- un(e) Nain(e) -> Dwarf
- un(e) Humain(e) -> Human

Renommez chaque fichier `character.json` en `character.name_of_the_character.json`

Ensuite, copiez les 3 fichiers dans le répertoire `04-generate-npc-descriptions` (nous les utiliserons plus tard)

## Questions ?

- Comment avoir quelque chose de plus aléatoire ?


## Quittez Docker Compose

[README](../README.md)








