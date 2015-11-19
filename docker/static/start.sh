#!/bin/bash

# 启动脚本

wget $1 -O /tmp/www.zip
unzip /tmp/www.zip -d /var/www/html

chmod -Rv 555 /var/www

nginx

while true
do
    sleep 50
done
