package main

import (
	"fmt"
	"strings"
)

func main() {
	multStatus := []string{"   ", "a    ", "b     ", "c", "d", "e", "f"}
	var newStatus []string
	for _, s := range multStatus {
		if v := strings.TrimSpace(s); v != "" {
			newStatus = append(newStatus, v)
		}
	}
	fmt.Println(newStatus)
}
