

Je suis un joueur. je peux etre un humain, un elf, un nain ou un magicien

## Carte et déplacement

Je suis sur une map où chaque case est un carré et une case représente un lieu, une pièce.
Je peux me déplacer de case en case.
Je pars de la case de coordonnées (0, 0).
Sur la case de coordonnées (0, 0), je peux me déplacer dans les 4 directions: haut, bas, gauche, droite.

Lorsque je me déplace dans une direction, je me retrouve sur la case voisine dans cette direction.

Je pourrais toujours revenir sur la case précédente.

L'ordinateur garde en mémoire les cases déjà visitées.

Lorsque j'arrive à la dixième case non visitée, le jeu s'arrête.

## Description des pièces

Lorsque j'arrive sur une nouvelle pièce, l'ordinateur me dit que je suis sur la case de coordonnées (x, y) où x et y sont les coordonnées de la case et décrit la piece. 
Lorsque je repasserais dans cette pièce, la description de la pièce sera la même.

## Personnages non joueurs

Il y a 3 personnages non joueurs sur la map. Les types possibles sont: un marchand, un garde, un sorcier.
Lorsque j'arrive sur une case, l'ordinateur decide si je rencontre un personnage non joueur.
Si je rencontre un personnage non joueur, l'ordinateur me demande si je veux discuter.
Si je veux discuter, l'ordinateur me dit un message.

Une fois un personnage non joueur rencontré, lorsque je reviens sur la case, je rencontre à nouveau le personnage non joueur. Et je peux à nouveau discuter avec lui.

## Objets

Lorsque que j'arrive sur une case, l'ordinateur peut décider si je trouve un objet (de type or, potion de regeneration). 

Si c'est une potion de regeneration, elle me redonne des points de vie  et ne sera plus disponible si je visite la piece à nouveau.

Si c'est de l'or, je peux le ramasser. Dans ce cas l'or est ajouté à mon inventaire et ne sera plus disponible si je visite la piece à nouveau.

## Monstres

Lorsque que j'arrive sur une case, l'ordinateur peut décider si il y a un monstre.
Il y a 5 types de monstre: un gobelin, un troll, un dragon, un loup, un ours.
Dans ce cas, je dois combattre le monstre ou fuir et retourner sur la case précédente. 
Le monstre a un nombre de points de vie et une force d'attaque. 
Je dois le combattre en lui infligeant des dégats. 
Lorsque le monstre n'a plus de points de vie, il meurt et je peux continuer mon exploration.
Si je reviens sur la case, le monstre n'est plus là.

## Inventaire et interface

Je peux à tout moment demander à l'ordinateur de m'afficher la map au format texte, mon inventaire et mes points de vie.
J'ai la possibilité de quitter le jeu à tout moment.

Toujours afficher la liste des commandes possibles.

Générer le code Go qui permet de jouer à ce jeu.
