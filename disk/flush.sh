#!/bin/bash

docker ps -a|awk -F ' '  'NR!=1{print $1}'|xargs docker rm
docker images|awk -F " " 'NR!=1{print$3}'|xargs docker rmi -f
