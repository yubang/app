#!/bin/bash

# 安装各种必须的软件
yum install unzip -y

# 安装nginx

yum -y install zlib zlib-devel openssl openssl--devel pcre pcre-devel

cd /tmp
wget --spider http://nginx.org/download/nginx-1.8.0.tar.gz
curl -o nginx-1.8.0.tar.gz http://nginx.org/download/nginx-1.8.0.tar.gz
tar -zxvf nginx-1.8.0.tar.gz
cd nginx-1.8.0
./configure --prefix=/usr/local/nginx
make
make install

# 安装node.js
cd /tmp
wget https://nodejs.org/dist/v4.2.2/node-v4.2.2-linux-x64.tar.gz
tar -zxvf node-v4.2.2-linux-x64.tar.gz
mv -v node-v4.2.2-linux-x64 /usr/local/node

mv -vf /tmp/nginx.conf /usr/local/nginx/conf/nginx.conf
