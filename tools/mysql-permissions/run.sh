#!/bin/bash

echo "Creating my.cnf"
echo '[client]' > ~/.my.cnf
echo 'user='"$MYSQL_USER" >> ~/.my.cnf
echo 'password='"$MYSQL_PASSWORD" >> ~/.my.cnf

cat ~/permissions.sql | sed -e "s/ROOT_PASSWORD/ROOT_PASSWORD/g" | mysql
exit 0
