#!/usr/bin/env bash

# 检查是否有环境变量
if [[ ! $GOPATH ]]; then
  echo "Need env GOPATH..." && exit 1
fi

ORG_PACKAGE=datahub.txzing.com  # 组织
OBJECT_NAME=mysql-gr-operator  # crd 对象名字

mkdir -p $GOPATH/src/$ORG_PACKAGE
PWD_PATH=`pwd`

if [[ ! -h $GOPATH/src/$ORG_PACKAGE/mysql-gr-crd ]]; then
  ln -s $PWD_PATH $GOPATH/src/$ORG_PACKAGE/$OBJECT_NAME
fi

# 代码生成的工作目录, 也就是项目路径
ROOT_PACKAGE="$ORG_PACKAGE/$OBJECT_NAME"
# API Group
CUSTOM_RESOURCE_NAME="txz_datahub"
# API Version
CUSTOM_RESOURCE_VERSION="v1"

TOOL_PATH=$GOPATH/src/k8s.io/code-generator

if [[ ! -d $TOOL_PATH ]]; then
  cd /tmp
  env GO111MODULE="" go get -d k8s.io/code-generator  # 下载源码
  cd $TOOL_PATH
  git checkout kubernetes-1.14.8  # 切换到对应的k8s版本
  go mod init  # 这里因为该分支使用Godep, 切换为go modules
  cd $PWD_PATH
fi

# 执行代码自动生成, pkg/client 是目标目录, pkg/apis 是类型定义目录
$TOOL_PATH/generate-groups.sh all $ROOT_PACKAGE/pkg/client \
  $ROOT_PACKAGE/pkg/apis \
  $CUSTOM_RESOURCE_NAME:$CUSTOM_RESOURCE_VERSION
