# 继承spark-app应用模版
FROM spark-app:python

LABEL maintainer="Tim King <timking.nf@foxmail.com>"

ENV SPARK_MASTER_NAME spark-master
ENV SPARK_MASTER_PORT 7077
# default: /app/app.py
ENV SPARK_APPLICATION_PYTHON_LOCATION /app/entrypoint.py
ENV SPARK_APPLICATION_ARGS ""

# spline require
ENV SCALA_VERSION 2.11.12
ENV SCALA_HOME /usr/share/scala

# 安装 scala, 并加scala 的jar包添加到spark jars目录下
RUN apk add --no-cache --virtual=.build-dependencies wget ca-certificates && \
    apk add --no-cache bash curl jq && \
    cd "/tmp" && \
    wget --no-verbose "https://downloads.typesafe.com/scala/${SCALA_VERSION}/scala-${SCALA_VERSION}.tgz" && \
    tar xzf "scala-${SCALA_VERSION}.tgz" && \
    mkdir "${SCALA_HOME}" && \
    rm "/tmp/scala-${SCALA_VERSION}/bin/"*.bat && \
    mv "/tmp/scala-${SCALA_VERSION}/bin" "/tmp/scala-${SCALA_VERSION}/lib" "${SCALA_HOME}" && \
    ln -s "${SCALA_HOME}/bin/"* "/usr/bin/" && \
    apk del .build-dependencies && \
    rm -rf "/tmp/"* && \
    cp "${SCALA_HOME}/lib/"* /spark/jars

CMD ["/bin/bash", "/submit.sh"]
