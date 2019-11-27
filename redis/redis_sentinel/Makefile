.PHONY: init
init:
	mkdir -p data/redis-1 data/redis-2 data/redis-3

.PHONY: up
up:
	docker-compose up -d

.PHONY: start
start:
	docker-compose start

.PHONY: down
down:
	docker-compose down

.PHONY: clean
clean:
	rm -rf data/redis-*/*

.PHONY: stop
stop:
	docker-compose stop

.PHONY: restart
restart:
	docker-compose restart

.PHONY: exec
exec:
	docker exec -it $(t) /bin/bash

.PHONY: reset
reset:
	cp -f cnf/redis-1/sentinel.conf.bak cnf/redis-1/sentinel.conf
	cp -f cnf/redis-1/sentinel.conf.bak cnf/redis-2/sentinel.conf
	cp -f cnf/redis-1/sentinel.conf.bak cnf/redis-3/sentinel.conf
	cp -f cnf/redis-1/redis.conf.bak cnf/redis-1/redis.conf
	cp -f cnf/redis-1/redis.conf.bak cnf/redis-2/redis.conf
	cp -f cnf/redis-1/redis.conf.bak cnf/redis-3/redis.conf
	echo "slaveof redis-1 6379" >> cnf/redis-2/redis.conf
	echo "slaveof redis-1 6379" >> cnf/redis-3/redis.conf
