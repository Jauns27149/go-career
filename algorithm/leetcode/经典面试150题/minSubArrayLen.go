/*
长度最小的子数组

给定一个含有 n 个正整数的数组和一个正整数 target 。
找出该数组中满足其总和大于等于 target 的长度最小的子数组
[numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0
*/
package main

import (
	"fmt"
)

type minSubArrayLenExample struct {
	nums   []int
	target int
	result int
}

func main() {
	examples := []minSubArrayLenExample{
		{[]int{12, 28, 83, 4, 25, 26, 25, 2, 25, 25, 25, 12}, 213, 8},
		{[]int{2, 16, 14, 15}, 20, 2},
		{[]int{1, 2, 3, 4, 5}, 15, 5},
		{[]int{2, 3, 1, 2, 4, 3}, 7, 2},
		{[]int{1, 2, 3, 4, 5}, 11, 3},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1}, 11, 0},
	}
	for _, e := range examples {
		fmt.Printf("计算：%v\n答案:%v\n-------\n", minSubArrayLen(e.target, e.nums), e.result)
	}
}

/*
1.
*/
func minSubArrayLen(target int, nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return calculate(target, nums, total)
}

func calculate(target int, nums []int, total int) int {
	if total >= target {
		l := calculate(target, nums[:len(nums)-1], total-nums[len(nums)-1])
		r := calculate(target, nums[1:], total-nums[0])
		if l != 0 && r != 0 {
			if l < r {
				return l
			} else {
				return r
			}
		}
		if l == 0 && r == 0 {
			return len(nums)
		}
		if l == 0 {
			return r
		}
		return l
	} else {
		return 0
	}
}
