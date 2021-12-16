#!/bin/bash
rm -rf build
mkdir -p build/slack_bot/src
cp ../go.* build
cp -rf src build/slack_bot
cp -rf ../store build
cp -rf ../openweather build

cp prepare_dependencies.sh build

docker rmi alekseikanash/slack_bot_service_weather
docker build \
         -t alekseikanash/slack_bot_service_weather .

rm -rf build