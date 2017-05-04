#!/bin/bash

nginx
cd /var/web
node index.js

while true
do
    sleep 50
done