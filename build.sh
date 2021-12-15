#!/bin/bash

docker image build -t gitnavi/api -f backend/api-prod.Dockerfile backend
docker image build -t gitnavi/job -f backend/job-prod.Dockerfile backend
docker image build -t gitnavi/database -f database/database.Dockerfile database

# docker image tag gitnavi/api registry.digitalocean.com/opeco17/gitnavi/api
# docker push registry.digitalocean.com/opeco17/gitnavi/api

# docker image tag gitnavi/api registry.digitalocean.com/opeco17/gitnavi/job
# docker push registry.digitalocean.com/opeco17/gitnavi/job

# docker image tag gitnavi/api registry.digitalocean.com/opeco17/gitnavi/database
# docker push registry.digitalocean.com/opeco17/gitnavi/database
