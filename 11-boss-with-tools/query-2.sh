#!/bin/bash 
SERVICE_URL="http://localhost:6666"
read -r -d '' DATA <<- EOM
{
  "question":"How can I exit this place?"
}
EOM

echo "Sending question: ${DATA} on ${SERVICE_URL}"
echo ""
# --silent

#curl --no-buffer ${SERVICE_URL}/api/chat \
#    -H "Content-Type: application/json" \
#    -d "${DATA}" 

docker run --rm --network host curlimages/curl:8.6.0 \
    --silent --no-buffer "${SERVICE_URL}/api/chat" \
    -H "Content-Type: application/json" \
    -d "${DATA}"

echo ""