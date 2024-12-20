#!/bin/bash 
SERVICE_URL="http://localhost:5051"

echo "Sending data to the AI service on ${SERVICE_URL}"
echo ""
# --silent

curl --no-buffer ${SERVICE_URL}/api/room/generate/name 

echo ""