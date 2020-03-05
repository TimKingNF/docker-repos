#!/usr/bin/env bash
set -ex

# export http_proxy="http://192.168.1.4:10087"
# export https_proxy="http://192.168.1.4:10087"

# change time zone
cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
timedatectl set-timezone Asia/Shanghai

mv /etc/apt/sources.list /etc/apt/sources.list.bak
cat >> /etc/apt/sources.list <<EOF
# 默认注释了源码镜像以提高 apt update 速度，如有需要可自行取消注释
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ bionic main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ bionic main restricted universe multiverse
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ bionic-updates main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ bionic-updates main restricted universe multiverse
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ bionic-backports main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ bionic-backports main restricted universe multiverse
deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ bionic-security main restricted universe multiverse
# deb-src https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ bionic-security main restricted universe multiverse
EOF

# 添加k8s源
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
cat > /etc/apt/sources.list.d/kubernetes.list <<EOF
# deb http://apt.kubernetes.io/ kubernetes-xenial main
deb http://mirrors.ustc.edu.cn/kubernetes/apt kubernetes-xenial mai
EOF

# 安装必要的程序
apt-get update
apt-get install -y wget vim docker.io kubeadm

# 写入host
echo 'set host name resolution'
cat >> /etc/hosts <<EOF
172.17.8.101 node1
172.17.8.102 node2
172.17.8.103 node3
EOF

# 设置docker加速
cat > /etc/docker/daemon.json <<EOF
{
  "registry-mirrors" : ["https://2wzhnbjj.mirror.aliyuncs.com"]
}
EOF
systemctl start docker
systemctl enable docker
systemctl enable kubelet
