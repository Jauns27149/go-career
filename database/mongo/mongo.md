MongoDB 是一种流行的非关系型数据库（NoSQL），特别适合处理大规模的数据集。它采用文档存储模型，数据以键值对的形式存储在文档中，每个文档对应一条记录。MongoDB 的设计目标是提供高性能、高可用性和易于扩展的数据存储解决方案。

# mongosh

- 连接mongo

  - mongosh : 连接本地默认数据库
  - mongosh "mongodb://localhost:27017" : 连接特定数据库

- 命令

  - show 
    - dbs : 查看所有数据库
    - collections : 查看当前数据库中的所有集合
  - use mydb : 连接到数据库

- 集合操作

  ```bash
  # 操作当前数据库下的user集合
  # 查询
  db.users.find() 
  # 插入
  db.users.insertOne({ name: "Alice", age: 30 }) 
  db.users.insertMany([
    { name: "Bob", age: 25 },
    { name: "Charlie", age: 28 }
  ])
  # 删除
  db.users.deleteOne({ name: "Alice" }) 
  db.users.deleteMany({ age: { $lt: 30 } })
  # 更新
  db.users.updateOne(
    { name: "Alice" },
    { $set: { age: 31 } }
  )
  db.users.updateMany(
    { age: { $lt: 30 } },
    { $set: { status: "young" } }
  )
  ```

  

# 运算符

## 比较运算符

| 符号 | 作用                                 |
| ---- | ------------------------------------ |
| $eq  | 等于（Equality）                     |
| $ne  | 不等于（Not equal）                  |
| $gt  | 大于（Greater than）                 |
| $gte | 大于等于（Greater than or equal to） |
| $lt  | 小于（Less than）                    |
| $lte | 小于等于（Less than or equal to）    |
| $in  | 在某个集合中（In）                   |
| $nin | 不在某个集合中（Not in）             |

## 逻辑运算符

| 符号 | 作用        |
| ---- | ----------- |
| $and | 与（AND）   |
| $or  | 或（OR）    |
| $not | 非（Not）   |
| $nor | 非或（NOR） |

## 正则表达式

| 符号     | 作用           |
| -------- | -------------- |
| $regex   | 正则表达式匹配 |
| $options | 正则表达式选项 |

## 字段存在性运算符

| 符号    | 作用         |
| ------- | ------------ |
| $exists | 字段是否存在 |
| $type   | 字段类型     |

## 数组运算符

| 符号       | 作用                 |
| ---------- | -------------------- |
| $all       | 匹配所有指定元素     |
| $elemMatch | 匹配数组中的一个文档 |
| $size      | 匹配数组的大小       |

### 更新运算符

| 符号   | 作用             |
| ------ | ---------------- |
| $set   | 设置字段的值     |
| $unset | 删除字段         |
| $inc   | 增加（增量更新） |
| $push  | 向数组中添加元素 |
| $pull  | 从数组中删除元素 |

