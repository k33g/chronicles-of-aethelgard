# Chronicles of Aethelgard

## Objectifs de ce workshop:

- Comprendre ce qu'est une application d'IA générative en utilisant Ollama et des petits LLM
- On utilisera directement l'API d'Ollama (qui est écrit en Go)
  - Donc tous les exemples sont en Go
  - Pas de panique, vous allez essentiellement jouer avec les informations à envoyer au(x) LLM(s)
  - Donc pas de code, ou presque (on va le lire et éventuellement le modifier)
- **ATTENTION: on va bosser sans GPU (si c'est possible!)**
  - Donc avec de trés petits LLMs
  - Pas forcément très précis, mais ils sont très pratiques pour travailler sur des petites architectures et on peut tout de même les éduquer un peu
  - On sauve la planète 🌍
- Pourquoi Docker ? 
  - Limiter les galères d'installation
  - J'💙 🐳 Compose

## En vrai que va-t-on faire ?

À l'aide de quelques exercices (j'en ai 17 🤪) nous allons progressivement voir comment fonctionne une application d'IA générative. Tous les principes que je vais vous montrer fonctionnent avec d'autres langages (Ollama propose un SDK JavaScript et un SDK Python et aussi une API REST). Si plus tard vous souhaitez vous mettre à LangChain(Python, JS, 4J), les principes restes identiques.

Le contexte: Et si on se faisait aider par l'IA pour créer un JdR en mode texte?

Donc on va développer des petits services que l'on appellera à partir d'une autre application : **un "jeu" 😂 façon "Dungeon Crawler"**.

Nous verrons comment:

### Parler avec les LLMs
> C'est là où vous allez un peu travailler

0. Si vous avez respecté les prérequis 😈: [`00-requirements`](00-requirements/README.md)
1. Générer la description d'une pièce dans un donjon: [`01-generate-room-description`](01-generate-room-description/README.md)
2. Générer des noms de personnages (en JSON): [`02-generate-names`](02-generate-names/README.md)
3. Pareil mais en mieux: [`03-generate-names-structured-output`](03-generate-names-structured-output/README.md)
4. Génerer une fiche de personnage: [`04-generate-npc-descriptions`](04-generate-npc-descriptions/README.md)
5. Créer un micro-service pour obtenir une description de pièce dans un donjon: [`05-room-services`](05-room-services/README.md)
6. Créer un service pour envoyer une question à un PNJ: [`06-bot-npc`](06-bot-npc/README.md)
7. Créer un service pour envoyer une question à un PNJ avec gestion de la mémoire conversationnelle: [`07-bot-with-memory`](07-bot-with-memory/README.md)
8. Créer un service pour envoyer une question à un PNJ + **RAG**: [`08-bot-with-rag`](08-bot-with-rag/README.md)
9. Créer un service qui donnera une description de monstre en fonction de son nom: [`09-monsters`](09-monsters/README.md)
10. Il faut un Boss de fin de niveau avec une forte personnalité: [`10-end-level-boss`](10-end-level-boss/README.md)
11. Le Boss de fin de niveau doit pouvoir détecter que vous lui donnez les mots magiques pour sortir du donjon: [`11-boss-with-tools`](11-boss-with-tools/README.md)

### Comment ré-utiliser tout ça
> Là, c'est plus du mode démo

12. Exemple d'application appelant les services de description de salles du donjon: [`12-call-room-services`](12-call-room-services/README.md)
13. De la même façon, voyons comment utiliser la description des monstres: [`13-call-monster-service`](13-call-monster-service/README.md)
14. Il va falloir faire la même chose avec les bots PNJ (avec quelques adaptations): [`14-chat-with-bot-services`](14-chat-with-bot-services/README.md)
15. Alors ça serait sympa d'avoir une vraie conversation (taper la question dans un terminal, attendre la réponse, re taper une nouvelle question, ...): [`15-chat-with-bot-services`](15-chat-with-bot-services/README.md)
16. Nous allons mettre en place le même principe avec le Boss de fin de niveau: [`16-chat-with-boss`](16-chat-with-boss/README.md)

Et à la fin nous essaieront de regrouper tout ça dans le mini "Dungeon Crawler"

16. Le "jeu": [`rpg`](rpg/README.md)

# 🎉 The End!

## Dans la prochaine version de ce workshop

- Re-écrire le "Dungeon Crawler" en plusieurs langages
- Créer un serveur MCP pour gérer les combats, points de vie, ...
- Transformer les bots en bots Discord
- ...
