#!/usr/bin/env sh

mysql="root:root@tcp(mysql:3306)/ark_admin"
table=$1
dir=$2

goctl model mysql datasource -url="${mysql}" -table="${table}" -c -dir ${dir} -home /server/dev/goctl