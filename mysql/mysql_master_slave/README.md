# Mysql Master Slave

mysql 主从复制, k8s 配置案例

## Requirements

* kubernetes
* docker

## Install

```bash
kubectl create -f conf.yaml # 创建ConfigMap
kubectl create -f svc.yaml  # 创建Service
kubectl create -f master_slave.yaml # 创建StatefulSet
```

如果需要使用 StorageClass , 则需要先下载 ceph 插件. ()

```bash
# 使用ceph插件 . 注意需要本机支持安装 ceph
kubectl apply -f https://raw.githubusercontent.com/rook/rook/v1.2.0/cluster/examples/kubernetes/ceph/common.yaml
kubectl apply -f https://raw.githubusercontent.com/rook/rook/v1.2.0/cluster/examples/kubernetes/ceph/operator.yaml
kubectl create -f storage_class.yaml # 创建PV相关
# 然后打开 master_slave.yaml 中的 storageClassName 注释
# 启动 StatefulSet
```

## Tests

启动完成之后, 执行如下命令, 测试集群是否配置成功

```bash
# 往主库写数据
kubectl run mysql-client --image=mysql:5.7 -i --rm --restart=Never -- \
mysql -h mysql-master-slave-0.mysql <<EOF
CREATE DATABASE test;
CREATE TABLE test.messages (message VARCHAR(250));
INSERT INTO test.messages VALUES ('hello');
EOF

# 通过读取服务读取
kubectl run mysql-client --image=mysql:5.7 -i -t --rm --restart=Never -- \
mysql -h mysql-read -e "SELECT * FROM test.messages"
```



