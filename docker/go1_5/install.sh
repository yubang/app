#!/bin/bash

# 安装各种必须的软件

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

# 安装go
cd /tmp
wget https://storage.googleapis.com/golang/go1.5.1.linux-amd64.tar.gz
tar -zxvf go1.5.1.linux-amd64.tar.gz
mv -v go /usr/local/go


# 配置 go环境
mkdir -v /var/gopath
echo "" >> ~/.bashrc
echo "export GOPATH=/var/gopath" >> ~/.bashrc
source ~/.bashrc

# 
yum install -y bzr

mv -vf /tmp/nginx.conf /usr/local/nginx/conf/nginx.conf
