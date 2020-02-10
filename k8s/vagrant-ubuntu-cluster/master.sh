#!/usr/bin/env bash

# 在宿主机上执行
if [[ `hostname` != "node1" ]]; then
  exit 1
fi

proxy="http://192.168.3.125:10087/"

# 挂代理拉取镜像
sudo -- bash -c "
mv /etc/default/docker /etc/default/docker.bak
mkdir -p /etc/systemd/system/docker.service.d
echo '
Environment=\"HTTPS_PROXY=$proxy\"
Environment=\"HTTP_PROXY=$proxy\"
' >> /etc/systemd/system/docker.service.d/http-proxy.conf
systemctl daemon-reload
systemctl restart docker
kubeadm config images pull --config=/vagrant/init-default.yaml
"

# kubeadm init
sudo kubeadm init --config=/vagrant/init-default.yaml

# 配置
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config

# 挂代理安装网络插件
# version="$(kubectl version | base64 | tr -d '\n')"
sudo -- bash -c "
export http_proxy=$proxy
export https_proxy=$proxy
export HTTP_PROXY=$proxy
export HTTPS_PROXY=$proxy
git clone https://github.com/coredns/deployment.git
"

cd deployment/kubernetes
./deploy.sh | kubectl apply -f -


