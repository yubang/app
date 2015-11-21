#!/bin/bash

# 安装各种必须的软件
yum groupinstall "Development tools" -y
yum install zlib-devel -y
yum install bzip2-devel -y
yum install openssl-devel -y
yum install ncurses-devel -y
yum install sqlite-devel -y

yum install wget -y
yum install tar -y

cd /tmp
wget https://www.python.org/ftp/python/3.5.0/Python-3.5.0.tar.xz
tar xvf Python-3.5.0.tar.xz
cd Python-3.5.0
./configure --prefix=/usr/local
make && make install

yum -y install zlib zlib-devel openssl openssl--devel pcre pcre-devel

cd /tmp
wget --spider http://nginx.org/download/nginx-1.8.0.tar.gz
curl -o nginx-1.8.0.tar.gz http://nginx.org/download/nginx-1.8.0.tar.gz
tar -zxvf nginx-1.8.0.tar.gz
cd nginx-1.8.0
./configure --prefix=/usr/local/nginx
make
make install

yum -y install mysql-devel

pip3 install gunicorn

mv -vf /tmp/nginx.conf /usr/local/nginx/conf/nginx.conf


