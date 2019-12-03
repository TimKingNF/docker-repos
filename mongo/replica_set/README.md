# Mongo Replica Set

Mongo DB Replica Set 案例

## Requirement

* docker
* docker-compose
* make

## Install

执行以下命令

```bash
make init
make run
make create_cluster
make connect # 连接集群
```

连接集群之后, 注意如果当前连接的是 Seconday 节点, 默认是不可读的, 需要手动执行 `rs.slaveOk()` , 只在当前会话生效