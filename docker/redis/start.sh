#!/bin/bash

echo "requirepass $1" >> /etc/redis.conf
redis-server /etc/redis.conf --port 80 
