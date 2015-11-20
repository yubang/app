#!/bin/bash

# 安装各种必须的软件
apt-get install wget -y
wget http://mirrors.163.com/.help/sources.list.trusty -O /etc/apt/sources.list
apt-get update

apt-get install nginx -y
apt-get install unzip -y
apt-get install apache2 -y
apt-get install php5 -y
apt-get install php5-gd -y
apt-get install php5-curl -y
apt-get install php5-mysql -y

