#!/bin/bash 
SERVICE_URL="http://localhost:5052"
read -r -d '' DATA <<- EOM
{
  "question":"Who am I?"
}
EOM

echo "Sending question: ${DATA} on ${SERVICE_URL}"
echo ""
# --silent

curl --no-buffer ${SERVICE_URL}/api/chat \
    -H "Content-Type: application/json" \
    -d "${DATA}" 

echo ""