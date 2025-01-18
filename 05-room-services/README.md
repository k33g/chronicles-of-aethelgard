# Room service

Faire un service qui permet d'obtenir une description de pièce à partir d'un nom.


Niveau noms aléatoires, c'est bof mais on ne peux pas tout avoir. on veut gardder un petit modele
par contre on peut aider le modèle.

ça sera l'exercice

ajouter des instructions pour aider le modèle à générer des noms de pièces
penser à l'enlever ensuite dans le code

```bash
docker run --rm --network host curlimages/curl:8.6.0 \
    --silent --no-buffer "http://localhost:5050/api/room/generate/name" 

docker run --rm --network host curlimages/curl:8.6.0 \
    --silent --no-buffer "http://localhost:5050/api/room/generate/description" \
    -H "Content-Type: application/json" \
    -d '{"room_name":"Minion Lair"}'

```