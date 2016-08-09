#!/bin/bash

set -e

DIR=$(PWD)/dist
mkdir -p $DIR

version=$(cat ./version.go | grep "const VERSION" | awk '{print $NF}' | sed 's/"//g')
os=$(go env GOOS)
arch=$(go env GOARCH)

for os in linux darwin; do
    name="stache-$version.$os-$arch"
    path="$DIR/$name"
    echo "Building $name"
    GOOS=$os GOARCH=$arch CGO_ENABLED=0 go build -o $path
    gzip < $path > $path.gz
done
