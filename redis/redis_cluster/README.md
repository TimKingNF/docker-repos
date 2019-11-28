# Redis Cluster

Redis Cluster 方案案例

## Requirements

* docker
* docker-compose
* make
* redis>5.0
* vagrant (OSX)

## Install

> Currently Redis Cluster does not support NATted environments and in general environments where IP addresses or TCP ports are remapped.

因为 Redis Cluster 目前不支持 NATted 环境 和 IP 或 TCP 端口被重新映射的环境. 所以要使用 docker 运行 Redis Cluster 只能使用 --net=host 主机模式. OSX 上的 Docker 不支持Host模式. 所以在OSX上运行, 需要多一个步骤

```bash
make build  # download vagrant box and create vistual machine
make ssh  # vagrant ssh
cd /vagrant
```

接下来的执行都是一致的

```bash
make init
make reset
make up
```

启动完成之后使用 `docker ps` 确认所有 redis 节点都已经启动

最后执行 `make create_cluster` 创建Redis Cluster

注意连接时需要使用 `redis-cli -c` 表示启用 Cluster Mode

### Other Commands

```bash
# 重新分片
redis-cli --cluster reshard 127.0.0.1:27000

# 检查集群状态
redis-cli --cluster check 127.0.0.1:27000

# 让某个节点进入下线状态
redis-cli -p 7002 debug segfault

# 集群添加空的新节点
redis-cli --cluster add-node 127.0.0.1:27006 127.0.0.1:27000

# 添加一个从节点
redis-cli --cluster add-node 127.0.0.1:27006 127.0.0.1:27000 --cluster-slave

# 删除一个节点
redis-cli --cluster del-node 127.0.0.1:7000 `<node-id>`

```

## Reference

* https://redis.io/topics/cluster-tutorial

