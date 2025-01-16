# Chronicles of Aethelgard

à la fois un moyen de jouer avec Ollama (sans GPU) et d'utiliser Docker Compose.

faire essentiellement travailler les gens sur les prompts 

à la fin il faudra tout mettre dans un même compose file

<!-- Construire le compose file au fur et a mesure
on evite la partie data externe et le network partagé 

-->

## Remarques

Expliquer pourquoi je ne fais pas de go build / ou dire que normalement je devrair le faire

## Pré-requis

### Docker

### Ollama
> https://hub.docker.com/r/ollama/ollama

```bash
docker pull ollama/ollama:0.5.4
```

LLMs:

- qwen2.5:0.5a
- qwen2.5:1.5b
- qwen2.5:3b
- qwen2:1.5b-instruct

ou un compose file à lancer à l'avance

### Golang

```bash
docker pull golang:1.23.1-alpine
```

localement et l'image

docker pull python:3.13-alpine

- Installer Go (une version récente - `j'utilise go1.23.1`) 
- Ollama (version `0.5.4`)
- Installer Docker



https://hub.docker.com/r/ollama/ollama


## Initialisation

```bash
docker volume create ollama_shared_data
```


there is a map for a role playing game : 10 * 10 cells. Randomly place three  orcs, two dragons
... and five treasures on the map

Entre le moment oú j'ai posté sur le CFP et maintenant, il s'est passé beaucoup de choses (évolution d'Ollama et des modèles + mes connaissances). 

Le jeu va être très simple. L'objectif étant de vous faire comprendre les fonctionnements de l'IA générative et de vous donner des idées pour vos propres projets.

## Pourquoi Go ?

Pour utiliser directement l'API d'Ollama.
La logique utilisée est la même que pour les autres langages.

## Bébés LLMs

Pas forcément très précis, mais ils sont très pratiques pour travailler sur des petites architectures et on peut tout de même les éduquer un peu.


## Principes et objectifs du workshop

- Comprendre comment fonctionne les applications d'IA générative
- Travailler sur les prompts
- On ne codera quasiment pas
- On va créer tous les éléments / outils pour aider à jouer à un jeu de rôle


#### 12-call-room-services

commencer l'assemblage avec compose
