

api地址 : `http://localhost:1323/swagger/index.html`

# 注释格式

| 注释   | 作用                            | 例子                    |
| ------ | ------------------------------- | ----------------------- |
| Router | 方法URL                         | // @Router /hello [get] |
| Tags   | 方法标签                        | // @Tags hello          |
| host   | 运行API的主机（主机名或IP地址） | // @host localhost:8080 |

goStack项目生成API命令：

```bash
swag.exe init -g .\main\main.go --parseDependency --exclude .\knife4go\,.\common\secure\utils\,.\metadata\,.\scheduler\
```



