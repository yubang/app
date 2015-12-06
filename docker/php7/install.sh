#!/bin/bash

# 安装各种必须的软件
yum groupinstall "Development tools" -y
yum install wget -y
yum install tar -y

yum install -y httpd httpd-devel
yum install -y libxml2-devel 
yum install -y openssl
yum install -y openssl-devel
yum -y install curl-devel
yum install libjpeg.x86_64 libpng.x86_64 freetype.x86_64 libjpeg-devel.x86_64 libpng-devel.x86_64 freetype-devel.x86_64 -y
yum install libjpeg-devel -y

cd /tmp
wget ftp://mcrypt.hellug.gr/pub/crypto/mcrypt/attic/libmcrypt/libmcrypt-2.5.7.tar.gz 
tar -zxvf libmcrypt-2.5.7.tar.gz 
cd libmcrypt-2.5.7 
./configure
make
make install 

echo /usr/local/lib >> /etc/ld.so.conf.d/local.conf 
ldconfig -v 

cd /tmp
wget http://cn2.php.net/get/php-7.0.0.tar.gz/from/this/mirror
tar -zxvf mirror

cd php-7.0.0/

./configure \
--prefix=/usr/local/php7 \
--with-apxs2=/usr/sbin/apxs \
--exec-prefix=/usr/local/php7 \
--bindir=/usr/local/php7/bin \
--sbindir=/usr/local/php7/sbin \
--includedir=/usr/local/php7/include \
--libdir=/usr/local/php7/lib/php \
--mandir=/usr/local/php7/php/man \
--with-config-file-path=/usr/local/php7/etc \
--with-mysql-sock=/var/run/mysql/mysql.sock \
--with-mcrypt=/usr/include \
--with-mhash \
--with-openssl \
--with-mysql=shared,mysqlnd \
--with-mysqli=shared,mysqlnd \
--with-pdo-mysql=shared,mysqlnd \
--with-gd \
--with-iconv \
--with-zlib \
--enable-zip \
--enable-inline-optimization \
--disable-debug \
--disable-rpath \
--enable-shared \
--enable-xml \
--enable-bcmath \
--enable-shmop \
--enable-sysvsem \
--enable-mbregex \
--enable-mbstring \
--enable-ftp \
--enable-gd-native-ttf \
--enable-pcntl \
--enable-sockets \
--with-xmlrpc \
--enable-soap \
--without-pear \
--with-gettext \
--enable-session \
--with-curl \
--with-jpeg-dir \
--with-freetype-dir \
--enable-opcache \
--enable-fpm \
--enable-fastcgi \
--with-fpm-user=php \
--with-fpm-group=php \
--without-gdbm \
--disable-fileinfo

make
make install

echo '<FilesMatch \.php$>' >> /etc/httpd/conf/httpd.conf
echo 'SetHandler application/x-httpd-php' >> /etc/httpd/conf/httpd.conf
echo '</FilesMatch>' >> /etc/httpd/conf/httpd.conf
