#!/bin/bash 
SERVICE_URL="http://localhost:6666"
read -r -d '' DATA <<- EOM
{
  "question":"What is the best pizza in the world?"
}
EOM

echo "Sending question: ${DATA} on ${SERVICE_URL}"
echo ""
# --silent

curl --no-buffer ${SERVICE_URL}/api/chat \
    -H "Content-Type: application/json" \
    -d "${DATA}" 

echo ""