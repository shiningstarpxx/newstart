/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: double_pointer_lt611.go
 * @Version: 1.0.0
 * @Data: 2021/8/4 9:53 AM
 */
package leetcode

import (
	"math"
	"sort"
)

func triangleNumber(nums []int) int {
	return binarySearchTwoPointers(nums)
}

func naiveTwoPointers(nums []int) int {
	sort.Ints(nums)
	ans := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			k := j + 1
			for k < len(nums) {
				if nums[i] + nums[j] <= nums[k] {
					break
				}
				k++
			}
			ans += k - j - 1
		}
	}
	return ans
}

func binarySearchTwoPointers(nums []int) int {
	sort.Ints(nums)
	ans := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			k := sort.SearchInts(nums, nums[i] + nums[j])
			ans += int(math.Max(float64(k - j - 1), float64(0)))
		}
	}
	return ans
}