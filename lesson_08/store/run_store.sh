#!/bin/bash
docker run -P -p 9000:9000/tcp --rm --name "store"  -d store_service