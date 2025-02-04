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

## os

### create

```go
func Create(name string) (file *File, err error)

/*
Create采用模式0666（任何人都可读写，不可执行）创建一个名为name的文件，如果文件已存在会截断它（为空文件）。
如果成功，返回的文件对象可用于I/O；对应的文件描述符具有O_RDWR模式。如果出错，错误底层类型是*PathError。
```

### *File

####  Readdirnames

```go
func (f *File) Readdirnames(n int) (names []string, err error)
/*
Readdir读取目录f的内容，返回一个有n个成员的[]string，切片成员为目录中文件对象的名字，采用目录顺序。
对本函数的下一次调用会返回上一次调用剩余未读取的内容的信息。

如果n>0，Readdir函数会返回一个最多n个成员的切片。这时，如果Readdir返回一个空切片，它会返回一个非nil的错误说明原因。
如果到达了目录f的结尾，返回值err会是io.EOF。

如果n<=0，Readdir函数返回目录中剩余所有文件对象的名字构成的切片。
此时，如果Readdir调用成功（读取所有内容直到结尾），它会返回该切片和nil的错误值。
如果在到达结尾前遇到错误，会返回之前成功读取的名字构成的切片和该错误。
```

## os/exec

### Command

```go
func Command(name string, arg ...string) *Cmd
/*
函数返回一个*Cmd，用于使用给出的参数执行name指定的程序。返回值只设定了Path和Args两个参数。
如果name不含路径分隔符，将使用LookPath获取完整路径；否则直接使用name。参数arg不应包含命令名。
```

### *Cmd

#### OutPut

```go
func (c *Cmd) Output() ([]byte, error)
/*
执行命令并返回标准输出的切片。
```



#### Run

```go
func (c *Cmd) Run() error
/*
Run执行c包含的命令，并阻塞直到完成。
如果命令成功执行，stdin、stdout、stderr的转交没有问题，并且返回状态码为0，方法的返回值为nil；
如果命令没有执行或者执行失败，会返回*ExitError类型的错误；否则返回的error可能是表示I/O问题。
```

