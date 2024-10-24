package main

import "fmt"

/*
快手外包面试题目
在一个递增数列nums中,寻找是否存在target值，如果存在返回下标，如果不存在返回-1
*/
func main() {
	s := []sample{
		{[]int{0, 2, 4, 5, 8}, 2, 1},
		{[]int{1, 2, 3, 5, 6}, 4, -1},
	}
	for _, v := range s {
		fmt.Printf("计算:%v\t-> 结果:%v\n", find(v.nums, v.target), v.result)
	}
}

func find(nums []int, target int) int {
	l, r := 0, len(nums)-1
	if target < nums[l] || target > nums[r] {
		return -1
	}

	switch target {
	case nums[l]:
		return l
	case nums[r]:
		return r
	default:
		for l+1 < r {
			m := l + (r-l)/2
			if target == nums[m] {
				return m
			}
			if target < nums[m] {
				r = m
			}
			if target > nums[m] {
				l = m
			}
		}
	}
	return -1
}

type sample struct {
	nums   []int
	target int
	result int
}
