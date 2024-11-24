package main

import "fmt"

type ninWindow struct {
	s      string
	t      string
	answer string
}

func main() {
	examples := []ninWindow{
		{"ADOBECODEBANC", "ABC", "BANC"},
	}
	for _, ex := range examples {
		fmt.Println(minWindow(ex.s, ex.t), "\t->", ex.answer)
	}
}

func minWindow(s string, t string) string {
	mapt := map[rune]int{}
	for _, v := range t {
		mapt[v]--
	}
	var targetSlice []int
	flag := len(t)
	for i, v := range s {
		if c, ok := mapt[v]; ok {
			if c < 0 {
				flag--
			}
			mapt[v]++
			targetSlice = append(targetSlice, i)
		}
	}
	if flag > 0 {
		return ""
	}
	targetSlice = findMinWindow(s, mapt, targetSlice)
	return s[targetSlice[0] : targetSlice[len(targetSlice)-1]+1]
}

func findMinWindow(s string, mapt map[rune]int, slice []int) []int {
	end := len(slice) - 1
	startCount := mapt[rune(s[slice[0]])]
	endCount := mapt[rune(s[slice[end]])]
	if startCount == 0 && endCount == 0 {
		return slice
	}
	leftSlice, rightSlice := slice, slice
	if startCount > 0 {
		m := map[rune]int{}
		for k, v := range mapt {
			m[k] = v
		}
		m[rune(s[slice[0]])]--
		leftSlice = findMinWindow(s, m, slice[1:end+1])
	}
	if endCount > 0 {
		m := map[rune]int{}
		for k, v := range mapt {
			m[k] = v
		}
		m[rune(s[slice[end]])]--
		rightSlice = findMinWindow(s, m, slice[:end])
	}
	left := leftSlice[len(leftSlice)-1] - leftSlice[0] + 1
	right := rightSlice[len(rightSlice)-1] - rightSlice[0] + 1
	if left < right {
		return leftSlice
	} else {
		return rightSlice
	}
}
