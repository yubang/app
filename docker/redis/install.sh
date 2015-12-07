#!/bin/bash

cd /tmp
wget http://download.redis.io/releases/redis-3.0.5.tar.gz
tar -zxvf redis-3.0.5.tar.gz
cd redis-3.0.5
make
make install
cp redis.conf /etc/

