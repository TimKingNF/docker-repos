sh.status()
print("========== Put Some data")

sh.enableSharding("test") // 允许test数据库进行sharding
sh.shardCollection("test.t",{id:"hashed"}) // 对test.t集合以id列为shard key进行hashed sharding
// 通过db.t.getIndexes()可以看到自动为id列创建了索引。

// 其他分片方式
// sh.shardCollection("test.t",{id:1}) // 对test.t集合以id列为shard key进行ranged sharding

db = db.getSiblingDB('test')
// db.d.insert({id:1, name:"test"})

for (var i=1; i <= 1000; i++) {
  db.t.insert({id:i,name:"Leo"})  // 会将t集合的数据相对均衡的散列到两个分片中
}

print("========== Shard Stats")
// db.t.stats()  // 连接上 mongos 切换到test数据库,执行该命令查看test.t的分片信息
