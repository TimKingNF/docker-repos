#!/usr/bin/env bash

redis-server /usr/local/etc/redis/redis.conf

# 保证容器不退出
while true; do
  sleep 1s
done
