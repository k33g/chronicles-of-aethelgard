docker compose up --watch
docker exec -it rpg-game-1 go run main.go
docker exec -e TERM=xterm-256color -it rpg-game-1 go run main.go
