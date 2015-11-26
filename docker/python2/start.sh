#!/bin/bash

# 启动脚本
mkdir -p /var/www/html
wget $1 -O /tmp/www.zip
unzip /tmp/www.zip -d /var/www/html

chmod -Rv 555 /var/www

/usr/local/nginx/sbin/nginx

# 调用pip
cd /var/www/html
pip install -r requirements.txt

# 启动进程
gunicorn -b 127.0.0.1:8000 -w 5 index:app

while true
do
    sleep 50
done
