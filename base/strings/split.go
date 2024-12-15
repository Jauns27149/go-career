package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "a,b,c,"
	ss := strings.Split(s, ",")
	for i, v := range ss {
		fmt.Println(i, v)
	}
}
