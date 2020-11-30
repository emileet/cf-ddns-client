FROM archlinux:latest
MAINTAINER Emily Maré (emileet) <emileet@plsnobully.me>

WORKDIR /app

COPY data/records.json /app/data/
COPY cf-ddns-client /app/run

ENV API_TOKEN=VALUE \
    IPV6=0

VOLUME ["/app/data"]
CMD ["/app/run"]