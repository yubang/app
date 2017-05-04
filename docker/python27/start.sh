#!/bin/bash

nginx
cd /var/web
gunicorn -b 127.0.0.1:8000 -w 1 -k gevent index:app

while true
do
    sleep 50
done