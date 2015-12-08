#!/bin/bash

# 安装各种必须的软件


yum install zlib-devel -y
yum install bzip2-devel -y
yum install openssl-devel -y
yum install ncurses-devel -y
yum install sqlite-devel -y


cd /tmp
wget https://www.python.org/ftp/python/2.7.9/Python-2.7.9.tar.xz
tar xvf Python-2.7.9.tar.xz
cd Python-2.7.9
./configure --prefix=/usr/local
make && make install

cd /tmp
wget --no-check-certificate 'https://pypi.python.org/packages/source/s/setuptools/setuptools-0.7.2.tar.gz'
tar -xvf setuptools-0.7.2.tar.gz
cd setuptools-0.7.2
python2.7 setup.py install

cd /tmp 
wget --no-check-certificate 'https://pypi.python.org/packages/source/p/pip/pip-1.2.tar.gz'
tar -xvf pip-1.2.tar.gz
cd pip-1.2
python2.7 setup.py install

cd /tmp
yum -y install zlib zlib-devel openssl openssl--devel pcre pcre-devel
wget --spider http://nginx.org/download/nginx-1.8.0.tar.gz
wget http://nginx.org/download/nginx-1.8.0.tar.gz
tar -zxvf nginx-1.8.0.tar.gz
cd nginx-1.8.0
./configure --prefix=/usr/local/nginx
make
make install

yum -y install mysql-devel

pip install gunicorn

mv -vf /tmp/nginx.conf /usr/local/nginx/conf/nginx.conf


