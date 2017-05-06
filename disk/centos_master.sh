#!/bin/bash

# 安装redis
yum install epel-release -y
yum install redis -y

# 安装docker
yum install docker -y

# 配置镜像地址
echo '{ "insecure-registries":[":5000"] }' > /etc/docker/daemon.json

# 启动docker
systemctl start docker.service

# 启动web界面
chmod +x web
./web &

# 启动反向代理
chmod +x proxy
./proxy &
