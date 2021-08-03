/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: sweep_line_lt587.go
 * @Version: 1.0.0
 * @Data: 2021/8/3 2:07 PM
 */
package leetcode

import "math"

func findUnsortedSubarray(nums []int) int {
	c := math.MinInt32
	right := -1
	for i := 0; i < len(nums); i++ {
		if c <= nums[i] {
			c = nums[i]
		} else {
			right = i
		}
	}

	left := -1
	c = math.MaxInt32
	for i := len(nums)- 1; i >= 0; i-- {
		if c >= nums[i] {
			c = nums[i]
		} else {
			left = i
		}
	}

	if left == -1 {
		return 0
	}

	return right - left + 1
}
