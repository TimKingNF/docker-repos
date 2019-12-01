.PHONY: submit_test
submit_test:
	docker run -v $(shell pwd)/entrypoint_test.py:/app/entrypoint.py \
	  --name spark_test -e ENABLE_INIT_DAEMON=false \
	  --link spark-master:spark-master \
	  --net=docker-spark_default my_spark_app:spline_test

.PHONY: submit_dev
submit_dev:
	docker run -v $(shell pwd)/:/app \
	  --name spline_test \
	  --link spark-master:spark-master \
	  -e SPARK_SUBMIT_ARGS="--jars /app/jars/spline-sample-*.jar --deploy-mode client --driver-java-options=\"-Dspline.mode=REQUIRED -Dspline.persistence.factory=za.co.absa.spline.persistence.mongo.MongoPersistenceFactory -Dspline.mongodb.url=mongodb://root:root@192.168.1.6:27017/lineage\"" \
	  --net=docker-spark_default my_spark_app:spline_test \
	  /bin/bash /app/submit.sh

.PHONY: run
run:
	docker-compose up -d

.PHONY: submit
submit:
	docker run -v $(shell pwd)/:/app \
	  --name spline_test \
	  -e SPARK_MASTER_NAME=spline-spark-master \
	  -e SPARK_SUBMIT_ARGS="--jars /app/jars/spline-sample-*.jar --deploy-mode client --driver-java-options=\"-Dspline.mode=REQUIRED -Dspline.persistence.factory=za.co.absa.spline.persistence.mongo.MongoPersistenceFactory -Dspline.mongodb.url=mongodb://spline-mongo:27017/lineage\"" \
	  --net=spline_test_default my_spark_app:spline_test \
	  /bin/bash /app/submit.sh

.PHONY: build
build:
	docker build -t my_spark_app:spline_test .

.PHONY: clean
clean:
	docker rm spline_test

.PHONY: stop
stop:
	docker stop spline_test

.PHONY: exec
exec:
	docker exec -it spline_test /bin/bash
