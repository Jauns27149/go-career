# context

​	context 包非常适用于需要跨多个 goroutine 或函数调用传递请求相关的信息。

- 作用：
  - 基本的取消操作
  - 使用超时控制
  - 使用上下文传递值

go 命令

```bash
go <command> [arguments]
```

- arguments
  - `-gcflags=-l` : 禁用所有函数的内联优化

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

