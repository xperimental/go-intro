#!/usr/bin/env bash

VERSION=$(git rev-parse --short HEAD)
readonly VERSION

echo "-- BUILDING IMAGE"
docker build -t xperimental/go-intro .
docker tag xperimental/go-intro "xperimental/go-intro:$VERSION"

echo "-- RUNNING"
docker run --rm -itp 3999:3999 "xperimental/go-intro:$VERSION"

