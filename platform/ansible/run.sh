#!/bin/bash
ansible-playbook \
 --inventory inventory.ini \
 --user ec2-user \
 --private-key ~/.ssh/kubetwo_id_rsa \
 "$1"