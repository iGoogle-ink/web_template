#!/bin/bash
function build_docker {
#   $1 代表第一个参数
    TAG=$1
    if [[ -z "$TAG" ]]; then
        echo "image tag can not null"
        exit
    fi

    IMAGE_NAME="web_template"

    IMAGE_FULL_NAME="$IMAGE_NAME:$TAG"

    echo ${IMAGE_FULL_NAME}

    HAS_OLD_IMAGES=$(docker images|grep ${IMAGE_NAME}|grep ${TAG}|wc -l)
    echo ${HAS_OLD_IMAGES} ${IMAGE_FULL_NAME} "images"

    if [[ ${HAS_OLD_IMAGES} -ne "0" ]]; then
        echo "Remove docker image..."
        docker rmi -f ${IMAGE_FULL_NAME}
    fi

    echo "Building docker image..."
    docker build -t ${IMAGE_FULL_NAME} .

    # docker push ${IMAGE_FULL_NAME}
}

set -e

echo "Building application..."

go mod tidy

CGO_ENABLED=0 GOOS=linux go build cmd/main.go

build_docker $1

echo "Cleanup..."

rm main

docker image prune -f

echo "Done"