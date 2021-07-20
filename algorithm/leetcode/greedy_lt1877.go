/**
 * @Author: peixingxin
 * @Description:
 * @File: greedy_lt1877.go
 * @Version: 1.0.0
 * @Data: 2021/7/20 12:51 PM
 */
package leetcode

import (
	"math"
	"sort"
)

func minPairSum(nums []int) int {
	sort.Ints(nums)
	ans := 0
	for i := 0; i < len(nums) / 2; i++ {
		ans = int(math.Max(float64(ans), float64(nums[i] + nums[len(nums) - i - 1])))
	}
	return ans
}