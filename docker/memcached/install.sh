#!/bin/bash

cd /tmp
wget https://sourceforge.net/projects/levent/files/libevent/libevent-2.0/libevent-2.0.22-stable.tar.gz
tar -zxvf libevent-2.0.22-stable.tar.gz
cd libevent-2.0.22-stable
./configure --prefix=/usr 
make 
make install

cd /tmp
wget http://memcached.org/files/memcached-1.4.25.tar.gz
tar -zxvf memcached-1.4.25.tar.gz
cd memcached-1.4.25
 ./configure --with-libevent=/usr 
make
make install

