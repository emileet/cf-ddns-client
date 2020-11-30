FROM alpine:latest
MAINTAINER Emily Maré (emileet) <emileet@plsnobully.me>

WORKDIR /app

RUN wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.32-r0/glibc-2.32-r0.apk && \
    wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub && \
    apk add --no-cache -U glibc-2.32-r0.apk

COPY data/records.json /app/data/
COPY cf-ddns-client /app/run

ENV API_TOKEN=VALUE \
    IPV6=0

VOLUME ["/app/data"]
CMD ["/app/run"]