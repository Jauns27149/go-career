/*
长度最小的子数组

给定一个含有 n 个正整数的数组和一个正整数 target 。
找出该数组中满足其总和大于等于 target 的长度最小的子数组
[numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0
*/
package main

import (
	"fmt"
	"sort"
)

type minSubArrayLenExample struct {
	nums   []int
	target int
	result int
}

func main() {
	examples := []minSubArrayLenExample{
		{[]int{1, 2, 3, 4, 5}, 11, 3},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1}, 11, 0},
	}
	for _, e := range examples {
		fmt.Printf("计算：%v\n答案:%v\n-------\n", minSubArrayLen(e.target, e.nums), e.result)
	}
}

/*
1. 排序，计算 i 到 i+1 的差值存储在新的 []int 中
2. 从子数组最小长度开始循环
*/
func minSubArrayLen(target int, nums []int) int {
	sort.Ints(nums)

	for i := 1; i < target && i < len(nums); i++ {
		if i == 1 {
			if nums[len(nums)-1] < target || target < nums[0] {
				continue
			}
			for _, v := range nums {
				if v == target {
					return 1
				}
				if v > target {
					break
				}
			}
			continue
		}
		for j, v := range nums {
			
		}
	}
	return 0
}
