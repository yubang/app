#!/bin/bash

ip=`ifconfig|grep 'inet '|grep -v '127.0'|xargs|awk -F '[ :]' '{print $2}'`

# 安装git
yum install epel-release -y
yum install git -y

# 安装docker
yum install docker -y

# 配置加速器
sudo sed -i 's|other_args="|other_args="--registry-mirror=http://yubang.m.alauda.cn |g' /etc/sysconfig/docker
sudo sed -i "s|OPTIONS='|OPTIONS='--registry-mirror=http://stmbtfly.m.alauda.cn |g" /etc/sysconfig/docker
sudo sed -i 'N;s|\[Service\]\n|\[Service\]\nEnvironmentFile=-/etc/sysconfig/docker\n|g' /usr/lib/systemd/system/docker.service
sudo sed -i 's|fd://|fd:// $other_args |g' /usr/lib/systemd/system/docker.service
sudo systemctl daemon-reload
sudo service docker restart

# 启动docker
systemctl start docker.service

# 生成ssh密钥
ssh-keygen -t rsa -P "" -f ~/.ssh/id_rsa

# 启动私有镜像服务
docker run -d -p $ip:5000 -v registry

# 启动镜像打包进程
chmod +x imageWorker
./imageWorker &
