#!/bin/bash 
SERVICE_URL="http://localhost:5053"
read -r -d '' DATA <<- EOM
{
  "monster_name":"Giant Snake"
}
EOM

echo "Sending data to the AI service: ${DATA} on ${SERVICE_URL}"
echo ""
# --silent

curl --no-buffer ${SERVICE_URL}/api/monster/generate/description \
    -H "Content-Type: application/json" \
    -d "${DATA}" 

echo ""