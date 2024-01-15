#!/usr/bin/env bash

# 使用方法：
# ./genModel.sh lottery lottery
# ./genModel.sh lottery prize
# 再将./genModel下的文件剪切到对应服务的model目录里面，记得改package

#生成的表名
tables=$2
#表生成的genmodel目录 改名为model 避免每次复制粘贴都得改包名
modeldir=./model

# 数据库配置
host=127.0.0.1
port=33069
dbname=$1
username=root
passwd=PXDN93VRKUm8TeE7

echo "开始创建库：$dbname 的表：$2"
goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tables}" -dir="${modeldir}" -cache=true --style=goZero

sql2pb -go_package ./pb -host 127.0.0.1 -package pb -password PXDN93VRKUm8TeE7 -port 33069 -schema lottery -service_name lotter -user root >lottery.proto
