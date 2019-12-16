.PHONY: init
init:
	mkdir -p data/mysql-1 data/mysql-2 data/mysql-3

.PHONY: start
start:
	docker-compose start

.PHONY: stop
stop:
	docker-compose stop

.PHONY: up
up:
	docker-compose up -d

.PHONY: down
down:
	docker-compose down

.PHONY: clean
clean:
	rm -rf ./data/mysql-*/*

.PHONY: exec
exec:
	# docker exec t=mysql-1
	docker exec -it $(t) /bin/bash

.PHONY: backup
backup:
	mysqldump --protocol=tcp -hlocalhost -P36001 -uroot -proot test -r test.sql

.PHONY: k8s-init
k8s-init:
	# 创建模版configmap, 后续使用的时候需要替换
	# $START_GR_ON_BOOT  是否开机自动启用GR
	# $BOOTSTRAP_GR  是否自动初始化组
	# $LOCAL_ADDRESS 本地地址 host:port
	# $SEEDS 集群所有机器
	# $LOCAL_HOST 本地地址
	# $SERVER_ID mysql server id
	kubectl create configmap mysql-mgr --from-file=template.cnf=./cnf/mysql.cnf --from-file=init.sql=./01_init.sql \
	  --from-literal=mgr.port=24901

.PHONY: k8s-run
k8s-run:
	kubectl apply -f ./k8s-mysql-mgr.yaml

.PHONY: k8s-clean
k8s-clean:
	kubectl delete -f ./k8s-mysql-mgr.yaml
	kubectl delete configmap mysql-mgr
	kubectl delete pvc data-mysql-mgr-0 data-mysql-mgr-1 data-mysql-mgr-2
