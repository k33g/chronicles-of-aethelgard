#!/bin/bash 
docker image save -o backup-runtimes-image.tar chronicles-of-aethelgard_devcontainer-ai-workspace:latest
docker image save -o backup-llms-image.tar chronicles-of-aethelgard_devcontainer-ollama-service:latest
