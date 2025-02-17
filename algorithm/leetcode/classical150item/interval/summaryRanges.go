/*
给定一个无重复元素的有序整数数组 nums 。
返回恰好覆盖数组中所有数字的最小有序区间范围列表 。
也就是说，nums 的每个元素都恰好被某个区间范围所覆盖，并且不存在属于某个范围但不属于 nums 的数字 x 。
列表中的每个区间范围 [a,b] 应该按如下格式输出：
"a->b" ，如果 a != b
"a" ，如果 a == b

示例 1：
输入：nums = [0,1,2,4,5,7]
输出：["0->2","4->5","7"]
解释：区间范围是：
[0,2] --> "0->2"
[4,5] --> "4->5"
[7,7] --> "7"
示例 2：
输入：nums = [0,2,3,4,6,8,9]
输出：["0","2->4","6","8->9"]
解释：区间范围是：
[0,0] --> "0"
[2,4] --> "2->4"
[6,6] --> "6"
[8,9] --> "8->9"

提示：
0 <= nums.length <= 20
-231 <= nums[i] <= 231 - 1
nums 中的所有值都 互不相同
nums 按升序排列
*/
package main

import (
	"go-career/algorithm/leetcode"
	"strconv"
)

type examplSummaryRanges struct {
	nums   []int
	answer []string
}

func main() {
	s := []examplSummaryRanges{{leetcode.ToIntSlice("[0,1,2,4,5,7]"), []string{"0->2", "4->5", "7"}}}
	for _, v := range s {
		leetcode.Printf(summaryRanges(v.nums), v.answer)
	}
}
func summaryRanges(nums []int) []string {
	answer, start := make([]string, 0, len(nums)), 0
	for i, v := range nums {
		if i != len(nums)-1 && v+1 == nums[i+1] {
			continue
		}
		if i != start {
			answer = append(answer, strconv.Itoa(nums[start])+"->"+strconv.Itoa(v))
		} else {
			answer = append(answer, strconv.Itoa(v))
		}
		start = i + 1
	}
	return answer
}
