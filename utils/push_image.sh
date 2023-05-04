#!/bin/sh

if [ $# -ne 1 ]; then
  echo "First argument: image tag to build and push" 1>&2
  exit 1
fi

IMAGE_PREFIX=opeco17/gitnavi
IMAGE_TAG=${1}

cd backend

docker image build -t ${IMAGE_PREFIX}-api:${IMAGE_TAG} -f api-prod.Dockerfile .
docker image build -t ${IMAGE_PREFIX}-job:${IMAGE_TAG} -f job-prod.Dockerfile .

docker push ${IMAGE_PREFIX}-api:${IMAGE_TAG}
docker push ${IMAGE_PREFIX}-job:${IMAGE_TAG}
