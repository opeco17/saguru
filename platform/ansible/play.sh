#!/bin/bash
ansible-playbook \
 --inventory inventory.ini \
 --user ec2-user \
 --private-key ~/.ssh/gitnavi_id_rsa \
 "$1"