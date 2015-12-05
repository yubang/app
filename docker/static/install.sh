#!/bin/bash

# 安装各种必须的软件
yum groupinstall "Development tools" -y
yum install wget -y
yum install tar -y

cd /tmp
yum -y install zlib zlib-devel openssl openssl--devel pcre pcre-devel
wget --spider http://nginx.org/download/nginx-1.8.0.tar.gz
wget http://nginx.org/download/nginx-1.8.0.tar.gz
tar -zxvf nginx-1.8.0.tar.gz
cd nginx-1.8.0
./configure --prefix=/usr/local/nginx
make
make install

mv -vf /tmp/nginx.conf /usr/local/nginx/conf/nginx.conf


