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


## 方法

### find()

查询集合中的文档，并且可以通过传递一个查询文档作为参数来实现条件查询。查询文档是一个包含查询条件的 BSON（或 JSON 样式）对象。

#### 简单条件查询

```javascript
db.instance_snapshot.find({status:"FREEZED"}).count()


const (
	PendingSnapshot     InstanceSnapshotStatus = "PENDING"
	FreezedSnapshot     InstanceSnapshotStatus = "FREEZED"
	SnapshotIngSnapshot InstanceSnapshotStatus = "SNAPSHOTING"
	AvailableSnapshot   InstanceSnapshotStatus = "AVAILABLE"
	ErrorSnapshot       InstanceSnapshotStatus = "ERROR"
	DeletingSnapshot    InstanceSnapshotStatus = "DELETING"
	DeletedSnapshot     InstanceSnapshotStatus = "DELETED"
	Restoring           InstanceSnapshotStatus = "RESTORING"
	CreateImage         InstanceSnapshotStatus = "UPLOADINGIMAGE"
	ErrRebuildError     int                    = 5 // rebuild instance.
)

```



```javascript
db.collection.find( <query>, <projection>, <options> )
/*
查询	使用查询运算符指定选择筛选器。要返回集合中的所有文档，请省略此参数或传递空文档 ({})。
投影	指定与查询过滤匹配的文档中要返回的字段。要返回匹配文档中的所有字段，请省略此参数。有关详细信息，请参阅投影。
选项	为查询指定其他选项。这些选项可修改查询行为以及结果的返回方式。有关详细信息，请参阅选项。
```

#### 使用比较操作符

```javascript
// 查询 field 大于 10 的所有文档
db.collection.find({ field: { $gt: 10 } })

// 查询 field 不等于 "value" 的所有文档
db.collection.find({ field: { $ne: "value" } })

// 查询 field 在 ["value1", "value2"] 中的所有文档
db.collection.find({ field: { $in: ["value1", "value2"] } })
```

#### 使用逻辑操作符

```javascript
// 使用 $and 操作符查询满足多个条件的文档
db.collection.find({ $and: [{ field1: "value1" }, { field2: "value2" }] })

// 使用 $or 操作符查询满足任一条件的文档
db.collection.find({ $or: [{ field1: "value1" }, { field2: "value2" }] })
```

#### 使用正则表达式

```javascript
// 查询 field 包含 "pattern" 的所有文档
db.collection.find({ field: /pattern/ })

// 查询 field 以 "start" 开头的所有文档
db.collection.find({ field: /^start/ })
```

#### 使用数组查询操作符

```javascript
// 查询 arrayField 包含 [1, 2, 3] 的所有文档
db.collection.find({ arrayField: { $all: [1, 2, 3] } })

// 查询 arrayField 中至少有一个元素匹配条件的对象
db.collection.find({ arrayField: { $elemMatch: { key: "value" } } })

// 查询 arrayField 的大小为 3 的所有文档
db.collection.find({ arrayField: { $size: 3 } })
```

#### 组合条件

```javascript
// 查询 status 为 "A" 并且 qty 大于 10 或者 items 数组包含 "paper" 的文档
db.collection.find({
  status: "A",
  $or: [
    { qty: { $gt: 10 } },
    { items: "paper" }
  ]
})
```

### sort()

对查询结果进行排序，传递一个包含字段名和排序顺序的对象给它。正数（1）表示升序，负数（-1）表示降序。

```
# 单字段排序
db.collection.find().sort({ field: 1 }) // 升序
db.collection.find().sort({ field: -1 }) // 降序
# 多字段排序
db.collection.find().sort({ field1: 1, field2: -1 })
```

### count()

查询返回文档数目。

```javascript
db.collection.find(query).count()
```

复杂的统计需求，比如分组统计、条件统计等，应该使用聚合管道。聚合管道允许你处理数据流，应用一系列操作符来计算统计数据。

```javascript
db.orders.aggregate([
  { $group: { _id: "$status", count: { $sum: 1 } } }
])
# $group阶段按status字段分组，然后用$sum计算每个状态出现的次数。
```

### limit()

