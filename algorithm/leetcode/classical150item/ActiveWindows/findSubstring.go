package main

import "fmt"

type finSubStringExample struct {
	s      string
	words  []string
	answer []int
}

func main() {
	examples := []finSubStringExample{
		{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}, []int{8}},
		{"barfoothefoobarman", []string{"foo", "bar"}, []int{0, 9}},
	}
	for _, ex := range examples {
		fmt.Printf("calculate:\t%v\nanswer:\t\t%v\n------------------------------\n",
			findSubstring(ex.s, ex.words), ex.answer)
	}
}

func findSubstring(s string, words []string) []int {
	wordsMap := make(map[string]int)
	for _, word := range words {
		wordsMap[word]++
	}

	size := len(words[0])
	left, right := 0, 0
	ans := make([]int, 0)
	total := len(words)
	for left+size*len(words) <= len(s) {
		word := s[right : right+size]
		if count, ok := wordsMap[word]; ok {
			if count > 1 {
				wordsMap[word]--
			} else {
				delete(wordsMap, word)
			}
			right += size
			total--
			if total == 0 {
				total = len(words)
				ans = append(ans, left)
				for _, w := range words {
					wordsMap[w]++
				}
				left++
				right = left
			}
		} else {
			left++
			if total < len(words) {
				for ; total < len(words); total++ {
					wordsMap[s[right-size:right]]++
					right -= size
				}
				total = len(words)
			}
			right = left
		}
	}
	return ans
}
