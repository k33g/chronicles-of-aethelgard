

## Déplacement

Je suis un personnage. je peux etre un humain, un elf, un nain ou un magicien.
Je me déplace sur une map où chaque case est un carré et une case représente un lieu, une pièce.
Ma case de départ est la case de coordonnées (0, 0).
Je peux me déplacer de case en case avec les directions: haut, bas, gauche, droite.
Lorsque je me déplace dans une direction, je me retrouve sur la case voisine dans cette direction.

Par défaut, les cases sont vides et non découvertes.
Lorsque je découvre une case, l'ordinateur me dit que je suis sur la case de coordonnées (x, y) où x et y sont les coordonnées de la case et décrit la pièce.
La case découverte reste découverte.

Lorsque j'arrive à la case de coordonnées (10,10), le jeu s'arrête.

## Objets, monstres et personnages non joueurs

Lorsque j'arrive sur une case, L'ordinateur décide de placer:
- soit de l'or, que je peux ramasser et qui est ajouté à mon inventaire. L'or disparait de la case.
- soit une potion de régénération, qui me redonne des points de vie. La potion disparait de la case.
- soit un monstre, que je dois combattre ou fuir. Le monstre disparait de la case si je gagne le combat.
- soit un personnage non joueur, avec qui je peux discuter. Le personnage non joueur reste sur la case.
- soit rien. et la case reste vide si je reviens dessus.

Et l'ordinateur donne une description de la case(pièce).

L'ordinateur garde en mémoire les cases déjà visitées.
Lorsque je reviens sur une case, la description de la pièce reste la même.

## Inventaire et interface

Je peux à tout moment demander à l'ordinateur de m'afficher la map au format texte, mon inventaire et mes points de vie.

Je peux quitter le jeu à tout moment.

Toujours afficher la liste des commandes possibles.


## Générer le code Go qui permet de jouer à ce jeu.

Structurer le code en plusieurs fichiers (créer un package main et d'autres packages).

