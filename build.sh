#!/bin/bash

docker image build -t opeco17/gitnavi-api -f backend/api-prod.Dockerfile backend
docker image build -t opeco17/gitnavi-job -f backend/job-prod.Dockerfile backend

docker push opeco17/gitnavi-api
docker push opeco17/gitnavi-job

# docker image tag gitnavi/api registry.digitalocean.com/opeco17/gitnavi/api
# docker push registry.digitalocean.com/opeco17/gitnavi/api

# docker image tag gitnavi/job registry.digitalocean.com/opeco17/gitnavi/job
# docker push registry.digitalocean.com/opeco17/gitnavi/job
