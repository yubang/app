#!/bin/bash

# 安装docker
yum install epel-release -y
yum install docker -y

# 配置镜像地址
echo '{ "insecure-registries":[":5000"] }' > /etc/docker/daemon.json

# 启动docker
systemctl start docker.service

# 提示
echo -e "请在web控制台copy加入集群服务器命令，然后执行\n"