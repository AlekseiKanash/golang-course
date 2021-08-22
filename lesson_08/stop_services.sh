#!/bin/bash

docker stop web
docker stop store

docker network rm web-store