/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: dp_lt300.go
 * @Version: 1.0.0
 * @Data: 2021/7/26 12:23 PM
 */
package leetcode

import "math"

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)

	ans := 1
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[j] + 1)))
				ans = int(math.Max(float64(ans), float64(dp[i])))
			}
		}
	}
	return ans
}
