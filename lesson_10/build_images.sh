#!/bin/bash
pushd web
./build_docker.sh
popd

pushd slack_bot
./build_docker.sh
popd
