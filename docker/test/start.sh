#!/bin/bash

mkdir -v /var/web
echo "welcome to use paas !" > /var/web/index.html

nginx

while true
do
    sleep 50
done