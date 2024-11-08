FROM ollama/ollama:0.3.14

RUN <<EOF
/bin/sh -c "\
    /bin/ollama serve & \
    sleep 1 && \
    ollama pull qwen2.5:0.5b && \
    ollama pull qwen2.5:1.5b && \
    ollama pull qwen2.5-coder:1.5b && \
    ollama pull granite3-moe:1b	&& \
    ollama pull smollm2:360m && \
    ollama pull mxbai-embed-large:latest \

"
EOF

ENTRYPOINT ["/bin/ollama"]
EXPOSE 11434
CMD ["serve"]
