#!/usr/bin/env bash

# 在其他节点上执行
if [[ `hostname` == "node1" ]]; then
  exit 1
fi

proxy="http://192.168.1.3:10087/"

# 挂代理然后加入节点
sudo -- bash -c "
mv /etc/default/docker /etc/default/docker.bak
mkdir -p /etc/systemd/system/docker.service.d
echo '
Environment=\"HTTPS_PROXY=$proxy\"
Environment=\"HTTP_PROXY=$proxy\"
' >> /etc/systemd/system/docker.service.d/http-proxy.conf
systemctl daemon-reload
systemctl restart docker

# 创建 join-config.yaml
kubeadm config print join-defaults > /tmp/join.config.yaml
sed -i 's/kube-apiserver/172.17.8.101/g' /tmp/join.config.yaml
kubeadm join --config=/tmp/join.config.yaml
"