限制查询返回的文档数量，通常与查询操作一起使用，返回指定数量的文档，这对于分页显示结果或仅获取前N个条目非常有用。

```javascript
db.collection.find(query, projection).limit(number)
/*
query：这是可选的条件参数，用于匹配要检索的文档。
projection：这是可选的参数，用于指定返回哪些字段。
number：这是一个必需的整数参数，表示你想要限制的最大返回文档数量。
```

### drop()

```javascript
db.collection_name.drop()
// 删除整个集合（collection）及其所有相关数据、索引和结构的命令
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

# 备份与恢复

`mongodump` 和 `mongorestore`是MongoDB自带的命令行工具。

## 使用 `mongodump` 进行备份

`mongodump` 工具用于创建数据库的备份。它会将数据从MongoDB服务器导出到一个 BSON 文件中，可以针对整个数据库实例、特定数据库或特定集合进行备份。

- 备份所有数据库：
  ```shell
  mongodump
  ```

- 备份指定数据库：
  ```shell
  mongodump --db mydatabase
  ```

- 备份指定集合：
  ```shell
  mongodump --db mydatabase --collection mycollection
  mongodump --db instance --collection instance_backup
  
  
  mongodump --db instance --collection instance_backup \
            --username root --password test \
            --authenticationDatabase admin
  ```
  
- 将备份保存到指定目录：
  ```shell
  mongodump --out /path/to/backup
  ```

## 使用 `mongorestore` 进行恢复

`mongorestore` 工具用于将由 `mongodump` 创建的备份恢复到MongoDB服务器。

- 恢复所有数据库：
  ```shell
  mongorestore
  ```

- 恢复到指定数据库：
  ```shell
  mongorestore --db mydatabase /path/to/backup/mydatabase
  ```

- 恢复指定集合：
  ```shell
  mongorestore --db mydatabase --collection mycollection /path/to/backup/mydatabase/mycollection.bson
  ```

- 从指定目录恢复：
  ```shell
  mongorestore --dir=/path/to/backup
  ```

# aggregate()

在 MongoDB 的聚合框架中，各个阶段的顺序非常重要，因为每个阶段的操作都是基于前一个阶段输出的数据。理解不同阶段之间的顺序可以帮助你构建高效的查询，并确保数据按照预期的方式被处理。

| 符号            | 作用                                                         |
| --------------- | ------------------------------------------------------------ |
| $match          | 过滤文档，只让符合条件的文档通过                             |
| $project        | 选择或重命名字段，或者创建新的计算字段                       |
| $lookup         | 执行类似 SQL 中的 `JOIN` 操作，将来自其他集合的数据合并到当前文档中 |
| $unwind         | 将包含数组的字段拆分为多个文档，每个文档对应数组中的一个元素 |
| $group          | 根据指定的表达式对输入文档进行分组，并计算累积值             |
| $sort           | 根据指定的字段对文档进行排序                                 |
| `$skip  $limit` | 跳过一定数量的文档（`$skip`），或限制返回的文档数量（`$limit`） |
| $count          |                                                              |
|                 |                                                              |
|                 |                                                              |

`$facet`, `$redact`, `$replaceRoot`, `$addFields`, `$bucket`

## project

用于选择（或排除）文档中的字段，从而定义输出文档的结构。

### 操作符

| 操作符        | 作用                 |
| ------------- | -------------------- |
| $toDate       | 把时间戳转化为日期   |
| $dateToString | 把日期转换为特定格式 |
|               |                      |

```javascript
db.instance_snapshot.aggregate([
  {$match: {status: {$ne: "DELETED"}}},
  {$project: {time: {$dateToString: {format: "%Y-%m", date:{$toDate: "$created_at"}}}}},
	{$group: {_id: "$time", count: {$sum: 1}}},
  {$project: {count: 1}},
])
```

## group

根据指定的字段对文档进行分组，并且可以在每个分组的基础上执行聚合计算。使用 $group 可以帮助你汇总数据、计算统计数据（如总和、平均值、最大值、最小值等）、计数以及更多复杂的数据分析任务。

```javascript
{
  $group: {
    _id: <expression>, // 必需，用于定义分组条件
    <field1>: { <accumulator1> : <expression> },
    ...
    <fieldN>: { <accumulatorN> : <expression> }
  }
}
```

- `_id`: 定义了分组的标准，可以是一个字段名（例如 `"$status"`），也可以是一个表达式。如果 `_id` 设置为 `null`，则所有的输入文档将被合并成一个单一的输出文档。
- `<field>` 和 `<accumulator>`：除了 `_id` 外，你可以定义多个字段来保存每组的计算结果。这些字段通常结合累积器表达式来实现各种统计功能。

### 累积器表达式

| 表达式    | 作用                                                        |
| --------- | ----------------------------------------------------------- |
| $sum      | 对数值字段求和，或者用于计数（当与常量 1 或 -1 一起使用时） |
| $avg      | 计算平均值                                                  |
| $min      | 找出最小值                                                  |
| $max      | 找出最大值                                                  |
| $addToSet | 收集数组中的唯一值                                          |
| $push     | 收集数组中的所有值                                          |
| $first    | 获取分组中第一个文档的字段值                                |
| $last     | 获取分组中最后一个文档的字段值                              |

## sort

用于对输入文档进行排序。它可以在聚合管道的任何位置使用，但通常会在最后或接近最后的位置，以便在数据处理完成后再进行排序。$sort 使用一个包含字段和排序方向（升序或降序）的文档作为参数。

```javascript
{ $sort: { <field1>: <order>, <field2>: <order> ... } }
```

- `<field>`：指定要排序的字段名。
- `<order>`：可以是 `1` 或 `-1`，分别表示升序（ascending）和降序（descending）。

# find()

```bash
db.users.find({ age: { $gt: 20 } }, { name: 1, age: 1 })
```

```bash
db.collection.find(query, projection)

