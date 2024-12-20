#!/bin/bash 
source data/.env

read -r -d '' DATA <<- EOM
{
  "user":"why the sky is blue?"
}
EOM

echo "Sending data to the AI Bot: ${DATA} on ${BOT_HOST}"
echo ""
# --silent

curl --no-buffer ${BOT_HOST}/api/chat \
    -H "Content-Type: application/json" \
    -d "${DATA}" 

echo ""