package main

import (
	"fmt"
	"golang.org/x/net/context"
	"time"
)

func longRunningTask(ctx context.Context) {
	select {
	case <-time.After(time.Second * 5):
		fmt.Println("Task completed")
	case <-ctx.Done():
		fmt.Println("Task cancelled:", ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go longRunningTask(ctx)
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(10 * time.Second)
}
