#!/bin/bash

nginx
cd /var/web
ruby index.rb

while true
do
    sleep 50
done