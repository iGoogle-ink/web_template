#!/bin/bash

function build_docker {
#   $1 代表第一个参数
    TAG=$1

    IMAGE_NAME="web_template"

    IMAGE_FULL_NAME="$IMAGE_NAME:$TAG"

    HAS_OLD_IMAGES=$(docker images|grep $IMAGE_NAME|grep $TAG|wc -l)
    echo $HAS_OLD_IMAGES

    if [ $HAS_OLD_IMAGES -ne "0" ]; then
        echo "Remove docker image..."
        docker rmi -f $IMAGE_FULL_NAME
    fi

    echo "Building docker image..."
    docker build -t $IMAGE_FULL_NAME .

#    echo "Push image to registry"
#    docker push $IMAGE_FULL_NAME
}

set -e

echo "Building application..."

go mod tidy

CGO_ENABLED=0 GOOS=linux go build -o main .

build_docker "latest"

echo "Cleanup main..."

rm main

echo "Done"