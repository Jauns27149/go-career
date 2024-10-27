/*
百度外包面试遇到题目
启动两个协程交叉打印1～100，
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	job := sync.WaitGroup{}
	job.Add(2)
	first := make(chan int, 1)
	first <- 1
	second := make(chan int)

	go func() {
		for i := 0; i < 50; i++ {
			v := <-first
			fmt.Printf("第1个函数：->%v\n", v)
			second <- v + 1
		}
		close(second)
		job.Done()
	}()

	go func() {
		for v := range second {
			fmt.Printf("第2个函数：->%v\n", v)
			first <- v + 1
		}
		job.Done()
	}()

	job.Wait()
}
