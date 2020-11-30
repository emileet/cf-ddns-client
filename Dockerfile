FROM golang:latest
MAINTAINER Emily Maré (emileet) <emileet@plsnobully.me>

WORKDIR /app

COPY data/records.json /app/data/
COPY go.mod main.go /app/

RUN go mod download && \
    go build -o run .

ENV API_TOKEN=VALUE \
    IPV6=0

VOLUME ["/app/data"]
CMD ["/app/run"]