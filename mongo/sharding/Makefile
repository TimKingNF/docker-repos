.PHONY: init
init:
	mkdir -p data/mongo-1 data/mongo-2 data/mongo-3 data/mongo-4 data/mongo-5 data/mongo-6 data/mongo-7

.PHONY: run
run:
	docker-compose up -d

.PHONY: run_mongos
run_mongos:
	docker-compose -f mongos.yml up -d

.PHONY: down
down:
	docker-compose down

.PHONY: down_mongos
down_mongos:
	docker-compose -f mongos.yml down

.PHONY: stop
stop:
	docker-compose stop

.PHONY: clean
clean:
	rm -rf data/mongo-*/*

.PHONY: shard1_init
shard1_init:
	docker exec -it shard-1-mongo-1 mongo shard-1-mongo-1:27018/admin --eval 'rs.initiate({_id:"shard-1",members:[{_id:0,host:"shard-1-mongo-1:27018"}, {_id:1,host:"shard-1-mongo-2:27018"}, {_id:2,host:"shard-1-mongo-3:27018"}]})'

.PHONY: shard2_init
shard2_init:
	docker exec -it shard-2-mongo-1 mongo shard-2-mongo-1:27018/admin --eval 'rs.initiate({_id:"shard-2",members:[{_id:0,host:"shard-2-mongo-1:27018"}]})'

.PHONY: config_init
config_init:
	docker exec -it mongo-config-1 mongo mongo-config-1:27019/admin --eval 'rs.initiate({_id:"config",members:[{_id:0,host:"mongo-config-1:27019"}, {_id:1,host:"mongo-config-2:27019"}]})'

.PHONY: sharding_init
sharding_init:
	docker exec -it mongos mongo --port 27017 --eval 'sh.addShard("shard-1/shard-1-mongo-1:27018,shard-1-mongo-2:27018,shard-1-mongo-3:27018")'
	docker exec -it mongos mongo --port 27017 --eval 'sh.addShard("shard-2/shard-2-mongo-1:27018")'


.PHONY: connect_mongos
connect_mongos:
	docker exec -it mongos mongo --port 27017

.PHONY: put_data
put_data:
	docker cp test.js mongos:/
	docker exec -it mongos mongo --port 27017 /test.js
