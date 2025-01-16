#!/bin/bash 
SERVICE_URL="http://localhost:5050"
read -r -d '' DATA <<- EOM
{
  "room_name":"Mieval Caves"
}
EOM

echo "Sending data to the AI service: ${DATA} on ${SERVICE_URL}"
echo ""
# --silent

curl --no-buffer ${SERVICE_URL}/api/room/generate/description \
    -H "Content-Type: application/json" \
    -d "${DATA}" 

echo ""