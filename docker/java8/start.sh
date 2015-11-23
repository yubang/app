#!/bin/bash

# 启动脚本
wget $1 -O /tmp/www.zip
unzip /tmp/www.zip -d /usr/local/jetty/webapps

chmod -Rv 555 /usr/local/jetty/webapps/

/usr/local/nginx/sbin/nginx

# 启动进程
cd /usr/local/jetty/
/usr/local/java/bin/java -jar start.jar

while true
do
    sleep 50
done
