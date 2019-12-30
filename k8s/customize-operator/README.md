# Customize Operator

Operator 测试

## Requirements

* k8s 1.14.8

## Install

```bash
go mod download  # 下载依赖
bash generate_groups.sh. # 编译k8s client 代码
go build -o controller .  # 编译生成程序
```

最后启动编译好的程序

```bash
./controller -kubeconfig=$HOME/.kube/config -alsologtostderr=true
```

## Usage

```bash
kubectl apply -f crd/mysql-gr-operator.yaml  # 注册CRD
kubectl apply -f example/example-mysql-gr-operator.yaml  # 创建CRD对象
```

