# Spline Test

Spark 血缘关系组件 spline 测试

## Requirements

* [docker-spark](https://github.com/big-data-europe/docker-spark)
* docker
* docker-compose
* make
* [spline](https://absaoss.github.io/spline/)
* mongo

## Before Run

### docker-spark

在运行之前, 首先需要准备 docker-spark 环境. 这里我们直接使用开源项目 [docker-spark](https://github.com/big-data-europe/docker-spark) 作为spark基本环境. 

搭建步骤也很简单, 执行以下命令即可

```bash
git clone https://github.com/big-data-europe/docker-spark
cd docker-spark
docker-compose up
```

然后我们需要编译运一个运行 spark-python App 的基本镜像, [docker-spark](https://github.com/big-data-europe/docker-spark) 也给我们提供好了, 进入 `template/python` 目录, 执行以下命令

```bash
docker build --rm -t spark-app:python .  # 创建spark-app:python镜像
```

最后回到本项目根路径下, 执行以下命令, 测试 spark-python App 是否能够正常运行

```bash
make build
make submit_test
```

最后运行出现如下内容, 说明 spark-python App 运行成功

```bash
========================================
result=DataFrame[(1 + 1): int]
```

### mongodb

spline 需要支持2种持久化方式, hdfs, mongo. 这里我们直接使用 mongodb. 容器命名为 mongo, 并绑定到本机中, 执行以下命令

```bash
docker run --name mongo -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=root -p 27017:27017 -d mongo
```

最后创建我们所需的数据库 lineage, 并设置好读写用户

```bash
mongo -uroot -proot
> use lineage
> db.createUser({ user: "root", pwd: "root", roles: [{ role: "readWrite",db: "lineage" }] })
```

### Spline Web

前往 [spline](https://absaoss.github.io/spline/) 官网下载 [Spline Web UI executable JAR](https://search.maven.org/remotecontent?filepath=za/co/absa/spline/spline-web/0.3.9/spline-web-0.3.9-exec-war.jar), 放在 `jars` 目录下即可

## Run

### Python

spline python 的运行相对要复杂一点, 首先我们需要下载 spline 的源码, 编译我们所需的支持 python 的 jar 包.

```bash
git clone https://github.com/AbsaOSS/spline.git
cd spline
git checkout -b 0.3.9 release/0.3.9  # 切换到稳定版分支
mvn package -P spark-2.4,shade -DskipTests  # 根据spark版本执行编译
```

最后生成的我们需要的jar包文件. 文件路径为 `sample/target/spline-sample-0.3.9.jar` 

将 `spline-sample-0.3.9.jar` 复制到 `项目根路径/jars` 目录下. 

然后将 `data` 目录复制到 `docker-worker-1` 的 `/home` 目录下, 这一步是将测试数据放到 spark-worker-1 中, 因为我们没有使用 hdfs, 所以需要用这种方式让 `saprk-worker-1` 读取本地文件, 最后执行以下命令启动 spline web

```bash
cd spline_test  # 切换为项目根目录

docker cp ./data spark-worker-1:/home  # 可以进去spark-worker-1确认是否成功

java -jar ./jars/spline-web-0.3.9-exec-war.jar -Dspline.mongodb.url=mongodb://root:root@127.0.0.1:27017/lineage -Dspline.mode=REQUIRED -httpPort=8088  # 启动spline web
```

启动完成后访问[网页](http://127.0.0.1:8088)确认启动成功

最后执行 `make submit_dev` 观察是否执行成功. [spline网页](http://127.0.0.1:8088) 是否有显示已完成的作业的血缘关系图.

### Scala

### Python

## Quick Start

快速直接启动所有环境, 直接执行 `docker-compose up -d` 

下载 [docker-spark](https://github.com/big-data-europe/docker-spark) 编译 `spark-app:python` 镜像

然后执行 `make submit` 提交任务