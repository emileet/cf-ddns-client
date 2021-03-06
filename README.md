# cf-ddns-client
[![pipeline status](https://git.plsnobully.me/emileet/cf-ddns-client/badges/master/pipeline.svg)](https://git.plsnobully.me/emileet/cf-ddns-client/-/commits/master)

a go based cloudflare dynamic dns client for managing domains with dynamic ip addresses

## instructions

clone this repo and then build the application

```shell
go build
```

copy `example.env` to `.env` and specify an api token with permissions to edit dns records for relevant zones, as well as enabling ipv6 resolution if required
```shell
API_TOKEN=emi1337xo
IPV6=0
```

copy `data/example.json` to `data/records.json` and configure it with the desired records
```json
{
    "records": [
        {
            "name": "www.plsnobully.me",
            "zone": "plsnobully.me"
        }
    ]
}
```

now run the application
```shell
./cf-ddns-client
```

## docker

clone this repo and then build an image (ensure `data/records.json` exists)

```shell
docker build -t emileet/cf-ddns-client .
```

alternatively, pull the image from the pod.plsnobully.me container registry

```shell
docker pull pod.plsnobully.me/emileet/cf-ddns-client:latest
```

now spin up a container
```shell
docker run --detach \
  -e API_TOKEN=emi1337xo \
  -e IPV6=0 \
  --name cf-ddns-client \
  emileet/cf-ddns-client:latest
```