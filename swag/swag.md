

api地址 : `http://localhost:1323/swagger/index.html`

# 注释格式

| 注释   | 作用                                             | 例子                    |
| ------ | ------------------------------------------------ | ----------------------- |
| Router | 方法URL                                          | // @Router /hello [get] |
| Tags   | 方法标签                                         | // @Tags Example        |
| host   | 运行API的主机（主机名或IP地址）                  | // @host localhost:8080 |
| id     | 用于标识操作的唯一字符串,在所有API操作中必须唯一 | // @id hello            |

goStack项目生成API命令：

```bash
swag.exe init -g .\main\main.go --parseDependency --exclude .\knife4go\,.\common\secure\utils\,.\metadata\,.\scheduler\
```

# swag cli

- swag 
  - init
    - -g --generalInfo <path> : API通用信息人口， (默认: "main.go")
    -  --exclude value  <path> : 解析扫描时排除的目录，多个目录可用逗号分隔（默认：空）
    - --parseDependency : 是否解析依赖目录中的go源文件，默认不
    - -o --output <path> : 文件输出目录 (默认: "./docs")

# Knife4go

文档URL 

```bash
http://127.0.0.1:8080/knife/doc.html
```



