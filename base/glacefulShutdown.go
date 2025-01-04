package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 创建一个新的 HTTP 服务器
	server := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(handleRequest),
	}

	// 启动服务器
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("server ListenAndServe failed: %v", err)
		}
	}()

	// 监听中断信号（如 Ctrl+C）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// 创建一个带有超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 使用 Shutdown 方法优雅地关闭服务器
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server Shutdown failed: %v", err)
	}

	log.Println("Server gracefully stopped")
}

// handleRequest 是一个简单的请求处理函数
func handleRequest(w http.ResponseWriter, r *http.Request) {
	time.Sleep(10 * time.Minute)
	fmt.Fprintf(w, "Hello, World!")
}
