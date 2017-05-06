#!/bin/bash

# 获取内网ip
ip=`ifconfig|grep 'inet '|grep -v '127.0'|xargs|awk -F '[ :]' '{print $2}'`

setenforce 0

# 安装git
yum install epel-release -y
yum install git -y

# 安装docker
yum install docker -y

# 使用加速器
sudo sed -i 's|other_args="|other_args="--registry-mirror=http://yubang.m.alauda.cn |g' /etc/sysconfig/docker
sudo sed -i "s|OPTIONS='|OPTIONS='--registry-mirror=http://stmbtfly.m.alauda.cn |g" /etc/sysconfig/docker
sudo sed -i 'N;s|\[Service\]\n|\[Service\]\nEnvironmentFile=-/etc/sysconfig/docker\n|g' /usr/lib/systemd/system/docker.service
sudo sed -i 's|fd://|fd:// $other_args |g' /usr/lib/systemd/system/docker.service
sudo systemctl daemon-reload
sudo service docker restart

# 启动docker
systemctl start docker.service

# 下载基础镜像
docker pull registry.alauda.cn/yubang/paas_base_test
docker pull registry.alauda.cn/yubang/paas_base_applet_static
docker pull registry.alauda.cn/yubang/paas_base_golang
docker pull registry.alauda.cn/yubang/paas_base_java
docker pull registry.alauda.cn/yubang/paas_base_nodejs
docker pull registry.alauda.cn/yubang/paas_base_php
docker pull registry.alauda.cn/yubang/paas_base_python2
docker pull registry.alauda.cn/yubang/paas_base_python2_worker
docker pull registry.alauda.cn/yubang/paas_base_static
docker pull registry.alauda.cn/yubang/paas_base_python3
docker pull registry.alauda.cn/yubang/paas_base_ruby

# 生成ssh密钥
ssh-keygen -t rsa -P "" -f ~/.ssh/id_rsa

# 启动私有镜像服务
docker run -d -p $ip:5000:5000 --restart=always registry

# 启动镜像打包进程
chmod +x imageWorker
./imageWorker &
