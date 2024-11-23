FROM ollama/ollama:0.4.1

RUN <<EOF
/bin/sh -c "\
    /bin/ollama serve & \
    sleep 1 && \
    ollama pull qwen2.5:0.5b && \
    ollama pull qwen2.5:1.5b && \
    ollama pull mxbai-embed-large:latest
"
EOF

ENTRYPOINT ["/bin/ollama"]
EXPOSE 11434
CMD ["serve"]
