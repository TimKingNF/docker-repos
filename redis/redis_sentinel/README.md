# Redis Sentinel

Redis Sentinel (哨兵) 服务搭建

## Requirements

* docker
* docker-compose
* make
* redis-cli

## Install

```bash
make init
make reset
make up
```

## Tests

使用 `redis-cli -p 26379` 连接上 redis-1

使用 `redis-cli -p 36379 sentinel get-master-addr-by-name redis-master` 获取Redis Master 地址

使用 `redis-cli -p 36379 sentinel slaves redis-master` 获取 Redis Slave 信息

## k8s run

```bash
make k8s_run
```

