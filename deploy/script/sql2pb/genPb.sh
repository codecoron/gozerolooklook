#!/usr/bin/env bash

# 使用方法：
# ./genPb.sh lottery lottery
# ./genPb.sh lottery prize
# ./genPb.sh looklook_usercenter user_sponsor
# 再将./genPb下的文件剪切到对应服务的model目录里面，记得改package

#生成的表名
tables=$2

# 数据库配置
host=127.0.0.1
port=33069
dbname=$1
username=root
passwd=PXDN93VRKUm8TeE7

sql2pb -go_package ./pb -host "${host}" -package pb -password "${passwd}" -port "${port}" -schema "${dbname}" -table "${tables}" -service_name "${tables}" -user "${username}" > "${tables}".proto
