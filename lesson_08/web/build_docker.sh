#!/bin/bash
rm -rf build
mkdir -p build/web/src
cp ../go.* build
cp -rf src build/web

cp -rf ../proto build
cp ../prepare_dependencies.sh build
cp ../generate_proto.sh build

docker rmi web_service
docker build \
         -t alekseikanash/web_service .

rm -rf build