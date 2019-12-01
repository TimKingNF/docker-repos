#! coding=utf-8

""" spline 测试程序 """

import time
import os
from pyspark.sql import SparkSession

spark = SparkSession.builder \
    .getOrCreate()

# 这里依赖client所在机器, 即driver端含scala2.11的jar包
sc = spark.sparkContext
sc._jvm.za.co.absa.spline.core.SparkLineageInitializer.enableLineageTracking(spark._jsparkSession)

# 这一步是因为分布式的情况下, 提交任务client模式下和worker都需要有同样的文件
# 如果用cluser模式, 则会因为sc在spark-master中, 从而找不到entrypoint.py文件
os.system("cp -r /app/data /home")

wikidata_file = "file:///home/data/input/wikidata.csv"
domain_file = "file:///home/data/input/domain.csv"
output_path = "file:///home/data/output"

source_ds = spark.read \
        .option("header", "true") \
        .option("inferSchema", "true") \
        .csv(wikidata_file) \
        .alias("source") \
        .filter("total_response_size > 1000") \
        .filter("count_views > 10")

domain_ds = spark.read \
        .option("header", "true") \
        .option("inferSchema", "true") \
        .csv(domain_file) \
        .alias("mapping")

joined_ds = source_ds.join(domain_ds,
                           source_ds.domain_code == domain_ds.d_code,
                           'left_outer')
joined_ds = joined_ds.select(joined_ds.page_title.alias("page"),
                             joined_ds.d_name.alias("domain"),
                             joined_ds.count_views)

joined_ds.write.mode("overwrite").parquet(output_path)

