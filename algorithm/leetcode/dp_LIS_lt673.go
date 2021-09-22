/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: dp_LIS_lt673.go
 * @Version: 1.0.0
 * @Data: 2021/9/20 5:35 PM
 */
package leetcode

func findNumberOfLIS(nums []int) int {
	n := len(nums)

	dp := make([]int, n + 1)  // 长度
	cnt := make([]int, n + 1) // 长度下的count
	maxLen := 0

	for i := 0; i < n; i++ {
		dp[i] = 1
		cnt[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j] + 1 > dp[i] {
					dp[i] = dp[j] + 1
					cnt[i] = cnt[j]
				} else if dp[j] + 1 == dp[i] {
					cnt[i] += cnt[j]
				}
			}
		}

		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		if dp[i] == maxLen {
			ans += cnt[i]
		}
	}
	return ans
}
