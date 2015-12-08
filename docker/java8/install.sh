#!/bin/bash

# 安装各种必须的软件

yum install unzip -y
yum install which -y

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

# 安装java8
cd /tmp
wget --no-cookies --header "Cookie: oraclelicense=accept-securebackup-cookie;" http://download.oracle.com/otn-pub/java/jdk/8u65-b17/jdk-8u65-linux-x64.tar.gz
tar -zxvf jdk-8u65-linux-x64.tar.gz
mv -v jdk1.8.0_65 /usr/local/java

echo 'export PATH=$PATH:/usr/local/java/bin' >> ~/.bashrc
source ~/.bashrc

# 安装jetty
cd /tmp
wget http://download.eclipse.org/jetty/stable-9/dist/jetty-distribution-9.3.6.v20151106.tar.gz
tar -zxvf jetty-distribution-9.3.6.v20151106.tar.gz
mv -v jetty-distribution-9.3.6.v20151106 /usr/local/jetty
echo 'export JETTY_HOME=/usr/local/jetty' >> ~/.bashrc
source ~/.bashrc

mv -vf /tmp/nginx.conf /usr/local/nginx/conf/nginx.conf
