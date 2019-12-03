.PHONY: init
init:
	mkdir -p data/mongo-1 data/mongo-2 data/mongo-3

.PHONY: run
run:
	docker-compose up -d

.PHONY: down
down:
	docker-compose down

.PHONY: stop
stop:
	docker-compose stop

.PHONY: clean
clean:
	rm -rf data/mongo-*/*

.PHONY: create_cluster
create_cluster:
	docker exec -it mongo-1 mongo mongo-1:27017/admin --eval "rs.initiate()"
	docker exec -it mongo-1 mongo mongo-1:27017/admin --eval 'rs.add("mongo-2:27017")'
	docker exec -it mongo-1 mongo mongo-1:27017/admin --eval 'rs.add("mongo-3:27017")'

.PHONY: connect
connect:
	docker exec -it mongo-1 mongo "mongodb://mongo-1:27017,mongo-2:27017,mongo-3:27017/testdb?readPreference=nearest"
