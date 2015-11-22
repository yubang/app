#!/bin/bash

# 启动脚本
mkdir -p /var/www/html
wget $1 -O /tmp/www.zip
unzip /tmp/www.zip -d /var/www/html

chmod -Rv 555 /var/www

/usr/local/nginx/sbin/nginx

# 启动进程
cd /var/www/html
/usr/local/node/bin/npm install
/usr/local/node/bin/node index.js

while true
do
    sleep 50
done
