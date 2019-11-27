#!/usr/bin/env bash

# 启动 redis server, 会持续运行
redis-server /usr/local/etc/redis/redis.conf

sleep 5s

# 启动 redis sentinel
redis-sentinel /usr/local/etc/redis/sentinel.conf

# 保证容器不退出
while true; do
  sleep 1s
done
