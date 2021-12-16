#!/bin/bash

docker network create intra

docker stop store
docker run --net=intra -P -p 9000:9000/tcp -itd --rm --name "store" store_service

docker stop web
docker run --net=intra -P -p 8080:80/tcp -itd --rm --name "web" web_service
