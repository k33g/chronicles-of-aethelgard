#!/bin/bash 
SERVICE_URL="http://localhost:5051"
read -r -d '' DATA <<- EOM
{
  "question":"What is your main quote?"
}
EOM

echo "Sending question: ${DATA} on ${SERVICE_URL}"
echo ""
# --silent

curl --no-buffer ${SERVICE_URL}/api/chat \
    -H "Content-Type: application/json" \
    -d "${DATA}" 

echo ""