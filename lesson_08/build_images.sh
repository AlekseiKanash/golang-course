#!/bin/bash
pushd web
./build_docker.sh
popd

pushd store
./build_docker.sh
popd