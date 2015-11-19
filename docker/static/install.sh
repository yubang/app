#!/bin/bash

# 安装各种必须的软件

apt-get install wget -y
apt-get install nginx -y
apt-get install unzip -y

mkdir -pv /var/www/html

mv -fv /tmp/default /etc/nginx/sites-enabled/default
chmod -v 777 /etc/nginx/sites-enabled/default
