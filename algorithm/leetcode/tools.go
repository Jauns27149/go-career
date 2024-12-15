package leetcode

import (
	"fmt"
	"strconv"
)

// s := "[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]"
func To2intSlice(s string) [][]int {
	n := len(s) - 1
	var os []int
	var ss [][]int
	var b []byte
	for i := 1; i < n; i++ {
		r := s[i]
		switch r {
		case '[':
		case ',':
		case ']':
			ss = append(ss, os)
			os = []int{}
		default:
			b = append(b, s[i])
			if s[i+1] == ',' || s[i+1] == ']' {
				v, err := strconv.Atoi(string(b))
				if err != nil {
					panic(err)
				}
				os = append(os, v)
				b = []byte{}
			}
		}
	}
	return ss
}

func Printf(compute, answer interface{}) {
	fmt.Printf("%v\n%v\n-------------------\n", compute, answer)
}
