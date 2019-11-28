.PHONY: build
build:
	vagrant up

.PHONY: ssh
ssh:
	vagrant ssh

.PHONY: init
init:
	mkdir -p data/redis-1 data/redis-2 data/redis-3 data/redis-4 data/redis-5 \
	data/redis-6 data/redis-7
	mkdir -p cnf/redis-1 cnf/redis-2 cnf/redis-3 cnf/redis-4 cnf/redis-5 \
	cnf/redis-6 cnf/redis-7

.PHONY: reset
reset:
	cp -f ./redis.conf.bak cnf/redis-1/redis.conf
	echo "port 27000" >> cnf/redis-1/redis.conf
	cp -f ./redis.conf.bak cnf/redis-2/redis.conf
	echo "port 27001" >> cnf/redis-2/redis.conf
	cp -f ./redis.conf.bak cnf/redis-3/redis.conf
	echo "port 27002" >> cnf/redis-3/redis.conf
	cp -f ./redis.conf.bak cnf/redis-4/redis.conf
	echo "port 27003" >> cnf/redis-4/redis.conf
	cp -f ./redis.conf.bak cnf/redis-5/redis.conf
	echo "port 27004" >> cnf/redis-5/redis.conf
	cp -f ./redis.conf.bak cnf/redis-6/redis.conf
	echo "port 27005" >> cnf/redis-6/redis.conf
	cp -f ./redis.conf.bak cnf/redis-7/redis.conf
	echo "port 27006" >> cnf/redis-7/redis.conf

.PHONY: up
up:
	docker-compose up -d

.PHONY: create_cluster
create_cluster:
	redis-cli --cluster create 127.0.0.1:27000 127.0.0.1:27001 127.0.0.1:27002 127.0.0.1:27003 127.0.0.1:27004 127.0.0.1:27005 --cluster-replicas 1

.PHONY: down
down:
	docker-compose down

.PHONY: stop
stop:
	docker-compose stop

.PHONY: clean
clean:
	rm -rf data/redis-*/*
