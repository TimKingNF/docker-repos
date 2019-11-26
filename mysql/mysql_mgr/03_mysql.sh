#!/usr/bin/env bash

# mysql binary
mysql=/usr/local/Cellar/mysql-client/8.0.18/bin/mysql
host=localhost
user=root
passwd=root
port=3306
tmp_sql=./tmp.sql

if [ "$1" = "mysql-2" ]; then
  port=36002
elif [ "$1" = "mysql-3" ]; then
  port=36003
else
  echo "参数错误"; exit 1
fi

mysql_master_cmd="${mysql} --protocol=tcp -h${host} -P36001 -u${user} -p${passwd}"
mysql_cmd="${mysql} --protocol=tcp -h${host} -P${port} -u${user} -p${passwd}"
cmd_ret=`${mysql_cmd} -sN -e"select @@global.gtid_executed;" 2>/dev/null`
gtid=${cmd_ret:0:36}
min=${cmd_ret:37:1}
max=${cmd_ret:39:1}

run_sql=""
for i in $(seq ${min} ${max}); do
  run_sql+="SET GTID_NEXT='${gtid}:${i}';begin;commit;"
done
run_sql+="SET GTID_NEXT='AUTOMATIC';"

# mysql master 执行
echo ${run_sql} > ${tmp_sql}
${mysql_master_cmd} < ${tmp_sql} 2>/dev/null

# mysql slave 启动 MGR
${mysql_cmd} -e" START GROUP_REPLICATION;" 2>/dev/null

rm -rf ${tmp_sql}

