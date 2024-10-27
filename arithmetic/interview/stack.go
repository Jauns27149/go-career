package main

import (
	"bytes"
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 启动一个goroutine来模拟并发执行
	go func() {
		for {
			select {
			case <-time.After(1 * time.Second):
				fmt.Println("Goroutine is running...")
			}
		}
	}()

	// 等待一段时间，让goroutine运行起来
	time.Sleep(2 * time.Second)

	// 获取所有goroutine的调用栈
	buf := make([]byte, 1024)
	n := runtime.Stack(buf, true)
	fmt.Println("Current goroutine stack trace:")
	fmt.Println(string(buf[:n]))

	// 获取所有goroutine的调用栈信息
	printAllGoroutineStacks()
}

func printAllGoroutineStacks() {
	buf := make([]byte, 1024)
	for i := 0; ; i++ {
		n := runtime.Stack(buf, true)
		if n < 1024 {
			break
		}
		buf = make([]byte, 2*len(buf))
	}

	// 获取所有goroutine的调用栈信息
	buf = buf[:runtime.Stack(buf, true)]
	stacks := splitStacks(buf)

	for _, stack := range stacks {
		fmt.Println(stack)
	}
}

// splitStacks 分割goroutine的调用栈信息
func splitStacks(buf []byte) []string {
	var stacks []string
	for _, stack := range bytes.Split(buf, []byte("goroutine ")) {
		if len(stack) > 0 {
			stacks = append(stacks, string(stack))
		}
	}
	return stacks
}

// 打印goroutine的调用栈信息
func printGoroutineStacks() {
	buf := make([]byte, 1024)
	n := runtime.Stack(buf, true)
	fmt.Println("Current goroutine stack trace:")
	fmt.Println(string(buf[:n]))
}
