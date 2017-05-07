# 镜像服务器配置

* 下载代码

```
wget https://github.com/yubang/app/archive/master.zip
unzip master.zip
cd app-master/disk
```

* 修改配置文件

```
vim config.json
修改Redis和ImageUrl的相关配置
```

* 执行安装脚本

```
/bin/bash centos_image.sh
```

* 修改docker配置，解决https错误问题

```
vim /etc/docker/daemon.json

#加上下面的内容：
{ "insecure-registries":["{ImageUrl字符串内容}"] }

# 重启docker
systemctl restart docker.service
```

上面是centos7的安装教程，其他系统请执行实现下面的流程：

* 下载代码，修改配置文件config.json
* 安装git和docker（版本大于1.12）
* 下载基础镜像（即config.json里面的Image列表镜像）
* 生成git密钥
* 修改docker配置，解决https错误问题
* 重启docker
* 启动imageWorker进程