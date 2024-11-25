package main

import "fmt"

type ninWindow struct {
	s      string
	t      string
	answer string
}

func main() {
	examples := []ninWindow{
		{"cabwefgewcwaefgcf", "cae", "cwae"},
		{"ADOBECODEBANC", "ABC", "BANC"},
	}
	for _, ex := range examples {
		fmt.Println(minWindow(ex.s, ex.t), "\t->", ex.answer)
	}
}

func minWindow(s string, t string) string {
	m := map[rune]int{}
	for _, v := range t {
		m[v]--
	}
	var is []int
	flag := len(t)
	for i, v := range s {
		if c, ok := m[v]; ok {
			if c < 0 {
				flag--
			}
			m[v]++
			is = append(is, i)
		}
	}
	if flag > 0 {
		return ""
	}

	start, end := 0, len(is)-1
	if v, _ := m[rune(s[is[start]])]; v == 0 {
		v, _ = m[rune(s[is[end]])]
		for v != 0 {
			m[rune(s[is[end]])]--
			end--
			v, _ = m[rune(s[is[end]])]
		}
		return s[is[start] : is[end]+1]
	}

	if v, _ := m[rune(s[is[end]])]; v == 0 {
		v, _ = m[rune(s[is[start]])]
		for v != 0 {
			m[rune(s[is[start]])]--
			start++
			v, _ = m[rune(s[is[start]])]
		}
		return s[is[start] : is[end]+1]
	}
	tm := map[rune]int{}
	for k, v := range m {
		tm[k] = v
	}

	v, _ := m[rune(s[is[start]])]
	for v != 0 {
		m[rune(s[is[start]])]--
		start++
		v, _ = m[rune(s[is[start]])]
	}
	v, _ = m[rune(s[is[end]])]
	for v != 0 {
		m[rune(s[is[end]])]--
		end--
		v, _ = m[rune(s[is[end]])]
	}
	answerS := s[is[start] : is[end]+1]

	start, end = 0, len(is)-1
	v, _ = tm[rune(s[is[end]])]
	for v != 0 {
		tm[rune(s[is[end]])]--
		end--
		v, _ = tm[rune(s[is[end]])]
	}
	v, _ = tm[rune(s[is[start]])]
	for v != 0 {
		m[rune(s[is[start]])]--
		start++
		v, _ = tm[rune(s[is[start]])]
	}
	answerE := s[is[start] : is[end]+1]

	if len(answerE) < len(answerS) {
		return answerE
	} else {
		return answerS
	}

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
