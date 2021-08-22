#!/bin/bash
rm -rf build
mkdir -p build/store
cp ../go.* build
cp -rf src build/store

cp -rf ../proto build
cp ../prepare_dependencies.sh build
cp ../generate_proto.sh build

docker rmi store_service
docker build \
         -t store_service .