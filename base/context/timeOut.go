package main

import (
	"fmt"
	"golang.org/x/net/context"
	"time"
)

func longRunningTest(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Task completed")
	case <-ctx.Done():
		fmt.Println("Task time out:", ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	go longRunningTest(ctx)
	time.Sleep(6 * time.Second)
}
