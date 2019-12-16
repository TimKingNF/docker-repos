# Mysql Group Replication

Mysql Group Replication (组复制) 测试.

## Requirements

* docker
* docker-compose
* make
* mysql client

## Install

```bash
make init
make up
mysql -hlocalhost -P36001 -uroot -proot < ./02_mysql-1.sql

# 在数据库初始化的时候需要执行, 在初始化数据库之后有部分数据属于其他组的事务, 需要手动调整过来
bash 03_mysql.sh mysql-2
bash 03_mysql.sh mysql-3
```

中间可以通过 `docker logs mysql-1` 来查看启动日志

使用 `mysql --protocol=tcp -hlocalhost -P36001 -uroot -proot` 连上mysql.

## Tests

测试集群是否正常启动

```bash
mysql -hlocalhost -P36001 -uroot -proot -e"SELECT * FROM performance_schema.replication_group_members;"
```

结果如下所示

```bash
"CHANNEL_NAME"  "MEMBER_ID"     "MEMBER_HOST"   "MEMBER_PORT"   "MEMBER_STATE"  "MEMBER_ROLE"     "MEMBER_VERSION"
"group_replication_applier"     "2494deb7-1033-11ea-9b3a-0242ac1a0002"  "mysql-1" "3306"   "ONLINE"        "PRIMARY"       "8.0.18"
"group_replication_applier"     "25474049-1033-11ea-8344-0242ac1a0004"  "mysql-2" "3306"   "ONLINE"        "PRIMARY"       "8.0.18"
"group_replication_applier"     "254f685f-1033-11ea-a1b8-0242ac1a0003"  "mysql-3" "3306"   "ONLINE"        "PRIMARY"       "8.0.18"
```

可以看到有3个节点,  `MEMBER_STATE` 为 `ONLINE` , `MEMBER_ROLE`为`PRIMARY`, 说明 MGR 多主模式启动成功.

```sql
CREATE DATABASE test;
CREATE TABLE test.messages (id int auto_increment, message VARCHAR(250), PRIMARY KEY ( `id` ));
INSERT INTO test.messages(message) VALUES ('hello');
```

## Single Primary Mode

修改 `cnf/mysql-*/mysql.conf` 中的配置文件

```bash
loose-group_replication_single_primary_mode=on
loose-group_replication_enforce_update_everywhere_checks=off
```

所有节点执行 ` STOP GROUP_REPLICATION;` 停止组复制.

在主节点执行如下内容

```mysql
SET GLOBAL group_replication_bootstrap_group=ON;

START GROUP_REPLICATION;

SET GLOBAL group_replication_bootstrap_group=OFF;
```

在其他节点执行

```mysql
START GROUP_REPLICATION;
```

最后检查MGR集群是否启动成功.

## Run In k8s

```bash
make k8s-init
make k8s-run
```

