# 命令

| 命令                        | 作用                             |
| --------------------------- | -------------------------------- |
| docker ps                   | 列出当前正在运行的容器           |
| docker images               | 列出本地机器上的所有 Docker 镜像 |
| docker rm <NAMES>           | 删除一个或多个已停止的容器       |
| docker rmi                  | 删除一个或多个镜像               |
| docker build  -t name:tag . | 利用Dockerfile文件打包镜像       |
| docker start <ID or NAME>   | 启动容器                         |
| docker stop <containerName> | 停止容器运行                     |

## docker exec

```bash
docker exec [OPTIONS] CONTAINER COMMAND [ARG...]
# 在已经运行的 Docker 容器中执行命令
docker exec -it CONTAINER_NAME_OR_ID /bin/bash
#  进入容器并启动一个交互式的 Bash shell
```

- options
  - -d 或 --detach: 在后台运行命令
  - -i 或 --interactive: 保持 STDIN 打开，即使没有附加
  - -t 或 --tty: 分配一个伪TTY
  - -u 或 --user: 指定运行命令的用户
  - -e 或 --env: 设置环境变量

# Dockerfile

