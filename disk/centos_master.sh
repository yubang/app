#!/bin/bash

# 获取内网ip
ip=`ifconfig|grep 'inet '|grep -v '127.0'|xargs|awk -F '[ :]' '{print $2}'`

setenforce 0

# 安装redis
yum install epel-release -y
yum install redis -y
sed -i "s/bind 127.0.0.1/bind $ip/" /etc/redis.conf |grep bind|grep bind
systemctl start redis.service

# 安装docker
yum install docker -y

# 启动docker
systemctl start docker.service

# 创建集群
docker swarm init --advertise-addr $ip

# 启动web界面
chmod +x web
./web &

# 启动反向代理
chmod +x proxy
./proxy &
