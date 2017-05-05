#!/bin/bash

nginx
mv -v /var/web /usr/share/tomcat/webapps/ROOT
/usr/libexec/tomcat/server start

while true
do
    sleep 50
done