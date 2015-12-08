#!/bin/bash

sed -i "s/MECH=pam/MECH=shadow/g" /etc/sysconfig/saslauthd
service saslauthd start

cd /tmp
echo $2 > p.txt
saslpasswd2 -c -a memcached -p < p.txt $1 

service saslauthd restart


redis-server /etc/redis.conf 



