# context

​	context 包非常适用于需要跨多个 goroutine 或函数调用传递请求相关的信息。

- 作用：
  - 基本的取消操作
  - 使用超时控制
  - 使用上下文传递值

go 命令

```bash
$env:GOPRIVATE="work.ctyun.cn"
$env:GOPROXY="direct"
# 配置私库，powershell
```

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

## builtin

### max

```go
func max[T cmp.Ordered](x T, y ...T) T
/*	
max 函数是一个内置函数，对于可比较大小的类型（cmp.Ordered），它返回固定数量参数中的最大值。
必须至少提供一个参数。如果类型 T 是浮点数类型并且任意参数是 NaN（非数字），那么 max 函数将返回 NaN。
```

```go
func main() {
	result := max(10, 20, 30, 40, 50)
	fmt.Printf("最大值是: %d\n", result)
}
/*
最大值是: 50
```

### 接口类型转换为具体类型

- 类型断言： `value.(type)` 形式来尝试将接口类型的值转换为指定的具体类型。如果转换成功，返回转换后的值；如果失败，则接收两个会返回 type 的默认零值和flase，接受一个会panic.
- 类型切换： `switch value.(type)`，它会尝试匹配不同的类型，并在找到匹配的类型后执行相应的代码块。

## strings

### Split

```go
func Split(s, sep string) []string
/*sep作为分隔符,把s切割为[]string
fmt.Print(strings.Split("a,b,c", ",")) 
Output:
[a b c]
```

### Contains

```go
func Contains(s string, substr string) bool
```

```go
// Contains reports whether substr is within s
fmt.Println(strings. Contains("seafood", "foo")) 
// true
```

### ReplaceAll

```go
func ReplaceAll(s string, old string, new string) string
```

```go
fmt.Println(strings. ReplaceAll("oink oink oink", "oink", "moo")) 
// moo moo moo
```

### TrimSpace

```go
func TrimSpace(s string) string
```

```go
fmt.Println(strings. TrimSpace(" \t\n Hello, Gophers \n\t\r\n")) 
// Hello, Gophers
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



## sort

### Slice

```go
func Slice(x any, less func(i int, j int) bool)
/*	Slice 函数根据提供的比较函数对切片 x 进行排序。如果 x 不是一个切片，它会引发恐慌（panic）。
		排序不保证稳定性：相等的元素可能会从原始顺序中颠倒。若需要稳定的排序，请使用 SliceStable
```

```go
func main() {
	people := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("按年龄排序:", people)
}
/*
按名字排序: [{Alice 55} {Bob 75} {Gopher 7} {Vera 24}]
按年龄排序: [{Gopher 7} {Vera 24} {Alice 55} {Bob 75}]
```

