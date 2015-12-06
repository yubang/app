#!/bin/bash

wget $1 -O /tmp/www.zip
unzip /tmp/www.zip -d /var/www/html

chmod -Rv 555 /var/www

service httpd start

while true
do
    sleep 50
done
