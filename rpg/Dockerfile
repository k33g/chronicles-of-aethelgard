FROM golang:1.23.4-alpine 

ENV TERM=xterm-256color

WORKDIR /app
COPY go.mod .
RUN go mod download


