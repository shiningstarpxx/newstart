/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: binary_search_lt162.go
 * @Version: 1.0.0
 * @Data: 2021/9/15 4:15 PM
 */
package leetcode

import "math"

func findPeakElement(nums []int) int {
	get := func(i int) int {
		if i < 0 || i >= len(nums) {
			return math.MinInt32
		}
		return nums[i]
	}

	i, j := 0, len(nums) - 1

	for i <= j {
		mid := i + (j - i) / 2
		if get(mid) > get(mid - 1) && get(mid) > get(mid + 1) {
			return mid
		}
		if get(mid) < get(mid + 1) {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return 0
}
