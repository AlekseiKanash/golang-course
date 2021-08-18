#!/bin/bash
docker run -P -p 8080:80/tcp --rm --name "web"  -d rest_server