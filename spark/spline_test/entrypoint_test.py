#! coding=utf-8

""" spark 测试程序  """

from pyspark.sql import SparkSession

spark = SparkSession.builder \
    .getOrCreate()

ret = spark.sql("select 1+1")
print("=" * 40)
print("result=%s" % ret)
