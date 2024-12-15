# pip

-  pip install <package-name> : 安装包
- pip install --upgrade <package-name> : 升级包

# 关键字

| 关键字 | 作用                         |
| ------ | ---------------------------- |
| as     | 用于给对象或异常指定一个别名 |
| raise  | 手动引发（抛出）异常         |
| and    | 逻辑运算符and                |
| pass   | 占位语句                     |



# 符号

| 符号 | 作用 |
| ---- | ---- |
| \|   | 或   |
|      |      |
|      |      |

# 数据类型

| 英文    | 中文   | 概述                                                 |
| ------- | ------ | ---------------------------------------------------- |
| dict    | 字典   | 键值对数据结构                                       |
| int     | 整数   | 表示整数值，没有大小限制                             |
| float   | 浮点数 | 表示带小数点的数值，遵循 IEEE 754 双精度浮点数标准   |
| complex | 复数   | 由实部和虚部组成，用于表示复数                       |
| bool    | 布尔值 | 表示真或假，只有两个可能的值：`True` 和 `False`      |
| `str`   | 字符串 | 用于表示文本数据，支持 Unicode 编码                  |
| `list`  | 列表   | 有序、可变的序列，可以包含不同类型的元素             |
| `tuple` | 元组   | 有序、不可变的序列，通常用于存储异构数据             |
| `dict`  | 字典   | 键值对集合，键必须是唯一的且不可变，值可以是任意类型 |
| set     | 集合   | 无序、不重复的元素集合，支持集合操作如并集、交集等   |

# 语句

| 字符 | 作用                                             |
| ---- | ------------------------------------------------ |
| del  | 用于删除对象、变量、列表中的元素或字典中的键值对 |
|      |                                                  |
|      |                                                  |

# 配置文件

setup.cfg

setup.py 可以读取 setup.cfg 文件中的配置，并将其传递给 setuptools.setup() 函数。

# 标准包

## collections

### 类

#### defaultdict

```bash
一种特殊字典类型，它继承自内置的 dict 类。
defaultdict 在访问一个不存在的键时不会抛出 KeyError，而是自动创建一个默认值。
```



### 函数

#### Counter()

```python
# Counter 可以接受任何可迭代对象（如列表、字符串、元组等），并自动计算其中每个元素出现的次数
counter = Counter('abracadabra')
print(counter)  
# 输出: Counter({'a': 5, 'b': 2, 'r': 2, 'c': 1, 'd': 1})
```

# 内建函数

## sorted()

```py
sorted(iterable, *, key=None, reverse=False)
# 对可迭代对象（如列表、元组、字符串等）进行排序，返回一个新的已排序列表
numbers = [5, 2, 9, 1, 5, 6]
sorted_numbers = sorted(numbers)
print(sorted_numbers)  
# 输出: [1, 2, 5, 5, 6, 9]
```



