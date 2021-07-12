/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  11:59 AM  12/7/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

func subarraysWithKDistinct(nums []int, k int) int {
	return withDistinctK(nums, k) - withDistinctK(nums, k - 1)
}

func withDistinctK(nums []int, k int) int {
	n := len(nums)
	left := 0
	right := 0
	count := 0
	cache := map[int]int{}

	res := 0
	for right < n {
		if cache[nums[right]] == 0 {
			count++
		}
		cache[nums[right]]++
		right++

		for count > k {
			cache[nums[left]]--
			if cache[nums[left]] == 0 {
				count--
			}
			left++
		}
		res += right - left
	}

	return res
}
