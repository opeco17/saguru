#!/bin/sh

kubectl create secret generic envs --from-env-file=.prodenv
