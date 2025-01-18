# Chronicles of Aethelgard

## Objectifs de ce workshop:

- Comprendre ce qu'est une application d'IA générative en utilisant Ollama et des petits LLM
- On utilisera directement l'API d'Ollama (qui est écrit en Go)
  - Donc tous les exemples sont en Go
  - Pas de panique, vous allez essentiellement jouer avec les informations à envoyer au(x) LLM(s)
  - Donc pas de code, ou presque (on va le lire et éventuellement le modifier)
- ATTENTION: on va bosser sans GPU (si c'est possible)
  - Donc avec de trés petits LLMs
  - Pas forcément très précis, mais ils sont très pratiques pour travailler sur des petites architectures et on peut tout de même les éduquer un peu
- Pourquoi Docker ? Limiter les galères d'installation

## En vrai que va-t-on faire ?

À l'aide de quelques exercices (j'en ai 17 🤪) nous allons progressivement voir comment fonctionne une application d'IA générative. Tous les principes que je vais vous montrer fonctionnent avec d'autres langages (Ollama propose un SDK JavaScript et un SDK Python et aussi une API REST). Si plus tard vous souhaitez vous mettre à LangChain(Python, JS, 4J), les principes restes identiques.

Le contexte: Et si on se faisait aider par l'IA pour créer un JdR en mode texte?

Nous verrons comment:

0. Si vous avez respecté les prérequis 😈: [`00-requirements`](00-requirements/README.md)
1. Générer la description d'une pièce dans un donjon: [`01-generate-room-description`](01-generate-room-description/README.md)
2. Générer des noms de personnages (en JSON): [`02-generate-names`](02-generate-names/README.md)
3. Pareil mais en mieux: [`03-generate-names-structured-output`](03-generate-names-structured-output/README.md)
4. Génerer une fiche de personnage: [`04-generate-npc-descriptions`](04-generate-npc-descriptions/README.md)
🚧

Et à la fin nous essaieront de regrouper tout ça dans un mini "Dungeon Crawler"



