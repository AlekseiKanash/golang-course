#!/bin/bash

docker stop web
docker run  -P -p 8080:80/tcp -itd --rm --name "web" web_service
docker stop store
docker run -P -p 9000:9000/tcp -itd --rm --name "store" store_service

# docker network rm web-store
# docker network create web-store
# docker network connect web-store web
# docker network connect web-store store