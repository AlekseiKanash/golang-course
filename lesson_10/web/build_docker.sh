#!/bin/bash
rm -rf build
mkdir -p build/web/src
cp ../go.* build
cp -rf src build/web
cp -rf ../store build
cp -rf ../openweather build

cp prepare_dependencies.sh build

docker rmi alekseikanash/web_service_weather
docker build \
         -t alekseikanash/web_service_weather .

rm -rf build