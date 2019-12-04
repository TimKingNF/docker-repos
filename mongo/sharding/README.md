# MongoDB Sharding

MongoDB Sharding(分片) 测试案例

## Requirement

* docker
* docker-compose
* make

## Install

在本案例中, 我们构建的集群如下所示, 一共7个Container

* shard-1
  * shard-1-mongo-1
  * shard-1-mongo-2
  * shard-1-mongo-3
* shard-2
  * shard-2-mongo-1
* ConfigServer
  * mongo-config-1
  * mongo-config-2

* mongos

执行以下命令, 创建 Sharding 集群

```bash
make init  # 初始化文件目录
make run # 启动 shard-1, shard-2, ConfigServer 等6个容器
make shard1_init  # shard1 3个容器组成副本集
make shard2_init  # shard2 副本集初始化
make config_init  # ConfigServer副本集初始化
make run_mongos  # 启动 mongos 容器, 会自动连上 ConfigServer
make sharding_init  # 连接mongos, 添加shard-1和shard-2
```

最后我们执行 `make connect_mongos` , 执行 `sh.status()` 检查是否成功, 创建过程中, 可以执行 `rs.status()` 确定副本集是否配置成功

## Tests

最后我们测试分片策略, 执行 `make put_data` , 会创建 `test` 数据库, 并往其中的 `t` 集合写入1000条数据, 我们可以通过访问 shard-1, shard-2 , 执行 `db.t.find().count()` 查询 `t` 集合的写入数量