# query: （可选）查询条件，指定需要匹配的文档。如果为空 {}，则返回集合中所有的文档。
# projection: （可选）指定需要返回哪些字段。如果不提供，默认返回所有字段。
```

# update

## updateOne()

```bash
db.collection.updateOne(filter, update, options) # 用于更新集合中第一个符合查询条件的文档

# filter: 查询条件，表示要更新的文档
# update: 更新操作，指定如何更新文档
# options: （可选）一些附加选项，如 upsert
```

```javascript
db.users.updateOne(
  { name: "Alice" },               // 查询条件：查找 name 为 "Alice" 的文档
  { $set: { age: 30 } }            // 更新操作：将 age 设置为 30
)
```

## updateMany()

```bash
db.collection.updateMany(filter, update, options) #用于更新集合中所有符合条件的文档

# filter: 查询条件，表示要更新的文档
# update: 更新操作，指定如何更新文档
# options: （可选）一些附加选项，如 upsert
```

```javascript
db.users.updateOne(
  { name: "Alice" },               // 查询条件：查找 name 为 "Alice" 的文档
  { $set: { age: 30 } }            // 更新操作：将 age 设置为 30
)
```

## replaceOne()

```bash
db.collection.replaceOne(filter, replacement, options)
# 用于将符合条件的文档完全替换为新的文档。
# 这个操作将会删除原文档中的所有字段，然后替换为新文档（只有替换文档中指定的字段会保留）

# filter: 查询条件，表示要替换的文档
# replacement: 新文档，作为替代的文档
# options: （可选）一些附加选项，如 upsert
```

```javascript
db.users.replaceOne(
  { name: "Alice" },         // 查询条件：查找 name 为 "Alice" 的文档
  { name: "Alice", age: 30 } // 替换文档：将整个文档替换为{name: "Alice", age: 30}
)
```

## 更新操作符

| 操作符    | 作用                                             |
| --------- | ------------------------------------------------ |
| $set      | 更新指定字段的值。如果字段不存在，则会添加该字段 |
| $unset    | 删除指定字段                                     |
| $inc      | 增加指定字段的值                                 |
| $push     | 向数组中添加元素                                 |
| $addToSet | 向数组中添加元素，但不会添加重复值               |
| $pop      | 删除数组的第一个或最后一个元素                   |
| $rename   | 重命名字段                                       |

