# context

​	context 包非常适用于需要跨多个 goroutine 或函数调用传递请求相关的信息。

- 作用：
  - 基本的取消操作
  - 使用超时控制
  - 使用上下文传递值

go 命令

```bash

```

- - 

# 内联优化

​	当 Go 编译器识别出某个函数适合内联时（通常是短小且简单的函数），它不会生成一个独立的函数调用指令，而是直接将该函数的代码复制到所有调用它的位置。

```go
// 未优化前
func add(a, b int) int {
    return a + b
}

func main() {
    result := add(1, 2)
    fmt.Println(result)
}

// 优化后
func main() {
    result := 1 + 2 // 直接计算结果
    fmt.Println(result)
}
```

# go命令

```bash
go <command> [arguments]

arguments
	-gcflags=-l		禁用所有函数的内联优化
```

## build

```bash
go build -o janus
```

## env

| 变量名        | 描述                                                  |
| ------------- | ----------------------------------------------------- |
| `GOOS`        | 目标操作系统的名称（如 `linux`, `windows`, `darwin`） |
| `GOARCH`      | 目标架构（如 `amd64`, `arm64`）                       |
| `GOPATH`      | Go 工作区路径，默认为 `$HOME/go`                      |
| `GOROOT`      | Go 安装路径                                           |
| `CGO_ENABLED` | 是否启用 CGO（`1` 表示启用，`0` 表示禁用）            |
| `GO111MODULE` | 控制模块支持的行为（`on`, `off`, `auto`）             |
| `GOMODCACHE`  | 模块缓存路径                                          |
| `GOPROXY`     | Go 模块代理地址                                       |
| `GONOPROXY`   | 不通过代理获取模块的模式列表                          |
| `GONOSUMDB`   | 不检查校验和数据库的模块列表                          |

```bash
go env # 查看环境变量
	-w 设置环境变量值
```

# 标准包

## strings

### Split

```go
func Split(s, sep string) []string
/*sep作为分隔符,把s切割为[]string
fmt.Print(strings.Split("a,b,c", ",")) 
Output:
[a b c]
```

## slices

### Sort

```go
func Sort[S interface{ ~[] E }, E cmp.Ordered](x S)
/* 
smallInts := []int8{0, 42, -10, 8}
slices.Sort(smallInts) 
fmt.Println(smallInts) 
Output:
[-10 0 8 42]
```



