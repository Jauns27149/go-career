package main

import (
	"fmt"
	"golang.org/x/net/context"
)

func handler(ctx context.Context) {
	userID := ctx.Value("userID")
	fmt.Println("Request received for user:", userID)
}
func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "userID", "123")
	handler(ctx)
}
