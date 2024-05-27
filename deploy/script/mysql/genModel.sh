#!/bin/bash
  
# table name generated 
tables=$2

# genmodel directory
modeldir=./genModel

# database settings
host=127.0.0.1
port=33069
dbname=microservices_$1
username=root
passwd=PXDN93VRKUm8TeE7


echo "beginning create---> database: $dbname, table: $2"
goctl model mysql datasource --url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" --table="${tables}"  --dir="${modeldir}" --cache=true --style=goZero
