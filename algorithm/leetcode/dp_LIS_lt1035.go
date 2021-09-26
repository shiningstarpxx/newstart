/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: dp_LIS_lt1035.go
 * @Version: 1.0.0
 * @Data: 2021/9/26 8:34 PM
 */
package leetcode

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	m := len(nums1)
	n := len(nums2)

	dp := make([][]int, m + 1)
	for i := 0; i < m + 1; i++ {
		dp[i] = make([]int, n + 1)
	}

	for i := 0 ; i < m; i++ {
		for j := 0; j < n; j++ {
			if (nums1[i] == nums2[j]) {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max1035(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	return dp[m][n]
}

func max1035(a, b int) int {
	if a > b {
		return a
	}
	return b
}
