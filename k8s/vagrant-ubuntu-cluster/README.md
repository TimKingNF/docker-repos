#k8s-vagrant-ubuntu-cluster

k8s 集群搭建, 运行在ubuntu下

## Requirements

* vagrant

## Install

```bash
# vagrant plugin install
vagrant plugin install vagrant-disksize
vagrant plugin install vagrant-proxyconf

# ubuntu 1.18 box download
vagrant box add https://mirrors.tuna.tsinghua.edu.cn/ubuntu-cloud-images/bionic/current/bionic-server-cloudimg-amd64-vagrant.box --name ubuntu/bionic

vagrant up # 启动并安装虚拟机
```

然后执行 `vagrant ssh node1` 登陆 node1 , 执行如下命令

```bash
cd /vagrant
bash master.sh
```

待确认kubeadm 安装成功且网络配置插件启动完毕之后, 返回宿主机执行 `vagrant ssh node2` 登陆 node2, 启动slave

```bash
cd /vagrant
bash slave.sh
```

**Tips**

* 启动之后使用 `sudo passwd root` 修改root密码
* 如果需要启用代理, 可以修改 `master.sh` 与 `slave.sh` 中的 `proxy` 的值

最后执行命令 `kubectl get pod -n kube-system` 检查是否所有Pod是否都启动成功, 执行 `kubectl get nodes` 查看是否所有Node 都启动成功.

