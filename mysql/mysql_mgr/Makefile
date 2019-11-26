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
